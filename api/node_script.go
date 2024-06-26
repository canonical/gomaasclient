package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// NodeScript is an interface defining API behaviour for
// node scripts
type NodeScript interface {
	Get(name string, includeScript bool) (*entity.NodeScript, error)
	Update(name string, nodeScriptParams *entity.NodeScriptParams, script []byte) (*entity.NodeScript, error)
	Delete(name string) error
	AddTag(name string, tag string) (*entity.NodeScript, error)
	RemoveTag(name string, tag string) (*entity.NodeScript, error)
	Download(name string, revision int) ([]byte, error)
	Revert(name string, to int) (*entity.NodeScript, error)
}
