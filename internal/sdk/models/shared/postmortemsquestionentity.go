// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type PostMortemsQuestionEntity struct {
	ID               *string                             `json:"id,omitempty"`
	Title            *string                             `json:"title,omitempty"`
	Body             *string                             `json:"body,omitempty"`
	Tooltip          *string                             `json:"tooltip,omitempty"`
	Kind             *string                             `json:"kind,omitempty"`
	QuestionTypeID   *string                             `json:"question_type_id,omitempty"`
	IsRequired       *bool                               `json:"is_required,omitempty"`
	AvailableOptions []string                            `json:"available_options,omitempty"`
	Conversations    []ConversationsAPIEntitiesReference `json:"conversations,omitempty"`
}

func (o *PostMortemsQuestionEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PostMortemsQuestionEntity) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *PostMortemsQuestionEntity) GetBody() *string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *PostMortemsQuestionEntity) GetTooltip() *string {
	if o == nil {
		return nil
	}
	return o.Tooltip
}

func (o *PostMortemsQuestionEntity) GetKind() *string {
	if o == nil {
		return nil
	}
	return o.Kind
}

func (o *PostMortemsQuestionEntity) GetQuestionTypeID() *string {
	if o == nil {
		return nil
	}
	return o.QuestionTypeID
}

func (o *PostMortemsQuestionEntity) GetIsRequired() *bool {
	if o == nil {
		return nil
	}
	return o.IsRequired
}

func (o *PostMortemsQuestionEntity) GetAvailableOptions() []string {
	if o == nil {
		return nil
	}
	return o.AvailableOptions
}

func (o *PostMortemsQuestionEntity) GetConversations() []ConversationsAPIEntitiesReference {
	if o == nil {
		return nil
	}
	return o.Conversations
}
