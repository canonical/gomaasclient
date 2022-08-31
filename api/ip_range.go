package api

import (
	"github.com/maas/gomaasclient/entity"
)

type IPRange interface {
	Get(id int) (*entity.IPRange, error)
	Update(id int, params *entity.IPRangeParams) (*entity.IPRange, error)
	Delete(id int) error
}
