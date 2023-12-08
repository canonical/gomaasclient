package entity

// VLAN represents the MAAS VLAN endpoint.
type VLAN struct {
	RelayVLAN     *VLAN  `json:"relay_vlan,omitempty"`
	Name          string `json:"name,omitempty"`
	ExternalDHCP  string `json:"external_dhcp,omitempty"`
	Description   string `json:"description,omitempty"`
	PrimaryRack   string `json:"primary_rack,omitempty"`
	SecondaryRack string `json:"secondary_rack,omitempty"`
	Space         string `json:"space,omitempty"`
	Fabric        string `json:"fabric,omitempty"`
	ResourceURI   string `json:"resource_uri,omitempty"`
	MTU           int    `json:"mtu,omitempty"`
	FabricID      int    `json:"fabric_id,omitempty"`
	VID           int    `json:"vid,omitempty"`
	ID            int    `json:"id,omitempty"`
	DHCPOn        bool   `json:"dhcp_on,omitempty"`
}

// VLANParams contains the options for a POST request to the vlans endpoint.
// Only the VID field is required. If Space is empty or the string "undefined",
// the VLAN will be created in the 'undefined' space.
type VLANParams struct {
	Name          string `url:"name,omitempty"`
	Description   string `url:"description,omitempty"`
	Space         string `url:"space,omitempty"`
	PrimaryRack   string `url:"primary_rack,omitempty"`
	SecondaryRack string `url:"secondary_rack,omitempty"`
	VID           int    `url:"vid,omitempty"`
	MTU           int    `url:"mtu,omitempty"`
	RelayVLAN     int    `url:"relay_vlan,omitempty"`
	DHCPOn        bool   `url:"dhcp_on"`
}
