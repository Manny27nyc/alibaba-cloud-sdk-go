package emr

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

// DescribeClusterServiceConfig invokes the emr.DescribeClusterServiceConfig API synchronously
// api document: https://help.aliyun.com/api/emr/describeclusterserviceconfig.html
func (client *Client) DescribeClusterServiceConfig(request *DescribeClusterServiceConfigRequest) (response *DescribeClusterServiceConfigResponse, err error) {
	response = CreateDescribeClusterServiceConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeClusterServiceConfigWithChan invokes the emr.DescribeClusterServiceConfig API asynchronously
// api document: https://help.aliyun.com/api/emr/describeclusterserviceconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeClusterServiceConfigWithChan(request *DescribeClusterServiceConfigRequest) (<-chan *DescribeClusterServiceConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeClusterServiceConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeClusterServiceConfig(request)
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

// DescribeClusterServiceConfigWithCallback invokes the emr.DescribeClusterServiceConfig API asynchronously
// api document: https://help.aliyun.com/api/emr/describeclusterserviceconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeClusterServiceConfigWithCallback(request *DescribeClusterServiceConfigRequest, callback func(response *DescribeClusterServiceConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeClusterServiceConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeClusterServiceConfig(request)
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

// DescribeClusterServiceConfigRequest is the request struct for api DescribeClusterServiceConfig
type DescribeClusterServiceConfigRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	HostInstanceId  string           `position:"Query" name:"HostInstanceId"`
	TagValue        string           `position:"Query" name:"TagValue"`
	GroupId         string           `position:"Query" name:"GroupId"`
	ServiceName     string           `position:"Query" name:"ServiceName"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
	ConfigVersion   string           `position:"Query" name:"ConfigVersion"`
}

// DescribeClusterServiceConfigResponse is the response struct for api DescribeClusterServiceConfig
type DescribeClusterServiceConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Config    Config `json:"Config" xml:"Config"`
}

// CreateDescribeClusterServiceConfigRequest creates a request to invoke DescribeClusterServiceConfig API
func CreateDescribeClusterServiceConfigRequest() (request *DescribeClusterServiceConfigRequest) {
	request = &DescribeClusterServiceConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DescribeClusterServiceConfig", "emr", "openAPI")
	return
}

// CreateDescribeClusterServiceConfigResponse creates a response to parse from DescribeClusterServiceConfig response
func CreateDescribeClusterServiceConfigResponse() (response *DescribeClusterServiceConfigResponse) {
	response = &DescribeClusterServiceConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
