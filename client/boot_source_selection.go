package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// BootSourceSelection implements api.BootSourceSelection
type BootSourceSelection struct {
	APIClient APIClient
}

func (b *BootSourceSelection) client(bootSourceID int, id int) *APIClient {
	return b.APIClient.SubClient("boot-sources").
		SubClient(fmt.Sprintf("%v", bootSourceID)).
		SubClient("selections").
		SubClient(fmt.Sprintf("%v", id))
}

// Get fetches a BootSourceSelection for the given bootSourceID and BootSourceSelection id
func (b *BootSourceSelection) Get(ctx context.Context, bootSourceID int, id int) (*entity.BootSourceSelection, error) {
	bootSourceSelection := new(entity.BootSourceSelection)
	err := b.client(bootSourceID, id).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, bootSourceSelection)
	})

	return bootSourceSelection, err
}

// Update updates a given BootSourceSelection
func (b *BootSourceSelection) Update(ctx context.Context, bootSourceID int, id int, params *entity.BootSourceSelectionParams) (*entity.BootSourceSelection, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	bootSourceSelection := new(entity.BootSourceSelection)
	err = b.client(bootSourceID, id).Put(ctx, qsp, func(data []byte) error {
		return json.Unmarshal(data, bootSourceSelection)
	})

	return bootSourceSelection, err
}

// Delete deletes a given BootSourceSelection
func (b *BootSourceSelection) Delete(ctx context.Context, bootSourceID int, id int) error {
	return b.client(bootSourceID, id).Delete(ctx)
}
