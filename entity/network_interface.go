package entity

// NetworkInterface represents the MaaS Interface endpoint.
type NetworkInterface struct {
	Children        []string                       `json:"children,omitempty"`
	Discovered      []NetworkInterfaceDiscoveredIP `json:"discovered,omitempty"`
	EffectiveMTU    int                            `json:"effective_mtu,omitempty"`
	Enabled         bool                           `json:"enabled,omitempty"`
	FirmwareVersion string                         `json:"firmware_version,omitempty"`
	ID              int                            `json:"id,omitempty"`
	InterfaceSpeed  int                            `json:"interface_speed,omitempty"`
	LinkConnected   bool                           `json:"link_connected,omitempty"`
	Links           []NetworkInterfaceLink         `json:"links,omitempty"`
	LinkSpeed       int                            `json:"link_speed,omitempty"`
	MACAddress      string                         `json:"mac_address,omitempty"`
	Name            string                         `json:"name,omitempty"`
	NUMANode        int                            `json:"numa_node,omitempty"`
	Params          interface{}                    `json:"params,omitempty"`
	Parents         []string                       `json:"parents,omitempty"`
	Product         string                         `json:"product,omitempty"`
	ResourceURI     string                         `json:"resource_uri,omitempty"`
	SRIOVMaxVF      int                            `json:"sriov_max_vf,omitempty"`
	SystemID        string                         `json:"system_id,omitempty"`
	Tags            []string                       `json:"tags,omitempty"`
	Type            string                         `json:"type,omitempty"`
	Vendor          string                         `json:"vendor,omitempty"`
	VLAN            VLAN                           `json:"vlan,omitempty"`
}

// NetworkInterfaceDiscoveredIP is consumed by NetworkInterface{} and should not be used directly.
type NetworkInterfaceDiscoveredIP struct {
	IPAddress string `json:"ip_address,omitempty"`
	Subnet    Subnet `json:"subnet,omitempty"`
}

// NetworkInterfaceLink is consumed by NetworkInterface{} and should not be used directly.
type NetworkInterfaceLink struct {
	ID        int    `json:"id,omitempty"`
	IPAddress string `json:"ip_address,omitempty"`
	Mode      string `json:"mode,omitempty"`
	Subnet    Subnet `json:"subnet,omitempty"`
}

// NetworkInterfaceLinkParams is used with NetworkInterface.LinkSubnet().
// Mode must be one of (AUTO, DHCP, STATIC, LINK_UP). IPAddress is ignored
// unless mode is STATIC, and will be set automatically if empty. Force
// allows LINK_UP to be set when other links exist, allows links between
// different VLANs, and deletes all other links on the interface.
// DefaultGateway is ignored unless Mode is AUTO or STATIC.
// Note: You can parse an IP address into a net.IP via net.ParseIP(string).
type NetworkInterfaceLinkParams struct {
	DefaultGateway bool   `url:"default_gateway"`
	Force          bool   `url:"force,omitempty"`
	IPAddress      string `url:"ip_address,omitempty"`
	Mode           string `url:"mode,omitempty"`
	Subnet         int    `url:"subnet,omitempty"`
}

// NetworkInterfacePhysicalParams is the parameters for the NetworkInterfaces create_physical POST operation.
type NetworkInterfacePhysicalParams struct {
	AcceptRA       bool   `url:"accept_ra,omitempty"`
	Enabled        bool   `url:"enabled,omitempty"`
	InterfaceSpeed int    `url:"interface_speed,omitempty"`
	IPAddress      string `url:"ip_address,omitempty"`
	IPAssignment   string `url:"ip_assignment,omitempty"`
	LinkConnected  bool   `url:"link_connected,omitempty"`
	LinkSpeed      int    `url:"link_speed,omitempty"`
	MACAddress     string `url:"mac_address,omitempty"`
	MTU            int    `url:"mtu,omitempty"`
	Name           string `url:"name,omitempty"`
	NUMANode       int    `url:"numa_node,omitempty"`
	VLAN           int    `url:"vlan,omitempty"`
	Tags           string `url:"tags,omitempty"`
}

// NetworkInterfaceBondParams is the parameters for the NetworkInterfaces create_bond POST operation.
type NetworkInterfaceBondParams struct {
	AcceptRA           bool   `url:"accept_ra,omitempty"`
	BondDownDelay      int    `url:"bond_downdelay,omitempty"`
	BondLACPRate       string `url:"bond_lacp_rate,omitempty"`
	BondMiimon         int    `url:"bond_miimon,omitempty"`
	BondMode           string `url:"bond_mode,omitempty"`
	BondNumberGratARP  int    `url:"bond_num_grat_arp,omitempty"`
	BondUpDelay        int    `url:"bond_updelay,omitempty"`
	BondXMitHashPolicy string `url:"bond_xmit_hash_policy,omitempty"`
	InterfaceSpeed     int    `url:"interface_speed,omitempty"`
	IPAddress          string `url:"ip_address,omitempty"`
	IPAssignment       string `url:"ip_assignment,omitempty"`
	LinkConnected      bool   `url:"link_connected,omitempty"`
	LinkSpeed          int    `url:"link_speed,omitempty"`
	MACAddress         string `url:"mac_address,omitempty"`
	MTU                int    `url:"mtu,omitempty"`
	Name               string `url:"name,omitempty"`
	Parents            []int  `url:"parents,omitempty"`
	Tags               string `url:"tags,omitempty"`
	VLAN               int    `url:"vlan,omitempty"`
}

// NetworkInterfaceBridgeParams is the parameters for the NetworkInterfaces create_bridge POST operation.
type NetworkInterfaceBridgeParams struct {
	AcceptRA       bool   `url:"accept_ra,omitempty"`
	BridgeFD       int    `url:"bridge_fd,omitempty"`
	BridgeSTP      bool   `url:"bridge_stp,omitempty"`
	BridgeType     string `url:"bridge_type,omitempty"`
	InterfaceSpeed int    `url:"interface_speed,omitempty"`
	IPAddress      string `url:"ip_address,omitempty"`
	IPAssignment   string `url:"ip_assignment,omitempty"`
	LinkConnected  bool   `url:"link_connected,omitempty"`
	LinkSpeed      int    `url:"link_speed,omitempty"`
	MACAddress     string `url:"mac_address,omitempty"`
	MTU            int    `url:"mtu,omitempty"`
	Name           string `url:"name,omitempty"`
	Parents        []int  `url:"parents,omitempty"`
	Tags           string `url:"tags,omitempty"`
	VLAN           int    `url:"vlan,omitempty"`
}

// NetworkInterfaceVLANParams is the parameters for the NetworkInterfaces create_vlan POST operation.
type NetworkInterfaceVLANParams struct {
	AcceptRA       bool   `url:"accept_ra,omitempty"`
	InterfaceSpeed int    `url:"interface_speed,omitempty"`
	IPAddress      string `url:"ip_address,omitempty"`
	IPAssignment   string `url:"ip_assignment,omitempty"`
	LinkConnected  bool   `url:"link_connected,omitempty"`
	LinkSpeed      int    `url:"link_speed,omitempty"`
	MTU            int    `url:"mtu,omitempty"`
	Name           string `url:"name,omitempty"`
	Parents        []int  `url:"parents,omitempty"`
	Tags           string `url:"tags,omitempty"`
	VLAN           int    `url:"vlan,omitempty"`
}

// NetworkInterfaceUpdateParams is the parameters for the NetworkInterfaces update_bond POST operation.
type NetworkInterfaceUpdateParams struct {
	AcceptRA           bool   `url:"accept_ra,omitempty"`
	BondDownDelay      int    `url:"bond_downdelay,omitempty"`
	BondLACPRate       string `url:"bond_lacp_rate,omitempty"`
	BondMiimon         int    `url:"bond_miimon,omitempty"`
	BondMode           string `url:"bond_mode,omitempty"`
	BondNumberGratARP  int    `url:"bond_num_grat_arp,omitempty"`
	BondUpDelay        int    `url:"bond_updelay,omitempty"`
	BondXMitHashPolicy string `url:"bond_xmit_hash_policy,omitempty"`
	BridgeFD           int    `url:"bridge_fd,omitempty"`
	BridgeSTP          bool   `url:"bridge_stp,omitempty"`
	BridgeType         string `url:"bridge_type,omitempty"`
	Enabled            bool   `url:"enabled,omitempty"`
	InterfaceSpeed     int    `url:"interface_speed,omitempty"`
	IPAddress          string `url:"ip_address,omitempty"`
	IPAssignment       string `url:"ip_assignment,omitempty"`
	LinkConnected      bool   `url:"link_connected,omitempty"`
	LinkSpeed          int    `url:"link_speed,omitempty"`
	MACAddress         string `url:"mac_address,omitempty"`
	MTU                int    `url:"mtu,omitempty"`
	Name               string `url:"name,omitempty"`
	NUMANode           int    `url:"numa_node,omitempty"`
	Parents            []int  `url:"parents,omitempty"`
	Tags               string `url:"tags,omitempty"`
	VLAN               int    `url:"vlan,omitempty"`
}
