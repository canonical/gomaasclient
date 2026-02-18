package entity

// ReservedIP represents the MAAS Reserved IP endpoint
type ReservedIP struct {
	ID          int    `json:"id,omitempty"`
	Comment     string `json:"comment,omitempty"`
	IP          string `json:"ip,omitempty"`
	MACAddress  string `json:"mac_address,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
	Subnet      Subnet `json:"subnet,omitempty"`
}

type ReservedIPCreateParams struct {
	IP         string `url:"ip"`
	MACAddress string `url:"mac_address,omitempty"`
	Comment    string `url:"comment,omitempty"`
	Subnet     int    `url:"subnet,omitempty"`
}

type ReservedIPUpdateParams struct {
	Comment string `url:"comment"`
}
