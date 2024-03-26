package entity

import (
	"testing"

	"github.com/canonical/gomaasclient/test/helper"
)

func TestBlockDevicet(t *testing.T) {
	device := new(BlockDevice)
	devices := new([]BlockDevice)

	// Unmarshal sample data into the types
	if err := helper.TestdataFromJSON("maas/block_device.json", device); err != nil {
		t.Fatal(err)
	}

	if err := helper.TestdataFromJSON("maas/block_devices.json", devices); err != nil {
		t.Fatal(err)
	}
}
