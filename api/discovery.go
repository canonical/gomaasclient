package api

import "github.com/maas/gomaasclient/entity"

type Discovery interface {
	Get(id string) (*entity.Discovery, error)
}
