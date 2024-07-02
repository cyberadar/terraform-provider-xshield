// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// NetworkInterface definition Details network interface information for a specific asset
type NetworkInterface struct {
	Flags       []string `json:"flags,omitempty"`
	Ipaddresses []string `json:"ipaddresses"`
	Macaddress  *string  `json:"macaddress,omitempty"`
	Name        string   `json:"name"`
}

func (o *NetworkInterface) GetFlags() []string {
	if o == nil {
		return nil
	}
	return o.Flags
}

func (o *NetworkInterface) GetIpaddresses() []string {
	if o == nil {
		return []string{}
	}
	return o.Ipaddresses
}

func (o *NetworkInterface) GetMacaddress() *string {
	if o == nil {
		return nil
	}
	return o.Macaddress
}

func (o *NetworkInterface) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
