package dcdn

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

// DescribeDcdnCertificateDetail invokes the dcdn.DescribeDcdnCertificateDetail API synchronously
// api document: https://help.aliyun.com/api/dcdn/describedcdncertificatedetail.html
func (client *Client) DescribeDcdnCertificateDetail(request *DescribeDcdnCertificateDetailRequest) (response *DescribeDcdnCertificateDetailResponse, err error) {
	response = CreateDescribeDcdnCertificateDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnCertificateDetailWithChan invokes the dcdn.DescribeDcdnCertificateDetail API asynchronously
// api document: https://help.aliyun.com/api/dcdn/describedcdncertificatedetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDcdnCertificateDetailWithChan(request *DescribeDcdnCertificateDetailRequest) (<-chan *DescribeDcdnCertificateDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnCertificateDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnCertificateDetail(request)
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

// DescribeDcdnCertificateDetailWithCallback invokes the dcdn.DescribeDcdnCertificateDetail API asynchronously
// api document: https://help.aliyun.com/api/dcdn/describedcdncertificatedetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDcdnCertificateDetailWithCallback(request *DescribeDcdnCertificateDetailRequest, callback func(response *DescribeDcdnCertificateDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnCertificateDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnCertificateDetail(request)
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

// DescribeDcdnCertificateDetailRequest is the request struct for api DescribeDcdnCertificateDetail
type DescribeDcdnCertificateDetailRequest struct {
	*requests.RpcRequest
	CertName      string           `position:"Query" name:"CertName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// DescribeDcdnCertificateDetailResponse is the response struct for api DescribeDcdnCertificateDetail
type DescribeDcdnCertificateDetailResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Cert      string `json:"Cert" xml:"Cert"`
	Key       string `json:"Key" xml:"Key"`
	CertId    int64  `json:"CertId" xml:"CertId"`
	CertName  string `json:"CertName" xml:"CertName"`
}

// CreateDescribeDcdnCertificateDetailRequest creates a request to invoke DescribeDcdnCertificateDetail API
func CreateDescribeDcdnCertificateDetailRequest() (request *DescribeDcdnCertificateDetailRequest) {
	request = &DescribeDcdnCertificateDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnCertificateDetail", "", "")
	return
}

// CreateDescribeDcdnCertificateDetailResponse creates a response to parse from DescribeDcdnCertificateDetail response
func CreateDescribeDcdnCertificateDetailResponse() (response *DescribeDcdnCertificateDetailResponse) {
	response = &DescribeDcdnCertificateDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
