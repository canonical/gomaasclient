package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// BCaches declares the API operations related to MAAS `nodes/{systemID}/bcaches` endpoint.
type BCaches interface {
	Get(systemID string) ([]entity.BCache, error)
	Create(systemID string, params *entity.BCacheParams) (*entity.BCache, error)
}
