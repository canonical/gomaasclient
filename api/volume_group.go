package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// VolumeGroup is an interface providing API behaviour for
// volume groups
type VolumeGroup interface {
	Get(systemID string, id int) (*entity.VolumeGroup, error)
	Update(systemID string, id int, params *entity.VolumeGroupParams) (*entity.VolumeGroup, error)
	Delete(systemID string, id int) error
	CreateLogicalVolume(systemID string, id int, params *entity.LogicalVolumeParams) (*entity.LogicalVolume, error)
	DeleteLogicalVolume(ssytemID string, id int, logicalVolumeID int) error
}
