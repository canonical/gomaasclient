package api

import (
	"github.com/ionutbalutoiu/gomaasclient/entity"
)

type Spaces interface {
	Get() ([]entity.Space, error)
	Create(name string) (*entity.Space, error)
}
