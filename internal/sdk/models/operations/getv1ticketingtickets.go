// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"net/http"
)

// GetV1TicketingTicketsQueryParamTagMatchStrategy - A matching strategy for the tags provided
type GetV1TicketingTicketsQueryParamTagMatchStrategy string

const (
	GetV1TicketingTicketsQueryParamTagMatchStrategyAny      GetV1TicketingTicketsQueryParamTagMatchStrategy = "any"
	GetV1TicketingTicketsQueryParamTagMatchStrategyMatchAll GetV1TicketingTicketsQueryParamTagMatchStrategy = "match_all"
	GetV1TicketingTicketsQueryParamTagMatchStrategyExclude  GetV1TicketingTicketsQueryParamTagMatchStrategy = "exclude"
)

func (e GetV1TicketingTicketsQueryParamTagMatchStrategy) ToPointer() *GetV1TicketingTicketsQueryParamTagMatchStrategy {
	return &e
}
func (e *GetV1TicketingTicketsQueryParamTagMatchStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "any":
		fallthrough
	case "match_all":
		fallthrough
	case "exclude":
		*e = GetV1TicketingTicketsQueryParamTagMatchStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetV1TicketingTicketsQueryParamTagMatchStrategy: %v", v)
	}
}

// State - Filter tickets by state
type State string

const (
	StateOpen       State = "open"
	StateInProgress State = "in_progress"
	StateCancelled  State = "cancelled"
	StateDone       State = "done"
)

func (e State) ToPointer() *State {
	return &e
}
func (e *State) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "open":
		fallthrough
	case "in_progress":
		fallthrough
	case "cancelled":
		fallthrough
	case "done":
		*e = State(v)
		return nil
	default:
		return fmt.Errorf("invalid value for State: %v", v)
	}
}

type GetV1TicketingTicketsRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
	// A comma separated list of tags
	Tags *string `queryParam:"style=form,explode=true,name=tags"`
	// A matching strategy for the tags provided
	TagMatchStrategy *GetV1TicketingTicketsQueryParamTagMatchStrategy `queryParam:"style=form,explode=true,name=tag_match_strategy"`
	// Filter tickets assigned to this user id
	AssignedUser *string `queryParam:"style=form,explode=true,name=assigned_user"`
	// Filter tickets by state
	State *State `queryParam:"style=form,explode=true,name=state"`
}

func (o *GetV1TicketingTicketsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetV1TicketingTicketsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *GetV1TicketingTicketsRequest) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *GetV1TicketingTicketsRequest) GetTagMatchStrategy() *GetV1TicketingTicketsQueryParamTagMatchStrategy {
	if o == nil {
		return nil
	}
	return o.TagMatchStrategy
}

func (o *GetV1TicketingTicketsRequest) GetAssignedUser() *string {
	if o == nil {
		return nil
	}
	return o.AssignedUser
}

func (o *GetV1TicketingTicketsRequest) GetState() *State {
	if o == nil {
		return nil
	}
	return o.State
}

type GetV1TicketingTicketsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List all of the tickets that have been added to the organiation
	TicketingTicketEntity *shared.TicketingTicketEntity
}

func (o *GetV1TicketingTicketsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1TicketingTicketsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1TicketingTicketsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetV1TicketingTicketsResponse) GetTicketingTicketEntity() *shared.TicketingTicketEntity {
	if o == nil {
		return nil
	}
	return o.TicketingTicketEntity
}
