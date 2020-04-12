package outboundbot

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

// DownloadRecording invokes the outboundbot.DownloadRecording API synchronously
// api document: https://help.aliyun.com/api/outboundbot/downloadrecording.html
func (client *Client) DownloadRecording(request *DownloadRecordingRequest) (response *DownloadRecordingResponse, err error) {
	response = CreateDownloadRecordingResponse()
	err = client.DoAction(request, response)
	return
}

// DownloadRecordingWithChan invokes the outboundbot.DownloadRecording API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/downloadrecording.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DownloadRecordingWithChan(request *DownloadRecordingRequest) (<-chan *DownloadRecordingResponse, <-chan error) {
	responseChan := make(chan *DownloadRecordingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DownloadRecording(request)
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

// DownloadRecordingWithCallback invokes the outboundbot.DownloadRecording API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/downloadrecording.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DownloadRecordingWithCallback(request *DownloadRecordingRequest, callback func(response *DownloadRecordingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DownloadRecordingResponse
		var err error
		defer close(result)
		response, err = client.DownloadRecording(request)
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

// DownloadRecordingRequest is the request struct for api DownloadRecording
type DownloadRecordingRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	TaskId     string `position:"Query" name:"TaskId"`
}

// DownloadRecordingResponse is the response struct for api DownloadRecording
type DownloadRecordingResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	Success        bool           `json:"Success" xml:"Success"`
	Code           string         `json:"Code" xml:"Code"`
	Message        string         `json:"Message" xml:"Message"`
	HttpStatusCode int            `json:"HttpStatusCode" xml:"HttpStatusCode"`
	DownloadParams DownloadParams `json:"DownloadParams" xml:"DownloadParams"`
}

// CreateDownloadRecordingRequest creates a request to invoke DownloadRecording API
func CreateDownloadRecordingRequest() (request *DownloadRecordingRequest) {
	request = &DownloadRecordingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "DownloadRecording", "outboundbot", "openAPI")
	return
}

// CreateDownloadRecordingResponse creates a response to parse from DownloadRecording response
func CreateDownloadRecordingResponse() (response *DownloadRecordingResponse) {
	response = &DownloadRecordingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
