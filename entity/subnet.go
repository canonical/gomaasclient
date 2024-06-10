package entity

import (
	"net"
)

// Subnet represents the MAAS Subnet endpoint.
type Subnet struct {
	Space                     string   `json:"space,omitempty"`
	CIDR                      string   `json:"cidr,omitempty"`
	Name                      string   `json:"name,omitempty"`
	ResourceURI               string   `json:"resource_uri,omitempty"`
	Description               string   `json:"description,omitempty"`
	GatewayIP                 net.IP   `json:"gateway_ip,omitempty"`
	DNSServers                []net.IP `json:"dns_servers,omitempty"`
	DisabledBootArchitectures []string `json:"disabled_boot_architectures,omitempty"`
	VLAN                      VLAN     `json:"vlan,omitempty"`
	ID                        int      `json:"id,omitempty"`
	RDNSMode                  int      `json:"rdns_mode,omitempty"`
	AllowDNS                  bool     `json:"allow_dns,omitempty"`
	Managed                   bool     `json:"managed,omitempty"`
	ActiveDiscovery           bool     `json:"active_discovery,omitempty"`
	AllowProxy                bool     `json:"allow_proxy,omitempty"`
}

// SubnetParams contains the parameters for the POST operation on the Subnets endpoint.
type SubnetParams struct {
	CIDR        string   `url:"cidr"`
	Name        string   `url:"name,omitempty"`
	Description string   `url:"description,omitempty"`
	VLAN        string   `url:"vlan,omitempty"`
	Fabric      string   `url:"fabric,omitempty"`
	GatewayIP   string   `url:"gateway_ip,omitempty"`
	DNSServers  []string `url:"dns_servers,omitempty"`
	VID         int      `url:"vid,omitempty"`
	RDNSMode    int      `url:"rdns_mode"`
	AllowDNS    bool     `url:"allow_dns"`
	AllowProxy  bool     `url:"allow_proxy"`
	Managed     bool     `url:"managed"`
}
