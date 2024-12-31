// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SignalsAPITargetEntity struct {
	ID         *string `json:"id,omitempty"`
	Name       *string `json:"name,omitempty"`
	Type       *string `json:"type,omitempty"`
	IsPageable *bool   `json:"is_pageable,omitempty"`
}

func (o *SignalsAPITargetEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SignalsAPITargetEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *SignalsAPITargetEntity) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *SignalsAPITargetEntity) GetIsPageable() *bool {
	if o == nil {
		return nil
	}
	return o.IsPageable
}
