package api

import (
	"github.com/ionutbalutoiu/gomaasclient/entity"
)

type Space interface {
	Get(id int) (*entity.Space, error)
	Update(id int, name string) (*entity.Space, error)
	Delete(id int) error
}
