package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// Zone is an interface defining API behaviour for zones
type Zone interface {
	Get(ctx context.Context, name string) (*entity.Zone, error)
	Update(ctx context.Context, name string, params *entity.ZoneParams) (*entity.Zone, error)
	Delete(ctx context.Context, name string) error
}
