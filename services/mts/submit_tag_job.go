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

func (client *Client) SubmitTagJob(request *SubmitTagJobRequest) (response *SubmitTagJobResponse, err error) {
	response = CreateSubmitTagJobResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SubmitTagJobWithChan(request *SubmitTagJobRequest) (<-chan *SubmitTagJobResponse, <-chan error) {
	responseChan := make(chan *SubmitTagJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitTagJob(request)
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

func (client *Client) SubmitTagJobWithCallback(request *SubmitTagJobRequest, callback func(response *SubmitTagJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitTagJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitTagJob(request)
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

type SubmitTagJobRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Input                string           `position:"Query" name:"Input"`
	TagConfig            string           `position:"Query" name:"TagConfig"`
	UserData             string           `position:"Query" name:"UserData"`
	PipelineId           string           `position:"Query" name:"PipelineId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type SubmitTagJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

func CreateSubmitTagJobRequest() (request *SubmitTagJobRequest) {
	request = &SubmitTagJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "SubmitTagJob", "mts", "openAPI")
	return
}

func CreateSubmitTagJobResponse() (response *SubmitTagJobResponse) {
	response = &SubmitTagJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
