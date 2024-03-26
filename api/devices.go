package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// Devices is an interface for listing and creating
// Device objects
type Devices interface {
	Get() ([]entity.Device, error)
	Create(deviceParams *entity.DeviceCreateParams) (*entity.Device, error)
}
