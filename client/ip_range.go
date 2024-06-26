//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// IPRange implements api.IPRange
type IPRange struct {
	APIClient APIClient
}

func (i *IPRange) client(id int) APIClient {
	return i.APIClient.GetSubObject("ipranges").GetSubObject(fmt.Sprintf("%v", id))
}

// Get fetches a given IPRange
func (i *IPRange) Get(id int) (*entity.IPRange, error) {
	ipRange := new(entity.IPRange)
	err := i.client(id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, ipRange)
	})

	return ipRange, err
}

// Update updates a given IPRange
func (i *IPRange) Update(id int, params *entity.IPRangeParams) (*entity.IPRange, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	ipRange := new(entity.IPRange)
	err = i.client(id).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, ipRange)
	})

	return ipRange, err
}

// Delete deletes a given IPRange
func (i *IPRange) Delete(id int) error {
	return i.client(id).Delete()
}
