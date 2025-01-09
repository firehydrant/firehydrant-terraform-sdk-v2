// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/internal/utils"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
	"time"
)

type File struct {
	FileName string `multipartForm:"name=file"`
	Content  []byte `multipartForm:"content"`
}

func (o *File) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

func (o *File) GetContent() []byte {
	if o == nil {
		return []byte{}
	}
	return o.Content
}

type VoteDirection string

const (
	VoteDirectionUp   VoteDirection = "up"
	VoteDirectionDown VoteDirection = "down"
)

func (e VoteDirection) ToPointer() *VoteDirection {
	return &e
}
func (e *VoteDirection) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "up":
		fallthrough
	case "down":
		*e = VoteDirection(v)
		return nil
	default:
		return fmt.Errorf("invalid value for VoteDirection: %v", v)
	}
}

type PostV1IncidentsIncidentIDAttachmentsRequestBody struct {
	File          File           `multipartForm:"file"`
	Description   *string        `multipartForm:"name=description"`
	OccurredAt    *time.Time     `multipartForm:"name=occurred_at"`
	VoteDirection *VoteDirection `multipartForm:"name=vote_direction"`
}

func (p PostV1IncidentsIncidentIDAttachmentsRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PostV1IncidentsIncidentIDAttachmentsRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PostV1IncidentsIncidentIDAttachmentsRequestBody) GetFile() File {
	if o == nil {
		return File{}
	}
	return o.File
}

func (o *PostV1IncidentsIncidentIDAttachmentsRequestBody) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PostV1IncidentsIncidentIDAttachmentsRequestBody) GetOccurredAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.OccurredAt
}

func (o *PostV1IncidentsIncidentIDAttachmentsRequestBody) GetVoteDirection() *VoteDirection {
	if o == nil {
		return nil
	}
	return o.VoteDirection
}

type PostV1IncidentsIncidentIDAttachmentsRequest struct {
	IncidentID  string                                          `pathParam:"style=simple,explode=false,name=incident_id"`
	RequestBody PostV1IncidentsIncidentIDAttachmentsRequestBody `request:"mediaType=multipart/form-data"`
}

func (o *PostV1IncidentsIncidentIDAttachmentsRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *PostV1IncidentsIncidentIDAttachmentsRequest) GetRequestBody() PostV1IncidentsIncidentIDAttachmentsRequestBody {
	if o == nil {
		return PostV1IncidentsIncidentIDAttachmentsRequestBody{}
	}
	return o.RequestBody
}

type PostV1IncidentsIncidentIDAttachmentsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Allows adding image attachments to an incident
	IncidentAttachmentEntity *shared.IncidentAttachmentEntity
}

func (o *PostV1IncidentsIncidentIDAttachmentsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostV1IncidentsIncidentIDAttachmentsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostV1IncidentsIncidentIDAttachmentsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostV1IncidentsIncidentIDAttachmentsResponse) GetIncidentAttachmentEntity() *shared.IncidentAttachmentEntity {
	if o == nil {
		return nil
	}
	return o.IncidentAttachmentEntity
}