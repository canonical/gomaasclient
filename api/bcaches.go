package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// BCaches declares the API operations related to MAAS `nodes/{systemID}/bcaches` endpoint.
type BCaches interface {
	Get(ctx context.Context, systemID string) ([]entity.BCache, error)
	Create(ctx context.Context, systemID string, params *entity.BCacheParams) (*entity.BCache, error)
}
