package entity

// NetworkInterface represents the MAAS Interface endpoint.
type NetworkInterface struct {
	Params          interface{}                    `json:"params,omitempty"`
	Name            string                         `json:"name,omitempty"`
	Product         string                         `json:"product,omitempty"`
	MACAddress      string                         `json:"mac_address,omitempty"`
	FirmwareVersion string                         `json:"firmware_version,omitempty"`
	Vendor          string                         `json:"vendor,omitempty"`
	Type            string                         `json:"type,omitempty"`
	SystemID        string                         `json:"system_id,omitempty"`
	ResourceURI     string                         `json:"resource_uri,omitempty"`
	Links           []NetworkInterfaceLink         `json:"links,omitempty"`
	Parents         []string                       `json:"parents,omitempty"`
	Children        []string                       `json:"children,omitempty"`
	Tags            []string                       `json:"tags,omitempty"`
	Discovered      []NetworkInterfaceDiscoveredIP `json:"discovered,omitempty"`
	VLAN            VLAN                           `json:"vlan,omitempty"`
	NUMANode        int                            `json:"numa_node,omitempty"`
	LinkSpeed       int                            `json:"link_speed,omitempty"`
	EffectiveMTU    int                            `json:"effective_mtu,omitempty"`
	SRIOVMaxVF      int                            `json:"sriov_max_vf,omitempty"`
	InterfaceSpeed  int                            `json:"interface_speed,omitempty"`
	ID              int                            `json:"id,omitempty"`
	Enabled         bool                           `json:"enabled,omitempty"`
	LinkConnected   bool                           `json:"link_connected,omitempty"`
}

// NetworkInterfaceDiscoveredIP is consumed by NetworkInterface{} and should not be used directly.
type NetworkInterfaceDiscoveredIP struct {
	IPAddress string `json:"ip_address,omitempty"`
	Subnet    Subnet `json:"subnet,omitempty"`
}

// NetworkInterfaceLink is consumed by NetworkInterface{} and should not be used directly.
type NetworkInterfaceLink struct {
	IPAddress string `json:"ip_address,omitempty"`
	Mode      string `json:"mode,omitempty"`
	Subnet    Subnet `json:"subnet,omitempty"`
	ID        int    `json:"id,omitempty"`
}

// NetworkInterfaceLinkParams is used with NetworkInterface.LinkSubnet().
// Mode must be one of (AUTO, DHCP, STATIC, LINK_UP). IPAddress is ignored
// unless mode is STATIC, and will be set automatically if empty. Force
// allows LINK_UP to be set when other links exist, allows links between
// different VLANs, and deletes all other links on the interface.
// DefaultGateway is ignored unless Mode is AUTO or STATIC.
// Note: You can parse an IP address into a net.IP via net.ParseIP(string).
type NetworkInterfaceLinkParams struct {
	IPAddress      string `url:"ip_address,omitempty"`
	Mode           string `url:"mode,omitempty"`
	Subnet         int    `url:"subnet,omitempty"`
	DefaultGateway bool   `url:"default_gateway"`
	Force          bool   `url:"force,omitempty"`
}

// NetworkInterfacePhysicalParams is the parameters for the NetworkInterfaces create_physical POST operation.
type NetworkInterfacePhysicalParams struct {
	Name           string `url:"name,omitempty"`
	IPAddress      string `url:"ip_address,omitempty"`
	IPAssignment   string `url:"ip_assignment,omitempty"`
	MACAddress     string `url:"mac_address,omitempty"`
	Tags           string `url:"tags,omitempty"`
	InterfaceSpeed int    `url:"interface_speed,omitempty"`
	LinkSpeed      int    `url:"link_speed,omitempty"`
	MTU            int    `url:"mtu,omitempty"`
	NUMANode       int    `url:"numa_node,omitempty"`
	VLAN           int    `url:"vlan,omitempty"`
	Enabled        bool   `url:"enabled,omitempty"`
	LinkConnected  bool   `url:"link_connected,omitempty"`
	AcceptRA       bool   `url:"accept_ra,omitempty"`
}

// NetworkInterfaceBondParams is the parameters for the NetworkInterfaces create_bond POST operation.
type NetworkInterfaceBondParams struct {
	IPAddress          string `url:"ip_address,omitempty"`
	Tags               string `url:"tags,omitempty"`
	BondLACPRate       string `url:"bond_lacp_rate,omitempty"`
	BondMode           string `url:"bond_mode,omitempty"`
	Name               string `url:"name,omitempty"`
	MACAddress         string `url:"mac_address,omitempty"`
	BondXMitHashPolicy string `url:"bond_xmit_hash_policy,omitempty"`
	IPAssignment       string `url:"ip_assignment,omitempty"`
	Parents            []int  `url:"parents,omitempty"`
	InterfaceSpeed     int    `url:"interface_speed,omitempty"`
	LinkSpeed          int    `url:"link_speed,omitempty"`
	BondUpDelay        int    `url:"bond_updelay,omitempty"`
	MTU                int    `url:"mtu,omitempty"`
	BondNumberGratARP  int    `url:"bond_num_grat_arp,omitempty"`
	BondMiimon         int    `url:"bond_miimon,omitempty"`
	BondDownDelay      int    `url:"bond_downdelay,omitempty"`
	VLAN               int    `url:"vlan,omitempty"`
	AcceptRA           bool   `url:"accept_ra,omitempty"`
	LinkConnected      bool   `url:"link_connected,omitempty"`
}

// NetworkInterfaceBridgeParams is the parameters for the NetworkInterfaces create_bridge POST operation.
type NetworkInterfaceBridgeParams struct {
	IPAssignment   string `url:"ip_assignment,omitempty"`
	MACAddress     string `url:"mac_address,omitempty"`
	Tags           string `url:"tags,omitempty"`
	BridgeType     string `url:"bridge_type,omitempty"`
	IPAddress      string `url:"ip_address,omitempty"`
	Name           string `url:"name,omitempty"`
	Parents        []int  `url:"parents,omitempty"`
	LinkSpeed      int    `url:"link_speed,omitempty"`
	MTU            int    `url:"mtu,omitempty"`
	BridgeFD       int    `url:"bridge_fd,omitempty"`
	InterfaceSpeed int    `url:"interface_speed,omitempty"`
	VLAN           int    `url:"vlan,omitempty"`
	AcceptRA       bool   `url:"accept_ra,omitempty"`
	LinkConnected  bool   `url:"link_connected,omitempty"`
	BridgeSTP      bool   `url:"bridge_stp,omitempty"`
}

// NetworkInterfaceVLANParams is the parameters for the NetworkInterfaces create_vlan POST operation.
type NetworkInterfaceVLANParams struct {
	IPAddress      string `url:"ip_address,omitempty"`
	IPAssignment   string `url:"ip_assignment,omitempty"`
	Name           string `url:"name,omitempty"`
	Tags           string `url:"tags,omitempty"`
	Parents        []int  `url:"parents,omitempty"`
	InterfaceSpeed int    `url:"interface_speed,omitempty"`
	LinkSpeed      int    `url:"link_speed,omitempty"`
	MTU            int    `url:"mtu,omitempty"`
	VLAN           int    `url:"vlan,omitempty"`
	AcceptRA       bool   `url:"accept_ra,omitempty"`
	LinkConnected  bool   `url:"link_connected,omitempty"`
}

// NetworkInterfaceUpdateParams is the parameters for the NetworkInterfaces update_bond POST operation.
type NetworkInterfaceUpdateParams struct {
	BridgeType         string `url:"bridge_type,omitempty"`
	Tags               string `url:"tags,omitempty"`
	BondLACPRate       string `url:"bond_lacp_rate,omitempty"`
	BondMode           string `url:"bond_mode,omitempty"`
	Name               string `url:"name,omitempty"`
	MACAddress         string `url:"mac_address,omitempty"`
	BondXMitHashPolicy string `url:"bond_xmit_hash_policy,omitempty"`
	IPAssignment       string `url:"ip_assignment,omitempty"`
	IPAddress          string `url:"ip_address,omitempty"`
	Parents            []int  `url:"parents,omitempty"`
	InterfaceSpeed     int    `url:"interface_speed,omitempty"`
	BondUpDelay        int    `url:"bond_updelay,omitempty"`
	VLAN               int    `url:"vlan,omitempty"`
	BondDownDelay      int    `url:"bond_downdelay,omitempty"`
	BridgeFD           int    `url:"bridge_fd,omitempty"`
	BondMiimon         int    `url:"bond_miimon,omitempty"`
	LinkSpeed          int    `url:"link_speed,omitempty"`
	NUMANode           int    `url:"numa_node,omitempty"`
	MTU                int    `url:"mtu,omitempty"`
	BondNumberGratARP  int    `url:"bond_num_grat_arp,omitempty"`
	Enabled            bool   `url:"enabled,omitempty"`
	LinkConnected      bool   `url:"link_connected,omitempty"`
	BridgeSTP          bool   `url:"bridge_stp,omitempty"`
	AcceptRA           bool   `url:"accept_ra,omitempty"`
}
