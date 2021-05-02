package entity

import (
	"net"
)

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
	Mode           string `json:"mode"`
	Subnet         int    `json:"subnet"`
	IPAddress      net.IP `json:"ip_address"`
	Force          bool   `json:"force"`
	DefaultGateway net.IP `json:"default_gateway"`
}

// NetworkInterfacePhysical is the parameters for the NetworkInterfaces create_physical POST operation.
type NetworkInterfacePhysicalParams struct {
	Name       string `json:"name"`
	MACAddress string `json:"mac_address"`
	Tags       string `json:"tags"`
	VLAN       string `json:"vlan"`
	MTU        int    `json:"mtu"`
	AcceptRA   bool   `json:"accept_ra"`
	Autoconf   bool   `json:"autoconf"`
}

// NetworkInterfaceBond is the parameters for the NetworkInterfaces create_bond POST operation.
type NetworkInterfaceBondParams struct {
	NetworkInterfacePhysicalParams
	Parents            []int  `json:"parents"`
	BondMode           string `json:"bond_mode"`
	BondMiimon         int    `json:"bond_miimon"`
	BondDownDelay      int    `json:"bond_downdelay"`
	BondUpDelay        int    `json:"bond_updelay"`
	BondLACPRate       string `json:"bond_lacp_rate"`
	BondXMitHashPolicy string `json:"bond_xmit_hash_policy"`
	BondNumberGratARP  int    `json:"bond_num_grat_arp"`
}

// NetworkInterfaceBridge is the parameters for the NetworkInterfaces create_bridge POST operation.
type NetworkInterfaceBridgeParams struct {
	NetworkInterfacePhysicalParams
	Parent     int    `json:"parent"`
	Bridgetype string `json:"bridge_type"`
	BridgeSTP  bool   `json:"bridge_stp"`
	BridgeFD   int    `json:"bridge_fd"`
}

// NetworkInterfaceVLAN is the parameters for the NetworkInterfaces create_vlan POST operation.
type NetworkInterfaceVLANParams struct {
	Tags     []string `json:"tags"`
	VLAN     string   `json:"vlan"`
	Parent   int      `json:"parent"`
	MTU      int      `json:"mtu"`
	AcceptRA bool     `json:"accept_ra"`
	Autoconf bool     `json:"autoconf"`
}
