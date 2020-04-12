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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// GetTotalUserReport invokes the ehpc.GetTotalUserReport API synchronously
// api document: https://help.aliyun.com/api/ehpc/gettotaluserreport.html
func (client *Client) GetTotalUserReport(request *GetTotalUserReportRequest) (response *GetTotalUserReportResponse, err error) {
	response = CreateGetTotalUserReportResponse()
	err = client.DoAction(request, response)
	return
}

// GetTotalUserReportWithChan invokes the ehpc.GetTotalUserReport API asynchronously
// api document: https://help.aliyun.com/api/ehpc/gettotaluserreport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetTotalUserReportWithChan(request *GetTotalUserReportRequest) (<-chan *GetTotalUserReportResponse, <-chan error) {
	responseChan := make(chan *GetTotalUserReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTotalUserReport(request)
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

// GetTotalUserReportWithCallback invokes the ehpc.GetTotalUserReport API asynchronously
// api document: https://help.aliyun.com/api/ehpc/gettotaluserreport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetTotalUserReportWithCallback(request *GetTotalUserReportRequest, callback func(response *GetTotalUserReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTotalUserReportResponse
		var err error
		defer close(result)
		response, err = client.GetTotalUserReport(request)
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

// GetTotalUserReportRequest is the request struct for api GetTotalUserReport
type GetTotalUserReportRequest struct {
	*requests.RpcRequest
	EndTime   requests.Integer `position:"Query" name:"EndTime"`
	ClusterId string           `position:"Query" name:"ClusterId"`
	StartTime requests.Integer `position:"Query" name:"StartTime"`
}

// GetTotalUserReportResponse is the response struct for api GetTotalUserReport
type GetTotalUserReportResponse struct {
	*responses.BaseResponse
	RequestId string                   `json:"RequestId" xml:"RequestId"`
	Metrics   string                   `json:"Metrics" xml:"Metrics"`
	Data      DataInGetTotalUserReport `json:"Data" xml:"Data"`
}

// CreateGetTotalUserReportRequest creates a request to invoke GetTotalUserReport API
func CreateGetTotalUserReportRequest() (request *GetTotalUserReportRequest) {
	request = &GetTotalUserReportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "GetTotalUserReport", "ehs", "openAPI")
	return
}

// CreateGetTotalUserReportResponse creates a response to parse from GetTotalUserReport response
func CreateGetTotalUserReportResponse() (response *GetTotalUserReportResponse) {
	response = &GetTotalUserReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
