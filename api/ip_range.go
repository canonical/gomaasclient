package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// IPRange is an interface defining API behaviour for IP ranges
type IPRange interface {
	Get(ctx context.Context, id int) (*entity.IPRange, error)
	Update(ctx context.Context, id int, params *entity.IPRangeParams) (*entity.IPRange, error)
	Delete(ctx context.Context, id int) error
}
