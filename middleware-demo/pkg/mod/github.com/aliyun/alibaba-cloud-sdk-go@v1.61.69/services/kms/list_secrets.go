package kms

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

// ListSecrets invokes the kms.ListSecrets API synchronously
// api document: https://help.aliyun.com/api/kms/listsecrets.html
func (client *Client) ListSecrets(request *ListSecretsRequest) (response *ListSecretsResponse, err error) {
	response = CreateListSecretsResponse()
	err = client.DoAction(request, response)
	return
}

// ListSecretsWithChan invokes the kms.ListSecrets API asynchronously
// api document: https://help.aliyun.com/api/kms/listsecrets.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListSecretsWithChan(request *ListSecretsRequest) (<-chan *ListSecretsResponse, <-chan error) {
	responseChan := make(chan *ListSecretsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSecrets(request)
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

// ListSecretsWithCallback invokes the kms.ListSecrets API asynchronously
// api document: https://help.aliyun.com/api/kms/listsecrets.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListSecretsWithCallback(request *ListSecretsRequest, callback func(response *ListSecretsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSecretsResponse
		var err error
		defer close(result)
		response, err = client.ListSecrets(request)
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

// ListSecretsRequest is the request struct for api ListSecrets
type ListSecretsRequest struct {
	*requests.RpcRequest
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	FetchTags  string           `position:"Query" name:"FetchTags"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// ListSecretsResponse is the response struct for api ListSecrets
type ListSecretsResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	PageNumber int        `json:"PageNumber" xml:"PageNumber"`
	PageSize   int        `json:"PageSize" xml:"PageSize"`
	TotalCount int        `json:"TotalCount" xml:"TotalCount"`
	SecretList SecretList `json:"SecretList" xml:"SecretList"`
}

// CreateListSecretsRequest creates a request to invoke ListSecrets API
func CreateListSecretsRequest() (request *ListSecretsRequest) {
	request = &ListSecretsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "ListSecrets", "kms", "openAPI")
	return
}

// CreateListSecretsResponse creates a response to parse from ListSecrets response
func CreateListSecretsResponse() (response *ListSecretsResponse) {
	response = &ListSecretsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
