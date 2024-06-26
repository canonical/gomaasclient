//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// RAID Contains functionality for manipulating the RAID entity.
type RAID struct {
	APIClient APIClient
}

func (r *RAID) client(systemID string, id int) APIClient {
	return r.APIClient.GetSubObject("nodes").GetSubObject(systemID).GetSubObject("raid").GetSubObject(fmt.Sprintf("%v", id))
}

// Get RAID details.
func (r *RAID) Get(systemID string, id int) (*entity.RAID, error) {
	raid := new(entity.RAID)
	err := r.client(systemID, id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, raid)
	})

	return raid, err
}

// Update RAID.
func (r *RAID) Update(systemID string, id int, params *entity.RAIDUpdateParams) (*entity.RAID, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	raid := new(entity.RAID)
	err = r.client(systemID, id).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, raid)
	})

	return raid, err
}

// Delete RAID.
func (r *RAID) Delete(systemID string, id int) error {
	return r.client(systemID, id).Delete()
}
