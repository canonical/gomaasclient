package gmaw

import (
	"fmt"
	"net/url"

	"github.com/juju/gomaasapi"

	"github.com/ionutbalutoiu/gomaasclient/api/endpoint"
	"github.com/ionutbalutoiu/gomaasclient/maas"
)

// NewPod returns a pointer to a Pod.
func NewPod(client *gomaasapi.MAASObject) *Pod {
	return &Pod{client: client}
}

// Pod implements the maas.PodFetcher interface.
type Pod struct {
	client *gomaasapi.MAASObject
}

// callPost returns the raw response from the MAAS API and any errors.
func (m *Pod) callPost(id int, op string, qsp url.Values) ([]byte, error) {
	mc := m.client.GetSubObject("pods").GetSubObject(fmt.Sprintf("%v", id))
	res, err := mc.CallPost(op, qsp)
	if err != nil {
		return nil, err
	}

	return res.GetBytes()
}

// Get fulfills the maas.PodFetcher interface
func (m *Pod) Get(id int) ([]byte, error) {
	mc := m.client.GetSubObject("pods").GetSubObject(fmt.Sprintf("%v", id))
	res, err := mc.CallGet("", url.Values{})
	if err != nil {
		return nil, err
	}

	return res.GetBytes()
}

// Calls the update Pod API
func (m *Pod) Update(id int, params *endpoint.PodParams) ([]byte, error) {
	mc := m.client.GetSubObject("pods").GetSubObject(fmt.Sprintf("%v", id))
	maasObj, err := mc.Update(maas.ToQSP(params))
	if err != nil {
		return nil, err
	}

	return maasObj.MarshalJSON()
}

// Calls the delete Pod API
func (m *Pod) Delete(id int) error {
	mc := m.client.GetSubObject("pods").GetSubObject(fmt.Sprintf("%v", id))
	return mc.Delete()
}

// Compose fulfills the maas.PodFetcher interface
func (m *Pod) Compose(id int, params *endpoint.PodMachineParams) ([]byte, error) {
	qsp := maas.ToQSP(params)
	return m.callPost(id, "compose", qsp)
}

// Refresh fulfills the maas.PodFetcher interface
func (m *Pod) Refresh(id int) ([]byte, error) {
	return m.callPost(id, "refresh", url.Values{})
}
