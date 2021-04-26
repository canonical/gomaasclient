package api

import (
	"github.com/ionutbalutoiu/gomaasclient/api/endpoint"
)

// Subnets represents the MaaS Subnets endpoint
type Subnets interface {
	Get() ([]endpoint.Subnet, error)
	Post(*endpoint.SubnetParams) (*endpoint.Subnet, error)
}
