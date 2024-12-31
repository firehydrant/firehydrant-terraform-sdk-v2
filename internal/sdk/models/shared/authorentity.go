// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type AuthorEntity struct {
	ID     *string `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Source *string `json:"source,omitempty"`
	Email  *string `json:"email,omitempty"`
}

func (o *AuthorEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AuthorEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *AuthorEntity) GetSource() *string {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *AuthorEntity) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}
