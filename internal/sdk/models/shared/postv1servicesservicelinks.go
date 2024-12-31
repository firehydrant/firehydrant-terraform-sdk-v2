// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// Integration - The name of the service
type Integration string

const (
	IntegrationPagerDuty Integration = "pager_duty"
	IntegrationOpsgenie  Integration = "opsgenie"
	IntegrationVictorops Integration = "victorops"
)

func (e Integration) ToPointer() *Integration {
	return &e
}
func (e *Integration) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pager_duty":
		fallthrough
	case "opsgenie":
		fallthrough
	case "victorops":
		*e = Integration(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Integration: %v", v)
	}
}

// PostV1ServicesServiceLinks - Creates a service with the appropriate integration for each external service ID passed
type PostV1ServicesServiceLinks struct {
	// ID of the service
	ExternalServiceIds string `json:"external_service_ids"`
	// ID for the integration. This can be found by going to the edit page for the integration
	ConnectionID string `json:"connection_id"`
	// The name of the service
	Integration Integration `json:"integration"`
}

func (o *PostV1ServicesServiceLinks) GetExternalServiceIds() string {
	if o == nil {
		return ""
	}
	return o.ExternalServiceIds
}

func (o *PostV1ServicesServiceLinks) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PostV1ServicesServiceLinks) GetIntegration() Integration {
	if o == nil {
		return Integration("")
	}
	return o.Integration
}
