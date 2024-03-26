package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// Domain is an interface defining API behaviour for
// domain objects
type Domain interface {
	Get(id int) (*entity.Domain, error)
	SetDefault(id int) (*entity.Domain, error)
	Update(id int, params *entity.DomainParams) (*entity.Domain, error)
	Delete(id int) error
}
