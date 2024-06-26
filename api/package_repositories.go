package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// PackageRepositories is an interface for listing and creating
// Package Repository records
type PackageRepositories interface {
	Get() ([]entity.PackageRepository, error)
	Create(params *entity.PackageRepositoryParams) (*entity.PackageRepository, error)
}
