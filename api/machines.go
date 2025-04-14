package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// Machines is an interface for listing, creating, allocating,
// bulk-accepting enlistments and releasing Machine records
type Machines interface {
	Get(ctx context.Context, machinesParams *entity.MachinesParams) ([]entity.Machine, error)
	Create(ctx context.Context, machineParams *entity.MachineParams, powerParams map[string]interface{}) (*entity.Machine, error)
	Allocate(ctx context.Context, params *entity.MachineAllocateParams) (*entity.Machine, error)
	AcceptAll(ctx context.Context) error
	Release(ctx context.Context, systemID []string, comment string) error
	ListAllocated(ctx context.Context) ([]entity.Machine, error)
}
