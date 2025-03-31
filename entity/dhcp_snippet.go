package entity

type DHCPSnippet struct {
	Name        string `json:"name,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
	Value       string `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
	Subnet      Subnet `json:"subnet,omitempty"`
	Node        Node   `json:"node,omitempty"`
	ID          int    `json:"id,omitempty"`
	Global      bool   `json:"global_snippet"`
	Enabled     bool   `json:"enabled"`
}

type DHCPSnippetParams struct {
	Name        string `url:"name,omitempty"`
	Value       string `url:"value,omitempty"`
	Description string `url:"description,omitempty"`
	Node        string `url:"node,omitempty"`
	Subnet      int    `url:"subnet,omitempty"`
	Enabled     bool   `url:"enabled"`
	Global      bool   `url:"global_snippet"`
}
