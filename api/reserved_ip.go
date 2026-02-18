package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// ReservedIP is an interface defining API behavior for
// reserved IP objects
type ReservedIP interface {
	Get(id int) (*entity.ReservedIP, error)
	Update(id int, params *entity.ReservedIPUpdateParams) (*entity.ReservedIP, error)
	Delete(id int) error
}
