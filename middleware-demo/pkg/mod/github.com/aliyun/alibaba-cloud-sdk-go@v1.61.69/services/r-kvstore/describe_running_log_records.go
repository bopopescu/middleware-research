package r_kvstore

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

// DescribeRunningLogRecords invokes the r_kvstore.DescribeRunningLogRecords API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/describerunninglogrecords.html
func (client *Client) DescribeRunningLogRecords(request *DescribeRunningLogRecordsRequest) (response *DescribeRunningLogRecordsResponse, err error) {
	response = CreateDescribeRunningLogRecordsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRunningLogRecordsWithChan invokes the r_kvstore.DescribeRunningLogRecords API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describerunninglogrecords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRunningLogRecordsWithChan(request *DescribeRunningLogRecordsRequest) (<-chan *DescribeRunningLogRecordsResponse, <-chan error) {
	responseChan := make(chan *DescribeRunningLogRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRunningLogRecords(request)
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

// DescribeRunningLogRecordsWithCallback invokes the r_kvstore.DescribeRunningLogRecords API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describerunninglogrecords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRunningLogRecordsWithCallback(request *DescribeRunningLogRecordsRequest, callback func(response *DescribeRunningLogRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRunningLogRecordsResponse
		var err error
		defer close(result)
		response, err = client.DescribeRunningLogRecords(request)
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

// DescribeRunningLogRecordsRequest is the request struct for api DescribeRunningLogRecords
type DescribeRunningLogRecordsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	RoleType             string           `position:"Query" name:"RoleType"`
	NodeId               string           `position:"Query" name:"NodeId"`
	SQLId                requests.Integer `position:"Query" name:"SQLId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	DBName               string           `position:"Query" name:"DBName"`
}

// DescribeRunningLogRecordsResponse is the response struct for api DescribeRunningLogRecords
type DescribeRunningLogRecordsResponse struct {
	*responses.BaseResponse
	RequestId        string                           `json:"RequestId" xml:"RequestId"`
	InstanceId       string                           `json:"InstanceId" xml:"InstanceId"`
	StartTime        string                           `json:"StartTime" xml:"StartTime"`
	Engine           string                           `json:"Engine" xml:"Engine"`
	TotalRecordCount int                              `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageNumber       int                              `json:"PageNumber" xml:"PageNumber"`
	PageSize         int                              `json:"PageSize" xml:"PageSize"`
	PageRecordCount  int                              `json:"PageRecordCount" xml:"PageRecordCount"`
	Items            ItemsInDescribeRunningLogRecords `json:"Items" xml:"Items"`
}

// CreateDescribeRunningLogRecordsRequest creates a request to invoke DescribeRunningLogRecords API
func CreateDescribeRunningLogRecordsRequest() (request *DescribeRunningLogRecordsRequest) {
	request = &DescribeRunningLogRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeRunningLogRecords", "", "")
	return
}

// CreateDescribeRunningLogRecordsResponse creates a response to parse from DescribeRunningLogRecords response
func CreateDescribeRunningLogRecordsResponse() (response *DescribeRunningLogRecordsResponse) {
	response = &DescribeRunningLogRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
