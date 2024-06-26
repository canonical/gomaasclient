//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// PackageRepository implements api.PackageRepository
type PackageRepository struct {
	APIClient APIClient
}

func (p *PackageRepository) client(id int) APIClient {
	return p.APIClient.GetSubObject("package-repositories").GetSubObject(fmt.Sprintf("%v", id))
}

// Get fetches a packageRepository
func (p *PackageRepository) Get(id int) (*entity.PackageRepository, error) {
	packageRepository := new(entity.PackageRepository)
	err := p.client(id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, packageRepository)
	})

	return packageRepository, err
}

// Update updates a given PackageRepository
func (p *PackageRepository) Update(id int, packageRepositoryParams *entity.PackageRepositoryParams) (*entity.PackageRepository, error) {
	qsp, err := query.Values(packageRepositoryParams)
	if err != nil {
		return nil, err
	}

	packageRepository := new(entity.PackageRepository)
	err = p.client(id).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, packageRepository)
	})

	return packageRepository, err
}

// Delete deletes a given PackageRepository
func (p *PackageRepository) Delete(id int) error {
	return p.client(id).Delete()
}
