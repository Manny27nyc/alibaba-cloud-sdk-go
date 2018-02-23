package vpc

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

func (client *Client) DescribeBgpPeers(request *DescribeBgpPeersRequest) (response *DescribeBgpPeersResponse, err error) {
	response = CreateDescribeBgpPeersResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeBgpPeersWithChan(request *DescribeBgpPeersRequest) (<-chan *DescribeBgpPeersResponse, <-chan error) {
	responseChan := make(chan *DescribeBgpPeersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBgpPeers(request)
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

func (client *Client) DescribeBgpPeersWithCallback(request *DescribeBgpPeersRequest, callback func(response *DescribeBgpPeersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBgpPeersResponse
		var err error
		defer close(result)
		response, err = client.DescribeBgpPeers(request)
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

type DescribeBgpPeersRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	RouterId             string           `position:"Query" name:"RouterId"`
	BgpPeerId            string           `position:"Query" name:"BgpPeerId"`
	BgpGroupId           string           `position:"Query" name:"BgpGroupId"`
	IsDefault            requests.Boolean `position:"Query" name:"IsDefault"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type DescribeBgpPeersResponse struct {
	*responses.BaseResponse
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	TotalCount int      `json:"TotalCount" xml:"TotalCount"`
	PageNumber int      `json:"PageNumber" xml:"PageNumber"`
	PageSize   int      `json:"PageSize" xml:"PageSize"`
	BgpPeers   BgpPeers `json:"BgpPeers" xml:"BgpPeers"`
}

func CreateDescribeBgpPeersRequest() (request *DescribeBgpPeersRequest) {
	request = &DescribeBgpPeersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeBgpPeers", "vpc", "openAPI")
	return
}

func CreateDescribeBgpPeersResponse() (response *DescribeBgpPeersResponse) {
	response = &DescribeBgpPeersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
