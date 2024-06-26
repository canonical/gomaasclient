package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// RAID declares the API operations related to MAAS `nodes/{systemID}/raid/{id}` endpoint.
type RAID interface {
	Get(systemID string, id int) (raid *entity.RAID, err error)
	Update(systemID string, id int, params *entity.RAIDUpdateParams) (raid *entity.RAID, err error)
	Delete(systemID string, id int) (err error)
}
