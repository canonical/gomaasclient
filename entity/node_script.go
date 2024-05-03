package entity

type NodeScriptType int

// NodeScriptType referring from MAAS server
// https://github.com/canonical/maas/blob/deab73792a4fe839a2e84a926a6c728d510fc9ad/src/metadataserver/enum.py#L42
const (
	ScriptTypeCommissioning NodeScriptType = iota
	_
	ScriptTypeTesting
	ScriptTypeRelease
)

type NodeScriptHardwareType int

// NodeScriptHardwareType referring from MAAS server
// https://github.com/canonical/maas/blob/deab73792a4fe839a2e84a926a6c728d510fc9ad/src/metadataserver/enum.py#L126
const (
	ScriptHardwareTypeNode NodeScriptHardwareType = iota
	ScriptHardwareTypeCPU
	ScriptHardwareTypeMemory
	ScriptHardwareTypeStorage
	ScriptHardwareTypeNetwork
	ScriptHardwareTypeGPU
)

type NodeScriptParallel int

// NodeScriptParallel referring from MAAS server
// https://github.com/canonical/maas/blob/deab73792a4fe839a2e84a926a6c728d510fc9ad/src/metadataserver/enum.py#L146
const (
	ScriptParallelDisabled NodeScriptParallel = iota
	ScriptParallelInstance
	ScriptParallelAny
)

// NodeScript represents the MAAS Node Script endpoint.
type NodeScript struct {
	Packages                  map[string]interface{} `json:"packages,omitempty"`
	Results                   map[string]interface{} `json:"results,omitempty"`
	Parameters                map[string]interface{} `json:":parameters,omitempty"`
	TypeName                  string                 `json:"type_name,omitempty"`
	HardwareTypeName          string                 `json:"hardware_type_name,omitempty"`
	ParallelName              string                 `json:"parallel_name,omitempty"`
	Timeout                   string                 `json:"timeout,omitempty"`
	ResourceURI               string                 `json:"resource_uri,omitempty"`
	Name                      string                 `json:"name,omitempty"`
	Description               string                 `json:"description,omitempty"`
	Title                     string                 `json:"title,omitempty"`
	ForHardware               []string               `json:"for_hardware,omitempty"`
	Tags                      []string               `json:"tags,omitempty"`
	History                   []NodeScriptHistory    `json:"history,omitempty"`
	HardwareType              NodeScriptHardwareType `json:"hardware_type,omitempty"`
	Type                      NodeScriptType         `json:"type,omitempty"`
	Parallel                  NodeScriptParallel     `json:"parallel,omitempty"`
	ID                        int                    `json:"id,omitempty"`
	ApplyConfiguredNetworking bool                   `json:"apply_configured_networking,omitempty"`
	Default                   bool                   `json:"default,omitempty"`
	Recommission              bool                   `json:"recommission,omitempty"`
	MayReboot                 bool                   `json:"may_reboot,omitempty"`
	Destructive               bool                   `json:"destructive,omitempty"`
}

// NodeScriptHistory represents a NodeScript's "history" list item.
// This type should not be used directly.
type NodeScriptHistory struct {
	Comment string `json:"comment,omitempty"`
	Created string `json:"created,omitempty"`
	Data    string `json:"data,omitempty"`
	ID      int    `json:"id,omitempty"`
}

type NodeScriptParams struct {
	ScriptType                string   `url:"script_type,omitempty"`
	HardwareType              string   `url:"hardware_type,omitempty"`
	Parallel                  string   `url:"parallel,omitempty"`
	Packages                  string   `url:"packages,omitempty"`
	Timeout                   string   `url:"timeout,omitempty"`
	Comment                   string   `url:"comment,omitempty"`
	ForHardware               string   `url:"for_hardware,omitempty"`
	Name                      string   `url:"name,omitempty"`
	Title                     string   `url:"title,omitempty"`
	Description               string   `url:"description,omitempty"`
	Tags                      []string `url:"tags,omitempty"`
	ApplyConfiguredNetworking bool     `url:"apply_configured_networking,omitempty"`
	Destructive               bool     `url:"destructive,omitempty"`
	MayReboot                 bool     `url:"may_reboot,omitempty"`
	Recommission              bool     `url:"recommission,omitempty"`
}

type NodeScriptReadParams struct {
	Type          string `url:"type,omitempty"`
	HardwareType  string `url:"hardware_type,omitempty"`
	IncludeScript string `url:"include_script,omitempty"`
	Filters       string `url:"filters,omitempty"`
}
