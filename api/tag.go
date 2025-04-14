package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// Tag is an interface defining API behaviour for Tag objects
type Tag interface {
	Get(ctx context.Context, name string) (*entity.Tag, error)
	Update(ctx context.Context, name string, tagParams *entity.TagParams) (*entity.Tag, error)
	Delete(ctx context.Context, name string) error
	GetMachines(ctx context.Context, name string) ([]entity.Machine, error)
	AddMachines(ctx context.Context, name string, machineIds []string) error
	RemoveMachines(ctx context.Context, name string, machineIds []string) error
}
