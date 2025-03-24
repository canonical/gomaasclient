package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// Tags is an interface for listing and creating Tag objects
type Tags interface {
	Get(ctx context.Context) ([]entity.Tag, error)
	Create(ctx context.Context, tagParams *entity.TagParams) (*entity.Tag, error)
}
