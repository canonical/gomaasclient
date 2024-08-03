package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// DHCPSnippet is an interface defining API behaviour for
// DHCPSnippet objects
type DHCPSnippet interface {
	Get(id int) (*entity.DHCPSnippet, error)
	Update(id int, params *entity.DHCPSnippetParams) (*entity.DHCPSnippet, error)
	Delete(id int) error
}
