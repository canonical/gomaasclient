package api

import (
	"github.com/maas/gomaasclient/entity"
)

type Spaces interface {
	Get() ([]entity.Space, error)
	Create(name string) (*entity.Space, error)
}
