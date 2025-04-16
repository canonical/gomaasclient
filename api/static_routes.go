package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// StaticRoutes is an interface for listing and creating
// StaticRoutes records
type StaticRoutes interface {
	Get() ([]entity.StaticRoute, error)
	Create(params *entity.StaticRouteParams) (*entity.StaticRoute, error)
}
