package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// VLAN is an interface defining API behaviour for VLAN objects
type VLAN interface {
	Get(fabricID int, vid int) (*entity.VLAN, error)
	Update(fabricID int, vid int, params *entity.VLANParams) (*entity.VLAN, error)
	Delete(fabricID int, vid int) error
}
