package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// Spaces is an interface for listing and creating Space objects
type Spaces interface {
	Get(ctx context.Context) ([]entity.Space, error)
	Create(ctx context.Context, name string) (*entity.Space, error)
}
