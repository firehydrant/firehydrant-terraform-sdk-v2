// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// PatchV1IncidentsIncidentIDGenericChatMessagesMessageID - Update an existing generic chat message on an incident.
type PatchV1IncidentsIncidentIDGenericChatMessagesMessageID struct {
	Body string `json:"body"`
}

func (o *PatchV1IncidentsIncidentIDGenericChatMessagesMessageID) GetBody() string {
	if o == nil {
		return ""
	}
	return o.Body
}
