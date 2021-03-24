package mse

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

// DeleteEngineNamespace invokes the mse.DeleteEngineNamespace API synchronously
func (client *Client) DeleteEngineNamespace(request *DeleteEngineNamespaceRequest) (response *DeleteEngineNamespaceResponse, err error) {
	response = CreateDeleteEngineNamespaceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteEngineNamespaceWithChan invokes the mse.DeleteEngineNamespace API asynchronously
func (client *Client) DeleteEngineNamespaceWithChan(request *DeleteEngineNamespaceRequest) (<-chan *DeleteEngineNamespaceResponse, <-chan error) {
	responseChan := make(chan *DeleteEngineNamespaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteEngineNamespace(request)
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

// DeleteEngineNamespaceWithCallback invokes the mse.DeleteEngineNamespace API asynchronously
func (client *Client) DeleteEngineNamespaceWithCallback(request *DeleteEngineNamespaceRequest, callback func(response *DeleteEngineNamespaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteEngineNamespaceResponse
		var err error
		defer close(result)
		response, err = client.DeleteEngineNamespace(request)
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

// DeleteEngineNamespaceRequest is the request struct for api DeleteEngineNamespace
type DeleteEngineNamespaceRequest struct {
	*requests.RpcRequest
	ClusterId  string `position:"Query" name:"ClusterId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Id         string `position:"Query" name:"Id"`
}

// DeleteEngineNamespaceResponse is the response struct for api DeleteEngineNamespace
type DeleteEngineNamespaceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	HttpCode  string `json:"HttpCode" xml:"HttpCode"`
}

// CreateDeleteEngineNamespaceRequest creates a request to invoke DeleteEngineNamespace API
func CreateDeleteEngineNamespaceRequest() (request *DeleteEngineNamespaceRequest) {
	request = &DeleteEngineNamespaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "DeleteEngineNamespace", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteEngineNamespaceResponse creates a response to parse from DeleteEngineNamespace response
func CreateDeleteEngineNamespaceResponse() (response *DeleteEngineNamespaceResponse) {
	response = &DeleteEngineNamespaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
