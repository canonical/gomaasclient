package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// RAID Contains functionality for manipulating the RAID entity.
type RAID struct {
	APIClient APIClient
}

func (r *RAID) client(systemID string, id int) APIClient {
	return r.APIClient.GetSubObject("nodes").GetSubObject(systemID).GetSubObject("raid").GetSubObject(fmt.Sprintf("%v", id))
}

// Get RAID details.
func (r *RAID) Get(systemID string, id int) (raid *entity.RAID, err error) {
	raid = new(entity.RAID)
	err = r.client(systemID, id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, raid)
	})

	return
}

// Update RAID.
func (r *RAID) Update(systemID string, id int, params *entity.RAIDUpdateParams) (raid *entity.RAID, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}

	raid = new(entity.RAID)
	err = r.client(systemID, id).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, raid)
	})

	return
}

// Delete RAID.
func (r *RAID) Delete(systemID string, id int) error {
	return r.client(systemID, id).Delete()
}
