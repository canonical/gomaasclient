package entity

type SSHKey struct {
	ID          int    `json:"id,omitempty"`
	Key         string `json:"key,omitempty"`
	Keysource   string `json:"keysource,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
}
