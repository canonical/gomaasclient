package entity

import (
	"encoding/json"
	"testing"

	"github.com/maas/gomaasclient/test/helper"
)

func TestMachinet(t *testing.T) {
	machine := new(Machine)
	machines := new([]Machine)

	// Unmarshal sample data into the types
	if err := helper.TestdataFromJSON("maas/machine.json", machine); err != nil {
		t.Fatal(err)
	}
	if err := helper.TestdataFromJSON("maas/machines.json", machines); err != nil {
		t.Fatal(err)
	}
}

func TestMachineWithDeducedHwSync(t *testing.T) {
	testcases := map[string]struct {
		in  []byte
		out bool
	}{
		"api returns no hw_sync data": {
			in:  []byte(`{}`),
			out: false,
		},
		"api returns enable_hw_sync true": {
			in: []byte(`
        {
          "enable_hw_sync": true
        }`,
			),
			out: true,
		},
		"api returns enable_hw_sync false": {
			in: []byte(`
        {
          "enable_hw_sync": false
        }`,
			),
			out: false,
		},
		"deduced enable_hw_sync true": {
			in: []byte(`
        {
          "sync_interval": 900,
          "next_sync": "2023-05-11T21:15:04.208",
          "last_sync": "2023-05-11T21:00:04.208"
        }`,
			),
			out: true,
		},
		"deduced enable_hw_sync false": {
			in: []byte(`
        {
          "sync_interval": 0,
          "next_sync": null,
          "last_sync": null
        }`,
			),
			out: false,
		},
	}

	for name, tc := range testcases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			machine := Machine{}

			if err := json.Unmarshal(tc.in, &machine); err != nil {
				t.Fatal(err, string(tc.in))
			}
			if tc.out != machine.EnableHwSync {
				t.Fatalf("expected %v, got %v", tc.out, machine.EnableHwSync)
			}
		})
	}
}

func TestMAASTimeUnmarshalJSON(t *testing.T) {
	temp := struct {
		Time MAASTime
	}{}

	data := []byte(`{"Time": "2023-05-11T21:15:04.208"}`)
	if err := json.Unmarshal(data, &temp); err != nil {
		t.Fatal(err)
	}

	expected := "2023-05-11T21:15:04.208Z"
	if temp.Time.String() != expected {
		t.Fatalf("expected %v, got %v", expected, temp.Time.String())
	}
}
