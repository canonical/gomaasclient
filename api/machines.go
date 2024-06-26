package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// Machines is an interface for listing, creating, allocating,
// bulk-accepting enlistments and releasing Machine records
type Machines interface {
	Get(machinesParams *entity.MachinesParams) ([]entity.Machine, error)
	Create(machineParams *entity.MachineParams, powerParams map[string]interface{}) (*entity.Machine, error)
	Allocate(params *entity.MachineAllocateParams) (*entity.Machine, error)
	AcceptAll() error
	Release(systemID []string, comment string) error
	ListAllocated() ([]entity.Machine, error)
}
