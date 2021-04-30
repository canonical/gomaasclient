package client

import (
	"encoding/json"
	"net/url"

	"github.com/ionutbalutoiu/gomaasclient/entity"
)

// Contains functionality for manipulating the Pods entity.
type Pods struct {
	ApiClient ApiClient
}

func (p *Pods) client() ApiClient {
	return p.ApiClient.GetSubObject("pods")
}

func (p *Pods) Get() (pods []entity.Pod, err error) {
	err = p.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &pods)
	})
	return
}

func (m *Pods) Create(params *entity.PodParams) (pod *entity.Pod, err error) {
	pod = new(entity.Pod)
	err = m.client().Post("", ToQSP(params), func(data []byte) error {
		return json.Unmarshal(data, pod)
	})
	return
}
