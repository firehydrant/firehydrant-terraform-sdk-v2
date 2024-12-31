// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/internal/utils"
	"time"
)

// PostV1IncidentsLabels - Key:value pairs to track custom data for the incident
type PostV1IncidentsLabels struct {
}

type Impacts struct {
	// The type of impacted infrastructure. One of: environment, functionality, or service
	Type string `json:"type"`
	// The ID of the impacted infrastructure
	ID string `json:"id"`
	// The ID of the impact condition. Find these at /v1/severity_matrix/conditions
	ConditionID string `json:"condition_id"`
}

func (o *Impacts) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *Impacts) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Impacts) GetConditionID() string {
	if o == nil {
		return ""
	}
	return o.ConditionID
}

type Milestones struct {
	// The type/slug of the milestone. Must be one of the currently configured milestones.
	Type string `json:"type"`
	// An ISO8601 formatted string for when this milestone occurred. If you are removing a milestone, this field still needs to be set to some value
	OccurredAt time.Time `json:"occurred_at"`
}

func (m Milestones) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *Milestones) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Milestones) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *Milestones) GetOccurredAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.OccurredAt
}

type CustomFields struct {
	// The ID of the custom field you wish to set.
	FieldID string `json:"field_id"`
	// The value you wish to set on the custom field if the type of the field accepts string values
	ValueString *string `json:"value_string,omitempty"`
	// The value you wish to set on the custom field if the type of the field accepts array values
	ValueArray []string `json:"value_array,omitempty"`
}

func (o *CustomFields) GetFieldID() string {
	if o == nil {
		return ""
	}
	return o.FieldID
}

func (o *CustomFields) GetValueString() *string {
	if o == nil {
		return nil
	}
	return o.ValueString
}

func (o *CustomFields) GetValueArray() []string {
	if o == nil {
		return nil
	}
	return o.ValueArray
}

// PostV1Incidents - Create a new incident
type PostV1Incidents struct {
	Name                  string  `json:"name"`
	Summary               *string `json:"summary,omitempty"`
	CustomerImpactSummary *string `json:"customer_impact_summary,omitempty"`
	Description           *string `json:"description,omitempty"`
	Priority              *string `json:"priority,omitempty"`
	Severity              *string `json:"severity,omitempty"`
	SeverityConditionID   *string `json:"severity_condition_id,omitempty"`
	SeverityImpactID      *string `json:"severity_impact_id,omitempty"`
	// List of alert IDs that this incident should be associated to
	AlertIds []string `json:"alert_ids,omitempty"`
	// Key:value pairs to track custom data for the incident
	Labels *PostV1IncidentsLabels `json:"labels,omitempty"`
	// List of ids of Runbooks to attach to this incident. Foregoes any conditions these Runbooks may have guarding automatic attachment.
	RunbookIds []string `json:"runbook_ids,omitempty"`
	// List of tags for the incident
	TagList []string `json:"tag_list,omitempty"`
	// An array of impacted infrastructure
	Impacts []Impacts `json:"impacts,omitempty"`
	// An array of milestones to set on an incident. This can be used to create an already-resolved incident.
	Milestones []Milestones `json:"milestones,omitempty"`
	Restricted *bool        `json:"restricted,omitempty"`
	// IDs of teams you wish to assign to this incident.
	TeamIds []string `json:"team_ids,omitempty"`
	// An array of custom fields to set on the incident.
	CustomFields  []CustomFields `json:"custom_fields,omitempty"`
	ExternalLinks *string        `json:"external_links,omitempty"`
	// If true, the incident type values will not be copied to the incident. This is useful when creating an incident from an incident type, but you want to set the values manually.
	SkipIncidentTypeValues *bool `default:"false" json:"skip_incident_type_values"`
}

func (p PostV1Incidents) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PostV1Incidents) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PostV1Incidents) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PostV1Incidents) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *PostV1Incidents) GetCustomerImpactSummary() *string {
	if o == nil {
		return nil
	}
	return o.CustomerImpactSummary
}

func (o *PostV1Incidents) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PostV1Incidents) GetPriority() *string {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *PostV1Incidents) GetSeverity() *string {
	if o == nil {
		return nil
	}
	return o.Severity
}

func (o *PostV1Incidents) GetSeverityConditionID() *string {
	if o == nil {
		return nil
	}
	return o.SeverityConditionID
}

func (o *PostV1Incidents) GetSeverityImpactID() *string {
	if o == nil {
		return nil
	}
	return o.SeverityImpactID
}

func (o *PostV1Incidents) GetAlertIds() []string {
	if o == nil {
		return nil
	}
	return o.AlertIds
}

func (o *PostV1Incidents) GetLabels() *PostV1IncidentsLabels {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *PostV1Incidents) GetRunbookIds() []string {
	if o == nil {
		return nil
	}
	return o.RunbookIds
}

func (o *PostV1Incidents) GetTagList() []string {
	if o == nil {
		return nil
	}
	return o.TagList
}

func (o *PostV1Incidents) GetImpacts() []Impacts {
	if o == nil {
		return nil
	}
	return o.Impacts
}

func (o *PostV1Incidents) GetMilestones() []Milestones {
	if o == nil {
		return nil
	}
	return o.Milestones
}

func (o *PostV1Incidents) GetRestricted() *bool {
	if o == nil {
		return nil
	}
	return o.Restricted
}

func (o *PostV1Incidents) GetTeamIds() []string {
	if o == nil {
		return nil
	}
	return o.TeamIds
}

func (o *PostV1Incidents) GetCustomFields() []CustomFields {
	if o == nil {
		return nil
	}
	return o.CustomFields
}

func (o *PostV1Incidents) GetExternalLinks() *string {
	if o == nil {
		return nil
	}
	return o.ExternalLinks
}

func (o *PostV1Incidents) GetSkipIncidentTypeValues() *bool {
	if o == nil {
		return nil
	}
	return o.SkipIncidentTypeValues
}
