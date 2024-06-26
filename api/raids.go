package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// RAIDs declares the API operations related to MAAS `nodes/{systemID}/raids` endpoint.
type RAIDs interface {
	Get(systemID string) (raids []entity.RAID, err error)
	Create(systemID string, params *entity.RAIDCreateParams) (raid *entity.RAID, err error)
}
