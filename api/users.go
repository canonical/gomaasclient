package api

import "github.com/canonical/gomaasclient/entity"

// Users is an interface for listing and creating users
type Users interface {
	Get() ([]entity.User, error)
	Create(params *entity.UserParams) (*entity.User, error)
}
