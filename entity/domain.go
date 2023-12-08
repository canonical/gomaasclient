package entity

// Domain represents the MAAS Domain endpoint
type Domain struct {
	Name                string `json:"name,omitempty"`
	ResourceURI         string `json:"resource_uri,omitempty"`
	ID                  int    `json:"id,omitempty"`
	TTL                 int    `json:"ttl,omitempty"`
	ResourceRecordCount int    `json:"resource_record_count,omitempty"`
	Authoritative       bool   `json:"authoritative,omitempty"`
	IsDefault           bool   `json:"is_default,omitempty"`
}

type DomainParams struct {
	Name          string `url:"name,omitempty"`
	TTL           int    `url:"ttl,omitempty"`
	Authoritative bool   `url:"authoritative"`
}
