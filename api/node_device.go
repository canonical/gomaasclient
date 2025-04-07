package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// NodeDevice is an interface defining API behaviour for Node Devices
type NodeDevice interface {
	Get(ctx context.Context, systemID string, id int) (*entity.NodeDevice, error)
	Delete(ctx context.Context, systemID string, id int) error
}
