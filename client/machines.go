package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// Machines contains functionality for manipulating the Machines entity.
type Machines struct {
	APIClient APIClient
}

func (m *Machines) client() *APIClient {
	return m.APIClient.SubClient("machines")
}

// Get fetches a list machines.
func (m *Machines) Get(ctx context.Context, machinesParams *entity.MachinesParams) ([]entity.Machine, error) {
	qsp, err := query.Values(machinesParams)
	if err != nil {
		return nil, err
	}

	machines := make([]entity.Machine, 0)
	err = m.client().Get(ctx, "", qsp, func(data []byte) error {
		return json.Unmarshal(data, &machines)
	})

	return machines, err
}

// Create machine.
func (m *Machines) Create(ctx context.Context, machineParams *entity.MachineParams, powerParams map[string]interface{}) (*entity.Machine, error) {
	qsp, err := query.Values(machineParams)
	if err != nil {
		return nil, err
	}

	for k, v := range powerParamsToURLValues(powerParams) {
		// Since qsp.Add(k, v...) is not allowed
		qsp[k] = append(qsp[k], v...)
	}

	machine := new(entity.Machine)
	err = m.client().Post(ctx, "", qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// Allocate machine.
func (m *Machines) Allocate(ctx context.Context, params *entity.MachineAllocateParams) (*entity.Machine, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	machine := new(entity.Machine)
	err = m.client().Post(ctx, "allocate", qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// Release machine.
func (m *Machines) Release(ctx context.Context, systemIDs []string, comment string) error {
	qsp := url.Values{}
	for _, val := range systemIDs {
		qsp.Add("machines", val)
	}

	if comment != "" {
		qsp.Add("comment", comment)
	}

	return m.client().Post(ctx, "release", qsp, func(data []byte) error { return nil })
}

// AcceptAll accepts all pending enlistments
func (m *Machines) AcceptAll(ctx context.Context) error {
	qsp := url.Values{}
	return m.client().Post(ctx, "accept_all", qsp, func(data []byte) error { return nil })
}

// ListAllocated lists allocated machines
func (m *Machines) ListAllocated(ctx context.Context) ([]entity.Machine, error) {
	machines := make([]entity.Machine, 0)
	err := m.client().Get(ctx, "list_allocated", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &machines)
	})

	return machines, err
}
