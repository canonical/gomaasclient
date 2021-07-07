package api

import "github.com/ionutbalutoiu/gomaasclient/entity"

type User interface {
	Get(userName string) (*entity.User, error)
	Delete(userName string) error
}
