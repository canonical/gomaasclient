package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// RAID declares the API operations related to MAAS `nodes/{systemID}/raid/{id}` endpoint.
type RAID interface {
	Get(ctx context.Context, systemID string, id int) (raid *entity.RAID, err error)
	Update(ctx context.Context, systemID string, id int, params *entity.RAIDUpdateParams) (raid *entity.RAID, err error)
	Delete(ctx context.Context, systemID string, id int) (err error)
}
