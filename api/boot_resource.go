package api

import (
	"github.com/maas/gomaasclient/entity"
)

// BootResource is an interface defining API behaviour for
// boot resources
type BootResource interface {
	Get(id int) (*entity.BootResource, error)
	Delete(id int) error
}
