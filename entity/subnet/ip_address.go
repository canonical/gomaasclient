package subnet

import (
	"net"
)

// IPAddress represents an IP address from a Subnet's GetIPAddresses()
type IPAddress struct {
	Created     string      `json:"created,omitempty"`
	Updated     string      `json:"updated,omitempty"`
	User        string      `json:"user,omitempty"`
	IP          net.IP      `json:"ip,omitempty"`
	NodeSummary NodeSummary `json:"node_summary,omitempty"`
	AllocType   int         `json:"alloc_type,omitempty"`
}

// NodeSummary represents the optional node_summary from GetIPAddresses().
// This type should not be used directly.
type NodeSummary struct {
	SystemID    string `json:"system_id,omitempty"`
	FQDN        string `json:"fqdn,omitempty"`
	Hostname    string `json:"hostname,omitempty"`
	Via         string `json:"via,omitempty"`
	NodeType    int    `json:"node_type,omitempty"`
	IsContainer bool   `json:"is_container,omitempty"`
}
