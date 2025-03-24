package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// VolumeGroups is an interface for listing and creating
// VolumeGroup records
type VolumeGroups interface {
	Get(ctx context.Context, systemID string) ([]entity.VolumeGroup, error)
	Create(ctx context.Context, systemID string, params *entity.VolumeGroupCreateParams) (*entity.VolumeGroup, error)
}
