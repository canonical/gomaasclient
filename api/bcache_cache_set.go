package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// BCacheCacheSet declares the API operations related to MAAS `nodes/{systemID}/bcache-cache-set/{id}` endpoint.
type BCacheCacheSet interface {
	Get(ctx context.Context, systemID string, id int) (*entity.BCacheCacheSet, error)
	Update(ctx context.Context, systemID string, id int, params *entity.BCacheCacheSetParams) (*entity.BCacheCacheSet, error)
	Delete(ctx context.Context, systemID string, id int) error
}
