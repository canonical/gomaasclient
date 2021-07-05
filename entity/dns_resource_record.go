package entity

type DNSResourceRecord struct {
	ID          int    `json:"id,omitempty"`
	TTL         int    `json:"ttl,omitempty"`
	RRType      string `json:"rrtype,omitempty"`
	RRData      string `json:"rrdata,omitempty"`
	FQDN        string `json:"fqdn,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

type DNSResourceRecordParams struct {
	FQDN   string `url:"fqdn,omitempty"`
	Name   string `url:"name,omitempty"`
	Domain string `url:"domain,omitempty"`
	RRType string `url:"rrtype,omitempty"`
	RRData string `url:"rrdata,omitempty"`
	TTL    int    `url:"ttl"`
}
