package api

import "github.com/canonical/gomaasclient/entity"

type Discovery interface {
	Get(id string) (*entity.Discovery, error)
}
