package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// IPRange is an interface defining API behaviour for IP ranges
type IPRange interface {
	Get(id int) (*entity.IPRange, error)
	Update(id int, params *entity.IPRangeParams) (*entity.IPRange, error)
	Delete(id int) error
}
