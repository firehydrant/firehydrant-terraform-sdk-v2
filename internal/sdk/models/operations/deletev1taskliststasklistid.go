// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type DeleteV1TaskListsTaskListIDRequest struct {
	TaskListID string `pathParam:"style=simple,explode=false,name=task_list_id"`
}

func (o *DeleteV1TaskListsTaskListIDRequest) GetTaskListID() string {
	if o == nil {
		return ""
	}
	return o.TaskListID
}

type DeleteV1TaskListsTaskListIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Delete a task list
	TaskListEntity *shared.TaskListEntity
}

func (o *DeleteV1TaskListsTaskListIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteV1TaskListsTaskListIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteV1TaskListsTaskListIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteV1TaskListsTaskListIDResponse) GetTaskListEntity() *shared.TaskListEntity {
	if o == nil {
		return nil
	}
	return o.TaskListEntity
}