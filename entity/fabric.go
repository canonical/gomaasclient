package entity

type Fabric struct {
	ID          int    `json:"id,omitempty"`
	VLANs       []VLAN `json:"vlans,omitempty"`
	Name        string `json:"name,omitempty"`
	ClassType   string `json:"class_type,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

type FabricParams struct {
	Name        string `url:"name,omitempty"`
	Description string `url:"description,omitempty"`
	ClassType   string `url:"class_type,omitempty"`
}
