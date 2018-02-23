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

func (client *Client) EnableAlarm(request *EnableAlarmRequest) (response *EnableAlarmResponse, err error) {
	response = CreateEnableAlarmResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) EnableAlarmWithChan(request *EnableAlarmRequest) (<-chan *EnableAlarmResponse, <-chan error) {
	responseChan := make(chan *EnableAlarmResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableAlarm(request)
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

func (client *Client) EnableAlarmWithCallback(request *EnableAlarmRequest, callback func(response *EnableAlarmResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableAlarmResponse
		var err error
		defer close(result)
		response, err = client.EnableAlarm(request)
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

type EnableAlarmRequest struct {
	*requests.RpcRequest
	CallbyCmsOwner string `position:"Query" name:"callby_cms_owner"`
	Id             string `position:"Query" name:"Id"`
}

type EnableAlarmResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateEnableAlarmRequest() (request *EnableAlarmRequest) {
	request = &EnableAlarmRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2017-03-01", "EnableAlarm", "cms", "openAPI")
	return
}

func CreateEnableAlarmResponse() (response *EnableAlarmResponse) {
	response = &EnableAlarmResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
