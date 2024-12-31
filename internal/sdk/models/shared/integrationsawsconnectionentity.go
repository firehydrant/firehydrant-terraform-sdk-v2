// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionStatus string

const (
	ConnectionStatusPendingSetup       ConnectionStatus = "pending_setup"
	ConnectionStatusRoleAssumed        ConnectionStatus = "role_assumed"
	ConnectionStatusCantAssumeRole     ConnectionStatus = "cant_assume_role"
	ConnectionStatusInvalidPermissions ConnectionStatus = "invalid_permissions"
	ConnectionStatusValidated          ConnectionStatus = "validated"
)

func (e ConnectionStatus) ToPointer() *ConnectionStatus {
	return &e
}
func (e *ConnectionStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pending_setup":
		fallthrough
	case "role_assumed":
		fallthrough
	case "cant_assume_role":
		fallthrough
	case "invalid_permissions":
		fallthrough
	case "validated":
		*e = ConnectionStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionStatus: %v", v)
	}
}

// IntegrationsAwsConnectionEntity - Integrations_Aws_ConnectionEntity model
type IntegrationsAwsConnectionEntity struct {
	ID                *string           `json:"id,omitempty"`
	AwsAccountID      *string           `json:"aws_account_id,omitempty"`
	TargetArn         *string           `json:"target_arn,omitempty"`
	ExternalID        *string           `json:"external_id,omitempty"`
	ConnectionStatus  *ConnectionStatus `json:"connection_status,omitempty"`
	StatusText        *string           `json:"status_text,omitempty"`
	StatusDescription *string           `json:"status_description,omitempty"`
	EnvironmentID     *string           `json:"environment_id,omitempty"`
	EnvironmentName   *string           `json:"environment_name,omitempty"`
	Regions           []string          `json:"regions,omitempty"`
}

func (o *IntegrationsAwsConnectionEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *IntegrationsAwsConnectionEntity) GetAwsAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AwsAccountID
}

func (o *IntegrationsAwsConnectionEntity) GetTargetArn() *string {
	if o == nil {
		return nil
	}
	return o.TargetArn
}

func (o *IntegrationsAwsConnectionEntity) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *IntegrationsAwsConnectionEntity) GetConnectionStatus() *ConnectionStatus {
	if o == nil {
		return nil
	}
	return o.ConnectionStatus
}

func (o *IntegrationsAwsConnectionEntity) GetStatusText() *string {
	if o == nil {
		return nil
	}
	return o.StatusText
}

func (o *IntegrationsAwsConnectionEntity) GetStatusDescription() *string {
	if o == nil {
		return nil
	}
	return o.StatusDescription
}

func (o *IntegrationsAwsConnectionEntity) GetEnvironmentID() *string {
	if o == nil {
		return nil
	}
	return o.EnvironmentID
}

func (o *IntegrationsAwsConnectionEntity) GetEnvironmentName() *string {
	if o == nil {
		return nil
	}
	return o.EnvironmentName
}

func (o *IntegrationsAwsConnectionEntity) GetRegions() []string {
	if o == nil {
		return nil
	}
	return o.Regions
}
