package gmaw

import (
	"encoding/json"
	"net/url"

	"github.com/juju/gomaasapi"

	"github.com/ionutbalutoiu/gomaasclient/api/endpoint"
	"github.com/ionutbalutoiu/gomaasclient/maas"
)

// Subnets provides methods for the Subnets operations in the MaaS API.
// This type should be instantiated via NewSubnets(). It fulfills the
// api.Subnets interface.
type Subnets struct {
	client Client
}

// NewSubnets configures a new Subnets.
func NewSubnets(client *gomaasapi.MAASObject) *Subnets {
	c := client.GetSubObject("subnets")
	return &Subnets{client: Client{&c}}
}

// Get returns information about all of the configured subnets.
// This function returns an error if the gomaasapi returns an error or if
// the response cannot be decoded.
func (s *Subnets) Get() (subnets []endpoint.Subnet, err error) {
	err = s.client.Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &subnets)
	})
	return
}

// Post creates a new subnet and returns information about the new subnet.
// This function returns an error if the gomaasapi returns an error or if
// the response cannot be decoded.
func (s *Subnets) Post(p *endpoint.SubnetParams) (subnet *endpoint.Subnet, err error) {
	qsp := maas.ToQSP(p)
	subnet = new(endpoint.Subnet)
	err = s.client.Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, subnet)
	})
	return
}
