// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// PostV1ServiceDependencies - Creates a service dependency relationship between two services
type PostV1ServiceDependencies struct {
	ServiceID          string `json:"service_id"`
	ConnectedServiceID string `json:"connected_service_id"`
	// A note to describe the service dependency relationship
	Notes *string `json:"notes,omitempty"`
}

func (o *PostV1ServiceDependencies) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *PostV1ServiceDependencies) GetConnectedServiceID() string {
	if o == nil {
		return ""
	}
	return o.ConnectedServiceID
}

func (o *PostV1ServiceDependencies) GetNotes() *string {
	if o == nil {
		return nil
	}
	return o.Notes
}
