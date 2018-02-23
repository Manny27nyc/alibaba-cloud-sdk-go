package ecs

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

func (client *Client) AddTags(request *AddTagsRequest) (response *AddTagsResponse, err error) {
	response = CreateAddTagsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) AddTagsWithChan(request *AddTagsRequest) (<-chan *AddTagsResponse, <-chan error) {
	responseChan := make(chan *AddTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddTags(request)
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

func (client *Client) AddTagsWithCallback(request *AddTagsRequest, callback func(response *AddTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddTagsResponse
		var err error
		defer close(result)
		response, err = client.AddTags(request)
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

type AddTagsRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceType         string           `position:"Query" name:"ResourceType"`
	ResourceId           string           `position:"Query" name:"ResourceId"`
	Tag1Key              string           `position:"Query" name:"Tag.1.Key"`
	Tag2Key              string           `position:"Query" name:"Tag.2.Key"`
	Tag3Key              string           `position:"Query" name:"Tag.3.Key"`
	Tag4Key              string           `position:"Query" name:"Tag.4.Key"`
	Tag5Key              string           `position:"Query" name:"Tag.5.Key"`
	Tag1Value            string           `position:"Query" name:"Tag.1.Value"`
	Tag2Value            string           `position:"Query" name:"Tag.2.Value"`
	Tag3Value            string           `position:"Query" name:"Tag.3.Value"`
	Tag4Value            string           `position:"Query" name:"Tag.4.Value"`
	Tag5Value            string           `position:"Query" name:"Tag.5.Value"`
}

type AddTagsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateAddTagsRequest() (request *AddTagsRequest) {
	request = &AddTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "AddTags", "ecs", "openAPI")
	return
}

func CreateAddTagsResponse() (response *AddTagsResponse) {
	response = &AddTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
