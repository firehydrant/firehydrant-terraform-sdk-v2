// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ServiceTier - Integer representing service tier. Lower values represent higher criticality. If not specified the default value will be 5.
type ServiceTier int

const (
	ServiceTierZero  ServiceTier = 0
	ServiceTierOne   ServiceTier = 1
	ServiceTierTwo   ServiceTier = 2
	ServiceTierThree ServiceTier = 3
	ServiceTierFour  ServiceTier = 4
	ServiceTierFive  ServiceTier = 5
)

func (e ServiceTier) ToPointer() *ServiceTier {
	return &e
}
func (e *ServiceTier) UnmarshalJSON(data []byte) error {
	var v int
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 5:
		*e = ServiceTier(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ServiceTier: %v", v)
	}
}

type Functionalities struct {
	// If you are trying to create a new functionality and attach it to this service, set the summary key
	Summary *string `json:"summary,omitempty"`
	// If you are trying to reuse a functionality, you may set the ID to attach it to the service
	ID *string `json:"id,omitempty"`
}

func (o *Functionalities) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *Functionalities) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type Links struct {
	// Short name used to display and identify this link
	Name string `json:"name"`
	// URL
	HrefURL string `json:"href_url"`
	// An optional URL to an icon representing this link
	IconURL *string `json:"icon_url,omitempty"`
}

func (o *Links) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Links) GetHrefURL() string {
	if o == nil {
		return ""
	}
	return o.HrefURL
}

func (o *Links) GetIconURL() *string {
	if o == nil {
		return nil
	}
	return o.IconURL
}

// Owner - An object representing a Team that owns the service
type Owner struct {
	ID string `json:"id"`
}

func (o *Owner) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type Teams struct {
	ID string `json:"id"`
}

func (o *Teams) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type ExternalResources struct {
	RemoteID string `json:"remote_id"`
	// The integration slug for the external resource. Can be one of: github, opsgenie, pager_duty, victorops. Not required if the resource has already been imported.
	ConnectionType *string `json:"connection_type,omitempty"`
}

func (o *ExternalResources) GetRemoteID() string {
	if o == nil {
		return ""
	}
	return o.RemoteID
}

func (o *ExternalResources) GetConnectionType() *string {
	if o == nil {
		return nil
	}
	return o.ConnectionType
}

// PostV1Services - Creates a service for the organization, you may also create or attach functionalities to the service on create.
type PostV1Services struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	// A hash of label keys and values
	Labels map[string]string `json:"labels,omitempty"`
	// Integer representing service tier. Lower values represent higher criticality. If not specified the default value will be 5.
	ServiceTier *ServiceTier `json:"service_tier,omitempty"`
	// An array of functionalities
	Functionalities []Functionalities `json:"functionalities,omitempty"`
	// An array of links to associate with this service
	Links []Links `json:"links,omitempty"`
	// An object representing a Team that owns the service
	Owner *Owner `json:"owner,omitempty"`
	// An array of teams to attach to this service.
	Teams                 []Teams `json:"teams,omitempty"`
	AlertOnAdd            *bool   `json:"alert_on_add,omitempty"`
	AutoAddRespondingTeam *bool   `json:"auto_add_responding_team,omitempty"`
	// An array of external resources to attach to this service.
	ExternalResources []ExternalResources `json:"external_resources,omitempty"`
}

func (o *PostV1Services) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PostV1Services) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PostV1Services) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *PostV1Services) GetServiceTier() *ServiceTier {
	if o == nil {
		return nil
	}
	return o.ServiceTier
}

func (o *PostV1Services) GetFunctionalities() []Functionalities {
	if o == nil {
		return nil
	}
	return o.Functionalities
}

func (o *PostV1Services) GetLinks() []Links {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *PostV1Services) GetOwner() *Owner {
	if o == nil {
		return nil
	}
	return o.Owner
}

func (o *PostV1Services) GetTeams() []Teams {
	if o == nil {
		return nil
	}
	return o.Teams
}

func (o *PostV1Services) GetAlertOnAdd() *bool {
	if o == nil {
		return nil
	}
	return o.AlertOnAdd
}

func (o *PostV1Services) GetAutoAddRespondingTeam() *bool {
	if o == nil {
		return nil
	}
	return o.AutoAddRespondingTeam
}

func (o *PostV1Services) GetExternalResources() []ExternalResources {
	if o == nil {
		return nil
	}
	return o.ExternalResources
}
