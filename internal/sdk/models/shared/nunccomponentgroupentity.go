// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type NuncComponentGroupEntity struct {
	ID               *string `json:"id,omitempty"`
	ComponentGroupID *string `json:"component_group_id,omitempty"`
	Name             *string `json:"name,omitempty"`
	Position         *int    `json:"position,omitempty"`
}

func (o *NuncComponentGroupEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *NuncComponentGroupEntity) GetComponentGroupID() *string {
	if o == nil {
		return nil
	}
	return o.ComponentGroupID
}

func (o *NuncComponentGroupEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *NuncComponentGroupEntity) GetPosition() *int {
	if o == nil {
		return nil
	}
	return o.Position
}
