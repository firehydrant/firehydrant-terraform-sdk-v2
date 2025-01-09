// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type TeamEntity struct {
	CreatedAt            types.String                               `tfsdk:"created_at"`
	CreatedBy            *AuthorEntity                              `tfsdk:"created_by"`
	Description          types.String                               `tfsdk:"description"`
	Functionalities      []FunctionalityEntity                      `tfsdk:"functionalities"`
	ID                   types.String                               `tfsdk:"id"`
	Memberships          []MembershipEntity                         `tfsdk:"memberships"`
	MsTeamsChannel       *IntegrationsMicrosoftTeamsV2ChannelEntity `tfsdk:"ms_teams_channel"`
	Name                 types.String                               `tfsdk:"name"`
	OwnedFunctionalities []FunctionalityEntity                      `tfsdk:"owned_functionalities"`
	OwnedRunbooks        []SlimRunbookEntity                        `tfsdk:"owned_runbooks"`
	SignalsIcalURL       types.String                               `tfsdk:"signals_ical_url"`
	SlackChannel         *IntegrationsSlackSlackChannelEntity       `tfsdk:"slack_channel"`
	Slug                 types.String                               `tfsdk:"slug"`
	UpdatedAt            types.String                               `tfsdk:"updated_at"`
}