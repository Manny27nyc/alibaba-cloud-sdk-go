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

// PlanComponentTopo invokes the emr.PlanComponentTopo API synchronously
// api document: https://help.aliyun.com/api/emr/plancomponenttopo.html
func (client *Client) PlanComponentTopo(request *PlanComponentTopoRequest) (response *PlanComponentTopoResponse, err error) {
	response = CreatePlanComponentTopoResponse()
	err = client.DoAction(request, response)
	return
}

// PlanComponentTopoWithChan invokes the emr.PlanComponentTopo API asynchronously
// api document: https://help.aliyun.com/api/emr/plancomponenttopo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PlanComponentTopoWithChan(request *PlanComponentTopoRequest) (<-chan *PlanComponentTopoResponse, <-chan error) {
	responseChan := make(chan *PlanComponentTopoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PlanComponentTopo(request)
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

// PlanComponentTopoWithCallback invokes the emr.PlanComponentTopo API asynchronously
// api document: https://help.aliyun.com/api/emr/plancomponenttopo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PlanComponentTopoWithCallback(request *PlanComponentTopoRequest, callback func(response *PlanComponentTopoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PlanComponentTopoResponse
		var err error
		defer close(result)
		response, err = client.PlanComponentTopo(request)
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

// PlanComponentTopoRequest is the request struct for api PlanComponentTopo
type PlanComponentTopoRequest struct {
	*requests.RpcRequest
	ClusterType     string                          `position:"Query" name:"ClusterType"`
	ResourceOwnerId requests.Integer                `position:"Query" name:"ResourceOwnerId"`
	HostGroup       *[]PlanComponentTopoHostGroup   `position:"Query" name:"HostGroup"  type:"Repeated"`
	HostInfo        *[]PlanComponentTopoHostInfo    `position:"Query" name:"HostInfo"  type:"Repeated"`
	StackName       string                          `position:"Query" name:"StackName"`
	ClusterId       string                          `position:"Query" name:"ClusterId"`
	StackVersion    string                          `position:"Query" name:"StackVersion"`
	ServiceInfo     *[]PlanComponentTopoServiceInfo `position:"Query" name:"ServiceInfo"  type:"Repeated"`
}

// PlanComponentTopoHostGroup is a repeated param struct in PlanComponentTopoRequest
type PlanComponentTopoHostGroup struct {
	GroupType string `name:"GroupType"`
	NodeCount string `name:"NodeCount"`
	GroupName string `name:"GroupName"`
}

// PlanComponentTopoHostInfo is a repeated param struct in PlanComponentTopoRequest
type PlanComponentTopoHostInfo struct {
	HpHostBizId   string `name:"HpHostBizId"`
	HostGroupName string `name:"HostGroupName"`
}

// PlanComponentTopoServiceInfo is a repeated param struct in PlanComponentTopoRequest
type PlanComponentTopoServiceInfo struct {
	ServiceEcmVersion string `name:"ServiceEcmVersion"`
	ServiceVersion    string `name:"ServiceVersion"`
	ServiceName       string `name:"ServiceName"`
}

// PlanComponentTopoResponse is the response struct for api PlanComponentTopo
type PlanComponentTopoResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	HostComponentList HostComponentList `json:"HostComponentList" xml:"HostComponentList"`
}

// CreatePlanComponentTopoRequest creates a request to invoke PlanComponentTopo API
func CreatePlanComponentTopoRequest() (request *PlanComponentTopoRequest) {
	request = &PlanComponentTopoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "PlanComponentTopo", "emr", "openAPI")
	return
}

// CreatePlanComponentTopoResponse creates a response to parse from PlanComponentTopo response
func CreatePlanComponentTopoResponse() (response *PlanComponentTopoResponse) {
	response = &PlanComponentTopoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
