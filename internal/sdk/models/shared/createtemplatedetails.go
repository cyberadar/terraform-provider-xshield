// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type TemplateType string

const (
	TemplateTypeApplicationTemplate TemplateType = "application-template"
	TemplateTypeBlockTemplate       TemplateType = "block-template"
)

func (e TemplateType) ToPointer() *TemplateType {
	return &e
}
func (e *TemplateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "application-template":
		fallthrough
	case "block-template":
		*e = TemplateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TemplateType: %v", v)
	}
}

type CreateTemplateDetails struct {
	TemplateCategory    *string        `json:"templateCategory,omitempty"`
	TemplateDescription *string        `json:"templateDescription,omitempty"`
	TemplateName        *string        `json:"templateName,omitempty"`
	TemplatePaths       []MetadataPath `json:"templatePaths,omitempty"`
	TemplatePorts       []MetadataPort `json:"templatePorts,omitempty"`
	TemplateType        *TemplateType  `json:"templateType,omitempty"`
}

func (o *CreateTemplateDetails) GetTemplateCategory() *string {
	if o == nil {
		return nil
	}
	return o.TemplateCategory
}

func (o *CreateTemplateDetails) GetTemplateDescription() *string {
	if o == nil {
		return nil
	}
	return o.TemplateDescription
}

func (o *CreateTemplateDetails) GetTemplateName() *string {
	if o == nil {
		return nil
	}
	return o.TemplateName
}

func (o *CreateTemplateDetails) GetTemplatePaths() []MetadataPath {
	if o == nil {
		return nil
	}
	return o.TemplatePaths
}

func (o *CreateTemplateDetails) GetTemplatePorts() []MetadataPort {
	if o == nil {
		return nil
	}
	return o.TemplatePorts
}

func (o *CreateTemplateDetails) GetTemplateType() *TemplateType {
	if o == nil {
		return nil
	}
	return o.TemplateType
}
