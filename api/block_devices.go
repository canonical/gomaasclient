package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// BlockDevices is an interface for listing and creating
// BlockDevice records
type BlockDevices interface {
	Get(systemID string) ([]entity.BlockDevice, error)
	Create(systemID string, params *entity.BlockDeviceParams) (*entity.BlockDevice, error)
}
