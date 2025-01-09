// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// PostV1RunbooksType - Deprecated, but still required. Please just use 'incident'
type PostV1RunbooksType string

const (
	PostV1RunbooksTypeIncident       PostV1RunbooksType = "incident"
	PostV1RunbooksTypeGeneral        PostV1RunbooksType = "general"
	PostV1RunbooksTypeInfrastructure PostV1RunbooksType = "infrastructure"
	PostV1RunbooksTypeIncidentRole   PostV1RunbooksType = "incident_role"
)

func (e PostV1RunbooksType) ToPointer() *PostV1RunbooksType {
	return &e
}
func (e *PostV1RunbooksType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "incident":
		fallthrough
	case "general":
		fallthrough
	case "infrastructure":
		fallthrough
	case "incident_role":
		*e = PostV1RunbooksType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostV1RunbooksType: %v", v)
	}
}

// PostV1RunbooksOwner - An object representing a Team that owns the runbook
type PostV1RunbooksOwner struct {
	ID string `json:"id"`
}

func (o *PostV1RunbooksOwner) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type AttachmentRule struct {
	// The JSON logic for the attaching the runbook
	Logic string `json:"logic"`
	// The user data for the rule
	UserData *string `json:"user_data,omitempty"`
}

func (o *AttachmentRule) GetLogic() string {
	if o == nil {
		return ""
	}
	return o.Logic
}

func (o *AttachmentRule) GetUserData() *string {
	if o == nil {
		return nil
	}
	return o.UserData
}

type Rule struct {
	// The JSON logic for the rule
	Logic string `json:"logic"`
	// The user data for the rule
	UserData *string `json:"user_data,omitempty"`
}

func (o *Rule) GetLogic() string {
	if o == nil {
		return ""
	}
	return o.Logic
}

func (o *Rule) GetUserData() *string {
	if o == nil {
		return nil
	}
	return o.UserData
}

type PostV1RunbooksSteps struct {
	// Name for step
	Name string `json:"name"`
	// ID of action to use for this step.
	ActionID string `json:"action_id"`
	Rule     *Rule  `json:"rule,omitempty"`
}

func (o *PostV1RunbooksSteps) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PostV1RunbooksSteps) GetActionID() string {
	if o == nil {
		return ""
	}
	return o.ActionID
}

func (o *PostV1RunbooksSteps) GetRule() *Rule {
	if o == nil {
		return nil
	}
	return o.Rule
}

// PostV1Runbooks - Create a new runbook for use with incidents.
type PostV1Runbooks struct {
	Name string `json:"name"`
	// Deprecated, but still required. Please just use 'incident'
	Type PostV1RunbooksType `json:"type"`
	// Deprecated. Use description
	Summary *string `json:"summary,omitempty"`
	// A longer description about the Runbook. Supports markdown format
	Description *string `json:"description,omitempty"`
	// Whether or not this runbook should be automatically attached to restricted incidents. Note that setting this to `true` will prevent it from being attached to public incidents, even manually. Defaults to `false`.
	AutoAttachToRestrictedIncidents *bool `json:"auto_attach_to_restricted_incidents,omitempty"`
	// Whether or not this runbook is a tutorial runbook
	Tutorial *bool `json:"tutorial,omitempty"`
	// An object representing a Team that owns the runbook
	Owner          *PostV1RunbooksOwner  `json:"owner,omitempty"`
	AttachmentRule *AttachmentRule       `json:"attachment_rule,omitempty"`
	Steps          []PostV1RunbooksSteps `json:"steps,omitempty"`
}

func (o *PostV1Runbooks) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PostV1Runbooks) GetType() PostV1RunbooksType {
	if o == nil {
		return PostV1RunbooksType("")
	}
	return o.Type
}

func (o *PostV1Runbooks) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *PostV1Runbooks) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PostV1Runbooks) GetAutoAttachToRestrictedIncidents() *bool {
	if o == nil {
		return nil
	}
	return o.AutoAttachToRestrictedIncidents
}

func (o *PostV1Runbooks) GetTutorial() *bool {
	if o == nil {
		return nil
	}
	return o.Tutorial
}

func (o *PostV1Runbooks) GetOwner() *PostV1RunbooksOwner {
	if o == nil {
		return nil
	}
	return o.Owner
}

func (o *PostV1Runbooks) GetAttachmentRule() *AttachmentRule {
	if o == nil {
		return nil
	}
	return o.AttachmentRule
}

func (o *PostV1Runbooks) GetSteps() []PostV1RunbooksSteps {
	if o == nil {
		return nil
	}
	return o.Steps
}
