package entity

type SSLKey struct {
	Key         string `json:"key,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
	ID          int    `json:"id,omitempty"`
}
