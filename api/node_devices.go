package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// NodeDevices is an interface for listing Node Devices objects
type NodeDevices interface {
	Get(ctx context.Context, systemID string, param *entity.NodeDeviceParams) ([]entity.NodeDevice, error)
}
