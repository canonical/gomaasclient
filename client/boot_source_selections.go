package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// BootSourceSelections implements api.BootSourceSelections
type BootSourceSelections struct {
	APIClient APIClient
}

func (b *BootSourceSelections) client(bootSourceID int) *APIClient {
	return b.APIClient.SubClient("boot-sources").
		SubClient(fmt.Sprintf("%v", bootSourceID)).
		SubClient("selections")
}

// Get fetches a list of BootSourceSelection objects
func (b *BootSourceSelections) Get(ctx context.Context, bootSourceID int) ([]entity.BootSourceSelection, error) {
	bootSourceSelections := make([]entity.BootSourceSelection, 0)
	err := b.client(bootSourceID).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &bootSourceSelections)
	})

	return bootSourceSelections, err
}

// Create creates a BootSourceSelection object
func (b *BootSourceSelections) Create(ctx context.Context, bootSourceID int, params *entity.BootSourceSelectionParams) (*entity.BootSourceSelection, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	bootSourceSelection := new(entity.BootSourceSelection)
	err = b.client(bootSourceID).Post(ctx, "", qsp, func(data []byte) error {
		return json.Unmarshal(data, bootSourceSelection)
	})

	return bootSourceSelection, err
}
