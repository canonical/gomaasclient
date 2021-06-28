package entity

type Space struct {
	ID          int      `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Subnets     []Subnet `json:"subnets,omitempty"`
	VLANs       []VLAN   `json:"vlans,omitempty"`
	ResourceURI string   `json:"resource_uri,omitempty"`
}
