package aliyuncvc

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

// QueryMeetingInfo invokes the aliyuncvc.QueryMeetingInfo API synchronously
// api document: https://help.aliyun.com/api/aliyuncvc/querymeetinginfo.html
func (client *Client) QueryMeetingInfo(request *QueryMeetingInfoRequest) (response *QueryMeetingInfoResponse, err error) {
	response = CreateQueryMeetingInfoResponse()
	err = client.DoAction(request, response)
	return
}

// QueryMeetingInfoWithChan invokes the aliyuncvc.QueryMeetingInfo API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/querymeetinginfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryMeetingInfoWithChan(request *QueryMeetingInfoRequest) (<-chan *QueryMeetingInfoResponse, <-chan error) {
	responseChan := make(chan *QueryMeetingInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMeetingInfo(request)
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

// QueryMeetingInfoWithCallback invokes the aliyuncvc.QueryMeetingInfo API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/querymeetinginfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryMeetingInfoWithCallback(request *QueryMeetingInfoRequest, callback func(response *QueryMeetingInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMeetingInfoResponse
		var err error
		defer close(result)
		response, err = client.QueryMeetingInfo(request)
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

// QueryMeetingInfoRequest is the request struct for api QueryMeetingInfo
type QueryMeetingInfoRequest struct {
	*requests.RpcRequest
	MeetingUUID string `position:"Body" name:"MeetingUUID"`
}

// QueryMeetingInfoResponse is the response struct for api QueryMeetingInfo
type QueryMeetingInfoResponse struct {
	*responses.BaseResponse
	ErrorCode   int         `json:"ErrorCode" xml:"ErrorCode"`
	Message     string      `json:"Message" xml:"Message"`
	Success     bool        `json:"Success" xml:"Success"`
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	MeetingInfo MeetingInfo `json:"MeetingInfo" xml:"MeetingInfo"`
}

// CreateQueryMeetingInfoRequest creates a request to invoke QueryMeetingInfo API
func CreateQueryMeetingInfoRequest() (request *QueryMeetingInfoRequest) {
	request = &QueryMeetingInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-09-19", "QueryMeetingInfo", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryMeetingInfoResponse creates a response to parse from QueryMeetingInfo response
func CreateQueryMeetingInfoResponse() (response *QueryMeetingInfoResponse) {
	response = &QueryMeetingInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}