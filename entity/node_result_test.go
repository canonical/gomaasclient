package entity

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/canonical/gomaasclient/test/helper"
)

var sampleNodeResults []NodeResult = []NodeResult{
	{
		SystemID:    "qmrdsx",
		TypeName:    "Release",
		StatusName:  "Passed",
		Started:     "Tue, 21 Oct 2025 15:17:12 +0000",
		Ended:       "Tue, 21 Oct 2025 15:17:12 +0000",
		LastPing:    "Tue, 21 Oct 2025 15:17:12 +0000",
		Runtime:     "0:00:00",
		ResourceURI: "/MAAS/api/2.0/nodes/qmrdsx/results/11/",
		Results: []NodeResultItem{
			{
				Stdout:           "SGVsbG8gV29ybGQK",
				StatusName:       "Passed",
				Updated:          "Tue, 21 Oct 2025 15:17:12 +0000",
				Result:           "",
				Stderr:           "",
				Started:          "Tue, 21 Oct 2025 15:17:12 +0000",
				Ended:            "Tue, 21 Oct 2025 15:17:12 +0000",
				Runtime:          "0:00:00",
				EstimatedRuntime: "0:00:00",
				Output:           "SGVsbG8gV29ybGQK",
				Created:          "Tue, 21 Oct 2025 15:16:19 +0000",
				Name:             "testname",
				Parameters:       []byte("{}"),
				Status:           PASSED,
				ID:               56,
				ScriptID:         32,
				ScriptRevisionID: 33,
				ExitStatus:       0,
				Starttime:        []byte("1761059832.783871"),
				Endtime:          []byte("1761059832.844518"),
				Suppressed:       false,
			},
		},
		Status: PASSED,
		Type:   RELEASE,
		ID:     11,
	},
	{
		SystemID:    "qmrdsx",
		TypeName:    "Release",
		StatusName:  "Aborted",
		Started:     "",
		Ended:       "",
		LastPing:    "",
		Runtime:     "",
		ResourceURI: "/MAAS/api/2.0/nodes/qmrdsx/results/196/",
		Results: []NodeResultItem{
			{
				StatusName:       "Aborted",
				Updated:          "Wed, 22 Oct 2025 14:27:01 +0000",
				Started:          "",
				Ended:            "",
				Runtime:          "",
				EstimatedRuntime: "Unknown",
				Created:          "Wed, 22 Oct 2025 14:27:01 +0000",
				Name:             "wipe-disks",
				Parameters:       []byte(`{"quick_erase":{"type":"boolean","value":true,"argument_format":"--quick-erase"}}`),
				Status:           ABORTED,
				ID:               871,
				ScriptID:         31,
				ScriptRevisionID: 0,
				ExitStatus:       0,
				Starttime:        []byte("\"\""),
				Endtime:          []byte("\"\""),
				Suppressed:       false,
			},
		},
		Status: ABORTED,
		Type:   RELEASE,
		ID:     196,
	},
}

func TestNodeResult(t *testing.T) {
	nrs := new([]NodeResult)

	// Unmarshal sample data into the types
	if err := helper.TestdataFromJSON("maas/node_results.json", nrs); err != nil {
		t.Fatal(err)
	}

	// Verify the values are correct
	if diff := cmp.Diff(&sampleNodeResults, nrs); diff != "" {
		t.Fatalf("json.Decode([]NodeResult) mismatch (-want +got):\n%s", diff)
	}
}
