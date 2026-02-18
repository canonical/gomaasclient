package client

import (
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// ReservedIPs implements api.ReservedIPs
type ReservedIPs struct {
	APIClient APIClient
}

func (r *ReservedIPs) client() APIClient {
	return r.APIClient.GetSubObject("reservedips")
}

// Get fetches a list of Reserved IP objects
func (r *ReservedIPs) Get() ([]entity.ReservedIP, error) {
	reservedIPs := make([]entity.ReservedIP, 0)
	err := r.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &reservedIPs)
	})

	return reservedIPs, err
}

// Create creates a new Reserved IP
func (r *ReservedIPs) Create(params *entity.ReservedIPCreateParams) (*entity.ReservedIP, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	reservedIP := new(entity.ReservedIP)
	err = r.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, reservedIP)
	})

	return reservedIP, err
}
