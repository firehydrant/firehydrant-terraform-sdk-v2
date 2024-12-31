// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// PatchV1SavedSearchesResourceTypeSavedSearchID - Update a specific saved search
type PatchV1SavedSearchesResourceTypeSavedSearchID struct {
	IsPrivate    *bool          `json:"is_private,omitempty"`
	Name         *string        `json:"name,omitempty"`
	FilterValues map[string]any `json:"filter_values,omitempty"`
}

func (o *PatchV1SavedSearchesResourceTypeSavedSearchID) GetIsPrivate() *bool {
	if o == nil {
		return nil
	}
	return o.IsPrivate
}

func (o *PatchV1SavedSearchesResourceTypeSavedSearchID) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PatchV1SavedSearchesResourceTypeSavedSearchID) GetFilterValues() map[string]any {
	if o == nil {
		return nil
	}
	return o.FilterValues
}
