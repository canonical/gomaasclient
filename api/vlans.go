package api

import (
	"github.com/ionutbalutoiu/gomaasclient/api/params"
	"github.com/ionutbalutoiu/gomaasclient/maas/entity"
)

// VLANs represents the MaaS Vlans endpoint
type VLANs interface {
	Get(fabricID int) ([]entity.VLAN, error)
	Post(fabricID int, params *params.VLAN) (*entity.VLAN, error)
}
