// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"time"
)

type PostV1ScheduledMaintenancesStatusPages struct {
	// The slug identifying the type of status page
	IntegrationSlug *string `json:"integration_slug,omitempty"`
	// The UUID of the status page to display this maintenance on
	ConnectionID string `json:"connection_id"`
}

func (o *PostV1ScheduledMaintenancesStatusPages) GetIntegrationSlug() *string {
	if o == nil {
		return nil
	}
	return o.IntegrationSlug
}

func (o *PostV1ScheduledMaintenancesStatusPages) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type PostV1ScheduledMaintenancesImpacts struct {
	// The type of impact
	Type string `json:"type"`
	// The id of impact
	ID string `json:"id"`
	// The id of the condition
	ConditionID string `json:"condition_id"`
}

func (o *PostV1ScheduledMaintenancesImpacts) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *PostV1ScheduledMaintenancesImpacts) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PostV1ScheduledMaintenancesImpacts) GetConditionID() string {
	if o == nil {
		return ""
	}
	return o.ConditionID
}

// PostV1ScheduledMaintenances - Create a new scheduled maintenance event
type PostV1ScheduledMaintenances struct {
	Name string `json:"name"`
	// ISO8601 timestamp for the start time of the scheduled maintenance
	StartsAt time.Time `json:"starts_at"`
	// ISO8601 timestamp for the end time of the scheduled maintenance
	EndsAt      time.Time `json:"ends_at"`
	Summary     *string   `json:"summary,omitempty"`
	Description *string   `json:"description,omitempty"`
	// A json object of label keys and values
	Labels map[string]string `json:"labels,omitempty"`
	// An array of status pages to display this maintenance on
	StatusPages []PostV1ScheduledMaintenancesStatusPages `json:"status_pages,omitempty"`
	// An array of impact/condition combinations
	Impacts []PostV1ScheduledMaintenancesImpacts `json:"impacts,omitempty"`
}

func (p PostV1ScheduledMaintenances) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PostV1ScheduledMaintenances) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PostV1ScheduledMaintenances) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PostV1ScheduledMaintenances) GetStartsAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartsAt
}

func (o *PostV1ScheduledMaintenances) GetEndsAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.EndsAt
}

func (o *PostV1ScheduledMaintenances) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *PostV1ScheduledMaintenances) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PostV1ScheduledMaintenances) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *PostV1ScheduledMaintenances) GetStatusPages() []PostV1ScheduledMaintenancesStatusPages {
	if o == nil {
		return nil
	}
	return o.StatusPages
}

func (o *PostV1ScheduledMaintenances) GetImpacts() []PostV1ScheduledMaintenancesImpacts {
	if o == nil {
		return nil
	}
	return o.Impacts
}
