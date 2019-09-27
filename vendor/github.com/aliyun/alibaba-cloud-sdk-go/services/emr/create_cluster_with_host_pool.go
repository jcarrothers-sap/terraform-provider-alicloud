package emr

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateClusterWithHostPool invokes the emr.CreateClusterWithHostPool API synchronously
// api document: https://help.aliyun.com/api/emr/createclusterwithhostpool.html
func (client *Client) CreateClusterWithHostPool(request *CreateClusterWithHostPoolRequest) (response *CreateClusterWithHostPoolResponse, err error) {
	response = CreateCreateClusterWithHostPoolResponse()
	err = client.DoAction(request, response)
	return
}

// CreateClusterWithHostPoolWithChan invokes the emr.CreateClusterWithHostPool API asynchronously
// api document: https://help.aliyun.com/api/emr/createclusterwithhostpool.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateClusterWithHostPoolWithChan(request *CreateClusterWithHostPoolRequest) (<-chan *CreateClusterWithHostPoolResponse, <-chan error) {
	responseChan := make(chan *CreateClusterWithHostPoolResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateClusterWithHostPool(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CreateClusterWithHostPoolWithCallback invokes the emr.CreateClusterWithHostPool API asynchronously
// api document: https://help.aliyun.com/api/emr/createclusterwithhostpool.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateClusterWithHostPoolWithCallback(request *CreateClusterWithHostPoolRequest, callback func(response *CreateClusterWithHostPoolResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateClusterWithHostPoolResponse
		var err error
		defer close(result)
		response, err = client.CreateClusterWithHostPool(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// CreateClusterWithHostPoolRequest is the request struct for api CreateClusterWithHostPool
type CreateClusterWithHostPoolRequest struct {
	*requests.RpcRequest
	ResourceOwnerId  requests.Integer                        `position:"Query" name:"ResourceOwnerId"`
	ClusterName      string                                  `position:"Query" name:"ClusterName"`
	EasEnable        requests.Boolean                        `position:"Query" name:"EasEnable"`
	HostInfo         *[]CreateClusterWithHostPoolHostInfo    `position:"Query" name:"HostInfo"  type:"Repeated"`
	RelatedClusterId string                                  `position:"Query" name:"RelatedClusterId"`
	ClusterType      string                                  `position:"Query" name:"ClusterType"`
	HostGroup        *[]CreateClusterWithHostPoolHostGroup   `position:"Query" name:"HostGroup"  type:"Repeated"`
	StackName        string                                  `position:"Query" name:"StackName"`
	StackVersion     string                                  `position:"Query" name:"StackVersion"`
	ServiceInfo      *[]CreateClusterWithHostPoolServiceInfo `position:"Query" name:"ServiceInfo"  type:"Repeated"`
	Config           *[]CreateClusterWithHostPoolConfig      `position:"Query" name:"Config"  type:"Repeated"`
}

// CreateClusterWithHostPoolHostInfo is a repeated param struct in CreateClusterWithHostPoolRequest
type CreateClusterWithHostPoolHostInfo struct {
	HpHostBizId          string                                           `name:"HpHostBizId"`
	HostName             string                                           `name:"HostName"`
	Role                 string                                           `name:"Role"`
	GroupId              string                                           `name:"GroupId"`
	PrivateIp            string                                           `name:"PrivateIp"`
	ServiceComponentInfo *[]CreateClusterWithHostPoolServiceComponentInfo `name:"ServiceComponentInfo" type:"Repeated"`
	HostGroupName        string                                           `name:"HostGroupName"`
}

// CreateClusterWithHostPoolHostGroup is a repeated param struct in CreateClusterWithHostPoolRequest
type CreateClusterWithHostPoolHostGroup struct {
	GroupType string `name:"GroupType"`
	GroupId   string `name:"GroupId"`
	GroupName string `name:"GroupName"`
}

// CreateClusterWithHostPoolServiceInfo is a repeated param struct in CreateClusterWithHostPoolRequest
type CreateClusterWithHostPoolServiceInfo struct {
	ServiceEcmVersion string `name:"ServiceEcmVersion"`
	ServiceVersion    string `name:"ServiceVersion"`
	ServiceName       string `name:"ServiceName"`
}

// CreateClusterWithHostPoolConfig is a repeated param struct in CreateClusterWithHostPoolRequest
type CreateClusterWithHostPoolConfig struct {
	ConfigKey   string `name:"ConfigKey"`
	FileName    string `name:"FileName"`
	ConfigValue string `name:"ConfigValue"`
	ServiceName string `name:"ServiceName"`
}

// CreateClusterWithHostPoolServiceComponentInfo is a repeated param struct in CreateClusterWithHostPoolRequest
type CreateClusterWithHostPoolServiceComponentInfo struct {
	ServiceEcmVersion string `name:"ServiceEcmVersion"`
	ComponentName     string `name:"ComponentName"`
	ServiceName       string `name:"ServiceName"`
}

// CreateClusterWithHostPoolResponse is the response struct for api CreateClusterWithHostPool
type CreateClusterWithHostPoolResponse struct {
	*responses.BaseResponse
	RequestId          string `json:"RequestId" xml:"RequestId"`
	ClusterId          string `json:"ClusterId" xml:"ClusterId"`
	WorkFlowInstanceId string `json:"WorkFlowInstanceId" xml:"WorkFlowInstanceId"`
	OperationId        string `json:"OperationId" xml:"OperationId"`
}

// CreateCreateClusterWithHostPoolRequest creates a request to invoke CreateClusterWithHostPool API
func CreateCreateClusterWithHostPoolRequest() (request *CreateClusterWithHostPoolRequest) {
	request = &CreateClusterWithHostPoolRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "CreateClusterWithHostPool", "emr", "openAPI")
	return
}

// CreateCreateClusterWithHostPoolResponse creates a response to parse from CreateClusterWithHostPool response
func CreateCreateClusterWithHostPoolResponse() (response *CreateClusterWithHostPoolResponse) {
	response = &CreateClusterWithHostPoolResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}