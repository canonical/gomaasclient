package entity

type StaticRoute struct {
	GatewayIP   string `json:"gateway_ip,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
	Destination Subnet `json:"destination,omitempty"`
	Source      Subnet `json:"source,omitempty"`
	ID          int    `json:"id,omitempty"`
	Metric      int    `json:"metric,omitempty"`
}

type StaticRouteParams struct {
	Source      string `url:"source,omitempty"`
	Destination string `url:"destination,omitempty"`
	GatewayIP   string `url:"gateway_ip,omitempty"`
	Metric      int    `url:"metric,omitempty"`
}
