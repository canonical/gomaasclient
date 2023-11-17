package entity

import "net"

type Device struct {
	SystemID     string             `json:"system_id,omitempty"`
	Hostname     string             `json:"hostname,omitempty"`
	Description  string             `json:"description,omitempty"`
	Domain       Domain             `json:"domain,omitempty"`
	FQDN         string             `json:"fqdn,omitempty"`
	Owner        string             `json:"owner,omitempty"`
	OwnerData    interface{}        `json:"owner_data,omitempty"`
	Parent       string             `json:"parent,omitempty"`
	TagNames     []string           `json:"tag_names,omitempty"`
	AddressTTL   int                `json:"address_ttl,omitempty"`
	IPAddresses  []net.IP           `json:"ip_addresses,omitempty"`
	InterfaceSet []NetworkInterface `json:"interface_set,omitempty"`
	Zone         Zone               `json:"zone,omitempty"`
	NodeType     int                `json:"node_type,omitempty"`
	NodeTypeName string             `json:"node_type_name,omitempty"`
	ResourceURI  string             `json:"resource_uri,omitempty"`
}

type DeviceCreateParams struct {
	Hostname     string   `url:"hostname,omitempty"`
	Description  string   `url:"description,omitempty"`
	Domain       string   `url:"domain,omitempty"`
	Parent       string   `url:"parent,omitempty"`
	Zone         string   `url:"zone,omitempty"`
	MacAddresses []string `url:"mac_addresses,omitempty"`
}

type DeviceUpdateParams struct {
	Hostname    string `url:"hostname,omitempty"`
	Description string `url:"description,omitempty"`
	Domain      string `url:"domain,omitempty"`
	Parent      string `url:"parent,omitempty"`
	Zone        string `url:"zone,omitempty"`
}
