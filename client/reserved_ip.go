//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// ReservedIP implements api.ReservedIP
type ReservedIP struct {
	APIClient APIClient
}

func (r *ReservedIP) client(id int) APIClient {
	return r.APIClient.GetSubObject(fmt.Sprintf("reservedips/%v", id))
}

// Get fetches a given Reserved IP
func (r *ReservedIP) Get(id int) (*entity.ReservedIP, error) {
	reservedIP := new(entity.ReservedIP)
	err := r.client(id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, reservedIP)
	})

	return reservedIP, err
}

// Update updates the given Reserved IP
func (r *ReservedIP) Update(id int, params *entity.ReservedIPUpdateParams) (*entity.ReservedIP, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	reservedIP := new(entity.ReservedIP)
	err = r.client(id).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, reservedIP)
	})

	return reservedIP, err
}

// Delete deletes a given Reserved IP
func (r *ReservedIP) Delete(id int) error {
	return r.client(id).Delete()
}
