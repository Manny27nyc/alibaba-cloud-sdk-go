package rds

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

func (client *Client) ImportDataForSQLServer(request *ImportDataForSQLServerRequest) (response *ImportDataForSQLServerResponse, err error) {
	response = CreateImportDataForSQLServerResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ImportDataForSQLServerWithChan(request *ImportDataForSQLServerRequest) (<-chan *ImportDataForSQLServerResponse, <-chan error) {
	responseChan := make(chan *ImportDataForSQLServerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ImportDataForSQLServer(request)
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

func (client *Client) ImportDataForSQLServerWithCallback(request *ImportDataForSQLServerRequest, callback func(response *ImportDataForSQLServerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ImportDataForSQLServerResponse
		var err error
		defer close(result)
		response, err = client.ImportDataForSQLServer(request)
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

type ImportDataForSQLServerRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	FileName             string           `position:"Query" name:"FileName"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type ImportDataForSQLServerResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ImportID  int    `json:"ImportID" xml:"ImportID"`
}

func CreateImportDataForSQLServerRequest() (request *ImportDataForSQLServerRequest) {
	request = &ImportDataForSQLServerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ImportDataForSQLServer", "rds", "openAPI")
	return
}

func CreateImportDataForSQLServerResponse() (response *ImportDataForSQLServerResponse) {
	response = &ImportDataForSQLServerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
