package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

type Discovery interface {
	Get(ctx context.Context, id string) (*entity.Discovery, error)
}
