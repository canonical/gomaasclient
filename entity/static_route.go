package entity

type StaticRoute struct {
	ID          int    `json:"id,omitempty"`
	Source      string `json:"source,omitempty"`
	Destination string `json:"destination,omitempty"`
	Gateway_ip  string `json:"gateway_ip,omitempty"`
	Metric      string `json:"metric,omitempty"`
}

type StaticRouteParams struct {
	Source      string `json:"source,omitempty"`
	Destination string `json:"destination,omitempty"`
	Gateway_ip  string `json:"gateway_ip,omitempty"`
	Metric      string `json:"metric,omitempty"`
}
