// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/internal/utils"
)

// BadRequest - A collection of codes that generally means the end user got something wrong in making the request
type BadRequest struct {
	Message              *string `json:"message,omitempty"`
	AdditionalProperties any     `additionalProperties:"true" json:"-"`
}

func (b BadRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BadRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *BadRequest) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *BadRequest) GetAdditionalProperties() any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}
