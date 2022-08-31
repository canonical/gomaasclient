package api

import (
	"github.com/maas/gomaasclient/entity"
)

type VMHost interface {
	Get(id int) (*entity.VMHost, error)
	Update(id int, params *entity.VMHostParams) (*entity.VMHost, error)
	Delete(id int) error
	Compose(id int, params *entity.VMHostMachineParams) (*entity.Machine, error)
	Refresh(id int) (*entity.VMHost, error)
	GetParameters(id int) (map[string]string, error)
}
