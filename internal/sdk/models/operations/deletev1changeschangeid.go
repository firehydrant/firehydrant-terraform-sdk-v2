// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteV1ChangesChangeIDRequest struct {
	ChangeID string `pathParam:"style=simple,explode=false,name=change_id"`
}

func (o *DeleteV1ChangesChangeIDRequest) GetChangeID() string {
	if o == nil {
		return ""
	}
	return o.ChangeID
}

type DeleteV1ChangesChangeIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteV1ChangesChangeIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteV1ChangesChangeIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteV1ChangesChangeIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
