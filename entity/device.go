package entity

import "net"

type Device struct {
	OwnerData           interface{}        `json:"owner_data,omitempty"`
	Zone                Zone               `json:"zone,omitempty"`
	FQDN                string             `json:"fqdn,omitempty"`
	Hostname            string             `json:"hostname,omitempty"`
	SystemID            string             `json:"system_id,omitempty"`
	Owner               string             `json:"owner,omitempty"`
	Description         string             `json:"description,omitempty"`
	Parent              string             `json:"parent,omitempty"`
	ResourceURI         string             `json:"resource_uri,omitempty"`
	NodeTypeName        string             `json:"node_type_name,omitempty"`
	WorkloadAnnotations map[string]string  `json:"workload_annotations,omitempty"`
	IPAddresses         []net.IP           `json:"ip_addresses,omitempty"`
	InterfaceSet        []NetworkInterface `json:"interface_set,omitempty"`
	TagNames            []string           `json:"tag_names,omitempty"`
	Domain              Domain             `json:"domain,omitempty"`
	NodeType            int                `json:"node_type,omitempty"`
	AddressTTL          int                `json:"address_ttl,omitempty"`
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
