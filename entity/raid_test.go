package entity

import (
	"testing"

	"github.com/canonical/gomaasclient/test/helper"
	"github.com/stretchr/testify/assert"
)

func TestRaid(t *testing.T) {
	raid := new(RAID)

	// Unmarshal sample data into the types
	if err := helper.TestdataFromJSON("maas/raid.json", raid); err != nil {
		t.Fatal(err)
	}

	t.Run("Devices", func(t *testing.T) {
		assert.Equal(t, 0, len(raid.SpareDevices))
		assert.Equal(t, 2, len(raid.Devices))
		assert.Equal(t, 124, raid.Devices[0].ID)
		assert.Equal(t, "qkhm8f", raid.Devices[0].SystemID)
	})

	t.Run("VirtualDevice", func(t *testing.T) {
		assert.Equal(t, "md0", raid.VirtualDevice.Name)
		assert.Equal(t, 63, raid.VirtualDevice.ID)
		assert.Equal(t, "qkhm8f", raid.VirtualDevice.SystemID)
	})
}
