package api

import (
	"github.com/maas/gomaasclient/entity"
)

// VLANs represents the MAAS Vlans endpoint
type VLANs interface {
	Get(fabricID int) ([]entity.VLAN, error)
	Create(fabricID int, params *entity.VLANParams) (*entity.VLAN, error)
}
