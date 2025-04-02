package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// RAIDs declares the API operations related to MAAS `nodes/{systemID}/raids` endpoint.
type RAIDs interface {
	Get(ctx context.Context, systemID string) (raids []entity.RAID, err error)
	Create(ctx context.Context, systemID string, params *entity.RAIDCreateParams) (raid *entity.RAID, err error)
}
