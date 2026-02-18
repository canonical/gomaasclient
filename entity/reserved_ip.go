package entity

// ReservedIP represents the MAAS Reserved IP endpoint
type ReservedIP struct {
	Comment     string `json:"comment,omitempty"`
	IP          string `json:"ip,omitempty"`
	MACAddress  string `json:"mac_address,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
	Subnet      Subnet `json:"subnet,omitempty"`
	ID          int    `json:"id,omitempty"`
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
