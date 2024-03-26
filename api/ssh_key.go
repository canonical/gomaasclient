package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// SSHKey is an interface defining API behaviour for SSHKey objects
type SSHKey interface {
	Get(id int) (*entity.SSHKey, error)
	Delete(id int) error
}
