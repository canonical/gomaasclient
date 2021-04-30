package client

import (
	"encoding/json"
	"net/url"

	"github.com/ionutbalutoiu/gomaasclient/entity"
)

// Contains functionality for manipulating the Machine entity.
type Machine struct {
	ApiClient ApiClient
}

func (m *Machine) client(systemID string) ApiClient {
	return m.ApiClient.GetSubObject("machines").GetSubObject(systemID)
}

// Get machine details.
func (m *Machine) Get(systemID string) (ma *entity.Machine, err error) {
	ma = new(entity.Machine)
	err = m.client(systemID).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, ma)
	})
	return
}

// Update machine.
func (m *Machine) Update(systemID string, machineParams *entity.MachineParams, powerParams map[string]string) (ma *entity.Machine, err error) {
	qsp := ToQSP(machineParams)
	for k, v := range powerParams {
		qsp.Add(k, v)
	}
	ma = new(entity.Machine)
	err = m.client(systemID).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, ma)
	})
	return
}

// Delete machine.
func (m *Machine) Delete(systemID string) error {
	return m.client(systemID).Delete()
}

// Commission machine.
func (m *Machine) Commission(systemID string, params *entity.MachineCommissionParams) (ma *entity.Machine, err error) {
	ma = new(entity.Machine)
	err = m.client(systemID).Post("commission", ToQSP(params), func(data []byte) error {
		return json.Unmarshal(data, ma)
	})
	return
}

// Deploy machine.
func (m *Machine) Deploy(systemID string, params *entity.MachineDeployParams) (ma *entity.Machine, err error) {
	ma = new(entity.Machine)
	err = m.client(systemID).Post("deploy", ToQSP(params), func(data []byte) error {
		return json.Unmarshal(data, ma)
	})
	return
}

// Lock machine.
func (m *Machine) Lock(systemID string, comment string) (ma *entity.Machine, err error) {
	qsp := make(url.Values)
	if comment != "" {
		qsp.Set("comment", comment)
	}
	ma = new(entity.Machine)
	err = m.client(systemID).Post("lock", qsp, func(data []byte) error {
		return json.Unmarshal(data, ma)
	})
	return
}
