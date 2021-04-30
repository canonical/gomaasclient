package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ionutbalutoiu/gomaasclient/entity"
)

// Contains functionality for manipulating the Pod entity.
type Pod struct {
	ApiClient ApiClient
}

func (p *Pod) client(id int) ApiClient {
	return p.ApiClient.GetSubObject("pods").GetSubObject(fmt.Sprintf("%v", id))
}

func (p *Pod) Get(id int) (pod *entity.Pod, err error) {
	pod = new(entity.Pod)
	err = p.client(id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, pod)
	})
	return
}

func (p *Pod) Update(id int, params *entity.PodParams) (pod *entity.Pod, err error) {
	pod = new(entity.Pod)
	err = p.client(id).Put(ToQSP(params), func(data []byte) error {
		return json.Unmarshal(data, pod)
	})
	return
}

func (p *Pod) Delete(id int) (err error) {
	err = p.client(id).Delete()
	return
}

func (p *Pod) Compose(id int, params *entity.PodMachineParams) (machine *entity.Machine, err error) {
	machine = new(entity.Machine)
	err = p.client(id).Post("compose", ToQSP(params), func(data []byte) error {
		return json.Unmarshal(data, machine)
	})
	return
}

func (p *Pod) Refresh(id int) (pod *entity.Pod, err error) {
	err = p.client(id).Post("refresh", url.Values{}, func(data []byte) error { return nil })
	return
}
