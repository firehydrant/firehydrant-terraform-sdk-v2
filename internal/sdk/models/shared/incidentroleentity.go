// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/internal/utils"
	"time"
)

// IncidentRoleEntity model
type IncidentRoleEntity struct {
	ID          *string    `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Summary     *string    `json:"summary,omitempty"`
	Description *string    `json:"description,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DiscardedAt *time.Time `json:"discarded_at,omitempty"`
}

func (i IncidentRoleEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *IncidentRoleEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *IncidentRoleEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *IncidentRoleEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *IncidentRoleEntity) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *IncidentRoleEntity) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *IncidentRoleEntity) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *IncidentRoleEntity) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *IncidentRoleEntity) GetDiscardedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DiscardedAt
}
