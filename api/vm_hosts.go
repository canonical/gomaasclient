package api

import (
	"github.com/ionutbalutoiu/gomaasclient/entity"
)

type VMHosts interface {
	Get() ([]entity.VMHost, error)
	Create(params *entity.VMHostParams) (*entity.VMHost, error)
}
