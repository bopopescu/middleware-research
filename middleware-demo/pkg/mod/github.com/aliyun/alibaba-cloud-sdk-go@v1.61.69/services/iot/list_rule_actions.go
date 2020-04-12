package iot

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

// ListRuleActions invokes the iot.ListRuleActions API synchronously
// api document: https://help.aliyun.com/api/iot/listruleactions.html
func (client *Client) ListRuleActions(request *ListRuleActionsRequest) (response *ListRuleActionsResponse, err error) {
	response = CreateListRuleActionsResponse()
	err = client.DoAction(request, response)
	return
}

// ListRuleActionsWithChan invokes the iot.ListRuleActions API asynchronously
// api document: https://help.aliyun.com/api/iot/listruleactions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRuleActionsWithChan(request *ListRuleActionsRequest) (<-chan *ListRuleActionsResponse, <-chan error) {
	responseChan := make(chan *ListRuleActionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRuleActions(request)
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

// ListRuleActionsWithCallback invokes the iot.ListRuleActions API asynchronously
// api document: https://help.aliyun.com/api/iot/listruleactions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRuleActionsWithCallback(request *ListRuleActionsRequest, callback func(response *ListRuleActionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRuleActionsResponse
		var err error
		defer close(result)
		response, err = client.ListRuleActions(request)
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

// ListRuleActionsRequest is the request struct for api ListRuleActions
type ListRuleActionsRequest struct {
	*requests.RpcRequest
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	RuleId        requests.Integer `position:"Query" name:"RuleId"`
}

// ListRuleActionsResponse is the response struct for api ListRuleActions
type ListRuleActionsResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	Success        bool           `json:"Success" xml:"Success"`
	Code           string         `json:"Code" xml:"Code"`
	ErrorMessage   string         `json:"ErrorMessage" xml:"ErrorMessage"`
	RuleActionList RuleActionList `json:"RuleActionList" xml:"RuleActionList"`
}

// CreateListRuleActionsRequest creates a request to invoke ListRuleActions API
func CreateListRuleActionsRequest() (request *ListRuleActionsRequest) {
	request = &ListRuleActionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "ListRuleActions", "Iot", "openAPI")
	return
}

// CreateListRuleActionsResponse creates a response to parse from ListRuleActions response
func CreateListRuleActionsResponse() (response *ListRuleActionsResponse) {
	response = &ListRuleActionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
