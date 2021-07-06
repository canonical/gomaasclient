package entity

import "net"

type IPAddress struct {
	AllocType     int                `json:"alloc_type,omitempty"`
	AllocTypeName string             `json:"alloc_type_name,omitempty"`
	Created       string             `json:"created,omitempty"`
	ResourceURI   string             `json:"resource_uri,omitempty"`
	IP            net.IP             `json:"ip,omitempty"`
	Subnet        Subnet             `json:"subnet,omitempty"`
	InterfaceSet  []NetworkInterface `json:"interface_set,omitempty"`
	Owner         User               `json:"owner,omitempty"`
}

type IPAddressesParams struct {
	All        bool   `url:"all,omitempty"`
	Discovered bool   `url:"discovered,omitempty"`
	Force      bool   `url:"force,omitempty"`
	Subnet     string `url:"subnet,omitempty"`
	IP         string `url:"ip,omitempty"`
	Owner      string `url:"owner,omitempty"`
}
