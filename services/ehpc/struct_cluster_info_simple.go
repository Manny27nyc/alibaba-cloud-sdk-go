package ehpc

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

type ClusterInfoSimple struct {
	Id              string         `json:"Id" xml:"Id"`
	RegionId        string         `json:"RegionId" xml:"RegionId"`
	ZoneId          string         `json:"ZoneId" xml:"ZoneId"`
	Name            string         `json:"Name" xml:"Name"`
	Description     string         `json:"Description" xml:"Description"`
	Status          string         `json:"Status" xml:"Status"`
	OsTag           string         `json:"OsTag" xml:"OsTag"`
	AccountType     string         `json:"AccountType" xml:"AccountType"`
	SchedulerType   string         `json:"SchedulerType" xml:"SchedulerType"`
	Count           int            `json:"Count" xml:"Count"`
	InstanceType    string         `json:"InstanceType" xml:"InstanceType"`
	LoginNodes      string         `json:"LoginNodes" xml:"LoginNodes"`
	CreateTime      string         `json:"CreateTime" xml:"CreateTime"`
	ImageOwnerAlias string         `json:"ImageOwnerAlias" xml:"ImageOwnerAlias"`
	ImageId         string         `json:"ImageId" xml:"ImageId"`
	Managers        Managers       `json:"Managers" xml:"Managers"`
	Computes        Computes       `json:"Computes" xml:"Computes"`
	TotalResources  TotalResources `json:"TotalResources" xml:"TotalResources"`
	UsedResources   UsedResources  `json:"UsedResources" xml:"UsedResources"`
}
