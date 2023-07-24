package api

import (
	"github.com/maas/gomaasclient/entity"
)

type Machines interface {
	Get() ([]entity.Machine, error)
	Create(machineParams *entity.MachineParams, powerParams map[string]string) (*entity.Machine, error)
	Allocate(params *entity.MachineAllocateParams) (*entity.Machine, error)
	AcceptAll() error
	Release(systemID []string, comment string) error
}
