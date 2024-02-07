package entity

// Discovery represents device discovery in MAAS
type Discovery struct {
	DiscoveryID     string   `json:"discovery_id,omitempty"`
	IP              string   `json:"ip,omitempty"`
	MACAddress      string   `json:"mac_address,omitempty"`
	LastSeen        string   `json:"last_seen,omitempty"`
	Hostname        string   `json:"hostname,omitempty"`
	FabricName      string   `json:"fabric_name,omitempty"`
	Vid             string   `json:"vid,omitempty"`
	MacOrganization string   `json:"mac_organization,omitempty"`
	ResourceURI     string   `json:"resource_uri,omitempty"`
	Observer        Observer `json:"observer,omitempty"`
}

type Observer struct {
	SystemID      string `json:"system_id,omitempty"`
	Hostname      string `json:"hostname,omitempty"`
	InterfaceName string `json:"interface_name,omitempty"`
	InterfaceID   int    `json:"interface_id,omitempty"`
}

type DiscoveryClearParams struct {
	All        bool `url:"all,omitempty"`
	MDNS       bool `url:"mdns,omitempty"`
	Neighbours bool `url:"neighbours,omitempty"`
}
