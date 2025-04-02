package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// BootSourceSelections is an interface for listing and creating
// BootSourceSelection objects
type BootSourceSelections interface {
	Get(ctx context.Context, bootSourceID int) ([]entity.BootSourceSelection, error)
	Create(ctx context.Context, bootSourceID int, params *entity.BootSourceSelectionParams) (*entity.BootSourceSelection, error)
}
