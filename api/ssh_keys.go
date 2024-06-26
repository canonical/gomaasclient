package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// SSHKeys is an interface for listing and creating SSHKey objects
type SSHKeys interface {
	Get() ([]entity.SSHKey, error)
	Create(key string) (*entity.SSHKey, error)
	Import(keysource string) ([]entity.SSHKey, error)
}
