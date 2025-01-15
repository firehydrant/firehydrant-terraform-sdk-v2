// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"time"
)

// FunctionalityEntity model
type FunctionalityEntity struct {
	ID          *string    `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Slug        *string    `json:"slug,omitempty"`
	Description *string    `json:"description,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	// An object of label key and values
	Labels map[string]string `json:"labels,omitempty"`
	// List of active incident guids
	ActiveIncidents []string `json:"active_incidents,omitempty"`
	// List of links attached to this functionality.
	Links []LinksEntity `json:"links,omitempty"`
	// TeamEntity model
	Owner                 *TeamEntity   `json:"owner,omitempty"`
	AlertOnAdd            *bool         `json:"alert_on_add,omitempty"`
	AutoAddRespondingTeam *bool         `json:"auto_add_responding_team,omitempty"`
	UpdatedBy             *AuthorEntity `json:"updated_by,omitempty"`
	// Information about known linkages to representations of services outside of FireHydrant.
	ExternalResources []ExternalResourceEntity `json:"external_resources,omitempty"`
	// List of teams attached to the functionality
	Teams []TeamEntity `json:"teams,omitempty"`
}

func (f FunctionalityEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *FunctionalityEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *FunctionalityEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *FunctionalityEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *FunctionalityEntity) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *FunctionalityEntity) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *FunctionalityEntity) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *FunctionalityEntity) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *FunctionalityEntity) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *FunctionalityEntity) GetActiveIncidents() []string {
	if o == nil {
		return nil
	}
	return o.ActiveIncidents
}

func (o *FunctionalityEntity) GetLinks() []LinksEntity {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *FunctionalityEntity) GetOwner() *TeamEntity {
	if o == nil {
		return nil
	}
	return o.Owner
}

func (o *FunctionalityEntity) GetAlertOnAdd() *bool {
	if o == nil {
		return nil
	}
	return o.AlertOnAdd
}

func (o *FunctionalityEntity) GetAutoAddRespondingTeam() *bool {
	if o == nil {
		return nil
	}
	return o.AutoAddRespondingTeam
}

func (o *FunctionalityEntity) GetUpdatedBy() *AuthorEntity {
	if o == nil {
		return nil
	}
	return o.UpdatedBy
}

func (o *FunctionalityEntity) GetExternalResources() []ExternalResourceEntity {
	if o == nil {
		return nil
	}
	return o.ExternalResources
}

func (o *FunctionalityEntity) GetTeams() []TeamEntity {
	if o == nil {
		return nil
	}
	return o.Teams
}
