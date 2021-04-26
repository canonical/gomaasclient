package api

import (
	"github.com/ionutbalutoiu/gomaasclient/api/endpoint"
)

// VLANs represents the MaaS Vlans endpoint
type VLANs interface {
	Get(fabricID int) ([]endpoint.VLAN, error)
	Post(fabricID int, params *endpoint.VLANParams) (*endpoint.VLAN, error)
}
