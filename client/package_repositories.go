//nolint:dupl // disable dupl check on client for now
package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// PackageRepositories implements api.PackageRepositories
type PackageRepositories struct {
	APIClient APIClient
}

func (p *PackageRepositories) client() *APIClient {
	return p.APIClient.SubClient("package-repositories")
}

// Get fetches a list of PackageRepositories
func (p *PackageRepositories) Get(ctx context.Context) ([]entity.PackageRepository, error) {
	packageRepositories := make([]entity.PackageRepository, 0)
	err := p.client().Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &packageRepositories)
	})

	return packageRepositories, err
}

// Create creates a new PackageRepository
func (p *PackageRepositories) Create(ctx context.Context, packageRepositoryParams *entity.PackageRepositoryParams) (*entity.PackageRepository, error) {
	qsp, err := query.Values(packageRepositoryParams)
	if err != nil {
		return nil, err
	}

	packageRepository := new(entity.PackageRepository)
	err = p.client().Post(ctx, "", qsp, func(data []byte) error {
		return json.Unmarshal(data, packageRepository)
	})

	return packageRepository, err
}
