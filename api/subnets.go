package api

import (
	"github.com/ionutbalutoiu/gomaasclient/entity"
)

// Subnets represents the MaaS Subnets endpoint
type Subnets interface {
	Get() ([]entity.Subnet, error)
	Post(*entity.SubnetParams) (*entity.Subnet, error)
}
