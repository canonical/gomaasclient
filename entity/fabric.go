package entity

type Fabric struct {
	Name        string `json:"name,omitempty"`
	ClassType   string `json:"class_type,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
	VLANs       []VLAN `json:"vlans,omitempty"`
	ID          int    `json:"id,omitempty"`
}

type FabricParams struct {
	Name        string `url:"name,omitempty"`
	Description string `url:"description,omitempty"`
	ClassType   string `url:"class_type,omitempty"`
}
