package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sync"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// VMHost contains functionality for manipulating the VMHost entity.
type VMHost struct {
	APIClient APIClient
	mutex     sync.Mutex
}

func (p *VMHost) client(id int) APIClient {
	return p.APIClient.GetSubObject("pods").GetSubObject(fmt.Sprintf("%v", id))
}

// Get fetches a given VMHost
func (p *VMHost) Get(id int) (*entity.VMHost, error) {
	vmHost := new(entity.VMHost)
	err := p.client(id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, vmHost)
	})

	return vmHost, err
}

// Update updates a given VMHost
func (p *VMHost) Update(id int, params *entity.VMHostParams) (*entity.VMHost, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	vmHost := new(entity.VMHost)
	err = p.client(id).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, vmHost)
	})

	return vmHost, err
}

// Delete deletes a given VMHost
func (p *VMHost) Delete(id int) error {
	return p.client(id).Delete()
}

// Compose composes a VM on the given VMHost
func (p *VMHost) Compose(id int, params *entity.VMHostMachineParams) (*entity.Machine, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	machine := new(entity.Machine)
	err = p.client(id).Post("compose", qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})

	return machine, err
}

// Refresh refreshes resource info and VM status of a given VMHost
func (p *VMHost) Refresh(id int) (*entity.VMHost, error) {
	vmHost := new(entity.VMHost)
	err := p.client(id).Post("refresh", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, vmHost)
	})

	return vmHost, err
}

// GetParameters fetches the configured parameters of a given VMHost
func (p *VMHost) GetParameters(id int) (map[string]string, error) {
	params := map[string]string{}
	err := p.client(id).Get("parameters", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &params)
	})

	return params, err
}
