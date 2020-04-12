package cloudapi

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

// DescribeAppSecurity invokes the cloudapi.DescribeAppSecurity API synchronously
// api document: https://help.aliyun.com/api/cloudapi/describeappsecurity.html
func (client *Client) DescribeAppSecurity(request *DescribeAppSecurityRequest) (response *DescribeAppSecurityResponse, err error) {
	response = CreateDescribeAppSecurityResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAppSecurityWithChan invokes the cloudapi.DescribeAppSecurity API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeappsecurity.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAppSecurityWithChan(request *DescribeAppSecurityRequest) (<-chan *DescribeAppSecurityResponse, <-chan error) {
	responseChan := make(chan *DescribeAppSecurityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAppSecurity(request)
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

// DescribeAppSecurityWithCallback invokes the cloudapi.DescribeAppSecurity API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeappsecurity.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAppSecurityWithCallback(request *DescribeAppSecurityRequest, callback func(response *DescribeAppSecurityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAppSecurityResponse
		var err error
		defer close(result)
		response, err = client.DescribeAppSecurity(request)
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

// DescribeAppSecurityRequest is the request struct for api DescribeAppSecurity
type DescribeAppSecurityRequest struct {
	*requests.RpcRequest
	SecurityToken string                    `position:"Query" name:"SecurityToken"`
	AppId         requests.Integer          `position:"Query" name:"AppId"`
	Tag           *[]DescribeAppSecurityTag `position:"Query" name:"Tag"  type:"Repeated"`
}

// DescribeAppSecurityTag is a repeated param struct in DescribeAppSecurityRequest
type DescribeAppSecurityTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// DescribeAppSecurityResponse is the response struct for api DescribeAppSecurity
type DescribeAppSecurityResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	AppKey       string `json:"AppKey" xml:"AppKey"`
	AppSecret    string `json:"AppSecret" xml:"AppSecret"`
	CreatedTime  string `json:"CreatedTime" xml:"CreatedTime"`
	ModifiedTime string `json:"ModifiedTime" xml:"ModifiedTime"`
	AppCode      string `json:"AppCode" xml:"AppCode"`
}

// CreateDescribeAppSecurityRequest creates a request to invoke DescribeAppSecurity API
func CreateDescribeAppSecurityRequest() (request *DescribeAppSecurityRequest) {
	request = &DescribeAppSecurityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeAppSecurity", "apigateway", "openAPI")
	return
}

// CreateDescribeAppSecurityResponse creates a response to parse from DescribeAppSecurity response
func CreateDescribeAppSecurityResponse() (response *DescribeAppSecurityResponse) {
	response = &DescribeAppSecurityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
