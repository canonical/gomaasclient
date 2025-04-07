package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// BlockDevice is an interface providing API behaviour for
// block devices
type BlockDevice interface {
	Get(ctx context.Context, systemID string, id int) (*entity.BlockDevice, error)
	Update(ctx context.Context, systemID string, id int, params *entity.BlockDeviceParams) (*entity.BlockDevice, error)
	Delete(ctx context.Context, systemID string, id int) error
	AddTag(ctx context.Context, systemID string, id int, tag string) (*entity.BlockDevice, error)
	RemoveTag(ctx context.Context, systemID string, id int, tag string) (*entity.BlockDevice, error)
	Format(ctx context.Context, systemID string, id int, fsType string) (*entity.BlockDevice, error)
	Unformat(ctx context.Context, systemID string, id int) (*entity.BlockDevice, error)
	Mount(ctx context.Context, systemID string, id int, mountPoint string, mountOptions string) (*entity.BlockDevice, error)
	Unmount(ctx context.Context, systemID string, id int) (*entity.BlockDevice, error)
	SetBootDisk(ctx context.Context, systemID string, id int) error
}
