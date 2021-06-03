package api

import (
	"github.com/ionutbalutoiu/gomaasclient/entity"
)

type Pod interface {
	Get(id int) (*entity.Pod, error)
	Update(id int, params *entity.PodParams) (*entity.Pod, error)
	Delete(id int) error
	Compose(id int, params *entity.PodMachineParams) (*entity.Machine, error)
	Refresh(id int) (*entity.Pod, error)
	GetParameters(id int) (*entity.PodParams, error)
}
