package api

import (
	"github.com/canonical/gomaasclientlient/entity"
)

// SSLKeys is an interface for listing and creating SSLKey objects
type SSLKeys interface {
	Get() ([]entity.SSLKey, error)
	Create(key string) (*entity.SSLKey, error)
}
