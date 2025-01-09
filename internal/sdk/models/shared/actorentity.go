// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Type string

const (
	TypeFirehydrantUser  Type = "firehydrant_user"
	TypeFirehydrantBot   Type = "firehydrant_bot"
	TypeAlertmanager     Type = "alertmanager"
	TypeAsana            Type = "asana"
	TypeAws              Type = "aws"
	TypeBugsnag          Type = "bugsnag"
	TypeCheckly          Type = "checkly"
	TypeCustomAlerts     Type = "custom_alerts"
	TypeDatadog          Type = "datadog"
	TypeShortcut         Type = "shortcut"
	TypeNewRelic         Type = "new_relic"
	TypeNunc             Type = "nunc"
	TypeGithub           Type = "github"
	TypeGiphy            Type = "giphy"
	TypeGoogleMeet       Type = "google_meet"
	TypeGoogleCalendar   Type = "google_calendar"
	TypeMicrosoftTeams   Type = "microsoft_teams"
	TypeMicrosoftTeamsV2 Type = "microsoft_teams_v2"
	TypeWebex            Type = "webex"
	TypeJiraCloud        Type = "jira_cloud"
	TypeJiraOnprem       Type = "jira_onprem"
	TypeOpsgenie         Type = "opsgenie"
	TypePagerDuty        Type = "pager_duty"
	TypeHoneycomb        Type = "honeycomb"
	TypePatchy           Type = "patchy"
	TypeServiceNow       Type = "service_now"
	TypeSignals          Type = "signals"
	TypeSlack            Type = "slack"
	TypeStatuspage       Type = "statuspage"
	TypeVictorops        Type = "victorops"
	TypeZendesk          Type = "zendesk"
	TypeZoom             Type = "zoom"
	TypeConfluenceCloud  Type = "confluence_cloud"
	TypeGoogleDocs       Type = "google_docs"
	TypeZoomV2           Type = "zoom_v2"
	TypeLinear           Type = "linear"
	TypeCortex           Type = "cortex"
)

func (e Type) ToPointer() *Type {
	return &e
}
func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "firehydrant_user":
		fallthrough
	case "firehydrant_bot":
		fallthrough
	case "alertmanager":
		fallthrough
	case "asana":
		fallthrough
	case "aws":
		fallthrough
	case "bugsnag":
		fallthrough
	case "checkly":
		fallthrough
	case "custom_alerts":
		fallthrough
	case "datadog":
		fallthrough
	case "shortcut":
		fallthrough
	case "new_relic":
		fallthrough
	case "nunc":
		fallthrough
	case "github":
		fallthrough
	case "giphy":
		fallthrough
	case "google_meet":
		fallthrough
	case "google_calendar":
		fallthrough
	case "microsoft_teams":
		fallthrough
	case "microsoft_teams_v2":
		fallthrough
	case "webex":
		fallthrough
	case "jira_cloud":
		fallthrough
	case "jira_onprem":
		fallthrough
	case "opsgenie":
		fallthrough
	case "pager_duty":
		fallthrough
	case "honeycomb":
		fallthrough
	case "patchy":
		fallthrough
	case "service_now":
		fallthrough
	case "signals":
		fallthrough
	case "slack":
		fallthrough
	case "statuspage":
		fallthrough
	case "victorops":
		fallthrough
	case "zendesk":
		fallthrough
	case "zoom":
		fallthrough
	case "confluence_cloud":
		fallthrough
	case "google_docs":
		fallthrough
	case "zoom_v2":
		fallthrough
	case "linear":
		fallthrough
	case "cortex":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

type ActorEntity struct {
	ID    *string `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
	Type  *Type   `json:"type,omitempty"`
}

func (o *ActorEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ActorEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ActorEntity) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *ActorEntity) GetType() *Type {
	if o == nil {
		return nil
	}
	return o.Type
}
