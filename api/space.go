package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// Space is an interface defining API behaviour for
// Space objects
type Space interface {
	Get(ctx context.Context, id int) (*entity.Space, error)
	Update(ctx context.Context, id int, name string) (*entity.Space, error)
	Delete(ctx context.Context, id int) error
}
