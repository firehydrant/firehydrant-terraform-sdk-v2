// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type IncidentAttachmentEntityStatus string

const (
	IncidentAttachmentEntityStatusPendingUpload IncidentAttachmentEntityStatus = "pending_upload"
	IncidentAttachmentEntityStatusUploaded      IncidentAttachmentEntityStatus = "uploaded"
)

func (e IncidentAttachmentEntityStatus) ToPointer() *IncidentAttachmentEntityStatus {
	return &e
}
func (e *IncidentAttachmentEntityStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pending_upload":
		fallthrough
	case "uploaded":
		*e = IncidentAttachmentEntityStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IncidentAttachmentEntityStatus: %v", v)
	}
}

// Versions - An object with keys that designate a specific version or size of the attachment
type Versions struct {
}

// IncidentAttachmentEntity model
type IncidentAttachmentEntity struct {
	FileName        *string                         `json:"file_name,omitempty"`
	FileContentType *string                         `json:"file_content_type,omitempty"`
	SignedURL       *string                         `json:"signed_url,omitempty"`
	MediaType       *string                         `json:"media_type,omitempty"`
	Description     *string                         `json:"description,omitempty"`
	ExternalID      *string                         `json:"external_id,omitempty"`
	FileSize        *int                            `json:"file_size,omitempty"`
	Status          *IncidentAttachmentEntityStatus `json:"status,omitempty"`
	// An object with keys that designate a specific version or size of the attachment
	Versions *Versions `json:"versions,omitempty"`
}

func (o *IncidentAttachmentEntity) GetFileName() *string {
	if o == nil {
		return nil
	}
	return o.FileName
}

func (o *IncidentAttachmentEntity) GetFileContentType() *string {
	if o == nil {
		return nil
	}
	return o.FileContentType
}

func (o *IncidentAttachmentEntity) GetSignedURL() *string {
	if o == nil {
		return nil
	}
	return o.SignedURL
}

func (o *IncidentAttachmentEntity) GetMediaType() *string {
	if o == nil {
		return nil
	}
	return o.MediaType
}

func (o *IncidentAttachmentEntity) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *IncidentAttachmentEntity) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *IncidentAttachmentEntity) GetFileSize() *int {
	if o == nil {
		return nil
	}
	return o.FileSize
}

func (o *IncidentAttachmentEntity) GetStatus() *IncidentAttachmentEntityStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *IncidentAttachmentEntity) GetVersions() *Versions {
	if o == nil {
		return nil
	}
	return o.Versions
}
