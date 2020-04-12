package onsmqtt

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

// SendMessage invokes the onsmqtt.SendMessage API synchronously
// api document: https://help.aliyun.com/api/onsmqtt/sendmessage.html
func (client *Client) SendMessage(request *SendMessageRequest) (response *SendMessageResponse, err error) {
	response = CreateSendMessageResponse()
	err = client.DoAction(request, response)
	return
}

// SendMessageWithChan invokes the onsmqtt.SendMessage API asynchronously
// api document: https://help.aliyun.com/api/onsmqtt/sendmessage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SendMessageWithChan(request *SendMessageRequest) (<-chan *SendMessageResponse, <-chan error) {
	responseChan := make(chan *SendMessageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendMessage(request)
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

// SendMessageWithCallback invokes the onsmqtt.SendMessage API asynchronously
// api document: https://help.aliyun.com/api/onsmqtt/sendmessage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SendMessageWithCallback(request *SendMessageRequest, callback func(response *SendMessageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SendMessageResponse
		var err error
		defer close(result)
		response, err = client.SendMessage(request)
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

// SendMessageRequest is the request struct for api SendMessage
type SendMessageRequest struct {
	*requests.RpcRequest
	NoPersistFlag requests.Boolean `position:"Query" name:"NoPersistFlag"`
	MqttTopic     string           `position:"Query" name:"MqttTopic"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	Payload       string           `position:"Query" name:"Payload"`
	ReceiptId     string           `position:"Query" name:"ReceiptId"`
}

// SendMessageResponse is the response struct for api SendMessage
type SendMessageResponse struct {
	*responses.BaseResponse
	MsgId     string `json:"MsgId" xml:"MsgId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSendMessageRequest creates a request to invoke SendMessage API
func CreateSendMessageRequest() (request *SendMessageRequest) {
	request = &SendMessageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OnsMqtt", "2019-12-11", "SendMessage", "onsmqtt", "openAPI")
	return
}

// CreateSendMessageResponse creates a response to parse from SendMessage response
func CreateSendMessageResponse() (response *SendMessageResponse) {
	response = &SendMessageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
