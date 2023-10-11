/*
 * Copyright (c) 2020, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
 *
 * WSO2 LLC. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package org.wso2.apk.enforcer.security;

import com.nimbusds.jwt.JWTClaimsSet;
import net.minidev.json.JSONObject;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.wso2.apk.enforcer.commons.exception.APISecurityException;
import org.wso2.apk.enforcer.commons.exception.EnforcerException;
import org.wso2.apk.enforcer.commons.logging.ErrorDetails;
import org.wso2.apk.enforcer.commons.logging.LoggingConstants;
import org.wso2.apk.enforcer.commons.model.ResourceConfig;
import org.wso2.apk.enforcer.constants.APIConstants;
import org.wso2.apk.enforcer.constants.APISecurityConstants;
import org.wso2.apk.enforcer.constants.GeneralErrorCodeConstants;
import org.wso2.apk.enforcer.dto.APIKeyValidationInfoDTO;
import org.wso2.apk.enforcer.models.API;
import org.wso2.apk.enforcer.models.Application;
import org.wso2.apk.enforcer.models.ApplicationKeyMapping;
import org.wso2.apk.enforcer.models.Subscription;
import org.wso2.apk.enforcer.subscription.SubscriptionDataHolder;
import org.wso2.apk.enforcer.subscription.SubscriptionDataStore;
import org.wso2.apk.enforcer.util.FilterUtils;

import java.util.List;
import java.util.Set;

/**
 * Does the subscription and scope validation.
 */
public class KeyValidator {
    private static final Logger log = LogManager.getLogger(KeyValidator.class);

    /**
     * Validate the scopes related to the given validationContext.
     *
     * @param validationContext the token validation context
     * @return true is the scopes are valid
     * @throws EnforcerException throws if token validation fails.
     *                           this will indicate the message body for the error response
     */
    public static boolean validateScopes(TokenValidationContext validationContext) throws APISecurityException {

        if (validationContext.isCacheHit()) {
            return true;
        }
        APIKeyValidationInfoDTO apiKeyValidationInfoDTO = validationContext.getValidationInfoDTO();

        if (apiKeyValidationInfoDTO == null) {
            log.error("Error while validating scopes. Key validation information has not been set.",
                    ErrorDetails.errorLog(LoggingConstants.Severity.MINOR, 6603));
            throw new APISecurityException(APIConstants.StatusCodes.UNAUTHENTICATED.getCode(),
                    APISecurityConstants.API_AUTH_GENERAL_ERROR,
                    "Error while validating scopes. Key validation information has not been set");
        }

        Set<String> scopesFromToken = apiKeyValidationInfoDTO.getScopes();

        List<ResourceConfig> matchedResources;
        // when it is a graphQL api multiple matching resources will be returned.
        matchedResources = validationContext.getMatchingResourceConfigs();

        boolean allScopesValidated = true;
        // failedResourcePath - used to identify resource paths with failed scope validation.
        String failedResourcePath = "";
        for (ResourceConfig matchedResource : matchedResources) {
            // scopesValidated - indicate scope has validated
            boolean scopesValidated = false;
            String resourcePath = matchedResource.getPath();
            String[] scopesToValidate = matchedResource.getScopes();
            for (String scope : scopesToValidate) {
                if (scopesFromToken.contains(scope)) {
                    scopesValidated = true;
                    break;
                }
            }
            if (scopesToValidate.length > 0 && !scopesValidated) {
                allScopesValidated = false;
                failedResourcePath = resourcePath;
                break;
            }
        }
        if (!allScopesValidated) {
            apiKeyValidationInfoDTO.setAuthorized(false);
            apiKeyValidationInfoDTO.setValidationStatus(APIConstants.KeyValidationStatus.INVALID_SCOPE);
            String message = "User is NOT authorized to access the Resource: " + failedResourcePath
                    + ". Scope validation failed.";
            throw new APISecurityException(APIConstants.StatusCodes.UNAUTHORIZED.getCode(),
                    APISecurityConstants.INVALID_SCOPE, message);
        }
        return true;
    }

    /**
     * Validate subscriptions for access tokens.
     *
     * @param uuid        uuid of the API
     * @param apiContext  API context, used for logging purposes and to extract the tenant domain
     * @param apiVersion  API version, used for logging purposes
     * @param consumerKey consumer key related to the token
     * @param keyManager  key manager related to the token
     * @return validation information about the request
     */
    public static APIKeyValidationInfoDTO validateSubscription(String uuid, String apiContext, String apiVersion,
                                                               String consumerKey, String envType, String keyManager) {
        log.debug("Before validating subscriptions");
        log.debug("Validation Info : { uuid : {}, context : {}, version : {}, consumerKey : {} }",
                uuid, apiContext, apiVersion, consumerKey);
        String apiTenantDomain = FilterUtils.getTenantDomainFromRequestURL(apiContext);
        if (apiTenantDomain == null) {
            apiTenantDomain = APIConstants.SUPER_TENANT_DOMAIN_NAME;
        }

        API api = null;
        ApplicationKeyMapping key = null;
        Application app = null;
        Subscription sub = null;

        SubscriptionDataStore datastore = SubscriptionDataHolder.getInstance()
                .getTenantSubscriptionStore(apiTenantDomain);
        //TODO add a check to see whether datastore is initialized an load data using rest api if it is not loaded
        // TODO: (VirajSalaka) Handle the scenario where the event is dropped.
        if (datastore != null) {
            api = datastore.getApiByContextAndVersion(uuid);
            if (api != null) {
                // TODO: (Sampath) Handle the scenario when App keys are generated properly and sent
//                key = datastore.getKeyMappingByKeyAndKeyManager(consumerKey, keyManager);
//                if (key != null) {
                app = datastore.getApplicationById(key.getApplicationUUID());
                if (app != null) {
                    sub = datastore.getSubscriptionById(app.getUUID(), api.getApiUUID());
                    if (sub != null) {
                        log.debug("All information is retrieved from the inmemory data store.");
                    } else {
                        log.info(
                                "Valid subscription not found for oauth access token. " +
                                        "application: {} app_UUID: {} API_name: {} API_UUID : {}",
                                app.getName(), app.getUUID(), api.getApiName(), api.getApiUUID());
                    }
                } else {
                    log.info("Application not found in the data store for uuid " + key.getApplicationUUID());
                }
//                } else {
//                    log.info("Application key mapping not found in the data store for id consumerKey " + consumerKey);
//                }
            } else {
                log.info("API not found in the data store for API UUID :" + uuid);
            }
        } else {
            log.error("Subscription data store is null for tenant domain " + apiTenantDomain);
        }

        APIKeyValidationInfoDTO infoDTO = new APIKeyValidationInfoDTO();
        if (api != null && app != null && sub != null) {
            validate(infoDTO, datastore, api, envType, app, sub);
        }
        if (!infoDTO.isAuthorized() && infoDTO.getValidationStatus() == 0) {
            //Scenario where validation failed and message is not set
            infoDTO.setValidationStatus(APIConstants.KeyValidationStatus.API_AUTH_RESOURCE_FORBIDDEN);
        }
        log.debug("After validating subscriptions");
        return infoDTO;
    }

    /**
     * Validate subscriptions for API keys.
     *
     * @param apiUuid    uuid of the API
     * @param apiContext API context, used for logging purposes and to extract the tenant domain
     * @param payload    JWT claims set extracted from the API key
     * @return validation information about the request
     */
    public static APIKeyValidationInfoDTO validateSubscription(String apiUuid, String apiContext,
                                                               JWTClaimsSet payload, String envType) {
        log.debug("Before validating subscriptions with API key. API_uuid: {}, context: {}", apiUuid, apiContext);
        String apiTenantDomain = FilterUtils.getTenantDomainFromRequestURL(apiContext);
        if (apiTenantDomain == null) {
            apiTenantDomain = APIConstants.SUPER_TENANT_DOMAIN_NAME;
        }

        API api = null;
        Application app = null;
        Subscription sub = null;

        SubscriptionDataStore datastore = SubscriptionDataHolder.getInstance()
                .getTenantSubscriptionStore(apiTenantDomain);
        //TODO add a check to see whether datastore is initialized an load data using rest api if it is not loaded
        // TODO: (VirajSalaka) Handle the scenario where the event is dropped.
        if (datastore != null) {
            api = datastore.getApiByContextAndVersion(apiUuid);
            if (api != null) {
                JSONObject appObject = (JSONObject) payload.getClaim(APIConstants.JwtTokenConstants.APPLICATION);
                String appUuid = appObject.getAsString("uuid");
                if (!appObject.isEmpty() && !appUuid.isEmpty()) {
                    app = datastore.getApplicationById(appUuid);
                    if (app != null) {
                        sub = datastore.getSubscriptionById(app.getUUID(), api.getApiUUID());
                        if (sub != null) {
                            log.debug("All information is retrieved from the in memory data store.");
                        } else {
                            log.info(
                                    "Valid subscription not found for API key. " +
                                            "application: {} app_UUID: {} API_name: {} API_UUID : {}",
                                    app.getName(), app.getUUID(), api.getApiName(), api.getApiUUID());
                        }
                    } else {
                        log.info("Application not found in the data store for uuid {}", appUuid);
                    }
                } else {
                    log.info("Application claim not found in jwt for uuid");
                }
            } else {
                log.info("API not found in the data store for API UUID :" + apiUuid);
            }
        } else {
            log.error("Subscription data store is null for tenant domain " + apiTenantDomain);
        }

        APIKeyValidationInfoDTO infoDTO = new APIKeyValidationInfoDTO();
        if (api != null && app != null && sub != null) {
            validate(infoDTO, datastore, api, envType, app, sub);
        }
        if (!infoDTO.isAuthorized() && infoDTO.getValidationStatus() == 0) {
            //Scenario where validation failed and message is not set
            infoDTO.setValidationStatus(APIConstants.KeyValidationStatus.API_AUTH_RESOURCE_FORBIDDEN);
        }
        log.debug("After validating subscriptions with API key.");
        return infoDTO;
    }

    private static void validate(APIKeyValidationInfoDTO infoDTO, SubscriptionDataStore datastore,
                                 API api, String keyType, Application app, Subscription sub) {
        String subscriptionStatus = sub.getSubscriptionStatus();
        if (APIConstants.SubscriptionStatus.BLOCKED.equals(subscriptionStatus)) {
            infoDTO.setValidationStatus(APIConstants.KeyValidationStatus.API_BLOCKED);
            infoDTO.setAuthorized(false);
            return;
        } else if (APIConstants.SubscriptionStatus.ON_HOLD.equals(subscriptionStatus)
                || APIConstants.SubscriptionStatus.REJECTED.equals(subscriptionStatus)) {
            infoDTO.setValidationStatus(APIConstants.KeyValidationStatus.SUBSCRIPTION_INACTIVE);
            infoDTO.setAuthorized(false);
            return;
        } else if (APIConstants.SubscriptionStatus.PROD_ONLY_BLOCKED.equals(subscriptionStatus)
                && !APIConstants.API_KEY_TYPE_SANDBOX.equals(keyType)) {
            infoDTO.setValidationStatus(APIConstants.KeyValidationStatus.API_BLOCKED);
            infoDTO.setType(keyType);
            infoDTO.setAuthorized(false);
            return;
        } else if (APIConstants.LifecycleStatus.BLOCKED.equals(api.getLcState())) {
            infoDTO.setValidationStatus(GeneralErrorCodeConstants.API_BLOCKED_CODE);
            infoDTO.setAuthorized(false);
            return;
        }
        infoDTO.setTier(sub.getPolicyId());
        infoDTO.setSubscriber(app.getSubName());
        //infoDTO.setApplicationId(app.getId());
        infoDTO.setApplicationUUID(app.getUUID());
        infoDTO.setApiName(api.getApiName());
        infoDTO.setApiVersion(api.getApiVersion());
        infoDTO.setApiPublisher(api.getApiProvider());
        infoDTO.setApplicationName(app.getName());
        infoDTO.setApplicationTier(app.getPolicy());
        infoDTO.setAppAttributes(app.getAttributes());
        infoDTO.setApiUUID(api.getApiUUID());
        infoDTO.setType(keyType);
        infoDTO.setSubscriberTenantDomain(app.getTenantDomain());

        // Todo: (Sampath) This must be implemented as a part rate plans implementation.
//        ApplicationPolicy appPolicy = datastore.getApplicationPolicyByName(app.getPolicy());
//        SubscriptionPolicy subPolicy = datastore.getSubscriptionPolicyByName(sub.getPolicyId());
//        ApiPolicy apiPolicy = datastore.getApiPolicyByName(api.getApiTier());
//        boolean isContentAware = appPolicy.isContentAware() || subPolicy.isContentAware() ||
//                (apiPolicy != null && apiPolicy.isContentAware());
//        infoDTO.setContentAware(isContentAware);
//        int spikeArrest = 0;
//        String apiLevelThrottlingKey = "api_level_throttling_key";
//
//        if (subPolicy.getRateLimitCount() > 0) {
//            spikeArrest = subPolicy.getRateLimitCount();
//        }
//
//        String spikeArrestUnit = null;
//
//        if (subPolicy.getRateLimitTimeUnit() != null) {
//            spikeArrestUnit = subPolicy.getRateLimitTimeUnit();
//        }
//        boolean stopOnQuotaReach = subPolicy.isStopOnQuotaReach();
//        int graphQLMaxDepth = Math.max(subPolicy.getGraphQLMaxDepth(), 0);
//        int graphQLMaxComplexity = Math.max(subPolicy.getGraphQLMaxComplexity(), 0);
//        List<String> list = new ArrayList<>();
//        list.add(apiLevelThrottlingKey);
//        infoDTO.setSpikeArrestLimit(spikeArrest);
//        infoDTO.setSpikeArrestUnit(spikeArrestUnit);
//        infoDTO.setStopOnQuotaReach(stopOnQuotaReach);
//        infoDTO.setGraphQLMaxDepth(graphQLMaxDepth);
//        infoDTO.setGraphQLMaxComplexity(graphQLMaxComplexity);
//        // We also need to set throttling data list associated with given API. This need to have
//        // policy id and
//        // condition id list for all throttling tiers associated with this API.
//        infoDTO.setThrottlingDataList(list);
        infoDTO.setAuthorized(true);
    }
}
