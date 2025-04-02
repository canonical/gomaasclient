package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// Zones is an interface for listing and creating Zone objects
type Zones interface {
	Get(ctx context.Context) ([]entity.Zone, error)
	Create(ctx context.Context, params *entity.ZoneParams) (*entity.Zone, error)
}
