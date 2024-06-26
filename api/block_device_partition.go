package api

import "github.com/canonical/gomaasclient/entity"

// BlockDevicePartition is an interface providing API behaviour for
// block device partitions
type BlockDevicePartition interface {
	Get(systemID string, blockDeviceID int, id int) (*entity.BlockDevicePartition, error)
	Delete(systemID string, blockDeviceID int, id int) error
	AddTag(systemID string, blockDeviceID int, id int, tag string) (*entity.BlockDevicePartition, error)
	RemoveTag(systemID string, blockDeviceID int, id int, tag string) (*entity.BlockDevicePartition, error)
	Format(systemID string, blockDeviceID int, id int, fsType string, label string) (*entity.BlockDevicePartition, error)
	Unformat(systemID string, blockDeviceID int, id int) (*entity.BlockDevicePartition, error)
	Mount(systemID string, blockDeviceID int, id int, mountPoint string, mountOptions string) (*entity.BlockDevicePartition, error)
	Unmount(systemID string, blockDeviceID int, id int) (*entity.BlockDevicePartition, error)
}
