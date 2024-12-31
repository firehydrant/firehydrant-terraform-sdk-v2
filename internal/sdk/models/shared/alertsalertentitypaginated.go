// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// AlertsAlertEntityPaginated - Alerts_AlertEntityPaginated model
type AlertsAlertEntityPaginated struct {
	Data       []AlertsAlertEntity `json:"data,omitempty"`
	Pagination *PaginationEntity   `json:"pagination,omitempty"`
}

func (o *AlertsAlertEntityPaginated) GetData() []AlertsAlertEntity {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *AlertsAlertEntityPaginated) GetPagination() *PaginationEntity {
	if o == nil {
		return nil
	}
	return o.Pagination
}
