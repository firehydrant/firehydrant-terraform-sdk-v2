// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// Encoding for submitted data
type Encoding string

const (
	EncodingTextYaml         Encoding = "text/yaml"
	EncodingApplicationXYaml Encoding = "application/x-yaml"
	EncodingApplicationJSON  Encoding = "application/json"
)

func (e Encoding) ToPointer() *Encoding {
	return &e
}
func (e *Encoding) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text/yaml":
		fallthrough
	case "application/x-yaml":
		fallthrough
	case "application/json":
		*e = Encoding(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Encoding: %v", v)
	}
}

// PostV1CatalogsCatalogIDIngest - Accepts catalog data in the configured format and asyncronously processes the data to incorporate changes into service catalog.
type PostV1CatalogsCatalogIDIngest struct {
	// Encoding for submitted data
	Encoding Encoding `json:"encoding"`
	// Service data
	Data string `json:"data"`
}

func (o *PostV1CatalogsCatalogIDIngest) GetEncoding() Encoding {
	if o == nil {
		return Encoding("")
	}
	return o.Encoding
}

func (o *PostV1CatalogsCatalogIDIngest) GetData() string {
	if o == nil {
		return ""
	}
	return o.Data
}
