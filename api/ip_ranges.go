package api

import (
	"github.com/maas/gomaasclient/entity"
)

type IPRanges interface {
	Get() ([]entity.IPRange, error)
	Create(params *entity.IPRangeParams) (*entity.IPRange, error)
}
