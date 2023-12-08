package entity

// ResourcePool represents the MAAS ResourcePool endpoint
type ResourcePool struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

type ResourcePoolParams struct {
	Name        string `url:"name,omitempty"`
	Description string `url:"description,omitempty"`
}
