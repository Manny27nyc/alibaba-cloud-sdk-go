package unimkt

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

// QueryMedia invokes the unimkt.QueryMedia API synchronously
func (client *Client) QueryMedia(request *QueryMediaRequest) (response *QueryMediaResponse, err error) {
	response = CreateQueryMediaResponse()
	err = client.DoAction(request, response)
	return
}

// QueryMediaWithChan invokes the unimkt.QueryMedia API asynchronously
func (client *Client) QueryMediaWithChan(request *QueryMediaRequest) (<-chan *QueryMediaResponse, <-chan error) {
	responseChan := make(chan *QueryMediaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMedia(request)
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

// QueryMediaWithCallback invokes the unimkt.QueryMedia API asynchronously
func (client *Client) QueryMediaWithCallback(request *QueryMediaRequest, callback func(response *QueryMediaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMediaResponse
		var err error
		defer close(result)
		response, err = client.QueryMedia(request)
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

// QueryMediaRequest is the request struct for api QueryMedia
type QueryMediaRequest struct {
	*requests.RpcRequest
	Business         string `position:"Query" name:"Business"`
	MediaId          string `position:"Query" name:"MediaId"`
	UserId           string `position:"Query" name:"UserId"`
	OriginSiteUserId string `position:"Query" name:"OriginSiteUserId"`
	Environment      string `position:"Query" name:"Environment"`
	AppName          string `position:"Query" name:"AppName"`
	TenantId         string `position:"Query" name:"TenantId"`
	UserSite         string `position:"Query" name:"UserSite"`
}

// QueryMediaResponse is the response struct for api QueryMedia
type QueryMediaResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Model     Model  `json:"Model" xml:"Model"`
}

// CreateQueryMediaRequest creates a request to invoke QueryMedia API
func CreateQueryMediaRequest() (request *QueryMediaRequest) {
	request = &QueryMediaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("UniMkt", "2018-12-12", "QueryMedia", "uniMkt", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryMediaResponse creates a response to parse from QueryMedia response
func CreateQueryMediaResponse() (response *QueryMediaResponse) {
	response = &QueryMediaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
