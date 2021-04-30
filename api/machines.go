package api

import (
	"github.com/ionutbalutoiu/gomaasclient/entity"
)

type Machines interface {
	Get() ([]entity.Machine, error)
	Create(machineParams *entity.MachineParams, powerParams map[string]string) (*entity.Machine, error)
	Allocate(params *entity.MachinesAllocateParams) (*entity.Machine, error)
	Release(systemID []string, comment string) error
}
