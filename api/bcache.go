package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// BCache declares the API operations related to MAAS `nodes/{systemID}/bcache/{id}` endpoint.
type BCache interface {
	Get(systemID string, id int) (*entity.BCache, error)
	Update(systemID string, id int, params *entity.BCacheParams) (*entity.BCache, error)
	Delete(systemID string, id int) error
}
