package aegis

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

// DescribeAnalysisLogs invokes the aegis.DescribeAnalysisLogs API synchronously
// api document: https://help.aliyun.com/api/aegis/describeanalysislogs.html
func (client *Client) DescribeAnalysisLogs(request *DescribeAnalysisLogsRequest) (response *DescribeAnalysisLogsResponse, err error) {
	response = CreateDescribeAnalysisLogsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAnalysisLogsWithChan invokes the aegis.DescribeAnalysisLogs API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeanalysislogs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAnalysisLogsWithChan(request *DescribeAnalysisLogsRequest) (<-chan *DescribeAnalysisLogsResponse, <-chan error) {
	responseChan := make(chan *DescribeAnalysisLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAnalysisLogs(request)
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

// DescribeAnalysisLogsWithCallback invokes the aegis.DescribeAnalysisLogs API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeanalysislogs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAnalysisLogsWithCallback(request *DescribeAnalysisLogsRequest, callback func(response *DescribeAnalysisLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAnalysisLogsResponse
		var err error
		defer close(result)
		response, err = client.DescribeAnalysisLogs(request)
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

// DescribeAnalysisLogsRequest is the request struct for api DescribeAnalysisLogs
type DescribeAnalysisLogsRequest struct {
	*requests.RpcRequest
	SourceIp    string           `position:"Query" name:"SourceIp"`
	Query       string           `position:"Query" name:"Query"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	From        requests.Integer `position:"Query" name:"From"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	To          requests.Integer `position:"Query" name:"To"`
	Reverse     requests.Boolean `position:"Query" name:"Reverse"`
}

// DescribeAnalysisLogsResponse is the response struct for api DescribeAnalysisLogs
type DescribeAnalysisLogsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Logs      Logs   `json:"Logs" xml:"Logs"`
}

// CreateDescribeAnalysisLogsRequest creates a request to invoke DescribeAnalysisLogs API
func CreateDescribeAnalysisLogsRequest() (request *DescribeAnalysisLogsRequest) {
	request = &DescribeAnalysisLogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeAnalysisLogs", "vipaegis", "openAPI")
	return
}

// CreateDescribeAnalysisLogsResponse creates a response to parse from DescribeAnalysisLogs response
func CreateDescribeAnalysisLogsResponse() (response *DescribeAnalysisLogsResponse) {
	response = &DescribeAnalysisLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
