package api

import (
	"github.com/ionutbalutoiu/gomaasclient/entity"
)

type Domains interface {
	Get() ([]entity.Domain, error)
	Create(params *entity.DomainParams) (*entity.Domain, error)
	SetSerial(serial int) error
}
