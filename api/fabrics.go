package api

import (
	"github.com/ionutbalutoiu/gomaasclient/entity"
)

type Fabrics interface {
	Get() ([]entity.Fabric, error)
	Create(fabricParams *entity.FabricParams) (*entity.Fabric, error)
}
