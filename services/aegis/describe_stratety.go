package aegis

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

func (client *Client) DescribeStratety(request *DescribeStratetyRequest) (response *DescribeStratetyResponse, err error) {
	response = CreateDescribeStratetyResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeStratetyWithChan(request *DescribeStratetyRequest) (<-chan *DescribeStratetyResponse, <-chan error) {
	responseChan := make(chan *DescribeStratetyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeStratety(request)
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

func (client *Client) DescribeStratetyWithCallback(request *DescribeStratetyRequest, callback func(response *DescribeStratetyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeStratetyResponse
		var err error
		defer close(result)
		response, err = client.DescribeStratety(request)
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

type DescribeStratetyRequest struct {
	*requests.RpcRequest
	SourceIp        string           `position:"Query" name:"SourceIp"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
}

type DescribeStratetyResponse struct {
	*responses.BaseResponse
	RequestId  string                   `json:"RequestId" xml:"RequestId"`
	Count      int                      `json:"Count" xml:"Count"`
	Strategies []DataInDescribeStratety `json:"Strategies" xml:"Strategies"`
}

func CreateDescribeStratetyRequest() (request *DescribeStratetyRequest) {
	request = &DescribeStratetyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeStratety", "vipaegis", "openAPI")
	return
}

func CreateDescribeStratetyResponse() (response *DescribeStratetyResponse) {
	response = &DescribeStratetyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
