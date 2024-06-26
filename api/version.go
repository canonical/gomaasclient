package api

import (
	"github.com/canonical/gomaasclientlient/entity"
)

// Version is an interface for listing MAAS version details
type Version interface {
	Get() (*entity.Version, error)
}
