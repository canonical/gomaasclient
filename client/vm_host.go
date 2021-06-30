package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sync"

	"github.com/google/go-querystring/query"
	"github.com/ionutbalutoiu/gomaasclient/entity"
)

// Contains functionality for manipulating the VMHost entity.
type VMHost struct {
	ApiClient ApiClient
	mutex     sync.Mutex
}

func (p *VMHost) client(id int) ApiClient {
	return p.ApiClient.GetSubObject("pods").GetSubObject(fmt.Sprintf("%v", id))
}

func (p *VMHost) Get(id int) (vmHost *entity.VMHost, err error) {
	vmHost = new(entity.VMHost)
	err = p.client(id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, vmHost)
	})
	return
}

func (p *VMHost) Update(id int, params *entity.VMHostParams) (vmHost *entity.VMHost, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}
	vmHost = new(entity.VMHost)
	err = p.client(id).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, vmHost)
	})
	return
}

func (p *VMHost) Delete(id int) (err error) {
	err = p.client(id).Delete()
	return
}

func (p *VMHost) Compose(id int, params *entity.VMHostMachineParams) (machine *entity.Machine, err error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	qsp, err := query.Values(params)
	if err != nil {
		return
	}
	machine = new(entity.Machine)
	err = p.client(id).Post("compose", qsp, func(data []byte) error {
		return json.Unmarshal(data, machine)
	})
	return
}

func (p *VMHost) Refresh(id int) (vmHost *entity.VMHost, err error) {
	err = p.client(id).Post("refresh", url.Values{}, func(data []byte) error { return nil })
	return
}

func (p *VMHost) GetParameters(id int) (vmHostParams *entity.VMHostParams, err error) {
	vmHostParams = new(entity.VMHostParams)
	err = p.client(id).Get("parameters", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, vmHostParams)
	})
	return
}
