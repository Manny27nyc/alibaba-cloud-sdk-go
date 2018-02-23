package cms

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

func (client *Client) PutMetricData(request *PutMetricDataRequest) (response *PutMetricDataResponse, err error) {
	response = CreatePutMetricDataResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) PutMetricDataWithChan(request *PutMetricDataRequest) (<-chan *PutMetricDataResponse, <-chan error) {
	responseChan := make(chan *PutMetricDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PutMetricData(request)
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

func (client *Client) PutMetricDataWithCallback(request *PutMetricDataRequest, callback func(response *PutMetricDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PutMetricDataResponse
		var err error
		defer close(result)
		response, err = client.PutMetricData(request)
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

type PutMetricDataRequest struct {
	*requests.RpcRequest
	CallbyCmsOwner string `position:"Query" name:"callby_cms_owner"`
	Body           string `position:"Query" name:"Body"`
}

type PutMetricDataResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

func CreatePutMetricDataRequest() (request *PutMetricDataRequest) {
	request = &PutMetricDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2017-03-01", "PutMetricData", "cms", "openAPI")
	return
}

func CreatePutMetricDataResponse() (response *PutMetricDataResponse) {
	response = &PutMetricDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
