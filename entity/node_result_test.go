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
				Status:           2,
				ID:               56,
				ScriptID:         32,
				ScriptRevisionID: 33,
				ExitStatus:       0,
				Starttime:        1.761059832783871e+09,
				Endtime:          1.761059832844518e+09,
				Suppressed:       false,
			},
		},
		Status: 2,
		Type:   3,
		ID:     11,
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
