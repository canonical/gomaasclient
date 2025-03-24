package api

import "context"

// MAASServer represents the MAAS Server endpoint for changing global configuration settings
type MAASServer interface {
	Get(ctx context.Context, name string) (value []byte, err error)
	Post(ctx context.Context, name, value string) error
}
