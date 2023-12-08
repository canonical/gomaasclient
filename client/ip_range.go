package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// IPRange implements api.IPRange
type IPRange struct {
	ApiClient ApiClient
}

func (i *IPRange) client(id int) ApiClient {
	return i.ApiClient.GetSubObject("ipranges").GetSubObject(fmt.Sprintf("%v", id))
}

// Get fetches a given IPRange
func (i *IPRange) Get(id int) (ipRange *entity.IPRange, err error) {
	ipRange = new(entity.IPRange)
	err = i.client(id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, ipRange)
	})

	return
}

// Update updates a given IPRange
func (i *IPRange) Update(id int, params *entity.IPRangeParams) (ipRange *entity.IPRange, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}

	ipRange = new(entity.IPRange)
	err = i.client(id).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, ipRange)
	})

	return
}

// Delete deletes a given IPRange
func (i *IPRange) Delete(id int) error {
	return i.client(id).Delete()
}
