package api

import "github.com/maas/gomaasclient/entity"

type Users interface {
	Get() ([]entity.User, error)
	Create(params *entity.UserParams) (*entity.User, error)
}
