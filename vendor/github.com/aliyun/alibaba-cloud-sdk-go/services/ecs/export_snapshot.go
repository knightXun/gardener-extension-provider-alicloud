package ecs

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

// ExportSnapshot invokes the ecs.ExportSnapshot API synchronously
// api document: https://help.aliyun.com/api/ecs/exportsnapshot.html
func (client *Client) ExportSnapshot(request *ExportSnapshotRequest) (response *ExportSnapshotResponse, err error) {
	response = CreateExportSnapshotResponse()
	err = client.DoAction(request, response)
	return
}

// ExportSnapshotWithChan invokes the ecs.ExportSnapshot API asynchronously
// api document: https://help.aliyun.com/api/ecs/exportsnapshot.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ExportSnapshotWithChan(request *ExportSnapshotRequest) (<-chan *ExportSnapshotResponse, <-chan error) {
	responseChan := make(chan *ExportSnapshotResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExportSnapshot(request)
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

// ExportSnapshotWithCallback invokes the ecs.ExportSnapshot API asynchronously
// api document: https://help.aliyun.com/api/ecs/exportsnapshot.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ExportSnapshotWithCallback(request *ExportSnapshotRequest, callback func(response *ExportSnapshotResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExportSnapshotResponse
		var err error
		defer close(result)
		response, err = client.ExportSnapshot(request)
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

// ExportSnapshotRequest is the request struct for api ExportSnapshot
type ExportSnapshotRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SnapshotId           string           `position:"Query" name:"SnapshotId"`
	OssBucket            string           `position:"Query" name:"OssBucket"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	RoleName             string           `position:"Query" name:"RoleName"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ExportSnapshotResponse is the response struct for api ExportSnapshot
type ExportSnapshotResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
}

// CreateExportSnapshotRequest creates a request to invoke ExportSnapshot API
func CreateExportSnapshotRequest() (request *ExportSnapshotRequest) {
	request = &ExportSnapshotRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ExportSnapshot", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateExportSnapshotResponse creates a response to parse from ExportSnapshot response
func CreateExportSnapshotResponse() (response *ExportSnapshotResponse) {
	response = &ExportSnapshotResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
