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

func (client *Client) CreateCommonBandwidthPackage(request *CreateCommonBandwidthPackageRequest) (response *CreateCommonBandwidthPackageResponse, err error) {
	response = CreateCreateCommonBandwidthPackageResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateCommonBandwidthPackageWithChan(request *CreateCommonBandwidthPackageRequest) (<-chan *CreateCommonBandwidthPackageResponse, <-chan error) {
	responseChan := make(chan *CreateCommonBandwidthPackageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCommonBandwidthPackage(request)
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

func (client *Client) CreateCommonBandwidthPackageWithCallback(request *CreateCommonBandwidthPackageRequest, callback func(response *CreateCommonBandwidthPackageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCommonBandwidthPackageResponse
		var err error
		defer close(result)
		response, err = client.CreateCommonBandwidthPackage(request)
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

type CreateCommonBandwidthPackageRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Name                 string           `position:"Query" name:"Name"`
	Description          string           `position:"Query" name:"Description"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	Bandwidth            requests.Integer `position:"Query" name:"Bandwidth"`
	Ratio                requests.Integer `position:"Query" name:"Ratio"`
	InternetChargeType   string           `position:"Query" name:"InternetChargeType"`
}

type CreateCommonBandwidthPackageResponse struct {
	*responses.BaseResponse
	RequestId          string `json:"RequestId" xml:"RequestId"`
	BandwidthPackageId string `json:"BandwidthPackageId" xml:"BandwidthPackageId"`
}

func CreateCreateCommonBandwidthPackageRequest() (request *CreateCommonBandwidthPackageRequest) {
	request = &CreateCommonBandwidthPackageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateCommonBandwidthPackage", "vpc", "openAPI")
	return
}

func CreateCreateCommonBandwidthPackageResponse() (response *CreateCommonBandwidthPackageResponse) {
	response = &CreateCommonBandwidthPackageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
