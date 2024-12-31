// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/internal/utils"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
	"time"
)

type ListChangeEventsRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
	// The id of a previously saved search.
	SavedSearchID *string `queryParam:"style=form,explode=true,name=saved_search_id"`
	// A text query for change events
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// A comma separated list of label key / values in the format of "key=value,key2=value2". To filter change events that have a key (with no specific value), omit the value
	Labels *string `queryParam:"style=form,explode=true,name=labels"`
	// A comma separated list of environment IDs
	Environments *string `queryParam:"style=form,explode=true,name=environments"`
	// A comma separated list of service IDs
	Services *string `queryParam:"style=form,explode=true,name=services"`
	// The start time to start returning change events from
	StartsAt *string `queryParam:"style=form,explode=true,name=starts_at"`
	// The end time to return change events up to
	EndsAt *time.Time `queryParam:"style=form,explode=true,name=ends_at"`
}

func (l ListChangeEventsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListChangeEventsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListChangeEventsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListChangeEventsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *ListChangeEventsRequest) GetSavedSearchID() *string {
	if o == nil {
		return nil
	}
	return o.SavedSearchID
}

func (o *ListChangeEventsRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListChangeEventsRequest) GetLabels() *string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *ListChangeEventsRequest) GetEnvironments() *string {
	if o == nil {
		return nil
	}
	return o.Environments
}

func (o *ListChangeEventsRequest) GetServices() *string {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *ListChangeEventsRequest) GetStartsAt() *string {
	if o == nil {
		return nil
	}
	return o.StartsAt
}

func (o *ListChangeEventsRequest) GetEndsAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndsAt
}

type ListChangeEventsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List change events for the organization. Note: Not all information is included on a change event like attachments and related changes. You must fetch a change event separately to retrieve all of the information about it
	ChangeEventSlimEntityPaginated *shared.ChangeEventSlimEntityPaginated
	// A collection of codes that generally means the end user got something wrong in making the request
	BadRequest *shared.BadRequest
	// A collection of codes that generally means the client was not authenticated correctly for the request they want to make
	Unauthorized *shared.Unauthorized
	// Status codes relating to the resource/entity they are requesting not being found or endpoints/routes not existing
	NotFound *shared.NotFound
	// Status codes relating to the client being rate limited by the server
	RateLimited *shared.RateLimited
	// A collection of status codes that generally mean the server failed in an unexpected way
	InternalServerError *shared.InternalServerError
	// Timeouts occurred with the request
	Timeout *shared.Timeout
}

func (o *ListChangeEventsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListChangeEventsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListChangeEventsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListChangeEventsResponse) GetChangeEventSlimEntityPaginated() *shared.ChangeEventSlimEntityPaginated {
	if o == nil {
		return nil
	}
	return o.ChangeEventSlimEntityPaginated
}

func (o *ListChangeEventsResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *ListChangeEventsResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *ListChangeEventsResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *ListChangeEventsResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *ListChangeEventsResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *ListChangeEventsResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}
