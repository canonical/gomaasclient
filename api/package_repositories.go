package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// PackageRepositories is an interface for listing and creating
// Package Repository records
type PackageRepositories interface {
	Get(ctx context.Context) ([]entity.PackageRepository, error)
	Create(ctx context.Context, params *entity.PackageRepositoryParams) (*entity.PackageRepository, error)
}
