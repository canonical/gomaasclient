package client

import (
	"encoding/json"
	"net/url"

	"github.com/ionutbalutoiu/gomaasclient/entity"
)

// Contains functionality for manipulating the Machines entity.
type Machines struct {
	ApiClient ApiClient
}

func (m *Machines) client() ApiClient {
	return m.ApiClient.GetSubObject("machines")
}

// List machines.
func (m *Machines) Get() (machines []entity.Machine, err error) {
	err = m.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &machines)
	})
	return
}

// Create machine.
func (m *Machines) Create(machineParams *entity.MachineParams, powerParams map[string]string) (ma *entity.Machine, err error) {
	qsp := ToQSP(machineParams)
	for k, v := range powerParams {
		qsp.Add(k, v)
	}
	ma = new(entity.Machine)
	err = m.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, ma)
	})
	return
}

// Allocate machine.
func (m *Machines) Allocate(params *entity.MachineAllocateParams) (ma *entity.Machine, err error) {
	ma = new(entity.Machine)
	err = m.client().Post("allocate", ToQSP(params), func(data []byte) error {
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
