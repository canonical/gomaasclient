package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// Space is an interface defining API behaviour for
// Space objects
type Space interface {
	Get(id int) (*entity.Space, error)
	Update(id int, name string) (*entity.Space, error)
	Delete(id int) error
}
