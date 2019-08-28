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

// DeleteHostPool invokes the emr.DeleteHostPool API synchronously
// api document: https://help.aliyun.com/api/emr/deletehostpool.html
func (client *Client) DeleteHostPool(request *DeleteHostPoolRequest) (response *DeleteHostPoolResponse, err error) {
	response = CreateDeleteHostPoolResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteHostPoolWithChan invokes the emr.DeleteHostPool API asynchronously
// api document: https://help.aliyun.com/api/emr/deletehostpool.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteHostPoolWithChan(request *DeleteHostPoolRequest) (<-chan *DeleteHostPoolResponse, <-chan error) {
	responseChan := make(chan *DeleteHostPoolResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteHostPool(request)
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

// DeleteHostPoolWithCallback invokes the emr.DeleteHostPool API asynchronously
// api document: https://help.aliyun.com/api/emr/deletehostpool.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteHostPoolWithCallback(request *DeleteHostPoolRequest, callback func(response *DeleteHostPoolResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteHostPoolResponse
		var err error
		defer close(result)
		response, err = client.DeleteHostPool(request)
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

// DeleteHostPoolRequest is the request struct for api DeleteHostPool
type DeleteHostPoolRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	BizId           string           `position:"Query" name:"BizId"`
}

// DeleteHostPoolResponse is the response struct for api DeleteHostPool
type DeleteHostPoolResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteHostPoolRequest creates a request to invoke DeleteHostPool API
func CreateDeleteHostPoolRequest() (request *DeleteHostPoolRequest) {
	request = &DeleteHostPoolRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DeleteHostPool", "emr", "openAPI")
	return
}

// CreateDeleteHostPoolResponse creates a response to parse from DeleteHostPool response
func CreateDeleteHostPoolResponse() (response *DeleteHostPoolResponse) {
	response = &DeleteHostPoolResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
