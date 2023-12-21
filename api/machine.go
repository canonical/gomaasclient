package api

import (
	"github.com/maas/gomaasclient/entity"
)

// Machine is an interface defining API behaviour for Machine objects
type Machine interface {
	Get(systemID string) (*entity.Machine, error)
	Update(systemID string, machineParams *entity.MachineParams, powerParams map[string]interface{}) (*entity.Machine, error)
	Delete(systemID string) error
	Commission(systemID string, params *entity.MachineCommissionParams) (*entity.Machine, error)
	Deploy(systemID string, params *entity.MachineDeployParams) (*entity.Machine, error)
	Release(systemID string, params *entity.MachineReleaseParams) (*entity.Machine, error)
	Lock(systemID string, comment string) (*entity.Machine, error)
	ClearDefaultGateways(systemID string) (*entity.Machine, error)
	GetPowerParameters(systemID string) (map[string]interface{}, error)
	PowerOn(systemID string) (*entity.Machine, error)
	PowerOff(systemID string) (*entity.Machine, error)
	GetPowerState(systemID string) (*entity.MachinePowerState, error)
}
