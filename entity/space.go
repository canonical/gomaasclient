package entity

type Space struct {
	Name        string   `json:"name,omitempty"`
	ResourceURI string   `json:"resource_uri,omitempty"`
	Subnets     []Subnet `json:"subnets,omitempty"`
	VLANs       []VLAN   `json:"vlans,omitempty"`
	ID          int      `json:"id,omitempty"`
}
