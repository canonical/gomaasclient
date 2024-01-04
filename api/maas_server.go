package api

// MAASServer represents the MAAS Server endpoint for changing global configuration settings
type MAASServer interface {
	Get(name string) (value []byte, err error)
	Post(name, value string) error
}
