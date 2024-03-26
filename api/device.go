package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// Device is an interface defining API behaviour for
// devices
type Device interface {
	Get(systemID string) (*entity.Device, error)
	Update(systemID string, deviceParams *entity.DeviceUpdateParams) (*entity.Device, error)
	Delete(systemID string) error
	SetWorkloadAnnotations(systemID string, params map[string]string) (*entity.Device, error)
}
