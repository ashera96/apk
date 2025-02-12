/*
 *  Copyright (c) 2021, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 *
 *  WSO2 Inc. licenses this file to you under the Apache License,
 *  Version 2.0 (the "License"); you may not use this file except
 *  in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY

 * * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.wso2.apk.enforcer.commons.dto;

import java.util.ArrayList;
import java.util.List;

/**
 * DTO class to hold GraphQL query analyzer response information.
 */
public class QueryAnalyzerResponseDTO {

    boolean isSuccess;
    List<String> errorList = new ArrayList<>();

    public boolean isSuccess() {
        return isSuccess;
    }

    public void setSuccess(boolean success) {
        isSuccess = success;
    }

    public List<String> getErrorList() {
        return errorList;
    }

    public void addErrorToList(String error) {
        this.errorList.add(error);
    }
}
