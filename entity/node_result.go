package entity

// ResultType referring from MAAS server
// https://github.com/maas/maas/blob/master/src/metadataserver/enum.py#L56
type ResultType int

const (
	COMMISSIONING ResultType = iota
	INSTALLATION
	TESTING
	RELEASE
)

type NodeResult struct {
	Node struct {
		SystemID string `json:"system_id,omitempty"`
	} `json:"node,omitempty"`
	Name         string     `json:"name,omitempty"`
	Created      string     `json:"created,omitempty"`
	Updated      string     `json:"updated,omitempty"`
	Data         string     `json:"data,omitempty"`
	ResourceURI  string     `json:"resource_uri,omitempty"`
	ID           int        `json:"id,omitempty"`
	ScriptResult int        `json:"script_result,omitempty"`
	ResultType   ResultType `json:"result_type,omitempty"`
}

type NodeResultParams struct {
	Name []string   `url:"name,omitempty"`
	Type ResultType `url:"type,omitempty"`
}
