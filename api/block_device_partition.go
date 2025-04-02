package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// BlockDevicePartition is an interface providing API behaviour for
// block device partitions
type BlockDevicePartition interface {
	Get(ctx context.Context, systemID string, blockDeviceID int, id int) (*entity.BlockDevicePartition, error)
	Delete(ctx context.Context, systemID string, blockDeviceID int, id int) error
	AddTag(ctx context.Context, systemID string, blockDeviceID int, id int, tag string) (*entity.BlockDevicePartition, error)
	RemoveTag(ctx context.Context, systemID string, blockDeviceID int, id int, tag string) (*entity.BlockDevicePartition, error)
	Format(ctx context.Context, systemID string, blockDeviceID int, id int, fsType string, label string) (*entity.BlockDevicePartition, error)
	Unformat(ctx context.Context, systemID string, blockDeviceID int, id int) (*entity.BlockDevicePartition, error)
	Mount(ctx context.Context, systemID string, blockDeviceID int, id int, mountPoint string, mountOptions string) (*entity.BlockDevicePartition, error)
	Unmount(ctx context.Context, systemID string, blockDeviceID int, id int) (*entity.BlockDevicePartition, error)
}
