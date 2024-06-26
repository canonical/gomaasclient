package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// Fabric is an interface defining API behaviour for
// fabric objects
type Fabric interface {
	Get(id int) (*entity.Fabric, error)
	Update(id int, fabricParams *entity.FabricParams) (*entity.Fabric, error)
	Delete(id int) error
}
