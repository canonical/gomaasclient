package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// VMHosts is an interface for listing and creating VMHost objects
type VMHosts interface {
	Get(ctx context.Context) ([]entity.VMHost, error)
	Create(ctx context.Context, params *entity.VMHostParams) (*entity.VMHost, error)
}
