package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// PackageRepository represents the MAAS Server Package Repository endpoint
type PackageRepository interface {
	Get(ctx context.Context, id int) (*entity.PackageRepository, error)
	Update(ctx context.Context, id int, params *entity.PackageRepositoryParams) (*entity.PackageRepository, error)
	Delete(ctx context.Context, id int) error
}
