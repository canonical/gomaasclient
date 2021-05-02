package entity

type Fabric struct {
	ID          int    `json:"id,omitempty"`
	VLANs       []VLAN `json:"vlans,omitempty"`
	Name        string `json:"name,omitempty"`
	ClassType   string `json:"class_type,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
}
