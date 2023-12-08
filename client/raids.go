package client

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"

	"github.com/maas/gomaasclient/entity"
)

// RAIDs contains functionality for manipulating the RAIDs entity.
type RAIDs struct {
	ApiClient ApiClient
}

func (r *RAIDs) client(systemID string) ApiClient {
	return r.ApiClient.GetSubObject("nodes").GetSubObject(systemID).GetSubObject("raids")
}

// Get RAIDs of a machine.
func (r *RAIDs) Get(systemID string) (raids []entity.RAID, err error) {
	err = r.client(systemID).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &raids)
	})

	return
}

// Create a RAID of a machine.
func (r *RAIDs) Create(systemID string, params *entity.RAIDCreateParams) (raid *entity.RAID, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}

	raid = new(entity.RAID)
	err = r.client(systemID).Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, raid)
	})

	return
}
