// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/internal/utils"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
	"time"
)

// Sort - Allows sorting comments by the time they were posted, ascending or descending.
type Sort string

const (
	SortAsc  Sort = "asc"
	SortDesc Sort = "desc"
)

func (e Sort) ToPointer() *Sort {
	return &e
}
func (e *Sort) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = Sort(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Sort: %v", v)
	}
}

type ListConversationCommentsRequest struct {
	// An ISO8601 timestamp that allows filtering for comments posted before the provided time.
	Before *time.Time `queryParam:"style=form,explode=true,name=before"`
	// An ISO8601 timestamp that allows filtering for comments posted after the provided time.
	After *time.Time `queryParam:"style=form,explode=true,name=after"`
	// Allows sorting comments by the time they were posted, ascending or descending.
	Sort           *Sort  `default:"asc" queryParam:"style=form,explode=true,name=sort"`
	ConversationID string `pathParam:"style=simple,explode=false,name=conversation_id"`
}

func (l ListConversationCommentsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListConversationCommentsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListConversationCommentsRequest) GetBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.Before
}

func (o *ListConversationCommentsRequest) GetAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *ListConversationCommentsRequest) GetSort() *Sort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListConversationCommentsRequest) GetConversationID() string {
	if o == nil {
		return ""
	}
	return o.ConversationID
}

type ListConversationCommentsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
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

func (o *ListConversationCommentsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListConversationCommentsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListConversationCommentsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListConversationCommentsResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *ListConversationCommentsResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *ListConversationCommentsResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *ListConversationCommentsResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *ListConversationCommentsResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *ListConversationCommentsResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}
