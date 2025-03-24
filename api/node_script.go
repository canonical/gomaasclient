package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// NodeScript is an interface defining API behaviour for
// node scripts
type NodeScript interface {
	Get(ctx context.Context, name string, includeScript bool) (*entity.NodeScript, error)
	Update(ctx context.Context, name string, nodeScriptParams *entity.NodeScriptParams, script []byte) (*entity.NodeScript, error)
	Delete(ctx context.Context, name string) error
	AddTag(ctx context.Context, name string, tag string) (*entity.NodeScript, error)
	RemoveTag(ctx context.Context, name string, tag string) (*entity.NodeScript, error)
	Download(ctx context.Context, name string, revision int) ([]byte, error)
	Revert(ctx context.Context, name string, to int) (*entity.NodeScript, error)
}
