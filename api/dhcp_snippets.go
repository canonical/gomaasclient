package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// DHCPSnippets is an interface for listing and creaing
// DHCPSnippet records
type DHCPSnippets interface {
	Get() ([]entity.DHCPSnippet, error)
	Create(params *entity.DHCPSnippetParams) (*entity.DHCPSnippet, error)
}
