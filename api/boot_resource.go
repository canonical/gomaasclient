package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// BootResource is an interface defining API behaviour for
// boot resources
type BootResource interface {
	Get(ctx context.Context, id int) (*entity.BootResource, error)
	Delete(ctx context.Context, id int) error
}
