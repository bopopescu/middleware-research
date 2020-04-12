package ecs

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

// StopInstance invokes the ecs.StopInstance API synchronously
// api document: https://help.aliyun.com/api/ecs/stopinstance.html
func (client *Client) StopInstance(request *StopInstanceRequest) (response *StopInstanceResponse, err error) {
	response = CreateStopInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// StopInstanceWithChan invokes the ecs.StopInstance API asynchronously
// api document: https://help.aliyun.com/api/ecs/stopinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StopInstanceWithChan(request *StopInstanceRequest) (<-chan *StopInstanceResponse, <-chan error) {
	responseChan := make(chan *StopInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopInstance(request)
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

// StopInstanceWithCallback invokes the ecs.StopInstance API asynchronously
// api document: https://help.aliyun.com/api/ecs/stopinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StopInstanceWithCallback(request *StopInstanceRequest, callback func(response *StopInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopInstanceResponse
		var err error
		defer close(result)
		response, err = client.StopInstance(request)
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

// StopInstanceRequest is the request struct for api StopInstance
type StopInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StoppedMode          string           `position:"Query" name:"StoppedMode"`
	Hibernate            requests.Boolean `position:"Query" name:"Hibernate"`
	ForceStop            requests.Boolean `position:"Query" name:"ForceStop"`
	ConfirmStop          requests.Boolean `position:"Query" name:"ConfirmStop"`
	DryRun               requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// StopInstanceResponse is the response struct for api StopInstance
type StopInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopInstanceRequest creates a request to invoke StopInstance API
func CreateStopInstanceRequest() (request *StopInstanceRequest) {
	request = &StopInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "StopInstance", "ecs", "openAPI")
	return
}

// CreateStopInstanceResponse creates a response to parse from StopInstance response
func CreateStopInstanceResponse() (response *StopInstanceResponse) {
	response = &StopInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
