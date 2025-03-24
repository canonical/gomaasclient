package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// VMHost is an interface defining API behaviour for VMHost objects
type VMHost interface {
	Get(ctx context.Context, id int) (*entity.VMHost, error)
	Update(ctx context.Context, id int, params *entity.VMHostParams) (*entity.VMHost, error)
	Delete(ctx context.Context, id int) error
	Compose(ctx context.Context, id int, params *entity.VMHostMachineParams) (*entity.Machine, error)
	Refresh(ctx context.Context, id int) (*entity.VMHost, error)
	GetParameters(ctx context.Context, id int) (map[string]string, error)
}
