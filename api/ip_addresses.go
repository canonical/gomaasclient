package api

import "github.com/maas/gomaasclient/entity"

type IPAddresses interface {
	Get(params *entity.IPAddressesParams) ([]entity.IPAddress, error)
	Release(params *entity.IPAddressesParams) error
	Reserve(params *entity.IPAddressesParams) (*entity.IPAddress, error)
}
