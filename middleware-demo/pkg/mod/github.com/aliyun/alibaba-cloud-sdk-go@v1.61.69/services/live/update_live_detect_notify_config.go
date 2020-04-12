package live

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

// UpdateLiveDetectNotifyConfig invokes the live.UpdateLiveDetectNotifyConfig API synchronously
// api document: https://help.aliyun.com/api/live/updatelivedetectnotifyconfig.html
func (client *Client) UpdateLiveDetectNotifyConfig(request *UpdateLiveDetectNotifyConfigRequest) (response *UpdateLiveDetectNotifyConfigResponse, err error) {
	response = CreateUpdateLiveDetectNotifyConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateLiveDetectNotifyConfigWithChan invokes the live.UpdateLiveDetectNotifyConfig API asynchronously
// api document: https://help.aliyun.com/api/live/updatelivedetectnotifyconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateLiveDetectNotifyConfigWithChan(request *UpdateLiveDetectNotifyConfigRequest) (<-chan *UpdateLiveDetectNotifyConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateLiveDetectNotifyConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateLiveDetectNotifyConfig(request)
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

// UpdateLiveDetectNotifyConfigWithCallback invokes the live.UpdateLiveDetectNotifyConfig API asynchronously
// api document: https://help.aliyun.com/api/live/updatelivedetectnotifyconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateLiveDetectNotifyConfigWithCallback(request *UpdateLiveDetectNotifyConfigRequest, callback func(response *UpdateLiveDetectNotifyConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateLiveDetectNotifyConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateLiveDetectNotifyConfig(request)
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

// UpdateLiveDetectNotifyConfigRequest is the request struct for api UpdateLiveDetectNotifyConfig
type UpdateLiveDetectNotifyConfigRequest struct {
	*requests.RpcRequest
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	NotifyUrl     string           `position:"Query" name:"NotifyUrl"`
}

// UpdateLiveDetectNotifyConfigResponse is the response struct for api UpdateLiveDetectNotifyConfig
type UpdateLiveDetectNotifyConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateLiveDetectNotifyConfigRequest creates a request to invoke UpdateLiveDetectNotifyConfig API
func CreateUpdateLiveDetectNotifyConfigRequest() (request *UpdateLiveDetectNotifyConfigRequest) {
	request = &UpdateLiveDetectNotifyConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "UpdateLiveDetectNotifyConfig", "live", "openAPI")
	return
}

// CreateUpdateLiveDetectNotifyConfigResponse creates a response to parse from UpdateLiveDetectNotifyConfig response
func CreateUpdateLiveDetectNotifyConfigResponse() (response *UpdateLiveDetectNotifyConfigResponse) {
	response = &UpdateLiveDetectNotifyConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
