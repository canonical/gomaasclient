package entity

import (
	"encoding/json"
	"testing"

	"github.com/canonical/gomaasclient/test/helper"
	"github.com/stretchr/testify/assert"
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

func TestMAASUnMarshal(t *testing.T) {
	machine := new(Machine)

	// Unmarshal sample data into the types
	if err := helper.TestdataFromJSON("maas/machine.json", machine); err != nil {
		t.Fatal(err)
	}

	t.Run("CacheSets", func(t *testing.T) {
		assert.Equal(t, 1, len(machine.CacheSets))
		assert.Equal(t, 5, machine.CacheSets[0].ID)
		assert.Equal(t, true, machine.CacheSets[0].Incomplete)
		assert.Equal(t, "dqatxp", machine.CacheSets[0].SystemID)
	})
	t.Run("VolumeGroups", func(t *testing.T) {
		assert.Equal(t, 1, len(machine.VolumeGroups))
		assert.Equal(t, 4, machine.VolumeGroups[0].ID)
		assert.Equal(t, true, machine.VolumeGroups[0].Incomplete)
		assert.Equal(t, "dqatxp", machine.VolumeGroups[0].SystemID)
	})
	t.Run("BCaches", func(t *testing.T) {
		assert.Equal(t, 1, len(machine.BCaches))
		assert.Equal(t, 3, machine.BCaches[0].ID)
		assert.Equal(t, true, machine.BCaches[0].Incomplete)
		assert.Equal(t, "3gf1k4", machine.BCaches[0].SystemID)
	})
	t.Run("RAIDs", func(t *testing.T) {
		assert.Equal(t, 1, len(machine.RAIDs))
		assert.Equal(t, 2, machine.RAIDs[0].ID)
		assert.Equal(t, true, machine.RAIDs[0].Incomplete)
		assert.Equal(t, "rdt4fa", machine.RAIDs[0].SystemID)
	})
	t.Run("SpecialFS", func(t *testing.T) {
		assert.Equal(t, 1, len(machine.SpecialFilesystems))
		assert.Equal(t, "ramfs", machine.SpecialFilesystems[0].FSType)
		assert.Equal(t, "special", machine.SpecialFilesystems[0].Label)
		assert.Equal(t, "0e8b2c73-c31f-4fd0-b3b9-8dde74406afe", machine.SpecialFilesystems[0].UUID)
		assert.Equal(t, "/MqbWGqzmWV/KrM409CMFO/CsDtuVxjpQ", machine.SpecialFilesystems[0].MountPoint)
		assert.Equal(t, "mount-options-552Rua", machine.SpecialFilesystems[0].MountOptions)
	})
}
