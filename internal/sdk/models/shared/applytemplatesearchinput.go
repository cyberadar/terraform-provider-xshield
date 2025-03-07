// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type ApplyTemplateSearchInput struct {
	Comment    *string  `json:"comment,omitempty"`
	Criteria   string   `json:"criteria"`
	TemplateID *string  `json:"templateId,omitempty"`
	Templates  []string `json:"templates,omitempty"`
}

func (o *ApplyTemplateSearchInput) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *ApplyTemplateSearchInput) GetCriteria() string {
	if o == nil {
		return ""
	}
	return o.Criteria
}

func (o *ApplyTemplateSearchInput) GetTemplateID() *string {
	if o == nil {
		return nil
	}
	return o.TemplateID
}

func (o *ApplyTemplateSearchInput) GetTemplates() []string {
	if o == nil {
		return nil
	}
	return o.Templates
}
