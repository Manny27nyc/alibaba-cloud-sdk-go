package nas

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

func (client *Client) CreateFileSystem(request *CreateFileSystemRequest) (response *CreateFileSystemResponse, err error) {
	response = CreateCreateFileSystemResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateFileSystemWithChan(request *CreateFileSystemRequest) (<-chan *CreateFileSystemResponse, <-chan error) {
	responseChan := make(chan *CreateFileSystemResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateFileSystem(request)
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

func (client *Client) CreateFileSystemWithCallback(request *CreateFileSystemRequest, callback func(response *CreateFileSystemResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateFileSystemResponse
		var err error
		defer close(result)
		response, err = client.CreateFileSystem(request)
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

type CreateFileSystemRequest struct {
	*requests.RpcRequest
	StorageType  string `position:"Query" name:"StorageType"`
	ProtocolType string `position:"Query" name:"ProtocolType"`
	Description  string `position:"Query" name:"Description"`
}

type CreateFileSystemResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	FileSystemId string `json:"FileSystemId" xml:"FileSystemId"`
}

func CreateCreateFileSystemRequest() (request *CreateFileSystemRequest) {
	request = &CreateFileSystemRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "CreateFileSystem", "nas", "openAPI")
	return
}

func CreateCreateFileSystemResponse() (response *CreateFileSystemResponse) {
	response = &CreateFileSystemResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
