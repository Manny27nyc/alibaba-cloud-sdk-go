package mts

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

func (client *Client) QueryAsrJobList(request *QueryAsrJobListRequest) (response *QueryAsrJobListResponse, err error) {
	response = CreateQueryAsrJobListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryAsrJobListWithChan(request *QueryAsrJobListRequest) (<-chan *QueryAsrJobListResponse, <-chan error) {
	responseChan := make(chan *QueryAsrJobListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryAsrJobList(request)
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

func (client *Client) QueryAsrJobListWithCallback(request *QueryAsrJobListRequest, callback func(response *QueryAsrJobListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryAsrJobListResponse
		var err error
		defer close(result)
		response, err = client.QueryAsrJobList(request)
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

type QueryAsrJobListRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JobIds               string           `position:"Query" name:"JobIds"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type QueryAsrJobListResponse struct {
	*responses.BaseResponse
	RequestId   string                       `json:"RequestId" xml:"RequestId"`
	NonExistIds NonExistIdsInQueryAsrJobList `json:"NonExistIds" xml:"NonExistIds"`
	JobList     JobListInQueryAsrJobList     `json:"JobList" xml:"JobList"`
}

func CreateQueryAsrJobListRequest() (request *QueryAsrJobListRequest) {
	request = &QueryAsrJobListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryAsrJobList", "mts", "openAPI")
	return
}

func CreateQueryAsrJobListResponse() (response *QueryAsrJobListResponse) {
	response = &QueryAsrJobListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
