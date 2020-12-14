package idrsservice

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

// ListTaskGroups invokes the idrsservice.ListTaskGroups API synchronously
func (client *Client) ListTaskGroups(request *ListTaskGroupsRequest) (response *ListTaskGroupsResponse, err error) {
	response = CreateListTaskGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// ListTaskGroupsWithChan invokes the idrsservice.ListTaskGroups API asynchronously
func (client *Client) ListTaskGroupsWithChan(request *ListTaskGroupsRequest) (<-chan *ListTaskGroupsResponse, <-chan error) {
	responseChan := make(chan *ListTaskGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTaskGroups(request)
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

// ListTaskGroupsWithCallback invokes the idrsservice.ListTaskGroups API asynchronously
func (client *Client) ListTaskGroupsWithCallback(request *ListTaskGroupsRequest, callback func(response *ListTaskGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTaskGroupsResponse
		var err error
		defer close(result)
		response, err = client.ListTaskGroups(request)
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

// ListTaskGroupsRequest is the request struct for api ListTaskGroups
type ListTaskGroupsRequest struct {
	*requests.RpcRequest
	PageSize  requests.Integer `position:"Query" name:"PageSize"`
	PageIndex requests.Integer `position:"Query" name:"PageIndex"`
	Status    string           `position:"Query" name:"Status"`
}

// ListTaskGroupsResponse is the response struct for api ListTaskGroups
type ListTaskGroupsResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateListTaskGroupsRequest creates a request to invoke ListTaskGroups API
func CreateListTaskGroupsRequest() (request *ListTaskGroupsRequest) {
	request = &ListTaskGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("idrsservice", "2020-06-30", "ListTaskGroups", "idrsservice", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListTaskGroupsResponse creates a response to parse from ListTaskGroups response
func CreateListTaskGroupsResponse() (response *ListTaskGroupsResponse) {
	response = &ListTaskGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}