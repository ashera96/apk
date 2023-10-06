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

package org.wso2.apk.enforcer.subscription;

import org.apache.commons.lang3.StringUtils;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.wso2.apk.enforcer.commons.dto.ClaimMappingDto;
import org.wso2.apk.enforcer.commons.dto.JWKSConfigurationDTO;
import org.wso2.apk.enforcer.commons.exception.EnforcerException;
import org.wso2.apk.enforcer.config.dto.ExtendedTokenIssuerDto;
import org.wso2.apk.enforcer.discovery.ApiListDiscoveryClient;
import org.wso2.apk.enforcer.discovery.ApplicationDiscoveryClient;
import org.wso2.apk.enforcer.discovery.ApplicationKeyMappingDiscoveryClient;
import org.wso2.apk.enforcer.discovery.ApplicationMappingDiscoveryClient;
import org.wso2.apk.enforcer.discovery.ApplicationPolicyDiscoveryClient;
import org.wso2.apk.enforcer.discovery.JWTIssuerDiscoveryClient;
import org.wso2.apk.enforcer.discovery.SubscriptionDiscoveryClient;
import org.wso2.apk.enforcer.discovery.SubscriptionPolicyDiscoveryClient;
import org.wso2.apk.enforcer.discovery.subscription.APIs;
import org.wso2.apk.enforcer.discovery.subscription.Certificate;
import org.wso2.apk.enforcer.discovery.subscription.JWTIssuer;
import org.wso2.apk.enforcer.models.API;
import org.wso2.apk.enforcer.models.ApiPolicy;
import org.wso2.apk.enforcer.models.Application;
import org.wso2.apk.enforcer.models.ApplicationKeyMapping;
import org.wso2.apk.enforcer.models.ApplicationMapping;
import org.wso2.apk.enforcer.models.ApplicationPolicy;
import org.wso2.apk.enforcer.models.Subscription;
import org.wso2.apk.enforcer.models.SubscriptionPolicy;
import org.wso2.apk.enforcer.security.jwt.validator.JWTValidator;
import org.wso2.apk.enforcer.util.TLSUtils;

import java.io.IOException;
import java.security.cert.CertificateException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

/**
 * Implementation of the subscription data store.
 */
public class SubscriptionDataStoreImpl implements SubscriptionDataStore {

    private static final Logger log = LogManager.getLogger(SubscriptionDataStoreImpl.class);
    private static final SubscriptionDataStoreImpl instance = new SubscriptionDataStoreImpl();

    /**
     * ENUM to hold type of policies.
     */
    public enum PolicyType {
        SUBSCRIPTION,
        APPLICATION,
        API
    }

    public static final String DELEM_PERIOD = ":";

    // Maps for keeping Subscription related details.
    private Map<String, ApplicationKeyMapping> applicationKeyMappingMap;
    private Map<String, ApplicationMapping> applicationMappingMap;
    private Map<String, Application> applicationMap;
    private Map<String, API> apiMap;
    private Map<String, ApiPolicy> apiPolicyMap;
    private Map<String, SubscriptionPolicy> subscriptionPolicyMap;
    private Map<String, ApplicationPolicy> appPolicyMap;
    private Map<String, Subscription> subscriptionMap;

    private Map<String, Map<String, JWTValidator>> jwtValidatorMap;

    SubscriptionDataStoreImpl() {

    }

    public static SubscriptionDataStoreImpl getInstance() {

        return instance;
    }

    public void initializeStore() {

        this.applicationKeyMappingMap = new ConcurrentHashMap<>();
        this.applicationMap = new ConcurrentHashMap<>();
        this.apiMap = new ConcurrentHashMap<>();
        this.subscriptionPolicyMap = new ConcurrentHashMap<>();
        this.appPolicyMap = new ConcurrentHashMap<>();
        this.apiPolicyMap = new ConcurrentHashMap<>();
        this.subscriptionMap = new ConcurrentHashMap<>();
        this.applicationMappingMap = new ConcurrentHashMap<>();
        this.jwtValidatorMap = new ConcurrentHashMap<>();
        initializeLoadingTasks();
    }

    @Override
    public Application getApplicationById(String appUUID) {

        return applicationMap.get(appUUID);
    }

    @Override
    public API getApiByContextAndVersion(String uuid) {

        return apiMap.get(uuid);
    }

    @Override
    public Subscription getSubscriptionById(String appId, String apiId) {

        return subscriptionMap.get(SubscriptionDataStoreUtil.getSubscriptionCacheKey(appId, apiId));
    }

    private void initializeLoadingTasks() {

        SubscriptionDiscoveryClient.getInstance().watchSubscriptions();
        ApplicationDiscoveryClient.getInstance().watchApplications();
        ApiListDiscoveryClient.getInstance().watchApiList();
        ApplicationPolicyDiscoveryClient.getInstance().watchApplicationPolicies();
        SubscriptionPolicyDiscoveryClient.getInstance().watchSubscriptionPolicies();
        ApplicationKeyMappingDiscoveryClient.getInstance().watchApplicationKeyMappings();
        ApplicationMappingDiscoveryClient.getInstance().watchApplicationMappings();
        JWTIssuerDiscoveryClient.getInstance().watchJWTIssuers();
    }

    public void addSubscriptions(List<org.wso2.apk.enforcer.discovery.subscription.Subscription> subscriptionList) {

        Map<String, Subscription> newSubscriptionMap = new ConcurrentHashMap<>();

        for (org.wso2.apk.enforcer.discovery.subscription.Subscription subscription : subscriptionList) {
            SubscribedAPI subscribedAPI = new SubscribedAPI();
            subscribedAPI.setName(subscription.getSubscribedApi().getName());
            subscribedAPI.setVersions(subscription.getSubscribedApi().getVersionsList());

            Subscription newSubscription = new Subscription();
            newSubscription.setSubscriptionId(subscription.getUuid());
            newSubscription.setSubscriptionStatus(subscription.getSubStatus());
            newSubscription.setOrganization(subscription.getOrganization());
            newSubscription.setSubscribedApi(subscribedAPI);
//            newSubscription.setTimeStamp(Long.parseLong(subscription.getTimeStamp()));

            newSubscriptionMap.put(newSubscription.getCacheKey(), newSubscription);
        }

        if (log.isDebugEnabled()) {
            log.debug("Total Subscriptions in new cache: {}", newSubscriptionMap.size());
        }
        this.subscriptionMap = newSubscriptionMap;
    }

    public void addApplications(List<org.wso2.apk.enforcer.discovery.subscription.Application> applicationList) {

        Map<String, Application> newApplicationMap = new ConcurrentHashMap<>();

        for (org.wso2.apk.enforcer.discovery.subscription.Application application : applicationList) {
            Application newApplication = new Application();
            newApplication.setUUID(application.getUuid());
            newApplication.setName(application.getName());
            application.getAttributesMap().forEach(newApplication::addAttribute);

            newApplicationMap.put(newApplication.getCacheKey(), newApplication);
        }
        if (log.isDebugEnabled()) {
            log.debug("Total Applications in new cache: {}", newApplicationMap.size());
        }
        this.applicationMap = newApplicationMap;
    }

    public void addApis(List<APIs> apisList) {

        Map<String, API> newApiMap = new ConcurrentHashMap<>();

        for (APIs api : apisList) {
            API newApi = new API();
            //newApi.setApiId(Integer.parseInt(api.getApiId()));
            newApi.setApiName(api.getName());
            newApi.setApiProvider(api.getProvider());
            newApi.setApiType(api.getApiType());
            newApi.setApiVersion(api.getVersion());
            newApi.setContext(api.getBasePath());
            newApi.setApiTier(api.getPolicy());
            newApi.setApiUUID(api.getUuid());
            newApi.setLcState(api.getLcState());
            newApiMap.put(newApi.getCacheKey(), newApi);
        }
        if (log.isDebugEnabled()) {
            log.debug("Total Apis in new cache: {}", newApiMap.size());
        }
        this.apiMap = newApiMap;
    }

    public void addApplicationPolicies(
            List<org.wso2.apk.enforcer.discovery.subscription.ApplicationPolicy> applicationPolicyList) {

        Map<String, ApplicationPolicy> newAppPolicyMap = new ConcurrentHashMap<>();

        for (org.wso2.apk.enforcer.discovery.subscription.ApplicationPolicy applicationPolicy :
                applicationPolicyList) {
            ApplicationPolicy newApplicationPolicy = new ApplicationPolicy();
            newApplicationPolicy.setId(applicationPolicy.getId());
            newApplicationPolicy.setQuotaType(applicationPolicy.getQuotaType());
            newApplicationPolicy.setTenantId(applicationPolicy.getTenantId());
            newApplicationPolicy.setTierName(applicationPolicy.getName());

            newAppPolicyMap.put(newApplicationPolicy.getCacheKey(), newApplicationPolicy);
        }
        if (log.isDebugEnabled()) {
            log.debug("Total Application Policies in new cache: {}", newAppPolicyMap.size());
        }
        this.appPolicyMap = newAppPolicyMap;
    }

    public void addSubscriptionPolicies(
            List<org.wso2.apk.enforcer.discovery.subscription.SubscriptionPolicy> subscriptionPolicyList) {

        Map<String, SubscriptionPolicy> newSubscriptionPolicyMap = new ConcurrentHashMap<>();

        for (org.wso2.apk.enforcer.discovery.subscription.SubscriptionPolicy subscriptionPolicy :
                subscriptionPolicyList) {
            SubscriptionPolicy newSubscriptionPolicy = new SubscriptionPolicy();
            newSubscriptionPolicy.setId(subscriptionPolicy.getId());
            newSubscriptionPolicy.setQuotaType(subscriptionPolicy.getQuotaType());
            newSubscriptionPolicy.setRateLimitCount(subscriptionPolicy.getRateLimitCount());
            newSubscriptionPolicy.setRateLimitTimeUnit(subscriptionPolicy.getRateLimitTimeUnit());
            newSubscriptionPolicy.setStopOnQuotaReach(subscriptionPolicy.getStopOnQuotaReach());
            newSubscriptionPolicy.setTenantId(subscriptionPolicy.getTenantId());
            newSubscriptionPolicy.setTierName(subscriptionPolicy.getName());
            newSubscriptionPolicy.setGraphQLMaxComplexity(subscriptionPolicy.getGraphQLMaxComplexity());
            newSubscriptionPolicy.setGraphQLMaxDepth(subscriptionPolicy.getGraphQLMaxDepth());

            newSubscriptionPolicyMap.put(newSubscriptionPolicy.getCacheKey(), newSubscriptionPolicy);
        }
        if (log.isDebugEnabled()) {
            log.debug("Total Subscription Policies in new cache: {}", newSubscriptionPolicyMap.size());
        }
        this.subscriptionPolicyMap = newSubscriptionPolicyMap;
    }

    public void addApplicationKeyMappings(
            List<org.wso2.apk.enforcer.discovery.subscription.ApplicationKeyMapping> applicationKeyMappingList) {

        Map<String, ApplicationKeyMapping> newApplicationKeyMappingMap = new ConcurrentHashMap<>();

        for (org.wso2.apk.enforcer.discovery.subscription.ApplicationKeyMapping applicationKeyMapping :
                applicationKeyMappingList) {
            ApplicationKeyMapping mapping = new ApplicationKeyMapping();
            mapping.setApplicationUUID(applicationKeyMapping.getApplicationUUID());
            mapping.setSecurityScheme(applicationKeyMapping.getSecurityScheme());
            mapping.setApplicationIdentifier(applicationKeyMapping.getApplicationIdentifier());
            mapping.setKeyType(applicationKeyMapping.getKeyType());
            mapping.setEnvId(applicationKeyMapping.getEnvID());
            newApplicationKeyMappingMap.put(mapping.getCacheKey(), mapping);
        }
        if (log.isDebugEnabled()) {
            log.debug("Total Application Key Mappings in new cache: {}", newApplicationKeyMappingMap.size());
        }
        this.applicationKeyMappingMap = newApplicationKeyMappingMap;
    }

    public void addApplicationMappings(
            List<org.wso2.apk.enforcer.discovery.subscription.ApplicationMapping> applicationMappingList) {

        Map<String, ApplicationMapping> newApplicationMappingMap = new ConcurrentHashMap<>();

        for (org.wso2.apk.enforcer.discovery.subscription.ApplicationMapping applicationMapping :
                applicationMappingList) {
            ApplicationMapping appMapping = new ApplicationMapping();
            appMapping.setUuid(applicationMapping.getUuid());
            appMapping.setApplicationRef(applicationMapping.getApplicationRef());
            appMapping.setSubscriptionRef(applicationMapping.getSubscriptionRef());
            newApplicationMappingMap.put(appMapping.getCacheKey(), appMapping);
        }
        if (log.isDebugEnabled()) {
            log.debug("Total Application Mappings in new cache: {}", newApplicationMappingMap.size());
        }
        this.applicationMappingMap = newApplicationMappingMap;
    }

    @Override
    public List<API> getMatchingAPIs(String name, String context, String version, String uuid) {

        List<API> apiList = new ArrayList<>();
        for (API api : apiMap.values()) {
            boolean isNameMatching = true;
            boolean isContextMatching = true;
            boolean isVersionMatching = true;
            boolean isUUIDMatching = true;
            if (StringUtils.isNotEmpty(name)) {
                isNameMatching = api.getApiName().contains(name);
            }
            if (StringUtils.isNotEmpty(context)) {
                isContextMatching = api.getContext().equals(context);
            }
            if (StringUtils.isNotEmpty(version)) {
                isVersionMatching = api.getApiVersion().equals(version);
            }
            if (StringUtils.isNotEmpty(uuid)) {
                isUUIDMatching = api.getApiUUID().equals(uuid);
            }
            if (isNameMatching && isContextMatching && isVersionMatching && isUUIDMatching) {
                apiList.add(api);
            }
        }
        return apiList;
    }

    @Override
    public API getMatchingAPI(String context, String version) {

        for (API api : apiMap.values()) {
            if (StringUtils.isNotEmpty(context) && StringUtils.isNotEmpty(version)) {
                if (api.getContext().equals(context) && api.getApiVersion().equals(version)) {
                    return api;
                }
            }
        }
        return null;
    }

    @Override
    public List<Application> getMatchingApplications(String name, String uuid) {

        List<Application> applicationList = new ArrayList<>();
        for (Application application : applicationMap.values()) {
            boolean isNameMatching = true;
            boolean isUUIDMatching = true;
            if (StringUtils.isNotEmpty(name)) {
                isNameMatching = application.getName().contains(name);
            }
            if (StringUtils.isNotEmpty(uuid)) {
                isUUIDMatching = application.getUUID().equals(uuid);
            }
            if (isNameMatching && isUUIDMatching) {
                applicationList.add(application);
            }
        }
        return applicationList;
    }

    @Override
    public List<ApplicationKeyMapping> getMatchingKeyMapping(String applicationUUID, String applicationIdentifier) {

        List<ApplicationKeyMapping> applicationKeyMappingList = new ArrayList<>();

        for (ApplicationKeyMapping applicationKeyMapping : applicationKeyMappingMap.values()) {
            boolean isApplicationIdentifierMatching = true;
            boolean isAppUUIDMatching = true;

            if (StringUtils.isNotEmpty(applicationUUID)) {
                isAppUUIDMatching = applicationKeyMapping.getApplicationUUID().equals(applicationUUID);
            }
            if (StringUtils.isNotEmpty(applicationIdentifier)) {
                isApplicationIdentifierMatching = applicationKeyMapping.getApplicationIdentifier()
                        .equals(applicationIdentifier);
            }
            if (isApplicationIdentifierMatching && isAppUUIDMatching) {
                applicationKeyMappingList.add(applicationKeyMapping);
            }
        }
        return applicationKeyMappingList;
    }

    @Override
    public List<Subscription> getMatchingSubscriptions(String uuid, String apiName, String apiVersion, String state) {

        List<Subscription> subscriptionList = new ArrayList<>();

        for (Subscription subscription : subscriptionMap.values()) {
            boolean isUUIDMatching = true;
            boolean isApiNameMatching = true;
            boolean isApiVersionMatching = true;
            boolean isStateMatching = true;
            if (StringUtils.isNotEmpty(uuid)) {
                isUUIDMatching = subscription.getSubscriptionId().equals(uuid);
            }
            if (StringUtils.isNotEmpty(apiName)) {
                isApiNameMatching = subscription.getSubscribedApi().getName().equals(apiName);
            }
            if (StringUtils.isNotEmpty(apiVersion)) {
                // Todo: Regex check
            }
            if (StringUtils.isNotEmpty(state)) {
                isStateMatching = subscription.getSubscriptionStatus().equals(state);
            }
            if (isUUIDMatching && isApiNameMatching && isApiVersionMatching && isStateMatching) {
                subscriptionList.add(subscription);
            }
        }
        return subscriptionList;
    }

//    @Override
//    public List<ApplicationMapping> getMatchingApplicationMapping

    @Override
    public void addJWTIssuers(List<JWTIssuer> jwtIssuers) {

        Map<String, Map<String, JWTValidator>> jwtValidatorMap = new ConcurrentHashMap<>();
        for (JWTIssuer jwtIssuer : jwtIssuers) {
            try {
                ExtendedTokenIssuerDto tokenIssuerDto = new ExtendedTokenIssuerDto(jwtIssuer.getIssuer());
                tokenIssuerDto.setName(jwtIssuer.getName());
                tokenIssuerDto.setConsumerKeyClaim(jwtIssuer.getConsumerKeyClaim());
                tokenIssuerDto.setScopesClaim(jwtIssuer.getScopesClaim());
                Certificate certificate = jwtIssuer.getCertificate();
                if (StringUtils.isNotEmpty(certificate.getJwks().getUrl())) {
                    JWKSConfigurationDTO jwksConfigurationDTO = new JWKSConfigurationDTO();
                    if (StringUtils.isNotEmpty(certificate.getJwks().getTls())) {
                        java.security.cert.Certificate tlsCertificate =
                                TLSUtils.getCertificateFromContent(certificate.getJwks().getTls());
                        jwksConfigurationDTO.setCertificate(tlsCertificate);
                    }
                    jwksConfigurationDTO.setUrl(certificate.getJwks().getUrl());
                    jwksConfigurationDTO.setEnabled(true);
                    tokenIssuerDto.setJwksConfigurationDTO(jwksConfigurationDTO);
                }
                if (StringUtils.isNotEmpty(certificate.getCertificate())) {
                    java.security.cert.Certificate signingCertificate =
                            TLSUtils.getCertificateFromContent(certificate.getCertificate());
                    tokenIssuerDto.setCertificate(signingCertificate);
                }
                Map<String, String> claimMappingMap = jwtIssuer.getClaimMappingMap();
                Map<String, ClaimMappingDto> claimMappingDtos = new HashMap<>();
                claimMappingMap.forEach((remoteClaim, localClaim) -> {
                    claimMappingDtos.put(remoteClaim, new ClaimMappingDto(remoteClaim, localClaim));
                });
                tokenIssuerDto.setClaimMappings(claimMappingDtos);
                JWTValidator jwtValidator = new JWTValidator(tokenIssuerDto);
                Map<String, JWTValidator> orgBasedJWTValidatorMap = new ConcurrentHashMap<>();
                if (jwtValidatorMap.containsKey(jwtIssuer.getOrganization())) {
                    orgBasedJWTValidatorMap = jwtValidatorMap.get(jwtIssuer.getOrganization());
                }
                orgBasedJWTValidatorMap.put(jwtIssuer.getIssuer(), jwtValidator);
                jwtValidatorMap.put(jwtIssuer.getOrganization(), orgBasedJWTValidatorMap);
                this.jwtValidatorMap = jwtValidatorMap;
            } catch (EnforcerException | CertificateException | IOException e) {
                log.error("Error occurred while configuring JWT Validator for issuer " + jwtIssuer.getIssuer(), e);
            }
        }
    }

    @Override
    public JWTValidator getJWTValidatorByIssuer(String issuer, String organization) {

        Map<String, JWTValidator> orgBaseJWTValidators = jwtValidatorMap.get(organization);
        if (orgBaseJWTValidators != null) {
            return orgBaseJWTValidators.get(issuer);
        }
        return null;
    }
}
