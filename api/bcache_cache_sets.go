package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// BCacheCacheSets declares the API operations related to MAAS `nodes/{systemID}/bcache-cache-sets` endpoint.
type BCacheCacheSets interface {
	Get(ctx context.Context, systemID string) ([]entity.BCacheCacheSet, error)
	Create(ctx context.Context, systemID string, params *entity.BCacheCacheSetParams) (*entity.BCacheCacheSet, error)
}
