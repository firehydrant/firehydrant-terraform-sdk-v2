// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Security struct {
	APIKey *string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

func (o *Security) GetAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.APIKey
}
