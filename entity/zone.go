package entity

// Zone represents the MAAS Zone endpoint
type Zone struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
	ID          int    `json:"id,omitempty"`
}
