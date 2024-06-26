package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// BootSourceSelections is an interface for listing and creating
// BootSourceSelection objects
type BootSourceSelections interface {
	Get(bootSourceID int) ([]entity.BootSourceSelection, error)
	Create(bootSourceID int, params *entity.BootSourceSelectionParams) (*entity.BootSourceSelection, error)
}
