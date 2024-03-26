package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// PackageRepository represents the MAAS Server Package Repository endpoint
type PackageRepository interface {
	Get(id int) (*entity.PackageRepository, error)
	Update(id int, params *entity.PackageRepositoryParams) (*entity.PackageRepository, error)
	Delete(id int) error
}
