// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type VulnerablePackageData struct {
	CriticalCveCount *int64   `json:"criticalCveCount,omitempty"`
	CvssScore        *float64 `json:"cvssScore,omitempty"`
	HighCveCount     *int64   `json:"highCveCount,omitempty"`
	Product          *string  `json:"product,omitempty"`
	Severity         *int64   `json:"severity,omitempty"`
	TotalCveCount    *int64   `json:"totalCveCount,omitempty"`
	Vendor           *string  `json:"vendor,omitempty"`
	Version          *string  `json:"version,omitempty"`
	VulnSource       *string  `json:"vulnSource,omitempty"`
}

func (o *VulnerablePackageData) GetCriticalCveCount() *int64 {
	if o == nil {
		return nil
	}
	return o.CriticalCveCount
}

func (o *VulnerablePackageData) GetCvssScore() *float64 {
	if o == nil {
		return nil
	}
	return o.CvssScore
}

func (o *VulnerablePackageData) GetHighCveCount() *int64 {
	if o == nil {
		return nil
	}
	return o.HighCveCount
}

func (o *VulnerablePackageData) GetProduct() *string {
	if o == nil {
		return nil
	}
	return o.Product
}

func (o *VulnerablePackageData) GetSeverity() *int64 {
	if o == nil {
		return nil
	}
	return o.Severity
}

func (o *VulnerablePackageData) GetTotalCveCount() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalCveCount
}

func (o *VulnerablePackageData) GetVendor() *string {
	if o == nil {
		return nil
	}
	return o.Vendor
}

func (o *VulnerablePackageData) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}

func (o *VulnerablePackageData) GetVulnSource() *string {
	if o == nil {
		return nil
	}
	return o.VulnSource
}
