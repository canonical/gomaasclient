package api

import (
	"github.com/maas/gomaasclient/entity"
)

type VMHosts interface {
	Get() ([]entity.VMHost, error)
	Create(params *entity.VMHostParams) (*entity.VMHost, error)
}
