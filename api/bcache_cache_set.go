package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// BCacheCacheSet declares the API operations related to MAAS `nodes/{systemID}/bcache-cache-set/{id}` endpoint.
type BCacheCacheSet interface {
	Get(systemID string, id int) (*entity.BCacheCacheSet, error)
	Update(systemID string, id int, params *entity.BCacheCacheSetParams) (*entity.BCacheCacheSet, error)
	Delete(systemID string, id int) error
}
