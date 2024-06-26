package entity

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/canonical/gomaasclient/test/helper"
)

var sampleVLAN VLAN = VLAN{
	VID:           10,
	MTU:           1500,
	DHCPOn:        false,
	FabricID:      0,
	Name:          "10",
	PrimaryRack:   "7xtf67",
	Space:         "internal",
	SecondaryRack: "76y7pg",
	ID:            5002,
	Fabric:        "fabric-0",
	ResourceURI:   "/MAAS/api/2.0/vlans/5002/",
}

var sampleVLANs []VLAN = []VLAN{
	{
		VID:           0,
		MTU:           1500,
		DHCPOn:        false,
		ExternalDHCP:  "",
		Space:         "undefined",
		FabricID:      10,
		SecondaryRack: "",
		Fabric:        "fabric-10",
		ID:            5014,
		PrimaryRack:   "",
		Name:          "untagged",
		ResourceURI:   "/MAAS/api/2.0/vlans/5014/",
	},
	{
		VID:          1,
		MTU:          1500,
		DHCPOn:       false,
		ExternalDHCP: "",
		RelayVLAN: &VLAN{
			VID:           0,
			MTU:           1500,
			DHCPOn:        false,
			ExternalDHCP:  "",
			Space:         "undefined",
			FabricID:      10,
			SecondaryRack: "",
			Fabric:        "fabric-10",
			ID:            5014,
			PrimaryRack:   "",
			Name:          "untagged",
			ResourceURI:   "/MAAS/api/2.0/vlans/5014/",
		},
		Space:         "undefined",
		FabricID:      10,
		SecondaryRack: "",
		Fabric:        "fabric-10",
		ID:            5015,
		PrimaryRack:   "",
		Name:          "untagged",
		ResourceURI:   "/MAAS/api/2.0/vlans/5015/",
	},
}

func TestVLAN(t *testing.T) {
	vlan := new(VLAN)
	vlans := new([]VLAN)

	// Unmarshal sample data into the types
	if err := helper.TestdataFromJSON("maas/vlan.json", vlan); err != nil {
		t.Fatal(err)
	}

	if err := helper.TestdataFromJSON("maas/vlans.json", vlans); err != nil {
		t.Fatal(err)
	}

	// Verify the values are correct
	if diff := cmp.Diff(&sampleVLAN, vlan); diff != "" {
		t.Fatalf("json.Decode(VLAN) mismatch (-want +got):\n%s", diff)
	}

	if diff := cmp.Diff(&sampleVLANs, vlans); diff != "" {
		t.Fatalf("json.Decode([]VLAN) mismatch (-want +got):\n%s", diff)
	}
}
