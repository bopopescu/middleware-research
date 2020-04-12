package ens

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

// AttachEnsInstances invokes the ens.AttachEnsInstances API synchronously
// api document: https://help.aliyun.com/api/ens/attachensinstances.html
func (client *Client) AttachEnsInstances(request *AttachEnsInstancesRequest) (response *AttachEnsInstancesResponse, err error) {
	response = CreateAttachEnsInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// AttachEnsInstancesWithChan invokes the ens.AttachEnsInstances API asynchronously
// api document: https://help.aliyun.com/api/ens/attachensinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AttachEnsInstancesWithChan(request *AttachEnsInstancesRequest) (<-chan *AttachEnsInstancesResponse, <-chan error) {
	responseChan := make(chan *AttachEnsInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AttachEnsInstances(request)
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

// AttachEnsInstancesWithCallback invokes the ens.AttachEnsInstances API asynchronously
// api document: https://help.aliyun.com/api/ens/attachensinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AttachEnsInstancesWithCallback(request *AttachEnsInstancesRequest, callback func(response *AttachEnsInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AttachEnsInstancesResponse
		var err error
		defer close(result)
		response, err = client.AttachEnsInstances(request)
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

// AttachEnsInstancesRequest is the request struct for api AttachEnsInstances
type AttachEnsInstancesRequest struct {
	*requests.RpcRequest
	Scripts    string `position:"Query" name:"Scripts"`
	Version    string `position:"Query" name:"Version"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// AttachEnsInstancesResponse is the response struct for api AttachEnsInstances
type AttachEnsInstancesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAttachEnsInstancesRequest creates a request to invoke AttachEnsInstances API
func CreateAttachEnsInstancesRequest() (request *AttachEnsInstancesRequest) {
	request = &AttachEnsInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "AttachEnsInstances", "ens", "openAPI")
	return
}

// CreateAttachEnsInstancesResponse creates a response to parse from AttachEnsInstances response
func CreateAttachEnsInstancesResponse() (response *AttachEnsInstancesResponse) {
	response = &AttachEnsInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
