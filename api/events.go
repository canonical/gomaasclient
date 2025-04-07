package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// Events is an interface for node events
type Events interface {
	Get(ctx context.Context, params *entity.EventParams) (*entity.EventsResp, error)
}
