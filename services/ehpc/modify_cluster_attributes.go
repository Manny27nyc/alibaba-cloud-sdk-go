package ehpc

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

func (client *Client) ModifyClusterAttributes(request *ModifyClusterAttributesRequest) (response *ModifyClusterAttributesResponse, err error) {
	response = CreateModifyClusterAttributesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyClusterAttributesWithChan(request *ModifyClusterAttributesRequest) (<-chan *ModifyClusterAttributesResponse, <-chan error) {
	responseChan := make(chan *ModifyClusterAttributesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyClusterAttributes(request)
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

func (client *Client) ModifyClusterAttributesWithCallback(request *ModifyClusterAttributesRequest, callback func(response *ModifyClusterAttributesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyClusterAttributesResponse
		var err error
		defer close(result)
		response, err = client.ModifyClusterAttributes(request)
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

type ModifyClusterAttributesRequest struct {
	*requests.RpcRequest
	ClusterId   string `position:"Query" name:"ClusterId"`
	Name        string `position:"Query" name:"Name"`
	Description string `position:"Query" name:"Description"`
}

type ModifyClusterAttributesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateModifyClusterAttributesRequest() (request *ModifyClusterAttributesRequest) {
	request = &ModifyClusterAttributesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2017-07-14", "ModifyClusterAttributes", "ehs", "openAPI")
	return
}

func CreateModifyClusterAttributesResponse() (response *ModifyClusterAttributesResponse) {
	response = &ModifyClusterAttributesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
