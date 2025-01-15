// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/internal/utils"
	"time"
)

type SignalsAPIRuleEntityNotificationPriorityOverride string

const (
	SignalsAPIRuleEntityNotificationPriorityOverrideHigh   SignalsAPIRuleEntityNotificationPriorityOverride = "HIGH"
	SignalsAPIRuleEntityNotificationPriorityOverrideMedium SignalsAPIRuleEntityNotificationPriorityOverride = "MEDIUM"
	SignalsAPIRuleEntityNotificationPriorityOverrideLow    SignalsAPIRuleEntityNotificationPriorityOverride = "LOW"
)

func (e SignalsAPIRuleEntityNotificationPriorityOverride) ToPointer() *SignalsAPIRuleEntityNotificationPriorityOverride {
	return &e
}
func (e *SignalsAPIRuleEntityNotificationPriorityOverride) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HIGH":
		fallthrough
	case "MEDIUM":
		fallthrough
	case "LOW":
		*e = SignalsAPIRuleEntityNotificationPriorityOverride(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SignalsAPIRuleEntityNotificationPriorityOverride: %v", v)
	}
}

type SignalsAPIRuleEntity struct {
	ID                           *string                                           `json:"id,omitempty"`
	Name                         *string                                           `json:"name,omitempty"`
	Expression                   *string                                           `json:"expression,omitempty"`
	TeamID                       *string                                           `json:"team_id,omitempty"`
	Target                       *SignalsAPITargetEntity                           `json:"target,omitempty"`
	CreatedBy                    *AuthorEntity                                     `json:"created_by,omitempty"`
	CreatedAt                    *time.Time                                        `json:"created_at,omitempty"`
	UpdatedAt                    *time.Time                                        `json:"updated_at,omitempty"`
	IncidentType                 *SuccinctEntity                                   `json:"incident_type,omitempty"`
	NotificationPriorityOverride *SignalsAPIRuleEntityNotificationPriorityOverride `json:"notification_priority_override,omitempty"`
}

func (s SignalsAPIRuleEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SignalsAPIRuleEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SignalsAPIRuleEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SignalsAPIRuleEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *SignalsAPIRuleEntity) GetExpression() *string {
	if o == nil {
		return nil
	}
	return o.Expression
}

func (o *SignalsAPIRuleEntity) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *SignalsAPIRuleEntity) GetTarget() *SignalsAPITargetEntity {
	if o == nil {
		return nil
	}
	return o.Target
}

func (o *SignalsAPIRuleEntity) GetCreatedBy() *AuthorEntity {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *SignalsAPIRuleEntity) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *SignalsAPIRuleEntity) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *SignalsAPIRuleEntity) GetIncidentType() *SuccinctEntity {
	if o == nil {
		return nil
	}
	return o.IncidentType
}

func (o *SignalsAPIRuleEntity) GetNotificationPriorityOverride() *SignalsAPIRuleEntityNotificationPriorityOverride {
	if o == nil {
		return nil
	}
	return o.NotificationPriorityOverride
}
