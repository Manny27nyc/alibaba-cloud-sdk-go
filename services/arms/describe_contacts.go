package arms

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

// DescribeContacts invokes the arms.DescribeContacts API synchronously
func (client *Client) DescribeContacts(request *DescribeContactsRequest) (response *DescribeContactsResponse, err error) {
	response = CreateDescribeContactsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeContactsWithChan invokes the arms.DescribeContacts API asynchronously
func (client *Client) DescribeContactsWithChan(request *DescribeContactsRequest) (<-chan *DescribeContactsResponse, <-chan error) {
	responseChan := make(chan *DescribeContactsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeContacts(request)
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

// DescribeContactsWithCallback invokes the arms.DescribeContacts API asynchronously
func (client *Client) DescribeContactsWithCallback(request *DescribeContactsRequest, callback func(response *DescribeContactsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeContactsResponse
		var err error
		defer close(result)
		response, err = client.DescribeContacts(request)
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

// DescribeContactsRequest is the request struct for api DescribeContacts
type DescribeContactsRequest struct {
	*requests.RpcRequest
	ContactName string           `position:"Query" name:"ContactName"`
	Size        requests.Integer `position:"Query" name:"Size"`
	Phone       string           `position:"Query" name:"Phone"`
	Page        requests.Integer `position:"Query" name:"Page"`
	Email       string           `position:"Query" name:"Email"`
}

// DescribeContactsResponse is the response struct for api DescribeContacts
type DescribeContactsResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	PageBean  PageBean `json:"PageBean" xml:"PageBean"`
}

// CreateDescribeContactsRequest creates a request to invoke DescribeContacts API
func CreateDescribeContactsRequest() (request *DescribeContactsRequest) {
	request = &DescribeContactsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "DescribeContacts", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeContactsResponse creates a response to parse from DescribeContacts response
func CreateDescribeContactsResponse() (response *DescribeContactsResponse) {
	response = &DescribeContactsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
