package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// Domain is an interface defining API behaviour for
// domain objects
type Domain interface {
	Get(ctx context.Context, id int) (*entity.Domain, error)
	SetDefault(ctx context.Context, id int) (*entity.Domain, error)
	Update(ctx context.Context, id int, params *entity.DomainParams) (*entity.Domain, error)
	Delete(ctx context.Context, id int) error
}
