package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// BlockDevice is an interface providing API behaviour for
// block devices
type BlockDevice interface {
	Get(systemID string, id int) (*entity.BlockDevice, error)
	Update(systemID string, id int, params *entity.BlockDeviceParams) (*entity.BlockDevice, error)
	Delete(systemID string, id int) error
	AddTag(systemID string, id int, tag string) (*entity.BlockDevice, error)
	RemoveTag(systemID string, id int, tag string) (*entity.BlockDevice, error)
	Format(systemID string, id int, fsType string) (*entity.BlockDevice, error)
	Unformat(systemID string, id int) (*entity.BlockDevice, error)
	Mount(systemID string, id int, mountPoint string, mountOptions string) (*entity.BlockDevice, error)
	Unmount(systemID string, id int) (*entity.BlockDevice, error)
	SetBootDisk(systemID string, id int) error
}
