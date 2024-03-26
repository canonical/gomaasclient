package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// Tag is an interface defining API behaviour for Tag objects
type Tag interface {
	Get(name string) (*entity.Tag, error)
	Update(name string, tagParams *entity.TagParams) (*entity.Tag, error)
	Delete(name string) error
	GetMachines(name string) ([]entity.Machine, error)
	AddMachines(name string, machineIds []string) error
	RemoveMachines(name string, machineIds []string) error
}
