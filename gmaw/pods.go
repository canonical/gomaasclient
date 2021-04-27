package gmaw

import (
	"net/url"

	"github.com/juju/gomaasapi"

	"github.com/ionutbalutoiu/gomaasclient/api/endpoint"
	"github.com/ionutbalutoiu/gomaasclient/maas"
)

// NewPods returns a pointer to a Pods.
func NewPods(client *gomaasapi.MAASObject) *Pods {
	return &Pods{client: client}
}

// Pods implements the maas.PodsFetcher interface.
type Pods struct {
	client *gomaasapi.MAASObject
}

// Get fulfills the maas.PodsFetcher interface
func (m *Pods) Get() ([]byte, error) {
	mc := m.client.GetSubObject("pods")
	res, err := mc.CallGet("", url.Values{})
	if err != nil {
		return nil, err
	}

	return res.GetBytes()
}

// Create fulfills the maas.PodsFetcher interface
func (m *Pods) Create(params *endpoint.PodParams) ([]byte, error) {
	mc := m.client.GetSubObject("pods")
	res, err := mc.CallPost("", maas.ToQSP(params))
	if err != nil {
		return nil, err
	}

	return res.GetBytes()
}
