package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// BlockDevicePartitions is an interface for listing and creating
// BlockDevicePartition records
type BlockDevicePartitions interface {
	Get(ctx context.Context, systemID string, blockDeviceID int) ([]entity.BlockDevicePartition, error)
	Create(ctx context.Context, systemID string, blockDeviceID int, params *entity.BlockDevicePartitionParams) (*entity.BlockDevicePartition, error)
}
