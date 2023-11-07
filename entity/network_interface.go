package entity

// NetworkInterface represents the MaaS Interface endpoint.
type NetworkInterface struct {
	VLAN            VLAN                           `json:"vlan,omitempty"`
	Children        []string                       `json:"children,omitempty"`
	Parents         []string                       `json:"parents,omitempty"`
	Tags            []string                       `json:"tags,omitempty"`
	Discovered      []NetworkInterfaceDiscoveredIP `json:"discovered,omitempty"`
	Links           []NetworkInterfaceLink         `json:"links,omitempty"`
	Name            string                         `json:"name,omitempty"`
	MACAddress      string                         `json:"mac_address,omitempty"`
	Product         string                         `json:"product,omitempty"`
	FirmwareVersion string                         `json:"firmware_version,omitempty"`
	SystemID        string                         `json:"system_id,omitempty"`
	Params          interface{}                    `json:"params,omitempty"`
	Type            string                         `json:"type,omitempty"`
	Vendor          string                         `json:"vendor,omitempty"`
	ResourceURI     string                         `json:"resource_uri,omitempty"`
	EffectiveMTU    int                            `json:"effective_mtu,omitempty"`
	ID              int                            `json:"id,omitempty"`
	Enabled         bool                           `json:"enabled,omitempty"`
	LinkConnected   bool                           `json:"link_connected,omitempty"`
	InterfaceSpeed  int                            `json:"interface_speed,omitempty"`
	LinkSpeed       int                            `json:"link_speed,omitempty"`
	NUMANode        int                            `json:"numa_node,omitempty"`
	SRIOVMaxVF      int                            `json:"sriov_max_vf,omitempty"`
}

// NetworkInterfaceDiscoveredIP is consumed by NetworkInterface{} and should not be used directly.
type NetworkInterfaceDiscoveredIP struct {
	Subnet    Subnet `json:"subnet,omitempty"`
	IPAddress string `json:"ip_address,omitempty"`
}

// NetworkInterfaceLink is consumed by NetworkInterface{} and should not be used directly.
type NetworkInterfaceLink struct {
	ID        int    `json:"id,omitempty"`
	Mode      string `json:"mode,omitempty"`
	Subnet    Subnet `json:"subnet,omitempty"`
	IPAddress string `json:"ip_address,omitempty"`
}

// NetworkInterfaceLinkParams is used with NetworkInterface.LinkSubnet().
// Mode must be one of (AUTO, DHCP, STATIC, LINK_UP). IPAddress is ignored
// unless mode is STATIC, and will be set automatically if empty. Force
// allows LINK_UP to be set when other links exist, allows links between
// different VLANs, and deletes all other links on the interface.
// DefaultGateway is ignored unless Mode is AUTO or STATIC.
// Note: You can parse an IP address into a net.IP via net.ParseIP(string).
type NetworkInterfaceLinkParams struct {
	Mode           string `url:"mode,omitempty"`
	Subnet         int    `url:"subnet,omitempty"`
	Force          bool   `url:"force,omitempty"`
	DefaultGateway bool   `url:"default_gateway"`
	IPAddress      string `url:"ip_address,omitempty"`
}

// NetworkInterfaceParams is the common parameters for the NetworkInterfaces create_* POST operation.
type NetworkInterfaceParams struct {
	MTU            int      `url:"mtu,omitempty"`
	AcceptRA       bool     `url:"accept_ra,omitempty"`
	IPAssignment   string   `url:"ip_assignment,omitempty"`
	IPAddress      string   `url:"ip_address,omitempty"`
	LinkConnected  bool     `url:"link_connected,omitempty"`
	InterfaceSpeed int      `url:"interface_speed,omitempty"`
	LinkSpeed      int      `url:"link_speed,omitempty"`
	VLAN           VLAN     `url:"vlan,omitempty"`
	Tags           []string `url:"tags,omitempty"`
	MACAddress     string   `url:"mac_address,omitempty"`
	Name           string   `url:"name,omitempty"`
}

// NetworkInterfacePhysicalParams is the parameters for the NetworkInterfaces create_physical POST operation.
type NetworkInterfacePhysicalParams struct {
	NetworkInterfaceParams
	Enabled  bool `url:"enabled,omitempty"`
	NUMANode int  `url:"numa_node,omitempty"`
}

// NetworkInterfaceBondParams is the parameters for the NetworkInterfaces create_bond POST operation.
type NetworkInterfaceBondParams struct {
	NetworkInterfaceParams
	Parents            []int  `url:"parents,omitempty"`
	BondMode           string `url:"bond_mode,omitempty"`
	BondMiimon         int    `url:"bond_miimon,omitempty"`
	BondDownDelay      int    `url:"bond_downdelay,omitempty"`
	BondUpDelay        int    `url:"bond_updelay,omitempty"`
	BondLACPRate       string `url:"bond_lacp_rate,omitempty"`
	BondXMitHashPolicy string `url:"bond_xmit_hash_policy,omitempty"`
	BondNumberGratARP  int    `url:"bond_num_grat_arp,omitempty"`
}

// NetworkInterfaceBridgeParams is the parameters for the NetworkInterfaces create_bridge POST operation.
type NetworkInterfaceBridgeParams struct {
	NetworkInterfaceParams
	Parents    []int  `url:"parents,omitempty"`
	BridgeType string `url:"bridge_type,omitempty"`
	BridgeSTP  bool   `url:"bridge_stp,omitempty"`
	BridgeFD   int    `url:"bridge_fd,omitempty"`
}

// NetworkInterfaceVLANParams is the parameters for the NetworkInterfaces create_vlan POST operation.
type NetworkInterfaceVLANParams struct {
	NetworkInterfaceParams
}

// NetworkInterfaceUpdateParams is the parameters for the NetworkInterfaces update_bond POST operation.
type NetworkInterfaceUpdateParams struct {
	NetworkInterfaceBondParams
	NetworkInterfaceBridgeParams
	NetworkInterfaceVLANParams
}
