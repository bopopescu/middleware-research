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

// ModifyDCdnDomainSchdmByProperty invokes the dcdn.ModifyDCdnDomainSchdmByProperty API synchronously
// api document: https://help.aliyun.com/api/dcdn/modifydcdndomainschdmbyproperty.html
func (client *Client) ModifyDCdnDomainSchdmByProperty(request *ModifyDCdnDomainSchdmByPropertyRequest) (response *ModifyDCdnDomainSchdmByPropertyResponse, err error) {
	response = CreateModifyDCdnDomainSchdmByPropertyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDCdnDomainSchdmByPropertyWithChan invokes the dcdn.ModifyDCdnDomainSchdmByProperty API asynchronously
// api document: https://help.aliyun.com/api/dcdn/modifydcdndomainschdmbyproperty.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDCdnDomainSchdmByPropertyWithChan(request *ModifyDCdnDomainSchdmByPropertyRequest) (<-chan *ModifyDCdnDomainSchdmByPropertyResponse, <-chan error) {
	responseChan := make(chan *ModifyDCdnDomainSchdmByPropertyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDCdnDomainSchdmByProperty(request)
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

// ModifyDCdnDomainSchdmByPropertyWithCallback invokes the dcdn.ModifyDCdnDomainSchdmByProperty API asynchronously
// api document: https://help.aliyun.com/api/dcdn/modifydcdndomainschdmbyproperty.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDCdnDomainSchdmByPropertyWithCallback(request *ModifyDCdnDomainSchdmByPropertyRequest, callback func(response *ModifyDCdnDomainSchdmByPropertyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDCdnDomainSchdmByPropertyResponse
		var err error
		defer close(result)
		response, err = client.ModifyDCdnDomainSchdmByProperty(request)
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

// ModifyDCdnDomainSchdmByPropertyRequest is the request struct for api ModifyDCdnDomainSchdmByProperty
type ModifyDCdnDomainSchdmByPropertyRequest struct {
	*requests.RpcRequest
	Property   string           `position:"Query" name:"Property"`
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyDCdnDomainSchdmByPropertyResponse is the response struct for api ModifyDCdnDomainSchdmByProperty
type ModifyDCdnDomainSchdmByPropertyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDCdnDomainSchdmByPropertyRequest creates a request to invoke ModifyDCdnDomainSchdmByProperty API
func CreateModifyDCdnDomainSchdmByPropertyRequest() (request *ModifyDCdnDomainSchdmByPropertyRequest) {
	request = &ModifyDCdnDomainSchdmByPropertyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "ModifyDCdnDomainSchdmByProperty", "", "")
	return
}

// CreateModifyDCdnDomainSchdmByPropertyResponse creates a response to parse from ModifyDCdnDomainSchdmByProperty response
func CreateModifyDCdnDomainSchdmByPropertyResponse() (response *ModifyDCdnDomainSchdmByPropertyResponse) {
	response = &ModifyDCdnDomainSchdmByPropertyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
