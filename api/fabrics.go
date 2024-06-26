package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// Fabrics is an interface for listing and creating Fabric records
type Fabrics interface {
	Get() ([]entity.Fabric, error)
	Create(fabricParams *entity.FabricParams) (*entity.Fabric, error)
}
