package gmaw

import (
	"encoding/json"
	"net/url"

	"github.com/juju/gomaasapi"

	"github.com/ionutbalutoiu/gomaasclient/api/endpoint"
	"github.com/ionutbalutoiu/gomaasclient/maas"
)

// NetworkInterfaces provides methods for the Interfaces operations in the MaaS API.
// This type should be instantiated via NewNetworkInterfaces(). It fulfills the
// api.NetworkInterfaces interface.
type NetworkInterfaces struct {
	c Client
}

// NewNetworkInterfaces configures a new NetworkInterfaces.
func NewNetworkInterfaces(client *gomaasapi.MAASObject) *NetworkInterfaces {
	c := client.GetSubObject("nodes")
	return &NetworkInterfaces{c: Client{&c}}
}

// client returns a Client with the MAASObject that correlates to the correct endpoint.
func (i *NetworkInterfaces) client(systemID string) Client {
	return i.c.GetSubObject(systemID).GetSubObject("interfaces")
}

// Get returns information about all of <systemID>'s configured interfaces.
// This function returns an error if the gomaasapi returns an error or if
// the response cannot be decoded.
func (i *NetworkInterfaces) Get(systemID string) (ifcs []endpoint.NetworkInterface, err error) {
	err = i.client(systemID).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &ifcs)
	})
	return
}

// CreateBond creates a new bond interface on <systemID>.
// This function returns an error if the gomaasapi returns an error or if
// the response cannot be decoded.
func (i *NetworkInterfaces) CreateBond(systemID string,
	p *endpoint.NetworkInterfaceBondParams) (ifc *endpoint.NetworkInterface, err error) {
	qsp := maas.ToQSP(p)
	ifc = new(endpoint.NetworkInterface)
	err = i.client(systemID).Post("create_bond", qsp, func(data []byte) error {
		return json.Unmarshal(data, ifc)
	})
	return
}

// CreateBridge creates a new bridge interface on <systemID>.
// This function returns an error if the gomaasapi returns an error or if
// the response cannot be decoded.
func (i *NetworkInterfaces) CreateBridge(systemID string,
	p *endpoint.NetworkInterfaceBridgeParams) (ifc *endpoint.NetworkInterface, err error) {
	qsp := maas.ToQSP(p)
	ifc = new(endpoint.NetworkInterface)
	err = i.client(systemID).Post("create_bridge", qsp, func(data []byte) error {
		return json.Unmarshal(data, ifc)
	})
	return
}

// CreatePhysical creates a new physical interface on <systemID>.
// This function returns an error if the gomaasapi returns an error or if
// the response cannot be decoded.
func (i *NetworkInterfaces) CreatePhysical(systemID string,
	p *endpoint.NetworkInterfacePhysicalParams) (ifc *endpoint.NetworkInterface, err error) {
	qsp := maas.ToQSP(p)
	ifc = new(endpoint.NetworkInterface)
	err = i.client(systemID).Post("create_physical", qsp, func(data []byte) error {
		return json.Unmarshal(data, ifc)
	})
	return
}

// CreateVLAN creates a new VLAN interface on <systemID>.
// This function returns an error if the gomaasapi returns an error or if
// the response cannot be decoded.
func (i *NetworkInterfaces) CreateVLAN(systemID string,
	p *endpoint.NetworkInterfaceVLANParams) (ifc *endpoint.NetworkInterface, err error) {
	qsp := maas.ToQSP(p)
	ifc = new(endpoint.NetworkInterface)
	err = i.client(systemID).Post("create_vlan", qsp, func(data []byte) error {
		return json.Unmarshal(data, ifc)
	})
	return
}
