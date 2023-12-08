package client

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// Machines contains functionality for manipulating the Machines entity.
type Machines struct {
	ApiClient ApiClient
}

func (m *Machines) client() ApiClient {
	return m.ApiClient.GetSubObject("machines")
}

// Get fetches a list machines.
func (m *Machines) Get() (machines []entity.Machine, err error) {
	err = m.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &machines)
	})

	return
}

// Create machine.
func (m *Machines) Create(machineParams *entity.MachineParams, powerParams map[string]interface{}) (ma *entity.Machine, err error) {
	qsp, err := query.Values(machineParams)
	if err != nil {
		return
	}

	for k, v := range powerParamsToURLValues(powerParams) {
		// Since qsp.Add(k, v...) is not allowed
		qsp[k] = append(qsp[k], v...)
	}

	ma = new(entity.Machine)
	err = m.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, ma)
	})

	return
}

// Allocate machine.
func (m *Machines) Allocate(params *entity.MachineAllocateParams) (ma *entity.Machine, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}

	ma = new(entity.Machine)
	err = m.client().Post("allocate", qsp, func(data []byte) error {
		return json.Unmarshal(data, ma)
	})

	return
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
