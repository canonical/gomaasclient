package entity

import (
	"net"
)

type IPRange struct {
	Type        string `json:"type"`
	Comment     string `json:"comment,omitempty"`
	ResourceURI string `json:"resource_uri"`
	User        User   `json:"user,omitempty"`
	StartIP     net.IP `json:"start_ip"`
	EndIP       net.IP `json:"end_ip"`
	Subnet      Subnet `json:"subnet"`
	ID          int    `json:"id"`
}

type IPRangeParams struct {
	Type    string `url:"type"`
	Subnet  string `url:"subnet"`
	StartIP string `url:"start_ip"`
	EndIP   string `url:"end_ip"`
	Comment string `url:"comment,omitempty"`
}
