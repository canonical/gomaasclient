package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// StaticRoute is an interface defining API behavior for
// StaticRoute objects
type StaticRoute interface {
	Get(id int) (*entity.StaticRoute, error)
	Update(id int, params *entity.StaticRouteParams) (*entity.StaticRoute, error)
	Delete(id int) error
}
