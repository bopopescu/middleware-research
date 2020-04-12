package reid

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

// ListStore invokes the reid.ListStore API synchronously
// api document: https://help.aliyun.com/api/reid/liststore.html
func (client *Client) ListStore(request *ListStoreRequest) (response *ListStoreResponse, err error) {
	response = CreateListStoreResponse()
	err = client.DoAction(request, response)
	return
}

// ListStoreWithChan invokes the reid.ListStore API asynchronously
// api document: https://help.aliyun.com/api/reid/liststore.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListStoreWithChan(request *ListStoreRequest) (<-chan *ListStoreResponse, <-chan error) {
	responseChan := make(chan *ListStoreResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListStore(request)
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

// ListStoreWithCallback invokes the reid.ListStore API asynchronously
// api document: https://help.aliyun.com/api/reid/liststore.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListStoreWithCallback(request *ListStoreRequest, callback func(response *ListStoreResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListStoreResponse
		var err error
		defer close(result)
		response, err = client.ListStore(request)
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

// ListStoreRequest is the request struct for api ListStore
type ListStoreRequest struct {
	*requests.RpcRequest
}

// ListStoreResponse is the response struct for api ListStore
type ListStoreResponse struct {
	*responses.BaseResponse
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Stores       Stores `json:"Stores" xml:"Stores"`
}

// CreateListStoreRequest creates a request to invoke ListStore API
func CreateListStoreRequest() (request *ListStoreRequest) {
	request = &ListStoreRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("reid", "2019-09-28", "ListStore", "1.1.2", "openAPI")
	return
}

// CreateListStoreResponse creates a response to parse from ListStore response
func CreateListStoreResponse() (response *ListStoreResponse) {
	response = &ListStoreResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
