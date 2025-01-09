// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type GetV1AlertsAlertIDRequest struct {
	AlertID string `pathParam:"style=simple,explode=false,name=alert_id"`
}

func (o *GetV1AlertsAlertIDRequest) GetAlertID() string {
	if o == nil {
		return ""
	}
	return o.AlertID
}

type GetV1AlertsAlertIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Retrieve a single alert
	AlertsAlertEntity *shared.AlertsAlertEntity
}

func (o *GetV1AlertsAlertIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1AlertsAlertIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1AlertsAlertIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1AlertsAlertIDResponse) GetAlertsAlertEntity() *shared.AlertsAlertEntity {
	if o == nil {
		return nil
	}
	return o.AlertsAlertEntity
}
