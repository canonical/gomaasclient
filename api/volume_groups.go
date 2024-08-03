package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// VolumeGroups is an interface for listing and creating
// VolumeGroup records
type VolumeGroups interface {
	Get(systemID string) ([]entity.VolumeGroup, error)
	Create(systemID string, params *entity.VolumeGroupParams) (*entity.VolumeGroup, error)
}
