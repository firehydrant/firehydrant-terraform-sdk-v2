// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type IncidentTypeEntityTemplateEntity struct {
	CustomFields          types.String                             `tfsdk:"custom_fields"`
	CustomerImpactSummary types.String                             `tfsdk:"customer_impact_summary"`
	Description           types.String                             `tfsdk:"description"`
	Impacts               []IncidentTypeEntityTemplateImpactEntity `tfsdk:"impacts"`
	IncidentName          types.String                             `tfsdk:"incident_name"`
	Labels                map[string]types.String                  `tfsdk:"labels"`
	Priority              types.String                             `tfsdk:"priority"`
	PrivateIncident       types.Bool                               `tfsdk:"private_incident"`
	RunbookIds            []types.String                           `tfsdk:"runbook_ids"`
	Severity              types.String                             `tfsdk:"severity"`
	Summary               types.String                             `tfsdk:"summary"`
	TagList               []types.String                           `tfsdk:"tag_list"`
	TeamIds               []types.String                           `tfsdk:"team_ids"`
}