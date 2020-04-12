package slb

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

// ModifyLoadBalancerInternetSpec invokes the slb.ModifyLoadBalancerInternetSpec API synchronously
// api document: https://help.aliyun.com/api/slb/modifyloadbalancerinternetspec.html
func (client *Client) ModifyLoadBalancerInternetSpec(request *ModifyLoadBalancerInternetSpecRequest) (response *ModifyLoadBalancerInternetSpecResponse, err error) {
	response = CreateModifyLoadBalancerInternetSpecResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyLoadBalancerInternetSpecWithChan invokes the slb.ModifyLoadBalancerInternetSpec API asynchronously
// api document: https://help.aliyun.com/api/slb/modifyloadbalancerinternetspec.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyLoadBalancerInternetSpecWithChan(request *ModifyLoadBalancerInternetSpecRequest) (<-chan *ModifyLoadBalancerInternetSpecResponse, <-chan error) {
	responseChan := make(chan *ModifyLoadBalancerInternetSpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyLoadBalancerInternetSpec(request)
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

// ModifyLoadBalancerInternetSpecWithCallback invokes the slb.ModifyLoadBalancerInternetSpec API asynchronously
// api document: https://help.aliyun.com/api/slb/modifyloadbalancerinternetspec.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyLoadBalancerInternetSpecWithCallback(request *ModifyLoadBalancerInternetSpecRequest, callback func(response *ModifyLoadBalancerInternetSpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyLoadBalancerInternetSpecResponse
		var err error
		defer close(result)
		response, err = client.ModifyLoadBalancerInternetSpec(request)
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

// ModifyLoadBalancerInternetSpecRequest is the request struct for api ModifyLoadBalancerInternetSpec
type ModifyLoadBalancerInternetSpecRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AutoPay              requests.Boolean `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            requests.Integer `position:"Query" name:"Bandwidth"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tags                 string           `position:"Query" name:"Tags"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
	InternetChargeType   string           `position:"Query" name:"InternetChargeType"`
	Ratio                requests.Integer `position:"Query" name:"Ratio"`
}

// ModifyLoadBalancerInternetSpecResponse is the response struct for api ModifyLoadBalancerInternetSpec
type ModifyLoadBalancerInternetSpecResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	OrderId   int64  `json:"OrderId" xml:"OrderId"`
}

// CreateModifyLoadBalancerInternetSpecRequest creates a request to invoke ModifyLoadBalancerInternetSpec API
func CreateModifyLoadBalancerInternetSpecRequest() (request *ModifyLoadBalancerInternetSpecRequest) {
	request = &ModifyLoadBalancerInternetSpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "ModifyLoadBalancerInternetSpec", "slb", "openAPI")
	return
}

// CreateModifyLoadBalancerInternetSpecResponse creates a response to parse from ModifyLoadBalancerInternetSpec response
func CreateModifyLoadBalancerInternetSpecResponse() (response *ModifyLoadBalancerInternetSpecResponse) {
	response = &ModifyLoadBalancerInternetSpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
