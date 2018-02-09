package ehpc

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

func (client *Client) EditJobTemplate(request *EditJobTemplateRequest) (response *EditJobTemplateResponse, err error) {
	response = CreateEditJobTemplateResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) EditJobTemplateWithChan(request *EditJobTemplateRequest) (<-chan *EditJobTemplateResponse, <-chan error) {
	responseChan := make(chan *EditJobTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EditJobTemplate(request)
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

func (client *Client) EditJobTemplateWithCallback(request *EditJobTemplateRequest, callback func(response *EditJobTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EditJobTemplateResponse
		var err error
		defer close(result)
		response, err = client.EditJobTemplate(request)
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

type EditJobTemplateRequest struct {
	*requests.RpcRequest
	TemplateId         string           `position:"Query" name:"TemplateId"`
	CommandLine        string           `position:"Query" name:"CommandLine"`
	Name               string           `position:"Query" name:"Name"`
	RunasUser          string           `position:"Query" name:"RunasUser"`
	Priority           requests.Integer `position:"Query" name:"Priority"`
	PackagePath        string           `position:"Query" name:"PackagePath"`
	StdoutRedirectPath string           `position:"Query" name:"StdoutRedirectPath"`
	StderrRedirectPath string           `position:"Query" name:"StderrRedirectPath"`
	ReRunnable         requests.Boolean `position:"Query" name:"ReRunnable"`
	ArrayRequest       string           `position:"Query" name:"ArrayRequest"`
	Variables          string           `position:"Query" name:"Variables"`
}

type EditJobTemplateResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TemplateId string `json:"TemplateId" xml:"TemplateId"`
}

func CreateEditJobTemplateRequest() (request *EditJobTemplateRequest) {
	request = &EditJobTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2017-07-14", "EditJobTemplate", "", "")
	return
}

func CreateEditJobTemplateResponse() (response *EditJobTemplateResponse) {
	response = &EditJobTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
