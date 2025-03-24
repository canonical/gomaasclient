package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// BootResources implements api.BootResources
type BootResources struct {
	APIClient APIClient
}

func (b *BootResources) client() *APIClient {
	return b.APIClient.SubClient("boot-resources")
}

// Get fetches a list of boot resources
func (b *BootResources) Get(ctx context.Context, params *entity.BootResourcesReadParams) ([]entity.BootResource, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	bootResources := make([]entity.BootResource, 0)
	err = b.client().Get(ctx, "", qsp, func(data []byte) error {
		return json.Unmarshal(data, &bootResources)
	})

	return bootResources, err
}

// Create creates a new boot source
func (b *BootResources) Create(ctx context.Context, params *entity.BootResourceParams) (*entity.BootResource, error) {
	return nil, fmt.Errorf("not implemented")
}

// Import imports boot resources to rack controllers
func (b *BootResources) Import(ctx context.Context) error {
	return b.client().Post(ctx, "import", url.Values{}, func(data []byte) error {
		return nil
	})
}

// IsImporting returns importing status of boot resources importing to rack controllers
func (b *BootResources) IsImporting(ctx context.Context) (bool, error) {
	isImporting := new(bool)
	err := b.client().Get(ctx, "is_importing", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, isImporting)
	})

	return *isImporting, err
}

// StopImport stops importing boot resources to rack controllers
func (b *BootResources) StopImport(ctx context.Context) error {
	return b.client().Post(ctx, "stop_import", url.Values{}, func(data []byte) error {
		return nil
	})
}
