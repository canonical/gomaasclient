package client

import (
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

func (b *BootResources) client() APIClient {
	return b.APIClient.GetSubObject("boot-resources")
}

// Get fetches a list of boot resources
func (b *BootResources) Get(params *entity.BootResourcesReadParams) ([]entity.BootResource, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	bootResources := make([]entity.BootResource, 0)
	err = b.client().Get("", qsp, func(data []byte) error {
		return json.Unmarshal(data, &bootResources)
	})

	return bootResources, err
}

// Create creates a new boot source
func (b *BootResources) Create(params *entity.BootResourceParams) (*entity.BootResource, error) {
	return nil, fmt.Errorf("not implemented")
}

// Import imports boot resources to rack controllers
func (b *BootResources) Import() error {
	return b.client().Post("import", url.Values{}, func(data []byte) error {
		return nil
	})
}

// IsImporting returns importing status of boot resources importing to rack controllers
func (b *BootResources) IsImporting() (bool, error) {
	isImporting := new(bool)
	err := b.client().Get("is-importing", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, isImporting)
	})

	return *isImporting, err
}

// StopImport stops importing boot resources to rack controllers
func (b *BootResources) StopImport() error {
	return b.client().Post("stop-import", url.Values{}, func(data []byte) error {
		return nil
	})
}
