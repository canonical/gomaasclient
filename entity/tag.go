package entity

type Tag struct {
	Name        string `json:"name,omitempty"`
	Definition  string `json:"definition,omitempty"`
	Comment     string `json:"comment,omitempty"`
	KernelOpts  string `json:"kernel_opts,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

type TagParams struct {
	Name       string `json:"name"`
	Definition string `json:"definition"`
	Comment    string `json:"comment"`
	KernelOpts string `json:"kernel_opts"`
}
