package api

import (
	"github.com/maas/gomaasclient/entity"
)

type Device interface {
	Get(systemID string) (*entity.Device, error)
	Update(systemID string, deviceParams *entity.DeviceUpdateParams) (*entity.Device, error)
	Delete(systemID string) error
}
