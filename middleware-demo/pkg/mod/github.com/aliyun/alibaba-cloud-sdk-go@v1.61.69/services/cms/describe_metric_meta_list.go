package cms

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

// DescribeMetricMetaList invokes the cms.DescribeMetricMetaList API synchronously
// api document: https://help.aliyun.com/api/cms/describemetricmetalist.html
func (client *Client) DescribeMetricMetaList(request *DescribeMetricMetaListRequest) (response *DescribeMetricMetaListResponse, err error) {
	response = CreateDescribeMetricMetaListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMetricMetaListWithChan invokes the cms.DescribeMetricMetaList API asynchronously
// api document: https://help.aliyun.com/api/cms/describemetricmetalist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMetricMetaListWithChan(request *DescribeMetricMetaListRequest) (<-chan *DescribeMetricMetaListResponse, <-chan error) {
	responseChan := make(chan *DescribeMetricMetaListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMetricMetaList(request)
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

// DescribeMetricMetaListWithCallback invokes the cms.DescribeMetricMetaList API asynchronously
// api document: https://help.aliyun.com/api/cms/describemetricmetalist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMetricMetaListWithCallback(request *DescribeMetricMetaListRequest, callback func(response *DescribeMetricMetaListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMetricMetaListResponse
		var err error
		defer close(result)
		response, err = client.DescribeMetricMetaList(request)
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

// DescribeMetricMetaListRequest is the request struct for api DescribeMetricMetaList
type DescribeMetricMetaListRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	Labels     string           `position:"Query" name:"Labels"`
	Namespace  string           `position:"Query" name:"Namespace"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	MetricName string           `position:"Query" name:"MetricName"`
}

// DescribeMetricMetaListResponse is the response struct for api DescribeMetricMetaList
type DescribeMetricMetaListResponse struct {
	*responses.BaseResponse
	RequestId  string                            `json:"RequestId" xml:"RequestId"`
	Success    bool                              `json:"Success" xml:"Success"`
	Code       string                            `json:"Code" xml:"Code"`
	Message    string                            `json:"Message" xml:"Message"`
	TotalCount string                            `json:"TotalCount" xml:"TotalCount"`
	Resources  ResourcesInDescribeMetricMetaList `json:"Resources" xml:"Resources"`
}

// CreateDescribeMetricMetaListRequest creates a request to invoke DescribeMetricMetaList API
func CreateDescribeMetricMetaListRequest() (request *DescribeMetricMetaListRequest) {
	request = &DescribeMetricMetaListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeMetricMetaList", "cms", "openAPI")
	return
}

// CreateDescribeMetricMetaListResponse creates a response to parse from DescribeMetricMetaList response
func CreateDescribeMetricMetaListResponse() (response *DescribeMetricMetaListResponse) {
	response = &DescribeMetricMetaListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
