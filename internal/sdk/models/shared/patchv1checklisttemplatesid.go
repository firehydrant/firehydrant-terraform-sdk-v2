// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type PatchV1ChecklistTemplatesIDChecks struct {
	// Specify the check ID when updating an already existing check
	ID *string `json:"id,omitempty"`
	// The description of the check
	Description *string `json:"description,omitempty"`
	// The name of the check
	Name string `json:"name"`
}

func (o *PatchV1ChecklistTemplatesIDChecks) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PatchV1ChecklistTemplatesIDChecks) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PatchV1ChecklistTemplatesIDChecks) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type PatchV1ChecklistTemplatesIDConnectedServices struct {
	ID string `json:"id"`
	// Set to `true` to remove checklist from service
	Remove *bool `json:"remove,omitempty"`
}

func (o *PatchV1ChecklistTemplatesIDConnectedServices) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PatchV1ChecklistTemplatesIDConnectedServices) GetRemove() *bool {
	if o == nil {
		return nil
	}
	return o.Remove
}

// PatchV1ChecklistTemplatesID - Update a checklist templates attributes
type PatchV1ChecklistTemplatesID struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	// An array of checks for the checklist template
	Checks []PatchV1ChecklistTemplatesIDChecks `json:"checks,omitempty"`
	// The ID of the Team that owns the checklist template
	TeamID *string `json:"team_id,omitempty"`
	// Array of service IDs to attach checklist template to
	ConnectedServices []PatchV1ChecklistTemplatesIDConnectedServices `json:"connected_services,omitempty"`
	// If set to true, any services tagged on the checklist that are not included in the given array will be removed. Set this to true if you want to do a replacement operation for the services
	RemoveRemainingConnectedServices *bool `json:"remove_remaining_connected_services,omitempty"`
}

func (o *PatchV1ChecklistTemplatesID) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PatchV1ChecklistTemplatesID) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PatchV1ChecklistTemplatesID) GetChecks() []PatchV1ChecklistTemplatesIDChecks {
	if o == nil {
		return nil
	}
	return o.Checks
}

func (o *PatchV1ChecklistTemplatesID) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *PatchV1ChecklistTemplatesID) GetConnectedServices() []PatchV1ChecklistTemplatesIDConnectedServices {
	if o == nil {
		return nil
	}
	return o.ConnectedServices
}

func (o *PatchV1ChecklistTemplatesID) GetRemoveRemainingConnectedServices() *bool {
	if o == nil {
		return nil
	}
	return o.RemoveRemainingConnectedServices
}
