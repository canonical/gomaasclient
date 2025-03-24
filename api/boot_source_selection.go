package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// BootSourceSelection is an interface defining API behaviour for
// boot source selections
type BootSourceSelection interface {
	Get(ctx context.Context, bootSourceID int, id int) (*entity.BootSourceSelection, error)
	Update(ctx context.Context, bootSourceID int, id int, params *entity.BootSourceSelectionParams) (*entity.BootSourceSelection, error)
	Delete(ctx context.Context, bootSourceID int, id int) error
}
