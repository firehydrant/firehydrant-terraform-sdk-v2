// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// PostV1Webhooks - Create a new webhook
type PostV1Webhooks struct {
	URL string `json:"url"`
}

func (o *PostV1Webhooks) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}
