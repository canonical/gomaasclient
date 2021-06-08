package api

import (
	"github.com/ionutbalutoiu/gomaasclient/entity"
)

type Machine interface {
	Get(systemID string) (*entity.Machine, error)
	Update(systemID string, machineParams *entity.MachineParams, powerParams map[string]string) (*entity.Machine, error)
	Delete(systemID string) error
	Commission(systemID string, params *entity.MachineCommissionParams) (*entity.Machine, error)
	Deploy(systemID string, params *entity.MachineDeployParams) (*entity.Machine, error)
	Lock(systemID string, comment string) (*entity.Machine, error)
	ClearDefaultGateways(systemID string) (*entity.Machine, error)
	GetPowerParameters(systemID string) (map[string]string, error)
}
