package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// Devices is an interface for listing and creating
// Device objects
type Devices interface {
	Get(ctx context.Context) ([]entity.Device, error)
	Create(ctx context.Context, deviceParams *entity.DeviceCreateParams) (*entity.Device, error)
}
