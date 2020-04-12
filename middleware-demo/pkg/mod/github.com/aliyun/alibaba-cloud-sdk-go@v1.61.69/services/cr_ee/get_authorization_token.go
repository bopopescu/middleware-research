package cr_ee

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

// GetAuthorizationToken invokes the cr.GetAuthorizationToken API synchronously
// api document: https://help.aliyun.com/api/cr/getauthorizationtoken.html
func (client *Client) GetAuthorizationToken(request *GetAuthorizationTokenRequest) (response *GetAuthorizationTokenResponse, err error) {
	response = CreateGetAuthorizationTokenResponse()
	err = client.DoAction(request, response)
	return
}

// GetAuthorizationTokenWithChan invokes the cr.GetAuthorizationToken API asynchronously
// api document: https://help.aliyun.com/api/cr/getauthorizationtoken.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAuthorizationTokenWithChan(request *GetAuthorizationTokenRequest) (<-chan *GetAuthorizationTokenResponse, <-chan error) {
	responseChan := make(chan *GetAuthorizationTokenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAuthorizationToken(request)
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

// GetAuthorizationTokenWithCallback invokes the cr.GetAuthorizationToken API asynchronously
// api document: https://help.aliyun.com/api/cr/getauthorizationtoken.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAuthorizationTokenWithCallback(request *GetAuthorizationTokenRequest, callback func(response *GetAuthorizationTokenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAuthorizationTokenResponse
		var err error
		defer close(result)
		response, err = client.GetAuthorizationToken(request)
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

// GetAuthorizationTokenRequest is the request struct for api GetAuthorizationToken
type GetAuthorizationTokenRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// GetAuthorizationTokenResponse is the response struct for api GetAuthorizationToken
type GetAuthorizationTokenResponse struct {
	*responses.BaseResponse
	GetAuthorizationTokenIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
	Code                           string `json:"Code" xml:"Code"`
	RequestId                      string `json:"RequestId" xml:"RequestId"`
	TempUsername                   string `json:"TempUsername" xml:"TempUsername"`
	AuthorizationToken             string `json:"AuthorizationToken" xml:"AuthorizationToken"`
	ExpireTime                     int64  `json:"ExpireTime" xml:"ExpireTime"`
}

// CreateGetAuthorizationTokenRequest creates a request to invoke GetAuthorizationToken API
func CreateGetAuthorizationTokenRequest() (request *GetAuthorizationTokenRequest) {
	request = &GetAuthorizationTokenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cr", "2018-12-01", "GetAuthorizationToken", "cr", "openAPI")
	return
}

// CreateGetAuthorizationTokenResponse creates a response to parse from GetAuthorizationToken response
func CreateGetAuthorizationTokenResponse() (response *GetAuthorizationTokenResponse) {
	response = &GetAuthorizationTokenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
