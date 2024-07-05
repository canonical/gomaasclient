package entity

type DHCPSnippet struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	ResourceURI string `json:"resource_uri,omitempty"`
	Value       string `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
	Enabled     bool   `json:"enabled"`
	Node        Node   `json:"node,omitempty"`
	Subnet      Subnet `json:"subnet,omitempty"`
	Global      bool   `json:"global_snippet"`
}

type DHCPSnippetParams struct {
	Name        string `url:"name,omitempty"`
	Value       string `url:"value,omitempty"`
	Description string `url:"description,omitempty"`
	Enabled     bool   `url:"enabled"`
	Node        string `url:"node,omitempty"`
	Subnet      int    `url:"subnet,omitempty"`
	Global      bool   `url:"global_snippet"`
}
