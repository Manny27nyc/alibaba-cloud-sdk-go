package config

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

// DeliveryChannel is a nested struct in config response
type DeliveryChannel struct {
	Status                              int    `json:"Status" xml:"Status"`
	ConfigurationItemChangeNotification bool   `json:"ConfigurationItemChangeNotification" xml:"ConfigurationItemChangeNotification"`
	NonCompliantNotification            bool   `json:"NonCompliantNotification" xml:"NonCompliantNotification"`
	DeliveryChannelName                 string `json:"DeliveryChannelName" xml:"DeliveryChannelName"`
	AggregatorId                        string `json:"AggregatorId" xml:"AggregatorId"`
	DeliveryChannelId                   string `json:"DeliveryChannelId" xml:"DeliveryChannelId"`
	DeliveryChannelAssumeRoleArn        string `json:"DeliveryChannelAssumeRoleArn" xml:"DeliveryChannelAssumeRoleArn"`
	AccountId                           string `json:"AccountId" xml:"AccountId"`
	DeliveryChannelType                 string `json:"DeliveryChannelType" xml:"DeliveryChannelType"`
	ConfigurationSnapshot               bool   `json:"ConfigurationSnapshot" xml:"ConfigurationSnapshot"`
	OversizedDataOSSTargetArn           string `json:"OversizedDataOSSTargetArn" xml:"OversizedDataOSSTargetArn"`
	DeliveryChannelTargetArn            string `json:"DeliveryChannelTargetArn" xml:"DeliveryChannelTargetArn"`
	Description                         string `json:"Description" xml:"Description"`
	DeliveryChannelCondition            string `json:"DeliveryChannelCondition" xml:"DeliveryChannelCondition"`
}
