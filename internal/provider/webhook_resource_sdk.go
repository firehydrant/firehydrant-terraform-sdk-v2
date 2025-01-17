// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
)

func (r *WebhookResourceModel) ToSharedPostV1Webhooks() *shared.PostV1Webhooks {
	var url string
	url = r.URL.ValueString()

	out := shared.PostV1Webhooks{
		URL: url,
	}
	return &out
}

func (r *WebhookResourceModel) ToSharedPatchV1WebhooksWebhookID() *shared.PatchV1WebhooksWebhookID {
	url := new(string)
	if !r.URL.IsUnknown() && !r.URL.IsNull() {
		*url = r.URL.ValueString()
	} else {
		url = nil
	}
	state := new(shared.PatchV1WebhooksWebhookIDState)
	if !r.State.IsUnknown() && !r.State.IsNull() {
		*state = shared.PatchV1WebhooksWebhookIDState(r.State.ValueString())
	} else {
		state = nil
	}
	out := shared.PatchV1WebhooksWebhookID{
		URL:   url,
		State: state,
	}
	return &out
}
