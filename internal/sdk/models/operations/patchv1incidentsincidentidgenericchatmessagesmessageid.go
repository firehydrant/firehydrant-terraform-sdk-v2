// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

type PatchV1IncidentsIncidentIDGenericChatMessagesMessageIDRequest struct {
	MessageID                                              string                                                        `pathParam:"style=simple,explode=false,name=message_id"`
	IncidentID                                             string                                                        `pathParam:"style=simple,explode=false,name=incident_id"`
	PatchV1IncidentsIncidentIDGenericChatMessagesMessageID shared.PatchV1IncidentsIncidentIDGenericChatMessagesMessageID `request:"mediaType=application/json"`
}

func (o *PatchV1IncidentsIncidentIDGenericChatMessagesMessageIDRequest) GetMessageID() string {
	if o == nil {
		return ""
	}
	return o.MessageID
}

func (o *PatchV1IncidentsIncidentIDGenericChatMessagesMessageIDRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *PatchV1IncidentsIncidentIDGenericChatMessagesMessageIDRequest) GetPatchV1IncidentsIncidentIDGenericChatMessagesMessageID() shared.PatchV1IncidentsIncidentIDGenericChatMessagesMessageID {
	if o == nil {
		return shared.PatchV1IncidentsIncidentIDGenericChatMessagesMessageID{}
	}
	return o.PatchV1IncidentsIncidentIDGenericChatMessagesMessageID
}

type PatchV1IncidentsIncidentIDGenericChatMessagesMessageIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Update an existing generic chat message on an incident.
	EventGenericChatMessageEntity *shared.EventGenericChatMessageEntity
}

func (o *PatchV1IncidentsIncidentIDGenericChatMessagesMessageIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchV1IncidentsIncidentIDGenericChatMessagesMessageIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchV1IncidentsIncidentIDGenericChatMessagesMessageIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchV1IncidentsIncidentIDGenericChatMessagesMessageIDResponse) GetEventGenericChatMessageEntity() *shared.EventGenericChatMessageEntity {
	if o == nil {
		return nil
	}
	return o.EventGenericChatMessageEntity
}
