package entity

import (
	"testing"

	"github.com/canonical/gomaasclient/test/helper"
)

func TestVolumeGroup(t *testing.T) {
	volumeGroup := new(VolumeGroup)
	volumeGroups := new([]VolumeGroup)

	// Unmarshal sample data into the types
	if err := helper.TestdataFromJSON("maas/volume_group.json", volumeGroup); err != nil {
		t.Fatal(err)
	}

	if err := helper.TestdataFromJSON("maas/volume_groups.json", volumeGroups); err != nil {
		t.Fatal(err)
	}
}
