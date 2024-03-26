package entity

// Node represents the MAAS Node endpoint.
type Node Machine

// NUMANode represents the MAAS numa node
// referred from https://github.com/canonical/maas/blob/deab73792a4fe839a2e84a926a6c728d510fc9ad/src/maasserver/api/machines.py#L110
type NUMANode struct {
	HugepagesSet []NUMANodeHugepages `json:"hugepages_set,omitempty"`
	Cores        []int               `json:"cores,omitempty"`
	Index        int                 `json:"index,omitempty"`
	Memory       int                 `json:"memory,omitempty"`
}

// NUMANodeHugepages represents the MAAS numa node hugepages
// referred from https://github.com/canonical/maas/blob/deab73792a4fe839a2e84a926a6c728d510fc9ad/src/maasserver/api/machines.py#L108
type NUMANodeHugepages struct {
	PageSize int `json:"page_size,omitempty"`
	Total    int `json:"total,omitempty"`
}
