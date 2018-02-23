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

func (client *Client) DescribeImportsForSQLServer(request *DescribeImportsForSQLServerRequest) (response *DescribeImportsForSQLServerResponse, err error) {
	response = CreateDescribeImportsForSQLServerResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeImportsForSQLServerWithChan(request *DescribeImportsForSQLServerRequest) (<-chan *DescribeImportsForSQLServerResponse, <-chan error) {
	responseChan := make(chan *DescribeImportsForSQLServerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeImportsForSQLServer(request)
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

func (client *Client) DescribeImportsForSQLServerWithCallback(request *DescribeImportsForSQLServerRequest, callback func(response *DescribeImportsForSQLServerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeImportsForSQLServerResponse
		var err error
		defer close(result)
		response, err = client.DescribeImportsForSQLServer(request)
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

type DescribeImportsForSQLServerRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ImportId             requests.Integer `position:"Query" name:"ImportId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	EndTime              string           `position:"Query" name:"EndTime"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type DescribeImportsForSQLServerResponse struct {
	*responses.BaseResponse
	RequestId         string                             `json:"RequestId" xml:"RequestId"`
	TotalRecordCounts int                                `json:"TotalRecordCounts" xml:"TotalRecordCounts"`
	PageNumber        int                                `json:"PageNumber" xml:"PageNumber"`
	SQLItemsCounts    int                                `json:"SQLItemsCounts" xml:"SQLItemsCounts"`
	Items             ItemsInDescribeImportsForSQLServer `json:"Items" xml:"Items"`
}

func CreateDescribeImportsForSQLServerRequest() (request *DescribeImportsForSQLServerRequest) {
	request = &DescribeImportsForSQLServerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeImportsForSQLServer", "rds", "openAPI")
	return
}

func CreateDescribeImportsForSQLServerResponse() (response *DescribeImportsForSQLServerResponse) {
	response = &DescribeImportsForSQLServerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
