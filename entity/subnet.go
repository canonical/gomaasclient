package entity

import (
	"net"
)

// Subnet represents the MAAS Subnet endpoint.
type Subnet struct {
	Name            string   `json:"name,omitempty"`
	VLAN            VLAN     `json:"vlan,omitempty"`
	CIDR            string   `json:"cidr,omitempty"`
	RDNSMode        int      `json:"rdns_mode,omitempty"`
	GatewayIP       net.IP   `json:"gateway_ip,omitempty"`
	DNSServers      []net.IP `json:"dns_servers,omitempty"`
	AllowDNS        bool     `json:"allow_dns,omitempty"`
	AllowProxy      bool     `json:"allow_proxy,omitempty"`
	ActiveDiscovery bool     `json:"active_discovery,omitempty"`
	Managed         bool     `json:"managed,omitempty"`
	ID              int      `json:"id,omitempty"`
	Space           string   `json:"space,omitempty"`
	ResourceURI     string   `json:"resource_uri,omitempty"`
}

// SubnetParams contains the parameters for the POST operation on the Subnets endpoint.
type SubnetParams struct {
	CIDR        string   `url:"cidr"`
	Name        string   `url:"name,omitempty"`
	Description string   `url:"description,omitempty"`
	VLAN        string   `url:"vlan,omitempty"`
	Fabric      string   `url:"fabric,omitempty"`
	VID         int      `url:"vid,omitempty"`
	GatewayIP   string   `url:"gateway_ip,omitempty"`
	DNSServers  []string `url:"dns_servers,omitempty"`
	RDNSMode    int      `url:"rdns_mode"`
	AllowDNS    bool     `url:"allow_dns"`
	AllowProxy  bool     `url:"allow_proxy"`
	Managed     bool     `url:"managed"`
}
