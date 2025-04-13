package entity

type StaticRoute struct {
	ID          int    `json:"id,omitempty"`
	Source      Subnet `json:"source,omitempty"`
	Destination Subnet `json:"destination,omitempty"`
	GatewayIP   string `json:"gateway_ip,omitempty"`
	Metric      int    `json:"metric,omitempty"`
}

type StaticRouteParams struct {
	Source      string `url:"source,omitempty"`
	Destination string `url:"destination,omitempty"`
	GatewayIP   string `url:"gateway_ip,omitempty"`
	Metric      int    `url:"metric,omitempty"`
}
