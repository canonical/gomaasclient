package api

import (
	"github.com/ionutbalutoiu/gomaasclient/api/params"
	"github.com/ionutbalutoiu/gomaasclient/maas/entity"
)

// Subnets represents the MaaS Subnets endpoint
type Subnets interface {
	Get() ([]entity.Subnet, error)
	Post(*params.Subnet) (*entity.Subnet, error)
}
