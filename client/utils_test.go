package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPowerParamsToURLValues(t *testing.T) {
	testcases := map[string]struct {
		in  map[string]interface{}
		out string
	}{
		"strings": {
			in: map[string]interface{}{
				"a": "b",
				"c": "d",
			},
			out: "a=b&c=d",
		},
		"strings array": {
			in: map[string]interface{}{
				"a": []interface{}{"b", "c"},
			},
			out: "a=b&a=c",
		},
		"numbers": {
			in: map[string]interface{}{
				"a": 1,
				"b": 2,
			},
			out: "a=1&b=2",
		},
		"numbers array": {
			in: map[string]interface{}{
				"a": []interface{}{1, 2},
			},
			out: "a=1&a=2",
		},
		"boolean": {
			in: map[string]interface{}{
				"a": true,
				"b": false,
			},
			out: "a=true&b=false",
		},
		"boolean array": {
			in: map[string]interface{}{
				"a": []interface{}{true, false},
			},
			out: "a=true&a=false",
		},
	}

	for name, tc := range testcases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			res := powerParamsToURLValues(tc.in).Encode()
			assert.Equal(t, tc.out, res)
		})
	}
}
