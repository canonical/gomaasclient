package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// Device is an interface defining API behaviour for
// devices
type Device interface {
	Get(ctx context.Context, systemID string) (*entity.Device, error)
	Update(ctx context.Context, systemID string, deviceParams *entity.DeviceUpdateParams) (*entity.Device, error)
	Delete(ctx context.Context, systemID string) error
	SetWorkloadAnnotations(ctx context.Context, systemID string, params map[string]string) (*entity.Device, error)
}
