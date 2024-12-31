// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// PostV1Environments - Creates an environment for the organization
type PostV1Environments struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

func (o *PostV1Environments) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PostV1Environments) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}
