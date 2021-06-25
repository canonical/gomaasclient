package api

import (
	"github.com/ionutbalutoiu/gomaasclient/entity"
)

type VLAN interface {
	Get(fabricID int, vid int) (*entity.VLAN, error)
	Update(fabricID int, vid int, params *entity.VLANParams) (*entity.VLAN, error)
	Delete(fabricID int, vid int) error
}
