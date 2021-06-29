package entity

import (
	"net"
)

type IPRange struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	StartIP     net.IP `json:"start_ip"`
	EndIP       net.IP `json:"end_ip"`
	Subnet      Subnet `json:"subnet"`
	User        User   `json:"user,omitempty"`
	Comment     string `json:"comment,omitempty"`
	ResourceURI string `json:"resource_uri"`
}

type IPRangeParams struct {
	Type    string `url:"type"`
	Subnet  string `url:"subnet"`
	StartIP string `url:"start_ip"`
	EndIP   string `url:"end_ip"`
	Comment string `url:"comment,omitempty"`
}
