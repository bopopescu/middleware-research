package iqa

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

// DescribeProject invokes the iqa.DescribeProject API synchronously
// api document: https://help.aliyun.com/api/iqa/describeproject.html
func (client *Client) DescribeProject(request *DescribeProjectRequest) (response *DescribeProjectResponse, err error) {
	response = CreateDescribeProjectResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeProjectWithChan invokes the iqa.DescribeProject API asynchronously
// api document: https://help.aliyun.com/api/iqa/describeproject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeProjectWithChan(request *DescribeProjectRequest) (<-chan *DescribeProjectResponse, <-chan error) {
	responseChan := make(chan *DescribeProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeProject(request)
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

// DescribeProjectWithCallback invokes the iqa.DescribeProject API asynchronously
// api document: https://help.aliyun.com/api/iqa/describeproject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeProjectWithCallback(request *DescribeProjectRequest, callback func(response *DescribeProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeProjectResponse
		var err error
		defer close(result)
		response, err = client.DescribeProject(request)
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

// DescribeProjectRequest is the request struct for api DescribeProject
type DescribeProjectRequest struct {
	*requests.RpcRequest
	ProjectId string `position:"Query" name:"ProjectId"`
}

// DescribeProjectResponse is the response struct for api DescribeProject
type DescribeProjectResponse struct {
	*responses.BaseResponse
	RequestId           string `json:"RequestId" xml:"RequestId"`
	ProjectType         string `json:"ProjectType" xml:"ProjectType"`
	ProjectId           string `json:"ProjectId" xml:"ProjectId"`
	ProjectName         string `json:"ProjectName" xml:"ProjectName"`
	CreateTime          int64  `json:"CreateTime" xml:"CreateTime"`
	DeployTime          int64  `json:"DeployTime" xml:"DeployTime"`
	ModelId             string `json:"ModelId" xml:"ModelId"`
	ModelName           string `json:"ModelName" xml:"ModelName"`
	QuestionCount       int    `json:"QuestionCount" xml:"QuestionCount"`
	DataStatus          string `json:"DataStatus" xml:"DataStatus"`
	TestServiceStatus   string `json:"TestServiceStatus" xml:"TestServiceStatus"`
	OnlineServiceStatus string `json:"OnlineServiceStatus" xml:"OnlineServiceStatus"`
	DeployAvailable     string `json:"DeployAvailable" xml:"DeployAvailable"`
}

// CreateDescribeProjectRequest creates a request to invoke DescribeProject API
func CreateDescribeProjectRequest() (request *DescribeProjectRequest) {
	request = &DescribeProjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("iqa", "2019-08-13", "DescribeProject", "iqa", "openAPI")
	return
}

// CreateDescribeProjectResponse creates a response to parse from DescribeProject response
func CreateDescribeProjectResponse() (response *DescribeProjectResponse) {
	response = &DescribeProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
