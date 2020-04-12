package ccc

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

// CreateInstance invokes the ccc.CreateInstance API synchronously
// api document: https://help.aliyun.com/api/ccc/createinstance.html
func (client *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
	response = CreateCreateInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateInstanceWithChan invokes the ccc.CreateInstance API asynchronously
// api document: https://help.aliyun.com/api/ccc/createinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateInstanceWithChan(request *CreateInstanceRequest) (<-chan *CreateInstanceResponse, <-chan error) {
	responseChan := make(chan *CreateInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateInstance(request)
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

// CreateInstanceWithCallback invokes the ccc.CreateInstance API asynchronously
// api document: https://help.aliyun.com/api/ccc/createinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateInstanceWithCallback(request *CreateInstanceRequest, callback func(response *CreateInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateInstanceResponse
		var err error
		defer close(result)
		response, err = client.CreateInstance(request)
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

// CreateInstanceRequest is the request struct for api CreateInstance
type CreateInstanceRequest struct {
	*requests.RpcRequest
	PhoneNumbers   *[]string        `position:"Query" name:"PhoneNumbers"  type:"Repeated"`
	UserObject     *[]string        `position:"Query" name:"UserObject"  type:"Repeated"`
	DomainName     string           `position:"Query" name:"DomainName"`
	PhoneNumber    string           `position:"Query" name:"PhoneNumber"`
	Description    string           `position:"Query" name:"Description"`
	StorageMaxDays requests.Integer `position:"Query" name:"StorageMaxDays"`
	AdminRamId     *[]string        `position:"Query" name:"AdminRamId"  type:"Repeated"`
	Name           string           `position:"Query" name:"Name"`
	StorageMaxSize requests.Integer `position:"Query" name:"StorageMaxSize"`
	DirectoryId    string           `position:"Query" name:"DirectoryId"`
}

// CreateInstanceResponse is the response struct for api CreateInstance
type CreateInstanceResponse struct {
	*responses.BaseResponse
	RequestId      string                   `json:"RequestId" xml:"RequestId"`
	Success        bool                     `json:"Success" xml:"Success"`
	Code           string                   `json:"Code" xml:"Code"`
	Message        string                   `json:"Message" xml:"Message"`
	HttpStatusCode int                      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Instance       InstanceInCreateInstance `json:"Instance" xml:"Instance"`
}

// CreateCreateInstanceRequest creates a request to invoke CreateInstance API
func CreateCreateInstanceRequest() (request *CreateInstanceRequest) {
	request = &CreateInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "CreateInstance", "", "")
	return
}

// CreateCreateInstanceResponse creates a response to parse from CreateInstance response
func CreateCreateInstanceResponse() (response *CreateInstanceResponse) {
	response = &CreateInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
