// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type MetadataColumnDescriptor struct {
	CoreTag      *bool                `json:"coreTag,omitempty"`
	DataType     *DatatypesType       `json:"dataType,omitempty"`
	DisplayName  *string              `json:"displayName,omitempty"`
	Facetable    *bool                `json:"facetable,omitempty"`
	InternalName *string              `json:"internalName,omitempty"`
	ListOfValues *bool                `json:"listOfValues,omitempty"`
	Multivalued  *bool                `json:"multivalued,omitempty"`
	Qualifier    *string              `json:"qualifier,omitempty"`
	Searchable   *bool                `json:"searchable,omitempty"`
	Sortable     *bool                `json:"sortable,omitempty"`
	Unit         *string              `json:"unit,omitempty"`
	UserDefined  *bool                `json:"userDefined,omitempty"`
	Values       []MetadataFieldValue `json:"values,omitempty"`
}

func (o *MetadataColumnDescriptor) GetCoreTag() *bool {
	if o == nil {
		return nil
	}
	return o.CoreTag
}

func (o *MetadataColumnDescriptor) GetDataType() *DatatypesType {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *MetadataColumnDescriptor) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *MetadataColumnDescriptor) GetFacetable() *bool {
	if o == nil {
		return nil
	}
	return o.Facetable
}

func (o *MetadataColumnDescriptor) GetInternalName() *string {
	if o == nil {
		return nil
	}
	return o.InternalName
}

func (o *MetadataColumnDescriptor) GetListOfValues() *bool {
	if o == nil {
		return nil
	}
	return o.ListOfValues
}

func (o *MetadataColumnDescriptor) GetMultivalued() *bool {
	if o == nil {
		return nil
	}
	return o.Multivalued
}

func (o *MetadataColumnDescriptor) GetQualifier() *string {
	if o == nil {
		return nil
	}
	return o.Qualifier
}

func (o *MetadataColumnDescriptor) GetSearchable() *bool {
	if o == nil {
		return nil
	}
	return o.Searchable
}

func (o *MetadataColumnDescriptor) GetSortable() *bool {
	if o == nil {
		return nil
	}
	return o.Sortable
}

func (o *MetadataColumnDescriptor) GetUnit() *string {
	if o == nil {
		return nil
	}
	return o.Unit
}

func (o *MetadataColumnDescriptor) GetUserDefined() *bool {
	if o == nil {
		return nil
	}
	return o.UserDefined
}

func (o *MetadataColumnDescriptor) GetValues() []MetadataFieldValue {
	if o == nil {
		return nil
	}
	return o.Values
}
