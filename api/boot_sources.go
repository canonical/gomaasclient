package api

import (
	"github.com/maas/gomaasclient/entity"
)

// BootSources is an interface for listing and creating
// BootSource objects
type BootSources interface {
	Get() ([]entity.BootSource, error)
	Create(params *entity.BootSourceParams) (*entity.BootSource, error)
}
