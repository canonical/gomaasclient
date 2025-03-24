package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// VolumeGroup is an interface providing API behaviour for
// volume groups
type VolumeGroup interface {
	Get(ctx context.Context, systemID string, id int) (*entity.VolumeGroup, error)
	Update(ctx context.Context, systemID string, id int, params *entity.VolumeGroupUpdateParams) (*entity.VolumeGroup, error)
	Delete(ctx context.Context, systemID string, id int) error
	CreateLogicalVolume(ctx context.Context, systemID string, id int, params *entity.LogicalVolumeParams) (*entity.BlockDevice, error)
	DeleteLogicalVolume(ctx context.Context, systemID string, id int, logicalVolumeID int) error
}
