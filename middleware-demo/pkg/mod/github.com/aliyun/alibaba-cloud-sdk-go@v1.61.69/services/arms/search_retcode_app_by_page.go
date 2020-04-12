package arms

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

// SearchRetcodeAppByPage invokes the arms.SearchRetcodeAppByPage API synchronously
// api document: https://help.aliyun.com/api/arms/searchretcodeappbypage.html
func (client *Client) SearchRetcodeAppByPage(request *SearchRetcodeAppByPageRequest) (response *SearchRetcodeAppByPageResponse, err error) {
	response = CreateSearchRetcodeAppByPageResponse()
	err = client.DoAction(request, response)
	return
}

// SearchRetcodeAppByPageWithChan invokes the arms.SearchRetcodeAppByPage API asynchronously
// api document: https://help.aliyun.com/api/arms/searchretcodeappbypage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SearchRetcodeAppByPageWithChan(request *SearchRetcodeAppByPageRequest) (<-chan *SearchRetcodeAppByPageResponse, <-chan error) {
	responseChan := make(chan *SearchRetcodeAppByPageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchRetcodeAppByPage(request)
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

// SearchRetcodeAppByPageWithCallback invokes the arms.SearchRetcodeAppByPage API asynchronously
// api document: https://help.aliyun.com/api/arms/searchretcodeappbypage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SearchRetcodeAppByPageWithCallback(request *SearchRetcodeAppByPageRequest, callback func(response *SearchRetcodeAppByPageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchRetcodeAppByPageResponse
		var err error
		defer close(result)
		response, err = client.SearchRetcodeAppByPage(request)
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

// SearchRetcodeAppByPageRequest is the request struct for api SearchRetcodeAppByPage
type SearchRetcodeAppByPageRequest struct {
	*requests.RpcRequest
	RetcodeAppName string           `position:"Query" name:"RetcodeAppName"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
}

// SearchRetcodeAppByPageResponse is the response struct for api SearchRetcodeAppByPage
type SearchRetcodeAppByPageResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	PageBean  PageBean `json:"PageBean" xml:"PageBean"`
}

// CreateSearchRetcodeAppByPageRequest creates a request to invoke SearchRetcodeAppByPage API
func CreateSearchRetcodeAppByPageRequest() (request *SearchRetcodeAppByPageRequest) {
	request = &SearchRetcodeAppByPageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "SearchRetcodeAppByPage", "arms", "openAPI")
	return
}

// CreateSearchRetcodeAppByPageResponse creates a response to parse from SearchRetcodeAppByPage response
func CreateSearchRetcodeAppByPageResponse() (response *SearchRetcodeAppByPageResponse) {
	response = &SearchRetcodeAppByPageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
