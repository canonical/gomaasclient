package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// NodeScripts is an interface for listing and creating
// node script objects
type NodeScripts interface {
	Get(ctx context.Context, nodeScriptParams *entity.NodeScriptReadParams) ([]entity.NodeScript, error)
	Create(ctx context.Context, nodeScriptParams *entity.NodeScriptParams, script []byte) (*entity.NodeScript, error)
}
