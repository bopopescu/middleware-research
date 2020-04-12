package imm

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

// FindImagesByTagNames invokes the imm.FindImagesByTagNames API synchronously
// api document: https://help.aliyun.com/api/imm/findimagesbytagnames.html
func (client *Client) FindImagesByTagNames(request *FindImagesByTagNamesRequest) (response *FindImagesByTagNamesResponse, err error) {
	response = CreateFindImagesByTagNamesResponse()
	err = client.DoAction(request, response)
	return
}

// FindImagesByTagNamesWithChan invokes the imm.FindImagesByTagNames API asynchronously
// api document: https://help.aliyun.com/api/imm/findimagesbytagnames.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FindImagesByTagNamesWithChan(request *FindImagesByTagNamesRequest) (<-chan *FindImagesByTagNamesResponse, <-chan error) {
	responseChan := make(chan *FindImagesByTagNamesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FindImagesByTagNames(request)
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

// FindImagesByTagNamesWithCallback invokes the imm.FindImagesByTagNames API asynchronously
// api document: https://help.aliyun.com/api/imm/findimagesbytagnames.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FindImagesByTagNamesWithCallback(request *FindImagesByTagNamesRequest, callback func(response *FindImagesByTagNamesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FindImagesByTagNamesResponse
		var err error
		defer close(result)
		response, err = client.FindImagesByTagNames(request)
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

// FindImagesByTagNamesRequest is the request struct for api FindImagesByTagNames
type FindImagesByTagNamesRequest struct {
	*requests.RpcRequest
	Project  string           `position:"Query" name:"Project"`
	Limit    requests.Integer `position:"Query" name:"Limit"`
	TagNames string           `position:"Query" name:"TagNames"`
	Marker   string           `position:"Query" name:"Marker"`
	SetId    string           `position:"Query" name:"SetId"`
}

// FindImagesByTagNamesResponse is the response struct for api FindImagesByTagNames
type FindImagesByTagNamesResponse struct {
	*responses.BaseResponse
	SetId      string   `json:"SetId" xml:"SetId"`
	NextMarker string   `json:"NextMarker" xml:"NextMarker"`
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	Images     []Frames `json:"Images" xml:"Images"`
}

// CreateFindImagesByTagNamesRequest creates a request to invoke FindImagesByTagNames API
func CreateFindImagesByTagNamesRequest() (request *FindImagesByTagNamesRequest) {
	request = &FindImagesByTagNamesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "FindImagesByTagNames", "imm", "openAPI")
	return
}

// CreateFindImagesByTagNamesResponse creates a response to parse from FindImagesByTagNames response
func CreateFindImagesByTagNamesResponse() (response *FindImagesByTagNamesResponse) {
	response = &FindImagesByTagNamesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
