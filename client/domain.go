package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// Domain implements api.Domain
type Domain struct {
	APIClient APIClient
}

func (d *Domain) client(id int) APIClient {
	return d.APIClient.GetSubObject(fmt.Sprintf("domains/%v", id))
}

// Get fetches a given Domain
func (d *Domain) Get(id int) (*entity.Domain, error) {
	domain := new(entity.Domain)
	err := d.client(id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, domain)
	})

	return domain, err
}

// SetDefault sets the given Domain as the default Domain
func (d *Domain) SetDefault(id int) (*entity.Domain, error) {
	domain := new(entity.Domain)
	err := d.client(id).Post("set_default", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, domain)
	})

	return domain, err
}

// Update updates the given Domain
func (d *Domain) Update(id int, params *entity.DomainParams) (*entity.Domain, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	domain := new(entity.Domain)
	err = d.client(id).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, domain)
	})

	return domain, err
}

// Delete deletes a given Domain
func (d *Domain) Delete(id int) error {
	return d.client(id).Delete()
}
