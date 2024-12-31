// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// PostMortemsReasonEntity - PostMortems_ReasonEntity model
type PostMortemsReasonEntity struct {
	ID            *string                             `json:"id,omitempty"`
	Summary       *string                             `json:"summary,omitempty"`
	Position      *int                                `json:"position,omitempty"`
	CreatedBy     *AuthorEntity                       `json:"created_by,omitempty"`
	Conversations []ConversationsAPIEntitiesReference `json:"conversations,omitempty"`
}

func (o *PostMortemsReasonEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PostMortemsReasonEntity) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *PostMortemsReasonEntity) GetPosition() *int {
	if o == nil {
		return nil
	}
	return o.Position
}

func (o *PostMortemsReasonEntity) GetCreatedBy() *AuthorEntity {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *PostMortemsReasonEntity) GetConversations() []ConversationsAPIEntitiesReference {
	if o == nil {
		return nil
	}
	return o.Conversations
}
