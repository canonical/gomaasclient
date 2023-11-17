package api

import (
	"github.com/maas/gomaasclient/entity"
)

type Devices interface {
	Get() ([]entity.Device, error)
	Create(deviceParams *entity.DeviceCreateParams) (*entity.Device, error)
}
