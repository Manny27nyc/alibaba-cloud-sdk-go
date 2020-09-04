package cloudcallcenter

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

// FallbackContactFlowVersion invokes the cloudcallcenter.FallbackContactFlowVersion API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/fallbackcontactflowversion.html
func (client *Client) FallbackContactFlowVersion(request *FallbackContactFlowVersionRequest) (response *FallbackContactFlowVersionResponse, err error) {
	response = CreateFallbackContactFlowVersionResponse()
	err = client.DoAction(request, response)
	return
}

// FallbackContactFlowVersionWithChan invokes the cloudcallcenter.FallbackContactFlowVersion API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/fallbackcontactflowversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FallbackContactFlowVersionWithChan(request *FallbackContactFlowVersionRequest) (<-chan *FallbackContactFlowVersionResponse, <-chan error) {
	responseChan := make(chan *FallbackContactFlowVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FallbackContactFlowVersion(request)
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

// FallbackContactFlowVersionWithCallback invokes the cloudcallcenter.FallbackContactFlowVersion API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/fallbackcontactflowversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FallbackContactFlowVersionWithCallback(request *FallbackContactFlowVersionRequest, callback func(response *FallbackContactFlowVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FallbackContactFlowVersionResponse
		var err error
		defer close(result)
		response, err = client.FallbackContactFlowVersion(request)
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

// FallbackContactFlowVersionRequest is the request struct for api FallbackContactFlowVersion
type FallbackContactFlowVersionRequest struct {
	*requests.RpcRequest
	InstanceId           string `position:"Query" name:"InstanceId"`
	ContactFlowVersionId string `position:"Query" name:"ContactFlowVersionId"`
}

// FallbackContactFlowVersionResponse is the response struct for api FallbackContactFlowVersion
type FallbackContactFlowVersionResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateFallbackContactFlowVersionRequest creates a request to invoke FallbackContactFlowVersion API
func CreateFallbackContactFlowVersionRequest() (request *FallbackContactFlowVersionRequest) {
	request = &FallbackContactFlowVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "FallbackContactFlowVersion", "", "")
	request.Method = requests.POST
	return
}

// CreateFallbackContactFlowVersionResponse creates a response to parse from FallbackContactFlowVersion response
func CreateFallbackContactFlowVersionResponse() (response *FallbackContactFlowVersionResponse) {
	response = &FallbackContactFlowVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}