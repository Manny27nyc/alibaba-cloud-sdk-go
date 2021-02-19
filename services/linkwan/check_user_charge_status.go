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

// CheckUserChargeStatus invokes the linkwan.CheckUserChargeStatus API synchronously
func (client *Client) CheckUserChargeStatus(request *CheckUserChargeStatusRequest) (response *CheckUserChargeStatusResponse, err error) {
	response = CreateCheckUserChargeStatusResponse()
	err = client.DoAction(request, response)
	return
}

// CheckUserChargeStatusWithChan invokes the linkwan.CheckUserChargeStatus API asynchronously
func (client *Client) CheckUserChargeStatusWithChan(request *CheckUserChargeStatusRequest) (<-chan *CheckUserChargeStatusResponse, <-chan error) {
	responseChan := make(chan *CheckUserChargeStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckUserChargeStatus(request)
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

// CheckUserChargeStatusWithCallback invokes the linkwan.CheckUserChargeStatus API asynchronously
func (client *Client) CheckUserChargeStatusWithCallback(request *CheckUserChargeStatusRequest, callback func(response *CheckUserChargeStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckUserChargeStatusResponse
		var err error
		defer close(result)
		response, err = client.CheckUserChargeStatus(request)
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

// CheckUserChargeStatusRequest is the request struct for api CheckUserChargeStatus
type CheckUserChargeStatusRequest struct {
	*requests.RpcRequest
	ApiProduct  string `position:"Body" name:"ApiProduct"`
	ApiRevision string `position:"Body" name:"ApiRevision"`
}

// CheckUserChargeStatusResponse is the response struct for api CheckUserChargeStatus
type CheckUserChargeStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateCheckUserChargeStatusRequest creates a request to invoke CheckUserChargeStatus API
func CreateCheckUserChargeStatusRequest() (request *CheckUserChargeStatusRequest) {
	request = &CheckUserChargeStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "CheckUserChargeStatus", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCheckUserChargeStatusResponse creates a response to parse from CheckUserChargeStatus response
func CreateCheckUserChargeStatusResponse() (response *CheckUserChargeStatusResponse) {
	response = &CheckUserChargeStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
