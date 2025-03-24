package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// BCache declares the API operations related to MAAS `nodes/{systemID}/bcache/{id}` endpoint.
type BCache interface {
	Get(ctx context.Context, systemID string, id int) (*entity.BCache, error)
	Update(ctx context.Context, systemID string, id int, params *entity.BCacheParams) (*entity.BCache, error)
	Delete(ctx context.Context, ystemID string, id int) error
}
