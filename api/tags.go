package api

import (
	"github.com/maas/gomaasclient/entity"
)

type Tags interface {
	Get() ([]entity.Tag, error)
	Create(tagParams *entity.TagParams) (*entity.Tag, error)
}
