package entity

type SSHKey struct {
	Key         string `json:"key,omitempty"`
	Keysource   string `json:"keysource,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
	ID          int    `json:"id,omitempty"`
}
