package client

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// Machines contains functionality for manipulating the Machines entity.
type Machines struct {
	APIClient APIClient
}

func (m *Machines) client() APIClient {
	return m.APIClient.GetSubObject("machines")
}

// Get fetches a list machines.
func (m *Machines) Get() ([]entity.Machine, error) {
	machines := make([]entity.Machine, 0)
	err := m.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &machines)
	})

	return machines, err
}

// Create machine.
func (m *Machines) Create(machineParams *entity.MachineParams, powerParams map[string]interface{}) (*entity.Machine, error) {
	qsp, err := query.Values(machineParams)
	if err != nil {
		return nil, err
	}

	for k, v := range powerParamsToURLValues(powerParams) {
		// Since qsp.Add(k, v...) is not allowed
		qsp[k] = append(qsp[k], v...)
	}

	machine := new(entity.Machine)
	err = m.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// Allocate machine.
func (m *Machines) Allocate(params *entity.MachineAllocateParams) (*entity.Machine, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	machine := new(entity.Machine)
	err = m.client().Post("allocate", qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// Release machine.
func (m *Machines) Release(systemIDs []string, comment string) error {
	qsp := url.Values{}
	for _, val := range systemIDs {
		qsp.Add("machines", val)
	}

	if comment != "" {
		qsp.Add("comment", comment)
	}

	return m.client().Post("release", qsp, func(data []byte) error { return nil })
}

// AcceptAll accepts all pending enlistments
func (m *Machines) AcceptAll() error {
	qsp := url.Values{}
	return m.client().Post("accept_all", qsp, func(data []byte) error { return nil })
}
