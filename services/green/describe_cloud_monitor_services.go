package green

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

// DescribeCloudMonitorServices invokes the green.DescribeCloudMonitorServices API synchronously
func (client *Client) DescribeCloudMonitorServices(request *DescribeCloudMonitorServicesRequest) (response *DescribeCloudMonitorServicesResponse, err error) {
	response = CreateDescribeCloudMonitorServicesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCloudMonitorServicesWithChan invokes the green.DescribeCloudMonitorServices API asynchronously
func (client *Client) DescribeCloudMonitorServicesWithChan(request *DescribeCloudMonitorServicesRequest) (<-chan *DescribeCloudMonitorServicesResponse, <-chan error) {
	responseChan := make(chan *DescribeCloudMonitorServicesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCloudMonitorServices(request)
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

// DescribeCloudMonitorServicesWithCallback invokes the green.DescribeCloudMonitorServices API asynchronously
func (client *Client) DescribeCloudMonitorServicesWithCallback(request *DescribeCloudMonitorServicesRequest, callback func(response *DescribeCloudMonitorServicesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCloudMonitorServicesResponse
		var err error
		defer close(result)
		response, err = client.DescribeCloudMonitorServices(request)
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

// DescribeCloudMonitorServicesRequest is the request struct for api DescribeCloudMonitorServices
type DescribeCloudMonitorServicesRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
	PageSize string `position:"Query" name:"PageSize"`
	Page     string `position:"Query" name:"Page"`
}

// DescribeCloudMonitorServicesResponse is the response struct for api DescribeCloudMonitorServices
type DescribeCloudMonitorServicesResponse struct {
	*responses.BaseResponse
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	TotalCount int      `json:"TotalCount" xml:"TotalCount"`
	Items      []string `json:"Items" xml:"Items"`
}

// CreateDescribeCloudMonitorServicesRequest creates a request to invoke DescribeCloudMonitorServices API
func CreateDescribeCloudMonitorServicesRequest() (request *DescribeCloudMonitorServicesRequest) {
	request = &DescribeCloudMonitorServicesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "DescribeCloudMonitorServices", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCloudMonitorServicesResponse creates a response to parse from DescribeCloudMonitorServices response
func CreateDescribeCloudMonitorServicesResponse() (response *DescribeCloudMonitorServicesResponse) {
	response = &DescribeCloudMonitorServicesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}