package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// Machine is an interface defining API behaviour for Machine objects
type Machine interface {
	Get(ctx context.Context, systemID string) (*entity.Machine, error)
	Update(ctx context.Context, systemID string, machineParams *entity.MachineParams, powerParams map[string]interface{}) (*entity.Machine, error)
	Delete(ctx context.Context, systemID string) error
	Commission(ctx context.Context, systemID string, params *entity.MachineCommissionParams) (*entity.Machine, error)
	Deploy(ctx context.Context, systemID string, params *entity.MachineDeployParams) (*entity.Machine, error)
	Release(ctx context.Context, systemID string, params *entity.MachineReleaseParams) (*entity.Machine, error)
	Lock(ctx context.Context, systemID string, comment string) (*entity.Machine, error)
	Unlock(ctx context.Context, systemID string, comment string) (*entity.Machine, error)
	ClearDefaultGateways(ctx context.Context, systemID string) (*entity.Machine, error)
	GetPowerParameters(ctx context.Context, systemID string) (map[string]interface{}, error)
	PowerOn(ctx context.Context, systemID string, params *entity.MachinePowerOnParams) (*entity.Machine, error)
	PowerOff(ctx context.Context, systemID string, params *entity.MachinePowerOffParams) (*entity.Machine, error)
	GetPowerState(ctx context.Context, systemID string) (*entity.MachinePowerState, error)
	SetWorkloadAnnotations(ctx context.Context, systemID string, params map[string]string) (*entity.Machine, error)
	RescueMode(ctx context.Context, systemID string) (*entity.Machine, error)
	ExitRescueMode(ctx context.Context, systemID string) (*entity.Machine, error)
	Abort(ctx context.Context, systemID string, comment string) (*entity.Machine, error)
	MarkBroken(ctx context.Context, systemID string, comment string) (*entity.Machine, error)
	MarkFixed(ctx context.Context, systemID string, comment string) (*entity.Machine, error)
	GetToken(ctx context.Context, systemID string) (*entity.MachineToken, error)
	Details(ctx context.Context, systemID string) (*entity.MachineDetails, error)
	RestoreDefaultConfiguration(ctx context.Context, systemID string) error
	RestoreNetworkingConfiguration(ctx context.Context, systemID string) error
	RestoreStorageConfiguration(ctx context.Context, systemID string) error
	GetCurtinConfig(ctx context.Context, systemID string) (map[string]interface{}, error)
}
