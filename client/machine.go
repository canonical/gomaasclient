package client

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
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
