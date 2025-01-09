// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsRequest struct {
	// Slack Connection UUID
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	Page         *int   `queryParam:"style=form,explode=true,name=page"`
	PerPage      *int   `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
