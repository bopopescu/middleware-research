package voicenavigator

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

// ExportConversationDetails invokes the voicenavigator.ExportConversationDetails API synchronously
// api document: https://help.aliyun.com/api/voicenavigator/exportconversationdetails.html
func (client *Client) ExportConversationDetails(request *ExportConversationDetailsRequest) (response *ExportConversationDetailsResponse, err error) {
	response = CreateExportConversationDetailsResponse()
	err = client.DoAction(request, response)
	return
}

// ExportConversationDetailsWithChan invokes the voicenavigator.ExportConversationDetails API asynchronously
// api document: https://help.aliyun.com/api/voicenavigator/exportconversationdetails.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ExportConversationDetailsWithChan(request *ExportConversationDetailsRequest) (<-chan *ExportConversationDetailsResponse, <-chan error) {
	responseChan := make(chan *ExportConversationDetailsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExportConversationDetails(request)
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

// ExportConversationDetailsWithCallback invokes the voicenavigator.ExportConversationDetails API asynchronously
// api document: https://help.aliyun.com/api/voicenavigator/exportconversationdetails.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ExportConversationDetailsWithCallback(request *ExportConversationDetailsRequest, callback func(response *ExportConversationDetailsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExportConversationDetailsResponse
		var err error
		defer close(result)
		response, err = client.ExportConversationDetails(request)
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

// ExportConversationDetailsRequest is the request struct for api ExportConversationDetails
type ExportConversationDetailsRequest struct {
	*requests.RpcRequest
	BeginTimeLeftRange  requests.Integer `position:"Query" name:"BeginTimeLeftRange"`
	CallingNumber       string           `position:"Query" name:"CallingNumber"`
	InstanceId          string           `position:"Query" name:"InstanceId"`
	BeginTimeRightRange requests.Integer `position:"Query" name:"BeginTimeRightRange"`
}

// ExportConversationDetailsResponse is the response struct for api ExportConversationDetails
type ExportConversationDetailsResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ExportTaskId string `json:"ExportTaskId" xml:"ExportTaskId"`
}

// CreateExportConversationDetailsRequest creates a request to invoke ExportConversationDetails API
func CreateExportConversationDetailsRequest() (request *ExportConversationDetailsRequest) {
	request = &ExportConversationDetailsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("VoiceNavigator", "2018-06-12", "ExportConversationDetails", "voicebot", "openAPI")
	return
}

// CreateExportConversationDetailsResponse creates a response to parse from ExportConversationDetails response
func CreateExportConversationDetailsResponse() (response *ExportConversationDetailsResponse) {
	response = &ExportConversationDetailsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
