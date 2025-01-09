// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type DeleteV1IncidentsIncidentIDImpactTypeIDPathParamType string

const (
	DeleteV1IncidentsIncidentIDImpactTypeIDPathParamTypeEnvironments    DeleteV1IncidentsIncidentIDImpactTypeIDPathParamType = "environments"
	DeleteV1IncidentsIncidentIDImpactTypeIDPathParamTypeFunctionalities DeleteV1IncidentsIncidentIDImpactTypeIDPathParamType = "functionalities"
	DeleteV1IncidentsIncidentIDImpactTypeIDPathParamTypeServices        DeleteV1IncidentsIncidentIDImpactTypeIDPathParamType = "services"
	DeleteV1IncidentsIncidentIDImpactTypeIDPathParamTypeCustomers       DeleteV1IncidentsIncidentIDImpactTypeIDPathParamType = "customers"
)

func (e DeleteV1IncidentsIncidentIDImpactTypeIDPathParamType) ToPointer() *DeleteV1IncidentsIncidentIDImpactTypeIDPathParamType {
	return &e
}
func (e *DeleteV1IncidentsIncidentIDImpactTypeIDPathParamType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "environments":
		fallthrough
	case "functionalities":
		fallthrough
	case "services":
		fallthrough
	case "customers":
		*e = DeleteV1IncidentsIncidentIDImpactTypeIDPathParamType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteV1IncidentsIncidentIDImpactTypeIDPathParamType: %v", v)
	}
}

type DeleteV1IncidentsIncidentIDImpactTypeIDRequest struct {
	IncidentID string                                               `pathParam:"style=simple,explode=false,name=incident_id"`
	Type       DeleteV1IncidentsIncidentIDImpactTypeIDPathParamType `pathParam:"style=simple,explode=false,name=type"`
	ID         string                                               `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteV1IncidentsIncidentIDImpactTypeIDRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *DeleteV1IncidentsIncidentIDImpactTypeIDRequest) GetType() DeleteV1IncidentsIncidentIDImpactTypeIDPathParamType {
	if o == nil {
		return DeleteV1IncidentsIncidentIDImpactTypeIDPathParamType("")
	}
	return o.Type
}

func (o *DeleteV1IncidentsIncidentIDImpactTypeIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteV1IncidentsIncidentIDImpactTypeIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Bad Request
	ErrorEntity *shared.ErrorEntity
}

func (o *DeleteV1IncidentsIncidentIDImpactTypeIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteV1IncidentsIncidentIDImpactTypeIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteV1IncidentsIncidentIDImpactTypeIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteV1IncidentsIncidentIDImpactTypeIDResponse) GetErrorEntity() *shared.ErrorEntity {
	if o == nil {
		return nil
	}
	return o.ErrorEntity
}