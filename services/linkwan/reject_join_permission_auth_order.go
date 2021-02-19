package linkwan

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

// RejectJoinPermissionAuthOrder invokes the linkwan.RejectJoinPermissionAuthOrder API synchronously
func (client *Client) RejectJoinPermissionAuthOrder(request *RejectJoinPermissionAuthOrderRequest) (response *RejectJoinPermissionAuthOrderResponse, err error) {
	response = CreateRejectJoinPermissionAuthOrderResponse()
	err = client.DoAction(request, response)
	return
}

// RejectJoinPermissionAuthOrderWithChan invokes the linkwan.RejectJoinPermissionAuthOrder API asynchronously
func (client *Client) RejectJoinPermissionAuthOrderWithChan(request *RejectJoinPermissionAuthOrderRequest) (<-chan *RejectJoinPermissionAuthOrderResponse, <-chan error) {
	responseChan := make(chan *RejectJoinPermissionAuthOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RejectJoinPermissionAuthOrder(request)
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

// RejectJoinPermissionAuthOrderWithCallback invokes the linkwan.RejectJoinPermissionAuthOrder API asynchronously
func (client *Client) RejectJoinPermissionAuthOrderWithCallback(request *RejectJoinPermissionAuthOrderRequest, callback func(response *RejectJoinPermissionAuthOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RejectJoinPermissionAuthOrderResponse
		var err error
		defer close(result)
		response, err = client.RejectJoinPermissionAuthOrder(request)
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

// RejectJoinPermissionAuthOrderRequest is the request struct for api RejectJoinPermissionAuthOrder
type RejectJoinPermissionAuthOrderRequest struct {
	*requests.RpcRequest
	OrderId     string `position:"Query" name:"OrderId"`
	ApiProduct  string `position:"Body" name:"ApiProduct"`
	ApiRevision string `position:"Body" name:"ApiRevision"`
}

// RejectJoinPermissionAuthOrderResponse is the response struct for api RejectJoinPermissionAuthOrder
type RejectJoinPermissionAuthOrderResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateRejectJoinPermissionAuthOrderRequest creates a request to invoke RejectJoinPermissionAuthOrder API
func CreateRejectJoinPermissionAuthOrderRequest() (request *RejectJoinPermissionAuthOrderRequest) {
	request = &RejectJoinPermissionAuthOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "RejectJoinPermissionAuthOrder", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRejectJoinPermissionAuthOrderResponse creates a response to parse from RejectJoinPermissionAuthOrder response
func CreateRejectJoinPermissionAuthOrderResponse() (response *RejectJoinPermissionAuthOrderResponse) {
	response = &RejectJoinPermissionAuthOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
