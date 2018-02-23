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

func (client *Client) QueryCustomMetricList(request *QueryCustomMetricListRequest) (response *QueryCustomMetricListResponse, err error) {
	response = CreateQueryCustomMetricListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryCustomMetricListWithChan(request *QueryCustomMetricListRequest) (<-chan *QueryCustomMetricListResponse, <-chan error) {
	responseChan := make(chan *QueryCustomMetricListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryCustomMetricList(request)
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

func (client *Client) QueryCustomMetricListWithCallback(request *QueryCustomMetricListRequest, callback func(response *QueryCustomMetricListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryCustomMetricListResponse
		var err error
		defer close(result)
		response, err = client.QueryCustomMetricList(request)
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

type QueryCustomMetricListRequest struct {
	*requests.RpcRequest
	GroupId    string `position:"Query" name:"GroupId"`
	MetricName string `position:"Query" name:"MetricName"`
	Dimension  string `position:"Query" name:"Dimension"`
	Md5        string `position:"Query" name:"Md5"`
	UUID       string `position:"Query" name:"UUID"`
	Page       string `position:"Query" name:"Page"`
	Size       string `position:"Query" name:"Size"`
}

type QueryCustomMetricListResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    string `json:"Result" xml:"Result"`
}

func CreateQueryCustomMetricListRequest() (request *QueryCustomMetricListRequest) {
	request = &QueryCustomMetricListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2017-03-01", "QueryCustomMetricList", "cms", "openAPI")
	return
}

func CreateQueryCustomMetricListResponse() (response *QueryCustomMetricListResponse) {
	response = &QueryCustomMetricListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
