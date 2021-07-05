package entity

type DNSResource struct {
	ID              int                 `json:"id,omitempty"`
	AddressTTL      int                 `json:"address_ttl,omitempty"`
	FQDN            string              `json:"fqdn,omitempty"`
	ResourceURI     string              `json:"resource_uri,omitempty"`
	IPAddresses     []IPAddress         `json:"ip_addresses,omitempty"`
	ResourceRecords []DNSResourceRecord `json:"resource_records,omitempty"`
}

type DNSResourceRecord struct {
	ID          int    `json:"id,omitempty"`
	TTL         int    `json:"ttl,omitempty"`
	RRType      string `json:"rrtype,omitempty"`
	RRData      string `json:"rrdata,omitempty"`
	FQDN        string `json:"fqdn,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

type DNSResourceParams struct {
	AddressTTL  int    `url:"address_ttl"`
	FQDN        string `url:"fqdn,omitempty"`
	Name        string `url:"name,omitempty"`
	Domain      string `url:"domain,omitempty"`
	IPAddresses string `url:"ip_addresses,omitempty"`
}
