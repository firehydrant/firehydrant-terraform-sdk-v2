// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type UpdateStatusPageImageFile struct {
	FileName string `multipartForm:"name=fileName"`
	Content  []byte `multipartForm:"content"`
}

func (o *UpdateStatusPageImageFile) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

func (o *UpdateStatusPageImageFile) GetContent() []byte {
	if o == nil {
		return []byte{}
	}
	return o.Content
}

type UpdateStatusPageImageRequestBody struct {
	File *UpdateStatusPageImageFile `multipartForm:"file,name=file"`
}

func (o *UpdateStatusPageImageRequestBody) GetFile() *UpdateStatusPageImageFile {
	if o == nil {
		return nil
	}
	return o.File
}

type UpdateStatusPageImageRequest struct {
	NuncConnectionID string                            `pathParam:"style=simple,explode=false,name=nunc_connection_id"`
	Type             string                            `pathParam:"style=simple,explode=false,name=type"`
	RequestBody      *UpdateStatusPageImageRequestBody `request:"mediaType=multipart/form-data"`
}

func (o *UpdateStatusPageImageRequest) GetNuncConnectionID() string {
	if o == nil {
		return ""
	}
	return o.NuncConnectionID
}

func (o *UpdateStatusPageImageRequest) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *UpdateStatusPageImageRequest) GetRequestBody() *UpdateStatusPageImageRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type UpdateStatusPageImageResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Add or replace an image attached to a FireHydrant status page
	NuncConnectionEntity *shared.NuncConnectionEntity
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

func (o *UpdateStatusPageImageResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateStatusPageImageResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateStatusPageImageResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateStatusPageImageResponse) GetNuncConnectionEntity() *shared.NuncConnectionEntity {
	if o == nil {
		return nil
	}
	return o.NuncConnectionEntity
}

func (o *UpdateStatusPageImageResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *UpdateStatusPageImageResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *UpdateStatusPageImageResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *UpdateStatusPageImageResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *UpdateStatusPageImageResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *UpdateStatusPageImageResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}
