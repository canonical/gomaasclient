package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/yaml.v3"
)

// Machine contains functionality for manipulating the Machine entity.
type Machine struct {
	APIClient APIClient
}

func (m *Machine) client(systemID string) *APIClient {
	return m.APIClient.SubClient("machines").SubClient(systemID)
}

// Get machine details.
func (m *Machine) Get(ctx context.Context, systemID string) (*entity.Machine, error) {
	machine := new(entity.Machine)
	err := m.client(systemID).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// Update machine.
func (m *Machine) Update(ctx context.Context, systemID string, machineParams *entity.MachineParams, powerParams map[string]interface{}) (*entity.Machine, error) {
	qsp, err := query.Values(machineParams)
	if err != nil {
		return nil, err
	}

	for k, v := range powerParamsToURLValues(powerParams) {
		// Since qsp.Add(k, v...) is not allowed
		qsp[k] = append(qsp[k], v...)
	}

	machine := new(entity.Machine)
	err = m.client(systemID).Put(ctx, qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// Delete machine.
func (m *Machine) Delete(ctx context.Context, systemID string) error {
	return m.client(systemID).Delete(ctx)
}

// Commission machine.
func (m *Machine) Commission(ctx context.Context, systemID string, params *entity.MachineCommissionParams) (*entity.Machine, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	machine := new(entity.Machine)
	err = m.client(systemID).Post(ctx, "commission", qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// Deploy machine.
func (m *Machine) Deploy(ctx context.Context, systemID string, params *entity.MachineDeployParams) (*entity.Machine, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	machine := new(entity.Machine)
	err = m.client(systemID).Post(ctx, "deploy", qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// Release machine.
func (m *Machine) Release(ctx context.Context, systemID string, params *entity.MachineReleaseParams) (*entity.Machine, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	machine := new(entity.Machine)
	err = m.client(systemID).Post(ctx, "release", qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// Lock machine.
func (m *Machine) Lock(ctx context.Context, systemID string, comment string) (*entity.Machine, error) {
	qsp := make(url.Values)
	if comment != "" {
		qsp.Set("comment", comment)
	}

	machine := new(entity.Machine)
	err := m.client(systemID).Post(ctx, "lock", qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// Unlock marks a machine as ‘Unlocked’, to allow changes
func (m *Machine) Unlock(ctx context.Context, systemID string, comment string) (*entity.Machine, error) {
	qsp := make(url.Values)
	if comment != "" {
		qsp.Set("comment", comment)
	}

	machine := new(entity.Machine)
	err := m.client(systemID).Post(ctx, "unlock", qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// ClearDefaultGateways clears default gateways.
func (m *Machine) ClearDefaultGateways(ctx context.Context, systemID string) (*entity.Machine, error) {
	machine := new(entity.Machine)
	err := m.client(systemID).Post(ctx, "clear_default_gateways", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// GetPowerParameters fetches the power parameters of a given Machine
func (m *Machine) GetPowerParameters(ctx context.Context, systemID string) (map[string]interface{}, error) {
	params := map[string]interface{}{}
	err := m.client(systemID).Get(ctx, "power_parameters", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &params)
	})

	return params, err
}

// PowerOn machine
func (m *Machine) PowerOn(ctx context.Context, systemID string, params *entity.MachinePowerOnParams) (*entity.Machine, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	machine := new(entity.Machine)
	err = m.client(systemID).Post(ctx, "power_on", qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// PowerOff machine
func (m *Machine) PowerOff(ctx context.Context, systemID string, params *entity.MachinePowerOffParams) (*entity.Machine, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	machine := new(entity.Machine)
	err = m.client(systemID).Post(ctx, "power_off", qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// GetPowerState of the machine
func (m *Machine) GetPowerState(ctx context.Context, systemID string) (*entity.MachinePowerState, error) {
	ps := new(entity.MachinePowerState)
	err := m.client(systemID).Get(ctx, "query_power_state", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, ps)
	})

	return ps, err
}

// SetWorkloadAnnotations add, modify or remove workload annotations for given Machine
func (m *Machine) SetWorkloadAnnotations(ctx context.Context, systemID string, params map[string]string) (*entity.Machine, error) {
	qsp := url.Values{}
	for k, v := range params {
		qsp.Add(k, v)
	}

	machine := new(entity.Machine)
	err := m.client(systemID).Post(ctx, "set_workload_annotations", qsp, func(data []byte) error {
		return json.Unmarshal(data, &machine)
	})

	return machine, err
}

// RescueMode begins the rescue mode process on a given machine
func (m *Machine) RescueMode(ctx context.Context, systemID string) (*entity.Machine, error) {
	machine := new(entity.Machine)
	err := m.client(systemID).Post(ctx, "rescue_mode", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &machine)
	})

	return machine, err
}

// ExitRescueMode exits the rescue mode process on a given machine
func (m *Machine) ExitRescueMode(ctx context.Context, systemID string) (*entity.Machine, error) {
	machine := new(entity.Machine)
	err := m.client(systemID).Post(ctx, "exit_rescue_mode", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &machine)
	})

	return machine, err
}

// Abort aborts machine's current operation
func (m *Machine) Abort(ctx context.Context, systemID string, comment string) (*entity.Machine, error) {
	qsp := make(url.Values)
	if comment != "" {
		qsp.Set("comment", comment)
	}

	machine := new(entity.Machine)
	err := m.client(systemID).Post(ctx, "abort", qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// MarkBroken marks a machine as ‘Broken’
func (m *Machine) MarkBroken(ctx context.Context, systemID string, comment string) (*entity.Machine, error) {
	qsp := make(url.Values)
	if comment != "" {
		qsp.Set("comment", comment)
	}

	machine := new(entity.Machine)
	err := m.client(systemID).Post(ctx, "mark_broken", qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// MarkFixed marks a machine as ‘Fixed’
func (m *Machine) MarkFixed(ctx context.Context, systemID string, comment string) (*entity.Machine, error) {
	qsp := make(url.Values)
	if comment != "" {
		qsp.Set("comment", comment)
	}

	machine := new(entity.Machine)
	err := m.client(systemID).Post(ctx, "mark_fixed", qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// GetToken gets the machine token for a given machine
func (m *Machine) GetToken(ctx context.Context, systemID string) (*entity.MachineToken, error) {
	machineToken := new(entity.MachineToken)
	err := m.client(systemID).Get(ctx, "get_token", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, machineToken)
	})

	return machineToken, err
}

// Details gets the details for a given machine
func (m *Machine) Details(ctx context.Context, systemID string) (*entity.MachineDetails, error) {
	machineDetails := new(entity.MachineDetails)
	err := m.client(systemID).Get(ctx, "details", url.Values{}, func(data []byte) error {
		return bson.Unmarshal(data, machineDetails)
	})

	return machineDetails, err
}

// RestoreDefaultConfiguration sets the configuration to default values for a given machine
func (m *Machine) RestoreDefaultConfiguration(ctx context.Context, systemID string) error {
	return m.client(systemID).Post(ctx, "restore_default_configuration", url.Values{}, func(data []byte) error { return nil })
}

// RestoreNetworkingConfiguration sets the network configuration to default values for a given machine
func (m *Machine) RestoreNetworkingConfiguration(ctx context.Context, systemID string) error {
	return m.client(systemID).Post(ctx, "restore_networking_configuration", url.Values{}, func(data []byte) error { return nil })
}

// RestoreStorageConfiguration sets the storage configuration to default values for a given machine
func (m *Machine) RestoreStorageConfiguration(ctx context.Context, systemID string) error {
	return m.client(systemID).Post(ctx, "restore_storage_configuration", url.Values{}, func(data []byte) error { return nil })
}

// GetCurtinConfig gets the curtin config for a given machine
func (m *Machine) GetCurtinConfig(ctx context.Context, systemID string) (map[string]interface{}, error) {
	curtinConfig := map[string]interface{}{}
	err := m.client(systemID).Get(ctx, "get_curtin_config", url.Values{}, func(data []byte) error {
		return yaml.Unmarshal(data, &curtinConfig)
	})

	return curtinConfig, err
}
