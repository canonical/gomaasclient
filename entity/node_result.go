package entity

import "encoding/json"

// ResultType referring from MAAS server
// https://github.com/maas/maas/blob/master/src/metadataserver/enum.py#L58
type ResultType int

const (
	COMMISSIONING ResultType = iota
	INSTALLATION
	TESTING
	RELEASE
)

// ResultStatus referring from MAAS server
// https://github.com/maas/maas/blob/master/src/metadataserver/enum.py#L73
type ResultStatus int

const (
	PENDING ResultStatus = iota
	RUNNING
	PASSED
	FAILED
	TIMEDOUT
	ABORTED
	DEGRADED
	INSTALLING
	FAILEDINSTALLING
	SKIPPED
	APPLYINGNETCONF
	FAILEDAPPLYINGNETCONF
)

type NodeResult struct {
	SystemID    string           `json:"system_id,omitempty"`
	TypeName    string           `json:"type_name,omitempty"`
	StatusName  string           `json:"status_name,omitempty"`
	Started     string           `json:"started,omitempty"`
	Ended       string           `json:"ended,omitempty"`
	LastPing    string           `json:"last_ping,omitempty"`
	Runtime     string           `json:"runtime,omitempty"`
	ResourceURI string           `json:"resource_uri,omitempty"`
	Results     []NodeResultItem `json:"results,omitempty"`
	Status      ResultStatus     `json:"status,omitempty"`
	Type        ResultType       `json:"type,omitempty"`
	ID          int              `json:"id,omitempty"`
}

type NodeResultItem struct {
	Created          string          `json:"created,omitempty"`
	Stderr           string          `json:"stderr,omitempty"`
	Stdout           string          `json:"stdout,omitempty"`
	Result           string          `json:"result,omitempty"`
	Name             string          `json:"name,omitempty"`
	Started          string          `json:"started,omitempty"`
	Ended            string          `json:"ended,omitempty"`
	Runtime          string          `json:"runtime,omitempty"`
	EstimatedRuntime string          `json:"estimated_runtime,omitempty"`
	Output           string          `json:"output,omitempty"`
	Updated          string          `json:"updated,omitempty"`
	StatusName       string          `json:"status_name,omitempty"`
	Parameters       json.RawMessage `json:"parameters,omitempty"`
	Starttime        json.RawMessage `json:"starttime,omitempty"`
	Endtime          json.RawMessage `json:"endtime,omitempty"`
	Status           ResultStatus    `json:"status,omitempty"`
	ID               int             `json:"id,omitempty"`
	ScriptID         int             `json:"script_id,omitempty"`
	ScriptRevisionID int             `json:"script_revision_id,omitempty"`
	ExitStatus       int             `json:"exit_status,omitempty"`
	Suppressed       bool            `json:"suppressed,omitempty"`
}

type NodeResultParams struct {
	Filters       string       `url:"filters,omitempty"`
	Type          ResultType   `url:"type,omitempty"`
	HardwareType  HardwareType `url:"hardware_type,omitempty"`
	IncludeOutput bool         `url:"include_output,omitempty"`
}
