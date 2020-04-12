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

// NodeInstance is a nested struct in emr response
type NodeInstance struct {
	RunConf        string `json:"RunConf" xml:"RunConf"`
	MaxRetry       int    `json:"MaxRetry" xml:"MaxRetry"`
	EndTime        int64  `json:"EndTime" xml:"EndTime"`
	StartTime      int64  `json:"StartTime" xml:"StartTime"`
	NodeName       string `json:"NodeName" xml:"NodeName"`
	ProjectId      string `json:"ProjectId" xml:"ProjectId"`
	Id             string `json:"Id" xml:"Id"`
	JobType        string `json:"JobType" xml:"JobType"`
	JobName        string `json:"JobName" xml:"JobName"`
	Type           string `json:"Type" xml:"Type"`
	JobId          string `json:"JobId" xml:"JobId"`
	FailAct        string `json:"FailAct" xml:"FailAct"`
	ClusterId      string `json:"ClusterId" xml:"ClusterId"`
	RetryInterval  int64  `json:"RetryInterval" xml:"RetryInterval"`
	ParamConf      string `json:"ParamConf" xml:"ParamConf"`
	ExternalId     string `json:"ExternalId" xml:"ExternalId"`
	EnvConf        string `json:"EnvConf" xml:"EnvConf"`
	GmtCreate      int64  `json:"GmtCreate" xml:"GmtCreate"`
	ExternalInfo   string `json:"ExternalInfo" xml:"ExternalInfo"`
	Retries        int    `json:"Retries" xml:"Retries"`
	GmtModified    int64  `json:"GmtModified" xml:"GmtModified"`
	ExternalStatus string `json:"ExternalStatus" xml:"ExternalStatus"`
	JobParams      string `json:"JobParams" xml:"JobParams"`
	HostName       string `json:"HostName" xml:"HostName"`
	Status         string `json:"Status" xml:"Status"`
	Pending        bool   `json:"pending" xml:"pending"`
}
