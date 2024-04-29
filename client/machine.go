package client

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/yaml.v3"
)

// Machine contains functionality for manipulating the Machine entity.
type Machine struct {
	APIClient APIClient
}

func (m *Machine) client(systemID string) APIClient {
	return m.APIClient.GetSubObject("machines").GetSubObject(systemID)
}

// Get machine details.
func (m *Machine) Get(systemID string) (*entity.Machine, error) {
	machine := new(entity.Machine)
	err := m.client(systemID).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// Update machine.
func (m *Machine) Update(systemID string, machineParams *entity.MachineParams, powerParams map[string]interface{}) (*entity.Machine, error) {
	qsp, err := query.Values(machineParams)
	if err != nil {
		return nil, err
	}

	for k, v := range powerParamsToURLValues(powerParams) {
		// Since qsp.Add(k, v...) is not allowed
		qsp[k] = append(qsp[k], v...)
	}

	machine := new(entity.Machine)
	err = m.client(systemID).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// Delete machine.
func (m *Machine) Delete(systemID string) error {
	return m.client(systemID).Delete()
}

// Commission machine.
func (m *Machine) Commission(systemID string, params *entity.MachineCommissionParams) (*entity.Machine, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	machine := new(entity.Machine)
	err = m.client(systemID).Post("commission", qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// Deploy machine.
func (m *Machine) Deploy(systemID string, params *entity.MachineDeployParams) (*entity.Machine, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	machine := new(entity.Machine)
	err = m.client(systemID).Post("deploy", qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// Release machine.
func (m *Machine) Release(systemID string, params *entity.MachineReleaseParams) (*entity.Machine, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	machine := new(entity.Machine)
	err = m.client(systemID).Post("release", qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// Lock machine.
func (m *Machine) Lock(systemID string, comment string) (*entity.Machine, error) {
	qsp := make(url.Values)
	if comment != "" {
		qsp.Set("comment", comment)
	}

	machine := new(entity.Machine)
	err := m.client(systemID).Post("lock", qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// Unlock marks a machine as ‘Unlocked’, to allow changes
func (m *Machine) Unlock(systemID string, comment string) (*entity.Machine, error) {
	qsp := make(url.Values)
	if comment != "" {
		qsp.Set("comment", comment)
	}

	machine := new(entity.Machine)
	err := m.client(systemID).Post("unlock", qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// ClearDefaultGateways clears default gateways.
func (m *Machine) ClearDefaultGateways(systemID string) (*entity.Machine, error) {
	machine := new(entity.Machine)
	err := m.client(systemID).Post("clear_default_gateways", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// GetPowerParameters fetches the power parameters of a given Machine
func (m *Machine) GetPowerParameters(systemID string) (map[string]interface{}, error) {
	params := map[string]interface{}{}
	err := m.client(systemID).Get("power_parameters", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &params)
	})

	return params, err
}

// PowerOn machine
func (m *Machine) PowerOn(systemID string, params *entity.MachinePowerOnParams) (*entity.Machine, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	machine := new(entity.Machine)
	err = m.client(systemID).Post("power_on", qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// PowerOff machine
func (m *Machine) PowerOff(systemID string, params *entity.MachinePowerOffParams) (*entity.Machine, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	machine := new(entity.Machine)
	err = m.client(systemID).Post("power_off", qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// GetPowerState of the machine
func (m *Machine) GetPowerState(systemID string) (*entity.MachinePowerState, error) {
	ps := new(entity.MachinePowerState)
	err := m.client(systemID).Get("query_power_state", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, ps)
	})

	return ps, err
}

// SetWorkloadAnnotations add, modify or remove workload annotations for given Machine
func (m *Machine) SetWorkloadAnnotations(systemID string, params map[string]string) (*entity.Machine, error) {
	qsp := url.Values{}
	for k, v := range params {
		qsp.Add(k, v)
	}

	machine := new(entity.Machine)
	err := m.client(systemID).Post("set_workload_annotations", qsp, func(data []byte) error {
		return json.Unmarshal(data, &machine)
	})

	return machine, err
}

// RescueMode begins the rescue mode process on a given machine
func (m *Machine) RescueMode(systemID string) (*entity.Machine, error) {
	machine := new(entity.Machine)
	err := m.client(systemID).Post("rescue_mode", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &machine)
	})

	return machine, err
}

// ExitRescueMode exits the rescue mode process on a given machine
func (m *Machine) ExitRescueMode(systemID string) (*entity.Machine, error) {
	machine := new(entity.Machine)
	err := m.client(systemID).Post("exit_rescue_mode", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &machine)
	})

	return machine, err
}

// Abort aborts machine's current operation
func (m *Machine) Abort(systemID string, comment string) (*entity.Machine, error) {
	qsp := make(url.Values)
	if comment != "" {
		qsp.Set("comment", comment)
	}

	machine := new(entity.Machine)
	err := m.client(systemID).Post("abort", qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// MarkBroken marks a machine as ‘Broken’
func (m *Machine) MarkBroken(systemID string, comment string) (*entity.Machine, error) {
	qsp := make(url.Values)
	if comment != "" {
		qsp.Set("comment", comment)
	}

	machine := new(entity.Machine)
	err := m.client(systemID).Post("mark_broken", qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// MarkFixed marks a machine as ‘Fixed’
func (m *Machine) MarkFixed(systemID string, comment string) (*entity.Machine, error) {
	qsp := make(url.Values)
	if comment != "" {
		qsp.Set("comment", comment)
	}

	machine := new(entity.Machine)
	err := m.client(systemID).Post("mark_fixed", qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// GetToken gets the machine token for a given machine
func (m *Machine) GetToken(systemID string) (*entity.MachineToken, error) {
	machineToken := new(entity.MachineToken)
	err := m.client(systemID).Get("get_token", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, machineToken)
	})

	return machineToken, err
}

// Details gets the details for a given machine
func (m *Machine) Details(systemID string) (*entity.MachineDetails, error) {
	machineDetails := new(entity.MachineDetails)
	err := m.client(systemID).Get("details", url.Values{}, func(data []byte) error {
		return bson.Unmarshal(data, machineDetails)
	})

	return machineDetails, err
}

// GetCurtinConfig gets the curtin config for a given machine
func (m *Machine) GetCurtinConfig(systemID string) (map[string]interface{}, error) {
	curtinConfig := map[string]interface{}{}
	err := m.client(systemID).Get("get_curtin_config", url.Values{}, func(data []byte) error {
		return yaml.Unmarshal(data, &curtinConfig)
	})

	return curtinConfig, err
}
