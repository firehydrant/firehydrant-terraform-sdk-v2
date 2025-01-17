// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type PatchV1WebhooksWebhookIDState string

const (
	PatchV1WebhooksWebhookIDStateActive   PatchV1WebhooksWebhookIDState = "active"
	PatchV1WebhooksWebhookIDStateInactive PatchV1WebhooksWebhookIDState = "inactive"
)

func (e PatchV1WebhooksWebhookIDState) ToPointer() *PatchV1WebhooksWebhookIDState {
	return &e
}
func (e *PatchV1WebhooksWebhookIDState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "active":
		fallthrough
	case "inactive":
		*e = PatchV1WebhooksWebhookIDState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PatchV1WebhooksWebhookIDState: %v", v)
	}
}

// PatchV1WebhooksWebhookID - Update a specific webhook
type PatchV1WebhooksWebhookID struct {
	URL   *string                        `json:"url,omitempty"`
	State *PatchV1WebhooksWebhookIDState `json:"state,omitempty"`
}

func (o *PatchV1WebhooksWebhookID) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *PatchV1WebhooksWebhookID) GetState() *PatchV1WebhooksWebhookIDState {
	if o == nil {
		return nil
	}
	return o.State
}
