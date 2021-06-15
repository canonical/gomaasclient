package api

import (
	"github.com/ionutbalutoiu/gomaasclient/entity"
)

type BlockDevices interface {
	Get(systemID string) ([]entity.BlockDevice, error)
	Create(systemID string, params *entity.BlockDeviceParams) (*entity.BlockDevice, error)
}
