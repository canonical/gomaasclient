package node

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStatus(t *testing.T) {
	tests := []struct {
		name string
		got  Status
		want Status
	}{
		{name: "new", got: StatusNew, want: 0},
		{name: "default", got: StatusDefault, want: 0},
		{name: "commissioning", got: StatusCommissioning, want: 1},
		{name: "ready", got: StatusReady, want: 4},
		{name: "deployed", got: StatusDeployed, want: 6},
		{name: "deploying", got: StatusDeploying, want: 9},
		{name: "allocated", got: StatusAllocated, want: 10},
	}

	for _, testCase := range tests {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			diff := cmp.Diff(tc.want, tc.got)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
