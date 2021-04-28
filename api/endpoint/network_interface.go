package endpoint

import "net"

// NetworkInterface represents the MaaS Interface endpoint.
type NetworkInterface struct {
	VLAN               VLAN                   `json:"vlan,omitempty"`
	Children           []string               `json:"children,omitempty"`
	Parents            []string               `json:"parents,omitempty"`
	Tags               []string               `json:"tags,omitempty"`
	Discovered         []NetworkInterfaceLink `json:"discovered,omitempty"`
	Links              []NetworkInterfaceLink `json:"links,omitempty"`
	Name               string                 `json:"name,omitempty"`
	MACAddress         string                 `json:"mac_address,omitempty"`
	Product            string                 `json:"product,omitempty"`
	FirmwareVersion    string                 `json:"firmware_version,omitempty"`
	SystemID           string                 `json:"system_id,omitempty"`
	Params             interface{}            `json:"params,omitempty"`
	Type               string                 `json:"type,omitempty"`
	Vendor             string                 `json:"vendor,omitempty"`
	ResourceURI        string                 `json:"resource_uri,omitempty"`
	BondXMitHashPolicy string                 `json:"bond_x_mit_hash_policy,omitempty"`
	BondMode           string                 `json:"bond_mode,omitempty"`
	MTU                string                 `json:"mtu,omitempty"`
	EffectiveMTU       int                    `json:"effective_mtu,omitempty"`
	ID                 int                    `json:"id,omitempty"`
	BridgeFD           int                    `json:"bridge_fd,omitempty"`
	BondMIIMon         int                    `json:"bond_mii_mon,omitempty"`
	BondDownDelay      int                    `json:"bond_down_delay,omitempty"`
	BondUpDelay        int                    `json:"bond_up_delay,omitempty"`
	BondLACPRate       int                    `json:"bond_lacp_rate,omitempty"`
	AcceptRA           bool                   `json:"accept_ra,omitempty"`
	Autoconf           bool                   `json:"autoconf,omitempty"`
	Enabled            bool                   `json:"enabled,omitempty"`
	BridgeSTP          bool                   `json:"bridge_stp,omitempty"`
}

// NetworkInterfaceLink is consumed by NetworkInterface{} and should not be used directly.
type NetworkInterfaceLink struct {
	ID        int    `json:"id,omitempty"`
	Mode      string `json:"mode,omitempty"`
	Subnet    Subnet `json:"subnet,omitempty"`
	IPAddress net.IP `json:"ip_address,omitempty"`
}

// NetworkInterfaceLinkParams is used with NetworkInterface.LinkSubnet().
// Mode must be one of (AUTO, DHCP, STATIC, LINK_UP). IPAddress is ignored
// unless mode is STATIC, and will be set automatically if empty. Force
// allows LINK_UP to be set when other links exist, allows links between
// different VLANs, and deletes all other links on the interface.
// DefaultGateway is ignored unless Mode is AUTO or STATIC.
// Note: You can parse an IP address into a net.IP via net.ParseIP(string).
type NetworkInterfaceLinkParams struct {
	Mode           string `json:"mode,omitempty"`
	Subnet         int    `json:"subnet,omitempty"`
	IPAddress      net.IP `json:"ip_address,omitempty"`
	Force          bool   `json:"force,omitempty"`
	DefaultGateway net.IP `json:"default_gateway,omitempty"`
}

// NetworkInterfacePhysical is the parameters for the NetworkInterfaces create_physical POST operation.
type NetworkInterfacePhysicalParams struct {
	Name       string   `json:"name,omitempty"`
	MACAddress string   `json:"mac_address,omitempty"`
	Tags       []string `json:"tags,omitempty"`
	VLAN       string   `json:"vlan,omitempty"`
	MTU        int      `json:"mtu,omitempty"`
	AcceptRA   bool     `json:"accept_ra,omitempty"`
	Autoconf   bool     `json:"autoconf,omitempty"`
}

// NetworkInterfaceBond is the parameters for the NetworkInterfaces create_bond POST operation.
type NetworkInterfaceBondParams struct {
	NetworkInterfacePhysicalParams
	Parents            []int  `json:"parents,omitempty"`
	BondMode           string `json:"bond_mode,omitempty"`
	BondMiimon         int    `json:"bond_miimon,omitempty"`
	BondDownDelay      int    `json:"bond_down_delay,omitempty"`
	BondUpDelay        int    `json:"bond_up_delay,omitempty"`
	BondLACPRate       string `json:"bond_lacp_rate,omitempty"`
	BondXMitHashPolicy string `json:"bond_x_mit_hash_policy,omitempty"`
	BondNumberGratARP  int    `json:"bond_number_grat_arp,omitempty"`
}

// NetworkInterfaceBridge is the parameters for the NetworkInterfaces create_bridge POST operation.
type NetworkInterfaceBridgeParams struct {
	NetworkInterfacePhysicalParams
	Parent    int  `json:"parent,omitempty"`
	BridgeSTP bool `json:"bridge_stp,omitempty"`
	BridgeFD  int  `json:"bridge_fd,omitempty"`
}

// NetworkInterfaceVLAN is the parameters for the NetworkInterfaces create_vlan POST operation.
type NetworkInterfaceVLANParams struct {
	Tags     []string `json:"tags,omitempty"`
	VLAN     string   `json:"vlan,omitempty"`
	Parent   int      `json:"parent,omitempty"`
	MTU      int      `json:"mtu,omitempty"`
	AcceptRA bool     `json:"accept_ra,omitempty"`
	Autoconf bool     `json:"autoconf,omitempty"`
}
