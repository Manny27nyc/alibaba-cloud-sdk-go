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

// DescibeImportsFromDatabase invokes the rds.DescibeImportsFromDatabase API synchronously
func (client *Client) DescibeImportsFromDatabase(request *DescibeImportsFromDatabaseRequest) (response *DescibeImportsFromDatabaseResponse, err error) {
	response = CreateDescibeImportsFromDatabaseResponse()
	err = client.DoAction(request, response)
	return
}

// DescibeImportsFromDatabaseWithChan invokes the rds.DescibeImportsFromDatabase API asynchronously
func (client *Client) DescibeImportsFromDatabaseWithChan(request *DescibeImportsFromDatabaseRequest) (<-chan *DescibeImportsFromDatabaseResponse, <-chan error) {
	responseChan := make(chan *DescibeImportsFromDatabaseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescibeImportsFromDatabase(request)
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

// DescibeImportsFromDatabaseWithCallback invokes the rds.DescibeImportsFromDatabase API asynchronously
func (client *Client) DescibeImportsFromDatabaseWithCallback(request *DescibeImportsFromDatabaseRequest, callback func(response *DescibeImportsFromDatabaseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescibeImportsFromDatabaseResponse
		var err error
		defer close(result)
		response, err = client.DescibeImportsFromDatabase(request)
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

// DescibeImportsFromDatabaseRequest is the request struct for api DescibeImportsFromDatabase
type DescibeImportsFromDatabaseRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	StartTime            string           `position:"Query" name:"StartTime"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	ImportId             requests.Integer `position:"Query" name:"ImportId"`
	Engine               string           `position:"Query" name:"Engine"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescibeImportsFromDatabaseResponse is the response struct for api DescibeImportsFromDatabase
type DescibeImportsFromDatabaseResponse struct {
	*responses.BaseResponse
	RequestId        string                            `json:"RequestId" xml:"RequestId"`
	PageNumber       int                               `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  int                               `json:"PageRecordCount" xml:"PageRecordCount"`
	TotalRecordCount int                               `json:"TotalRecordCount" xml:"TotalRecordCount"`
	Items            ItemsInDescibeImportsFromDatabase `json:"Items" xml:"Items"`
}

// CreateDescibeImportsFromDatabaseRequest creates a request to invoke DescibeImportsFromDatabase API
func CreateDescibeImportsFromDatabaseRequest() (request *DescibeImportsFromDatabaseRequest) {
	request = &DescibeImportsFromDatabaseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescibeImportsFromDatabase", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescibeImportsFromDatabaseResponse creates a response to parse from DescibeImportsFromDatabase response
func CreateDescibeImportsFromDatabaseResponse() (response *DescibeImportsFromDatabaseResponse) {
	response = &DescibeImportsFromDatabaseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
