package push

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

// QueryPushRecords invokes the push.QueryPushRecords API synchronously
// api document: https://help.aliyun.com/api/push/querypushrecords.html
func (client *Client) QueryPushRecords(request *QueryPushRecordsRequest) (response *QueryPushRecordsResponse, err error) {
	response = CreateQueryPushRecordsResponse()
	err = client.DoAction(request, response)
	return
}

// QueryPushRecordsWithChan invokes the push.QueryPushRecords API asynchronously
// api document: https://help.aliyun.com/api/push/querypushrecords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryPushRecordsWithChan(request *QueryPushRecordsRequest) (<-chan *QueryPushRecordsResponse, <-chan error) {
	responseChan := make(chan *QueryPushRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryPushRecords(request)
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

// QueryPushRecordsWithCallback invokes the push.QueryPushRecords API asynchronously
// api document: https://help.aliyun.com/api/push/querypushrecords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryPushRecordsWithCallback(request *QueryPushRecordsRequest, callback func(response *QueryPushRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryPushRecordsResponse
		var err error
		defer close(result)
		response, err = client.QueryPushRecords(request)
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

// QueryPushRecordsRequest is the request struct for api QueryPushRecords
type QueryPushRecordsRequest struct {
	*requests.RpcRequest
	StartTime string           `position:"Query" name:"StartTime"`
	Source    string           `position:"Query" name:"Source"`
	NextToken string           `position:"Query" name:"NextToken"`
	PageSize  requests.Integer `position:"Query" name:"PageSize"`
	Keyword   string           `position:"Query" name:"Keyword"`
	EndTime   string           `position:"Query" name:"EndTime"`
	Target    string           `position:"Query" name:"Target"`
	AppKey    requests.Integer `position:"Query" name:"AppKey"`
	PushType  string           `position:"Query" name:"PushType"`
}

// QueryPushRecordsResponse is the response struct for api QueryPushRecords
type QueryPushRecordsResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	NextToken string    `json:"NextToken" xml:"NextToken"`
	PageSize  int       `json:"PageSize" xml:"PageSize"`
	PushInfos PushInfos `json:"PushInfos" xml:"PushInfos"`
}

// CreateQueryPushRecordsRequest creates a request to invoke QueryPushRecords API
func CreateQueryPushRecordsRequest() (request *QueryPushRecordsRequest) {
	request = &QueryPushRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "QueryPushRecords", "cps", "openAPI")
	return
}

// CreateQueryPushRecordsResponse creates a response to parse from QueryPushRecords response
func CreateQueryPushRecordsResponse() (response *QueryPushRecordsResponse) {
	response = &QueryPushRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
