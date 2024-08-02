package entity

type DNSResource struct {
	FQDN            string              `json:"fqdn,omitempty"`
	ResourceURI     string              `json:"resource_uri,omitempty"`
	IPAddresses     []IPAddress         `json:"ip_addresses,omitempty"`
	ResourceRecords []DNSResourceRecord `json:"resource_records,omitempty"`
	ID              int                 `json:"id,omitempty"`
	AddressTTL      int                 `json:"address_ttl,omitempty"`
}

type DNSResourceParams struct {
	FQDN        string `url:"fqdn,omitempty"`
	Name        string `url:"name,omitempty"`
	Domain      string `url:"domain,omitempty"`
	IPAddresses string `url:"ip_addresses,omitempty"`
	AddressTTL  int    `url:"address_ttl"`
}

type DNSResourcesParams struct {
	Domain string `url:"domain,omitempty"`
	FQDN   string `url:"fqdn,omitempty"`
	Name   string `url:"name,omitempty"`
	RRType string `url:"rrtype,omitempty"`
	All    bool   `url:"all,omitempty"`
}
