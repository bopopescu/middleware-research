package webplus

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

// DescribePkgVersions invokes the webplus.DescribePkgVersions API synchronously
// api document: https://help.aliyun.com/api/webplus/describepkgversions.html
func (client *Client) DescribePkgVersions(request *DescribePkgVersionsRequest) (response *DescribePkgVersionsResponse, err error) {
	response = CreateDescribePkgVersionsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePkgVersionsWithChan invokes the webplus.DescribePkgVersions API asynchronously
// api document: https://help.aliyun.com/api/webplus/describepkgversions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePkgVersionsWithChan(request *DescribePkgVersionsRequest) (<-chan *DescribePkgVersionsResponse, <-chan error) {
	responseChan := make(chan *DescribePkgVersionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePkgVersions(request)
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

// DescribePkgVersionsWithCallback invokes the webplus.DescribePkgVersions API asynchronously
// api document: https://help.aliyun.com/api/webplus/describepkgversions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePkgVersionsWithCallback(request *DescribePkgVersionsRequest, callback func(response *DescribePkgVersionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePkgVersionsResponse
		var err error
		defer close(result)
		response, err = client.DescribePkgVersions(request)
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

// DescribePkgVersionsRequest is the request struct for api DescribePkgVersions
type DescribePkgVersionsRequest struct {
	*requests.RoaRequest
	PkgVersionLabel  string           `position:"Query" name:"PkgVersionLabel"`
	AppId            string           `position:"Query" name:"AppId"`
	PageSize         requests.Integer `position:"Query" name:"PageSize"`
	PkgVersionSearch string           `position:"Query" name:"PkgVersionSearch"`
	PageNumber       requests.Integer `position:"Query" name:"PageNumber"`
}

// DescribePkgVersionsResponse is the response struct for api DescribePkgVersions
type DescribePkgVersionsResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	Code        string      `json:"Code" xml:"Code"`
	Message     string      `json:"Message" xml:"Message"`
	PageNumber  int         `json:"PageNumber" xml:"PageNumber"`
	PageSize    int         `json:"PageSize" xml:"PageSize"`
	TotalCount  int         `json:"TotalCount" xml:"TotalCount"`
	PkgVersions PkgVersions `json:"PkgVersions" xml:"PkgVersions"`
}

// CreateDescribePkgVersionsRequest creates a request to invoke DescribePkgVersions API
func CreateDescribePkgVersionsRequest() (request *DescribePkgVersionsRequest) {
	request = &DescribePkgVersionsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("WebPlus", "2019-03-20", "DescribePkgVersions", "/pop/v1/wam/pkgVersion", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribePkgVersionsResponse creates a response to parse from DescribePkgVersions response
func CreateDescribePkgVersionsResponse() (response *DescribePkgVersionsResponse) {
	response = &DescribePkgVersionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
