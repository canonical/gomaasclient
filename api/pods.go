package api

import (
	"github.com/ionutbalutoiu/gomaasclient/entity"
)

type Pods interface {
	Get() ([]entity.Pod, error)
	Create(params *entity.PodParams) (*entity.Pod, error)
}
