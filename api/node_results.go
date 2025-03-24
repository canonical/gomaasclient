package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// NodeResults is an interface for node script results
type NodeResults interface {
	Get(ctx context.Context, params *entity.NodeResultParams) ([]entity.NodeResult, error)
}
