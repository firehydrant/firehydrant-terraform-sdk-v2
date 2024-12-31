// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// MsTeamsChannel - MS Teams channel identity for channel associated with this team
type MsTeamsChannel struct {
	ChannelID string `json:"channel_id"`
	MsTeamID  string `json:"ms_team_id"`
}

func (o *MsTeamsChannel) GetChannelID() string {
	if o == nil {
		return ""
	}
	return o.ChannelID
}

func (o *MsTeamsChannel) GetMsTeamID() string {
	if o == nil {
		return ""
	}
	return o.MsTeamID
}

type Memberships struct {
	UserID     *string `json:"user_id,omitempty"`
	ScheduleID *string `json:"schedule_id,omitempty"`
	// An incident role ID that this user will automatically assigned if this team is assigned to an incident
	IncidentRoleID *string `json:"incident_role_id,omitempty"`
}

func (o *Memberships) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

func (o *Memberships) GetScheduleID() *string {
	if o == nil {
		return nil
	}
	return o.ScheduleID
}

func (o *Memberships) GetIncidentRoleID() *string {
	if o == nil {
		return nil
	}
	return o.IncidentRoleID
}

// PostV1Teams - Create a new team
type PostV1Teams struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Slug        *string `json:"slug,omitempty"`
	// The Slack channel ID that this team is associated with
	SlackChannelID *string `json:"slack_channel_id,omitempty"`
	// MS Teams channel identity for channel associated with this team
	MsTeamsChannel *MsTeamsChannel `json:"ms_teams_channel,omitempty"`
	Memberships    []Memberships   `json:"memberships,omitempty"`
}

func (o *PostV1Teams) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PostV1Teams) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PostV1Teams) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *PostV1Teams) GetSlackChannelID() *string {
	if o == nil {
		return nil
	}
	return o.SlackChannelID
}

func (o *PostV1Teams) GetMsTeamsChannel() *MsTeamsChannel {
	if o == nil {
		return nil
	}
	return o.MsTeamsChannel
}

func (o *PostV1Teams) GetMemberships() []Memberships {
	if o == nil {
		return nil
	}
	return o.Memberships
}
