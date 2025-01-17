package das

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

// Data is a nested struct in das response
type Data struct {
	PageSize            int64                                      `json:"PageSize" xml:"PageSize"`
	Timestamp           int64                                      `json:"timestamp" xml:"timestamp"`
	TaskState           string                                     `json:"TaskState" xml:"TaskState"`
	BenchStep           string                                     `json:"BenchStep" xml:"BenchStep"`
	Source              string                                     `json:"Source" xml:"Source"`
	ErrMsg              string                                     `json:"ErrMsg" xml:"ErrMsg"`
	TaskId              string                                     `json:"TaskId" xml:"TaskId"`
	Role                string                                     `json:"Role" xml:"Role"`
	SqlCompleteReuse    string                                     `json:"SqlCompleteReuse" xml:"SqlCompleteReuse"`
	DtsJobId            string                                     `json:"DtsJobId" xml:"DtsJobId"`
	EndState            string                                     `json:"EndState" xml:"EndState"`
	SqlFileOnOss        string                                     `json:"SqlFileOnOss" xml:"SqlFileOnOss"`
	Port                int                                        `json:"Port" xml:"Port"`
	Error               string                                     `json:"Error" xml:"Error"`
	Message             string                                     `json:"Message" xml:"Message"`
	CallerUid           string                                     `json:"CallerUid" xml:"CallerUid"`
	ErrorMsg            string                                     `json:"ErrorMsg" xml:"ErrorMsg"`
	TableSchema         string                                     `json:"TableSchema" xml:"TableSchema"`
	DtsJobClass         string                                     `json:"DtsJobClass" xml:"DtsJobClass"`
	PageNo              int64                                      `json:"PageNo" xml:"PageNo"`
	Extra               string                                     `json:"Extra" xml:"Extra"`
	OriUuid             string                                     `json:"OriUuid" xml:"OriUuid"`
	AccountId           string                                     `json:"AccountId" xml:"AccountId"`
	ArchiveState        int                                        `json:"ArchiveState" xml:"ArchiveState"`
	Complete            bool                                       `json:"complete" xml:"complete"`
	TaskType            string                                     `json:"TaskType" xml:"TaskType"`
	Status              string                                     `json:"Status" xml:"Status"`
	Rate                int64                                      `json:"Rate" xml:"Rate"`
	ResultId            string                                     `json:"resultId" xml:"resultId"`
	WorkDir             string                                     `json:"WorkDir" xml:"WorkDir"`
	ArchiveJobId        string                                     `json:"ArchiveJobId" xml:"ArchiveJobId"`
	State               string                                     `json:"state" xml:"state"`
	External            string                                     `json:"External" xml:"External"`
	CreateTime          string                                     `json:"CreateTime" xml:"CreateTime"`
	DownloadUrl         string                                     `json:"DownloadUrl" xml:"DownloadUrl"`
	JarOnOss            string                                     `json:"JarOnOss" xml:"JarOnOss"`
	SqlFilePath         string                                     `json:"SqlFilePath" xml:"SqlFilePath"`
	InstanceId          string                                     `json:"InstanceId" xml:"InstanceId"`
	SrcPublicIp         string                                     `json:"SrcPublicIp" xml:"SrcPublicIp"`
	ExpireTime          string                                     `json:"ExpireTime" xml:"ExpireTime"`
	DtsJobName          string                                     `json:"DtsJobName" xml:"DtsJobName"`
	ClientJarPath       string                                     `json:"ClientJarPath" xml:"ClientJarPath"`
	BenchCmd            string                                     `json:"BenchCmd" xml:"BenchCmd"`
	TenantId            string                                     `json:"TenantId" xml:"TenantId"`
	JobId               string                                     `json:"JobId" xml:"JobId"`
	Ip                  string                                     `json:"Ip" xml:"Ip"`
	Version             string                                     `json:"Version" xml:"Version"`
	RocksDbPath         string                                     `json:"RocksDbPath" xml:"RocksDbPath"`
	BackupType          string                                     `json:"BackupType" xml:"BackupType"`
	RequestDuration     int64                                      `json:"RequestDuration" xml:"RequestDuration"`
	ParseFilePath       string                                     `json:"ParseFilePath" xml:"ParseFilePath"`
	ClientType          string                                     `json:"ClientType" xml:"ClientType"`
	ClientGatewayId     string                                     `json:"ClientGatewayId" xml:"ClientGatewayId"`
	ParseCmd            string                                     `json:"ParseCmd" xml:"ParseCmd"`
	OwnerId             string                                     `json:"OwnerId" xml:"OwnerId"`
	MetaFilePath        string                                     `json:"MetaFilePath" xml:"MetaFilePath"`
	Code                int                                        `json:"Code" xml:"Code"`
	Total               int64                                      `json:"Total" xml:"Total"`
	DstIp               string                                     `json:"DstIp" xml:"DstIp"`
	DtsJobStatus        string                                     `json:"DtsJobStatus" xml:"DtsJobStatus"`
	DstType             string                                     `json:"DstType" xml:"DstType"`
	IsFinish            bool                                       `json:"isFinish" xml:"isFinish"`
	VpcId               string                                     `json:"VpcId" xml:"VpcId"`
	Topic               string                                     `json:"Topic" xml:"Topic"`
	Token               string                                     `json:"Token" xml:"Token"`
	DstInstanceUuid     string                                     `json:"DstInstanceUuid" xml:"DstInstanceUuid"`
	DbLinkId            int64                                      `json:"DbLinkId" xml:"DbLinkId"`
	SrcInstanceUuid     string                                     `json:"SrcInstanceUuid" xml:"SrcInstanceUuid"`
	ArchiveFolder       string                                     `json:"ArchiveFolder" xml:"ArchiveFolder"`
	EcsInstanceId       string                                     `json:"EcsInstanceId" xml:"EcsInstanceId"`
	ArchiveOssTableName string                                     `json:"ArchiveOssTableName" xml:"ArchiveOssTableName"`
	MetaFileOnOss       string                                     `json:"MetaFileOnOss" xml:"MetaFileOnOss"`
	Fail                bool                                       `json:"fail" xml:"fail"`
	BackupId            string                                     `json:"BackupId" xml:"BackupId"`
	ErrorMessage        string                                     `json:"ErrorMessage" xml:"ErrorMessage"`
	Uuid                string                                     `json:"Uuid" xml:"Uuid"`
	NodeId              string                                     `json:"NodeId" xml:"NodeId"`
	DtsJobState         int                                        `json:"DtsJobState" xml:"DtsJobState"`
	UserId              string                                     `json:"UserId" xml:"UserId"`
	SrcInstanceArea     string                                     `json:"SrcInstanceArea" xml:"SrcInstanceArea"`
	MetaFileName        string                                     `json:"MetaFileName" xml:"MetaFileName"`
	Description         string                                     `json:"Description" xml:"Description"`
	ErrorCode           string                                     `json:"ErrorCode" xml:"ErrorCode"`
	BenchStepStatus     string                                     `json:"BenchStepStatus" xml:"BenchStepStatus"`
	SmartPressureTime   int                                        `json:"SmartPressureTime" xml:"SmartPressureTime"`
	Results             string                                     `json:"Results" xml:"Results"`
	SqlFileName         string                                     `json:"SqlFileName" xml:"SqlFileName"`
	SyncStatus          string                                     `json:"SyncStatus" xml:"SyncStatus"`
	DstPort             int                                        `json:"DstPort" xml:"DstPort"`
	StatusCode          string                                     `json:"StatusCode" xml:"StatusCode"`
	LoadCmd             string                                     `json:"LoadCmd" xml:"LoadCmd"`
	SubResults          SubResultsInGetHDMAliyunResourceSyncResult `json:"SubResults" xml:"SubResults"`
	KeyPrefixes         KeyPrefixes                                `json:"KeyPrefixes" xml:"KeyPrefixes"`
	List                ListInDescribeCacheAnalysisJobs            `json:"List" xml:"List"`
	BigKeys             BigKeysInCreateCacheAnalysisJob            `json:"BigKeys" xml:"BigKeys"`
	Result              []ResultItem                               `json:"result" xml:"result"`
}
