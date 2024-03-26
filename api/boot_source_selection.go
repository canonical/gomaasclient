package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// BootSourceSelection is an interface defining API behaviour for
// boot source selections
type BootSourceSelection interface {
	Get(bootSourceID int, id int) (*entity.BootSourceSelection, error)
	Update(bootSourceID int, id int, params *entity.BootSourceSelectionParams) (*entity.BootSourceSelection, error)
	Delete(bootSourceID int, id int) error
}
