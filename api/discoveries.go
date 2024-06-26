package api

import "github.com/canonical/gomaasclient/entity"

type Discoveries interface {
	Get() ([]entity.Discovery, error)
	GetByUnknownIP() ([]entity.Discovery, error)
	GetByUnknownMAC() ([]entity.Discovery, error)
	GetByUnknownIpAndMAC() ([]entity.Discovery, error)
	Clear(params *entity.DiscoveryClearParams) error
	ClearByMACAndIP(mac, ip string) error
}
