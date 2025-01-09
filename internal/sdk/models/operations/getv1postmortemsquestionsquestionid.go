// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetV1PostMortemsQuestionsQuestionIDRequest struct {
	QuestionID string `pathParam:"style=simple,explode=false,name=question_id"`
}

func (o *GetV1PostMortemsQuestionsQuestionIDRequest) GetQuestionID() string {
	if o == nil {
		return ""
	}
	return o.QuestionID
}

type GetV1PostMortemsQuestionsQuestionIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetV1PostMortemsQuestionsQuestionIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetV1PostMortemsQuestionsQuestionIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetV1PostMortemsQuestionsQuestionIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}