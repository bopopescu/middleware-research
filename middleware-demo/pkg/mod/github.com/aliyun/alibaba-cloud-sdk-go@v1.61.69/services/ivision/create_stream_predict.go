package ivision

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

// CreateStreamPredict invokes the ivision.CreateStreamPredict API synchronously
// api document: https://help.aliyun.com/api/ivision/createstreampredict.html
func (client *Client) CreateStreamPredict(request *CreateStreamPredictRequest) (response *CreateStreamPredictResponse, err error) {
	response = CreateCreateStreamPredictResponse()
	err = client.DoAction(request, response)
	return
}

// CreateStreamPredictWithChan invokes the ivision.CreateStreamPredict API asynchronously
// api document: https://help.aliyun.com/api/ivision/createstreampredict.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateStreamPredictWithChan(request *CreateStreamPredictRequest) (<-chan *CreateStreamPredictResponse, <-chan error) {
	responseChan := make(chan *CreateStreamPredictResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateStreamPredict(request)
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

// CreateStreamPredictWithCallback invokes the ivision.CreateStreamPredict API asynchronously
// api document: https://help.aliyun.com/api/ivision/createstreampredict.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateStreamPredictWithCallback(request *CreateStreamPredictRequest, callback func(response *CreateStreamPredictResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateStreamPredictResponse
		var err error
		defer close(result)
		response, err = client.CreateStreamPredict(request)
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

// CreateStreamPredictRequest is the request struct for api CreateStreamPredict
type CreateStreamPredictRequest struct {
	*requests.RpcRequest
	ClientToken           string           `position:"Query" name:"ClientToken"`
	AutoStart             string           `position:"Query" name:"AutoStart"`
	Notify                string           `position:"Query" name:"Notify"`
	Output                string           `position:"Query" name:"Output"`
	ShowLog               string           `position:"Query" name:"ShowLog"`
	StreamType            string           `position:"Query" name:"StreamType"`
	FaceGroupId           string           `position:"Query" name:"FaceGroupId"`
	StreamId              string           `position:"Query" name:"StreamId"`
	DetectIntervals       string           `position:"Query" name:"DetectIntervals"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	ProbabilityThresholds string           `position:"Query" name:"ProbabilityThresholds"`
	ModelIds              string           `position:"Query" name:"ModelIds"`
	ModelUserData         string           `position:"Query" name:"ModelUserData"`
}

// CreateStreamPredictResponse is the response struct for api CreateStreamPredict
type CreateStreamPredictResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	PredictId string `json:"PredictId" xml:"PredictId"`
}

// CreateCreateStreamPredictRequest creates a request to invoke CreateStreamPredict API
func CreateCreateStreamPredictRequest() (request *CreateStreamPredictRequest) {
	request = &CreateStreamPredictRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ivision", "2019-03-08", "CreateStreamPredict", "ivision", "openAPI")
	return
}

// CreateCreateStreamPredictResponse creates a response to parse from CreateStreamPredict response
func CreateCreateStreamPredictResponse() (response *CreateStreamPredictResponse) {
	response = &CreateStreamPredictResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
