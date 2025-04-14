package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// BlockDevices is an interface for listing and creating
// BlockDevice records
type BlockDevices interface {
	Get(ctx context.Context, systemID string) ([]entity.BlockDevice, error)
	Create(ctx context.Context, systemID string, params *entity.BlockDeviceParams) (*entity.BlockDevice, error)
}
