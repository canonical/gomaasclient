package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// NodeScripts is an interface for listing and creating
// node script objects
type NodeScripts interface {
	Get(nodeScriptParams *entity.NodeScriptReadParams) ([]entity.NodeScript, error)
	Create(nodeScriptParams *entity.NodeScriptParams, script []byte) (*entity.NodeScript, error)
}
