package api

import (
	"github.com/maas/gomaasclient/entity"
)

// BCacheCacheSets declares the API operations related to MAAS `nodes/{systemID}/bcache-cache-sets` endpoint.
type BCacheCacheSets interface {
	Get(systemID string) ([]entity.BCacheCacheSet, error)
	Create(systemID string, params *entity.BCacheCacheSetParams) (*entity.BCacheCacheSet, error)
}
