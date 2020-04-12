package eas

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

// GetServiceVersion invokes the eas.GetServiceVersion API synchronously
// api document: https://help.aliyun.com/api/eas/getserviceversion.html
func (client *Client) GetServiceVersion(request *GetServiceVersionRequest) (response *GetServiceVersionResponse, err error) {
	response = CreateGetServiceVersionResponse()
	err = client.DoAction(request, response)
	return
}

// GetServiceVersionWithChan invokes the eas.GetServiceVersion API asynchronously
// api document: https://help.aliyun.com/api/eas/getserviceversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetServiceVersionWithChan(request *GetServiceVersionRequest) (<-chan *GetServiceVersionResponse, <-chan error) {
	responseChan := make(chan *GetServiceVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetServiceVersion(request)
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

// GetServiceVersionWithCallback invokes the eas.GetServiceVersion API asynchronously
// api document: https://help.aliyun.com/api/eas/getserviceversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetServiceVersionWithCallback(request *GetServiceVersionRequest, callback func(response *GetServiceVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetServiceVersionResponse
		var err error
		defer close(result)
		response, err = client.GetServiceVersion(request)
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

// GetServiceVersionRequest is the request struct for api GetServiceVersion
type GetServiceVersionRequest struct {
	*requests.RoaRequest
	ServiceName string `position:"Path" name:"service_name"`
	Region      string `position:"Path" name:"region"`
}

// GetServiceVersionResponse is the response struct for api GetServiceVersion
type GetServiceVersionResponse struct {
	*responses.BaseResponse
}

// CreateGetServiceVersionRequest creates a request to invoke GetServiceVersion API
func CreateGetServiceVersionRequest() (request *GetServiceVersionRequest) {
	request = &GetServiceVersionRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("eas", "2018-05-22", "GetServiceVersion", "/api/services/[region]/[service_name]/version", "", "")
	request.Method = requests.GET
	return
}

// CreateGetServiceVersionResponse creates a response to parse from GetServiceVersion response
func CreateGetServiceVersionResponse() (response *GetServiceVersionResponse) {
	response = &GetServiceVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
