package baas

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

// DescribeFabricOrganizationChaincodes invokes the baas.DescribeFabricOrganizationChaincodes API synchronously
// api document: https://help.aliyun.com/api/baas/describefabricorganizationchaincodes.html
func (client *Client) DescribeFabricOrganizationChaincodes(request *DescribeFabricOrganizationChaincodesRequest) (response *DescribeFabricOrganizationChaincodesResponse, err error) {
	response = CreateDescribeFabricOrganizationChaincodesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFabricOrganizationChaincodesWithChan invokes the baas.DescribeFabricOrganizationChaincodes API asynchronously
// api document: https://help.aliyun.com/api/baas/describefabricorganizationchaincodes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFabricOrganizationChaincodesWithChan(request *DescribeFabricOrganizationChaincodesRequest) (<-chan *DescribeFabricOrganizationChaincodesResponse, <-chan error) {
	responseChan := make(chan *DescribeFabricOrganizationChaincodesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFabricOrganizationChaincodes(request)
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

// DescribeFabricOrganizationChaincodesWithCallback invokes the baas.DescribeFabricOrganizationChaincodes API asynchronously
// api document: https://help.aliyun.com/api/baas/describefabricorganizationchaincodes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFabricOrganizationChaincodesWithCallback(request *DescribeFabricOrganizationChaincodesRequest, callback func(response *DescribeFabricOrganizationChaincodesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFabricOrganizationChaincodesResponse
		var err error
		defer close(result)
		response, err = client.DescribeFabricOrganizationChaincodes(request)
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

// DescribeFabricOrganizationChaincodesRequest is the request struct for api DescribeFabricOrganizationChaincodes
type DescribeFabricOrganizationChaincodesRequest struct {
	*requests.RpcRequest
	OrganizationId string `position:"Body" name:"OrganizationId"`
	Location       string `position:"Body" name:"Location"`
}

// DescribeFabricOrganizationChaincodesResponse is the response struct for api DescribeFabricOrganizationChaincodes
type DescribeFabricOrganizationChaincodesResponse struct {
	*responses.BaseResponse
	RequestId string                  `json:"RequestId" xml:"RequestId"`
	Success   bool                    `json:"Success" xml:"Success"`
	ErrorCode int                     `json:"ErrorCode" xml:"ErrorCode"`
	Result    []OrganizationChaincode `json:"Result" xml:"Result"`
}

// CreateDescribeFabricOrganizationChaincodesRequest creates a request to invoke DescribeFabricOrganizationChaincodes API
func CreateDescribeFabricOrganizationChaincodesRequest() (request *DescribeFabricOrganizationChaincodesRequest) {
	request = &DescribeFabricOrganizationChaincodesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "DescribeFabricOrganizationChaincodes", "baas", "openAPI")
	return
}

// CreateDescribeFabricOrganizationChaincodesResponse creates a response to parse from DescribeFabricOrganizationChaincodes response
func CreateDescribeFabricOrganizationChaincodesResponse() (response *DescribeFabricOrganizationChaincodesResponse) {
	response = &DescribeFabricOrganizationChaincodesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
