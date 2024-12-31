// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type IncidentsConferenceBridgeEntityAttachments struct {
}

type IncidentsConferenceBridgeEntity struct {
	ID *string `json:"id,omitempty"`
	// A list of objects attached to this item. Can be one of: LinkEntity, CustomerSupportIssueEntity, or GenericAttachmentEntity
	Attachments []IncidentsConferenceBridgeEntityAttachments `json:"attachments,omitempty"`
}

func (o *IncidentsConferenceBridgeEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *IncidentsConferenceBridgeEntity) GetAttachments() []IncidentsConferenceBridgeEntityAttachments {
	if o == nil {
		return nil
	}
	return o.Attachments
}
