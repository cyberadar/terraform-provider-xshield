// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type MetadataPort struct {
	ListenPort         *int64             `json:"listenPort,omitempty"`
	ListenPortName     *string            `json:"listenPortName,omitempty"`
	ListenPortProtocol *int64             `json:"listenPortProtocol,omitempty"`
	ListenPortReviewed *MetadataPortState `json:"listenPortReviewed,omitempty"`
	ListenProcessNames []string           `json:"listenProcessNames,omitempty"`
	LpID               *string            `json:"lpId,omitempty"`
}

func (o *MetadataPort) GetListenPort() *int64 {
	if o == nil {
		return nil
	}
	return o.ListenPort
}

func (o *MetadataPort) GetListenPortName() *string {
	if o == nil {
		return nil
	}
	return o.ListenPortName
}

func (o *MetadataPort) GetListenPortProtocol() *int64 {
	if o == nil {
		return nil
	}
	return o.ListenPortProtocol
}

func (o *MetadataPort) GetListenPortReviewed() *MetadataPortState {
	if o == nil {
		return nil
	}
	return o.ListenPortReviewed
}

func (o *MetadataPort) GetListenProcessNames() []string {
	if o == nil {
		return nil
	}
	return o.ListenProcessNames
}

func (o *MetadataPort) GetLpID() *string {
	if o == nil {
		return nil
	}
	return o.LpID
}
