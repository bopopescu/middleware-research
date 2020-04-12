package qualitycheck

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

// UploadData invokes the qualitycheck.UploadData API synchronously
// api document: https://help.aliyun.com/api/qualitycheck/uploaddata.html
func (client *Client) UploadData(request *UploadDataRequest) (response *UploadDataResponse, err error) {
	response = CreateUploadDataResponse()
	err = client.DoAction(request, response)
	return
}

// UploadDataWithChan invokes the qualitycheck.UploadData API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/uploaddata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UploadDataWithChan(request *UploadDataRequest) (<-chan *UploadDataResponse, <-chan error) {
	responseChan := make(chan *UploadDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UploadData(request)
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

// UploadDataWithCallback invokes the qualitycheck.UploadData API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/uploaddata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UploadDataWithCallback(request *UploadDataRequest, callback func(response *UploadDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UploadDataResponse
		var err error
		defer close(result)
		response, err = client.UploadData(request)
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

// UploadDataRequest is the request struct for api UploadData
type UploadDataRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// UploadDataResponse is the response struct for api UploadData
type UploadDataResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateUploadDataRequest creates a request to invoke UploadData API
func CreateUploadDataRequest() (request *UploadDataRequest) {
	request = &UploadDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "UploadData", "", "")
	return
}

// CreateUploadDataResponse creates a response to parse from UploadData response
func CreateUploadDataResponse() (response *UploadDataResponse) {
	response = &UploadDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
