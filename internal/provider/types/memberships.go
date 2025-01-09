// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Memberships struct {
	DefaultIncidentRole *IncidentRoleEntity `tfsdk:"default_incident_role"`
	IncidentRoleID      types.String        `tfsdk:"incident_role_id"`
	Schedule            *ScheduleEntity     `tfsdk:"schedule"`
	ScheduleID          types.String        `tfsdk:"schedule_id"`
	User                *UserEntity         `tfsdk:"user"`
	UserID              types.String        `tfsdk:"user_id"`
}
