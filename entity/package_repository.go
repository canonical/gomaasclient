package entity

type PackageRepository struct {
	Name               string   `json:"name,omitempty"`
	URL                string   `json:"url,omitempty"`
	Key                string   `json:"key,omitempty"`
	ResourceURI        string   `json:"resource_uri,omitempty"`
	Distributions      []string `json:"distributions,omitempty"`
	DisabledPockets    []string `json:"disabled_pockets,omitempty"`
	DisabledComponents []string `json:"disabled_components,omitempty"`
	Components         []string `json:"components,omitempty"`
	Arches             []string `json:"arches,omitempty"`
	ID                 int      `json:"id,omitempty"`
	DisableSources     bool     `json:"disable_sources,omitempty"`
	Enabled            bool     `json:"enabled,omitempty"`
}

type PackageRepositoryParams struct {
	Name               string `url:"name,omitempty"`
	URL                string `url:"url,omitempty"`
	Distributions      string `url:"distributions,omitempty"`
	DisabledPockets    string `url:"disabled_pockets,omitempty"`
	DisabledComponents string `url:"disabled_components,omitempty"`
	Components         string `url:"components,omitempty"`
	Arches             string `url:"arches,omitempty"`
	Key                string `url:"key,omitempty"`
	DisableSources     bool   `url:"disable_sources,omitempty"`
	Enabled            bool   `url:"enabled,omitempty"`
}
