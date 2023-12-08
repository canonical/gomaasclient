package entity

import "net"

type IPAddress struct {
	AllocTypeName string             `json:"alloc_type_name,omitempty"`
	Created       string             `json:"created,omitempty"`
	ResourceURI   string             `json:"resource_uri,omitempty"`
	Owner         User               `json:"owner,omitempty"`
	IP            net.IP             `json:"ip,omitempty"`
	InterfaceSet  []NetworkInterface `json:"interface_set,omitempty"`
	Subnet        Subnet             `json:"subnet,omitempty"`
	AllocType     int                `json:"alloc_type,omitempty"`
}

type IPAddressesParams struct {
	Subnet     string `url:"subnet,omitempty"`
	IP         string `url:"ip,omitempty"`
	Owner      string `url:"owner,omitempty"`
	All        bool   `url:"all,omitempty"`
	Discovered bool   `url:"discovered,omitempty"`
	Force      bool   `url:"force,omitempty"`
}
