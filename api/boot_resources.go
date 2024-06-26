package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// BootResources is an interface for listing and creating
// BootResource objects
type BootResources interface {
	Get(params *entity.BootResourcesReadParams) ([]entity.BootResource, error)
	Create(params *entity.BootResourceParams) (*entity.BootResource, error)
	Import() error
	IsImporting() (bool, error)
	StopImport() error
}
