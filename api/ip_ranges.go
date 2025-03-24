package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// IPRanges is an interface for listing and creating IPRange records
type IPRanges interface {
	Get(ctx context.Context) ([]entity.IPRange, error)
	Create(ctx context.Context, params *entity.IPRangeParams) (*entity.IPRange, error)
}
