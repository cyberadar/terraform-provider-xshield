// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type VulnerabilitiesRequestData struct {
	PackageString *string           `json:"packageString,omitempty"`
	Pagination    *PaginationConfig `json:"pagination,omitempty"`
}

func (o *VulnerabilitiesRequestData) GetPackageString() *string {
	if o == nil {
		return nil
	}
	return o.PackageString
}

func (o *VulnerabilitiesRequestData) GetPagination() *PaginationConfig {
	if o == nil {
		return nil
	}
	return o.Pagination
}
