package gmaw

import (
	"net/url"

	"github.com/juju/gomaasapi"

	"github.com/ionutbalutoiu/gomaasclient/api/endpoint"
	"github.com/ionutbalutoiu/gomaasclient/maas"
)

// Machine implements the maas.MachineFetcher interface.
// It exposes the functionality of the MAAS Machine endpoint (eg machines/{systemid}),
// and can be used with as many machines as desired since each method call takes the
// systemID as a parameter. Multiple instances of this type are only necessary to
// support multiple clients.
type Machine struct {
	client *gomaasapi.MAASObject
}

// NewMachine returns a pointer to a Machine.
func NewMachine(client *gomaasapi.MAASObject) *Machine {
	return &Machine{client: client}
}

// callPost returns the raw response from the MAAS API and any errors.
// This method creates the appropriate MAASObject for the API call, invokes the
// CallPost function, and returns the GetBytes() method of the response. It will
// return a nil byte array if CallPost returns an error.
func (m *Machine) callPost(systemID, op string, qsp url.Values) ([]byte, error) {
	mc := m.client.GetSubObject("machines").GetSubObject(systemID)
	res, err := mc.CallPost(op, qsp)
	if err != nil {
		return nil, err
	}

	return res.GetBytes()
}

// Get fulfills the maas.MachineFetcher interface
func (m *Machine) Get(systemID string) ([]byte, error) {
	mc := m.client.GetSubObject("machines").GetSubObject(systemID)
	res, err := mc.CallGet("", url.Values{})
	if err != nil {
		return nil, err
	}

	return res.GetBytes()
}

func (m *Machine) Update(systemID string, params *endpoint.MachineParams) ([]byte, error) {
	mc := m.client.GetSubObject("machines").GetSubObject(systemID)
	maasObj, err := mc.Update(maas.ToQSP(params))
	if err != nil {
		return nil, err
	}

	return maasObj.MarshalJSON()
}

func (m *Machine) Delete(systemID string) error {
	mc := m.client.GetSubObject("machines").GetSubObject(systemID)
	return mc.Delete()
}

// Commission fulfills the maas.MachineFetcher interface
func (m *Machine) Commission(systemID string, params *endpoint.MachineCommissionParams) ([]byte, error) {
	qsp := maas.ToQSP(params)
	return m.callPost(systemID, "commission", qsp)
}

// Deploy fulfills the maas.MachineFetcher interface
func (m *Machine) Deploy(systemID string, params *endpoint.MachineDeployParams) ([]byte, error) {
	qsp := maas.ToQSP(params)
	return m.callPost(systemID, "deploy", qsp)
}

// Lock fulfills the maas.MachineFetcher interface
func (m *Machine) Lock(systemID, comment string) ([]byte, error) {
	qsp := make(url.Values)
	if comment != "" {
		qsp.Set("comment", comment)
	}
	return m.callPost(systemID, "lock", qsp)
}
