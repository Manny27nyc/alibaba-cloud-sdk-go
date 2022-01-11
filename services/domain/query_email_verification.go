package domain

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

// QueryEmailVerification invokes the domain.QueryEmailVerification API synchronously
func (client *Client) QueryEmailVerification(request *QueryEmailVerificationRequest) (response *QueryEmailVerificationResponse, err error) {
	response = CreateQueryEmailVerificationResponse()
	err = client.DoAction(request, response)
	return
}

// QueryEmailVerificationWithChan invokes the domain.QueryEmailVerification API asynchronously
func (client *Client) QueryEmailVerificationWithChan(request *QueryEmailVerificationRequest) (<-chan *QueryEmailVerificationResponse, <-chan error) {
	responseChan := make(chan *QueryEmailVerificationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryEmailVerification(request)
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

// QueryEmailVerificationWithCallback invokes the domain.QueryEmailVerification API asynchronously
func (client *Client) QueryEmailVerificationWithCallback(request *QueryEmailVerificationRequest, callback func(response *QueryEmailVerificationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryEmailVerificationResponse
		var err error
		defer close(result)
		response, err = client.QueryEmailVerification(request)
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

// QueryEmailVerificationRequest is the request struct for api QueryEmailVerification
type QueryEmailVerificationRequest struct {
	*requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	Email        string `position:"Query" name:"Email"`
}

// QueryEmailVerificationResponse is the response struct for api QueryEmailVerification
type QueryEmailVerificationResponse struct {
	*responses.BaseResponse
	RequestId           string `json:"RequestId" xml:"RequestId"`
	GmtCreate           string `json:"GmtCreate" xml:"GmtCreate"`
	GmtModified         string `json:"GmtModified" xml:"GmtModified"`
	Email               string `json:"Email" xml:"Email"`
	UserId              string `json:"UserId" xml:"UserId"`
	EmailVerificationNo string `json:"EmailVerificationNo" xml:"EmailVerificationNo"`
	TokenSendTime       string `json:"TokenSendTime" xml:"TokenSendTime"`
	VerificationStatus  int    `json:"VerificationStatus" xml:"VerificationStatus"`
	VerificationTime    string `json:"VerificationTime" xml:"VerificationTime"`
	SendIp              string `json:"SendIp" xml:"SendIp"`
	ConfirmIp           string `json:"ConfirmIp" xml:"ConfirmIp"`
}

// CreateQueryEmailVerificationRequest creates a request to invoke QueryEmailVerification API
func CreateQueryEmailVerificationRequest() (request *QueryEmailVerificationRequest) {
	request = &QueryEmailVerificationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "QueryEmailVerification", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryEmailVerificationResponse creates a response to parse from QueryEmailVerification response
func CreateQueryEmailVerificationResponse() (response *QueryEmailVerificationResponse) {
	response = &QueryEmailVerificationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
