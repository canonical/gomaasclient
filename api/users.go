package api

import "github.com/ionutbalutoiu/gomaasclient/entity"

type Users interface {
	Get() ([]entity.User, error)
	Create(params *entity.UserParams) (*entity.User, error)
}
