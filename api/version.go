package api

import (
	"github.com/maas/gomaasclient/entity"
)

// Version is an interface for listing MAAS version details
type Version interface {
	Get() (*entity.Version, error)
}
