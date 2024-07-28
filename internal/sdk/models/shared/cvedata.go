// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CVEData struct {
	CveID       *string  `json:"cveID,omitempty"`
	Cvssscore   *float64 `json:"cvssscore,omitempty"`
	Description *string  `json:"description,omitempty"`
	ExploitURL  *string  `json:"exploitURL,omitempty"`
	Product     *string  `json:"product,omitempty"`
	Severity    *int64   `json:"severity,omitempty"`
	Vendor      *string  `json:"vendor,omitempty"`
	Version     *string  `json:"version,omitempty"`
}

func (o *CVEData) GetCveID() *string {
	if o == nil {
		return nil
	}
	return o.CveID
}

func (o *CVEData) GetCvssscore() *float64 {
	if o == nil {
		return nil
	}
	return o.Cvssscore
}

func (o *CVEData) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CVEData) GetExploitURL() *string {
	if o == nil {
		return nil
	}
	return o.ExploitURL
}

func (o *CVEData) GetProduct() *string {
	if o == nil {
		return nil
	}
	return o.Product
}

func (o *CVEData) GetSeverity() *int64 {
	if o == nil {
		return nil
	}
	return o.Severity
}

func (o *CVEData) GetVendor() *string {
	if o == nil {
		return nil
	}
	return o.Vendor
}

func (o *CVEData) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}
