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

func (client *Client) DescribeRenewalPrice(request *DescribeRenewalPriceRequest) (response *DescribeRenewalPriceResponse, err error) {
	response = CreateDescribeRenewalPriceResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeRenewalPriceWithChan(request *DescribeRenewalPriceRequest) (<-chan *DescribeRenewalPriceResponse, <-chan error) {
	responseChan := make(chan *DescribeRenewalPriceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRenewalPrice(request)
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

func (client *Client) DescribeRenewalPriceWithCallback(request *DescribeRenewalPriceRequest, callback func(response *DescribeRenewalPriceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRenewalPriceResponse
		var err error
		defer close(result)
		response, err = client.DescribeRenewalPrice(request)
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

type DescribeRenewalPriceRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceType         string           `position:"Query" name:"ResourceType"`
	ResourceId           string           `position:"Query" name:"ResourceId"`
	Period               requests.Integer `position:"Query" name:"Period"`
	PriceUnit            string           `position:"Query" name:"PriceUnit"`
}

type DescribeRenewalPriceResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	PriceInfo PriceInfo `json:"PriceInfo" xml:"PriceInfo"`
}

func CreateDescribeRenewalPriceRequest() (request *DescribeRenewalPriceRequest) {
	request = &DescribeRenewalPriceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeRenewalPrice", "ecs", "openAPI")
	return
}

func CreateDescribeRenewalPriceResponse() (response *DescribeRenewalPriceResponse) {
	response = &DescribeRenewalPriceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
