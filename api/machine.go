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
	Unlock(systemID string, comment string) (*entity.Machine, error)
	ClearDefaultGateways(systemID string) (*entity.Machine, error)
	GetPowerParameters(systemID string) (map[string]interface{}, error)
	PowerOn(systemID string, params *entity.MachinePowerOnParams) (*entity.Machine, error)
	PowerOff(systemID string, params *entity.MachinePowerOffParams) (*entity.Machine, error)
	GetPowerState(systemID string) (*entity.MachinePowerState, error)
	SetWorkloadAnnotations(systemID string, params map[string]string) (*entity.Machine, error)
	RescueMode(systemID string) (*entity.Machine, error)
	ExitRescueMode(systemID string) (*entity.Machine, error)
	Abort(systemID string, comment string) (*entity.Machine, error)
	MarkBroken(systemID string, comment string) (*entity.Machine, error)
	MarkFixed(systemID string, comment string) (*entity.Machine, error)
	GetToken(systemID string) (*entity.MachineToken, error)
}
