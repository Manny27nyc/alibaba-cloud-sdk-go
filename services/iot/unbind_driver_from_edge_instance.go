package iot

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

// UnbindDriverFromEdgeInstance invokes the iot.UnbindDriverFromEdgeInstance API synchronously
// api document: https://help.aliyun.com/api/iot/unbinddriverfromedgeinstance.html
func (client *Client) UnbindDriverFromEdgeInstance(request *UnbindDriverFromEdgeInstanceRequest) (response *UnbindDriverFromEdgeInstanceResponse, err error) {
	response = CreateUnbindDriverFromEdgeInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// UnbindDriverFromEdgeInstanceWithChan invokes the iot.UnbindDriverFromEdgeInstance API asynchronously
// api document: https://help.aliyun.com/api/iot/unbinddriverfromedgeinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnbindDriverFromEdgeInstanceWithChan(request *UnbindDriverFromEdgeInstanceRequest) (<-chan *UnbindDriverFromEdgeInstanceResponse, <-chan error) {
	responseChan := make(chan *UnbindDriverFromEdgeInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnbindDriverFromEdgeInstance(request)
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

// UnbindDriverFromEdgeInstanceWithCallback invokes the iot.UnbindDriverFromEdgeInstance API asynchronously
// api document: https://help.aliyun.com/api/iot/unbinddriverfromedgeinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnbindDriverFromEdgeInstanceWithCallback(request *UnbindDriverFromEdgeInstanceRequest, callback func(response *UnbindDriverFromEdgeInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnbindDriverFromEdgeInstanceResponse
		var err error
		defer close(result)
		response, err = client.UnbindDriverFromEdgeInstance(request)
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

// UnbindDriverFromEdgeInstanceRequest is the request struct for api UnbindDriverFromEdgeInstance
type UnbindDriverFromEdgeInstanceRequest struct {
	*requests.RpcRequest
	InstanceId    string `position:"Query" name:"InstanceId"`
	DriverId      string `position:"Query" name:"DriverId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
}

// UnbindDriverFromEdgeInstanceResponse is the response struct for api UnbindDriverFromEdgeInstance
type UnbindDriverFromEdgeInstanceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateUnbindDriverFromEdgeInstanceRequest creates a request to invoke UnbindDriverFromEdgeInstance API
func CreateUnbindDriverFromEdgeInstanceRequest() (request *UnbindDriverFromEdgeInstanceRequest) {
	request = &UnbindDriverFromEdgeInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "UnbindDriverFromEdgeInstance", "Iot", "openAPI")
	return
}

// CreateUnbindDriverFromEdgeInstanceResponse creates a response to parse from UnbindDriverFromEdgeInstance response
func CreateUnbindDriverFromEdgeInstanceResponse() (response *UnbindDriverFromEdgeInstanceResponse) {
	response = &UnbindDriverFromEdgeInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}