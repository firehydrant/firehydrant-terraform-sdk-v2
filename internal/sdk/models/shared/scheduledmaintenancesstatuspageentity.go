// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type ScheduledMaintenancesStatusPageEntity struct {
	ID              *string `json:"id,omitempty"`
	IntegrationID   *string `json:"integration_id,omitempty"`
	IntegrationSlug *string `json:"integration_slug,omitempty"`
	IntegrationName *string `json:"integration_name,omitempty"`
}

func (o *ScheduledMaintenancesStatusPageEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ScheduledMaintenancesStatusPageEntity) GetIntegrationID() *string {
	if o == nil {
		return nil
	}
	return o.IntegrationID
}

func (o *ScheduledMaintenancesStatusPageEntity) GetIntegrationSlug() *string {
	if o == nil {
		return nil
	}
	return o.IntegrationSlug
}

func (o *ScheduledMaintenancesStatusPageEntity) GetIntegrationName() *string {
	if o == nil {
		return nil
	}
	return o.IntegrationName
}
