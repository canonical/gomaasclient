package entity

import (
	"testing"

	"github.com/canonical/gomaasclient/test/helper"
	"github.com/google/go-cmp/cmp"
)

var sampleBootResource BootResource = BootResource{
	Sets: map[string]BootResourceSet{
		"20240828": {
			Files: map[string]BootResourceSetFile{
				"root.tgz": {
					Filename: "root.tgz",
					Filetype: "root-tgz",
					SHA256:   "5b341be98b05d409a41ce8c2d881dbd80df9de07da7613d6722d4f0b5e99e0f1",
					Size:     935047294,
					Complete: true,
				},
			},
			Version:  "20240828",
			Label:    "uploaded",
			Size:     935047294,
			Complete: true,
		},
	},
	Type:         "Uploaded",
	Name:         "alma9",
	Architecture: "amd64/generic",
	BaseImage:    "rhel/9",
	ResourceURI:  "/MAAS/api/2.0/boot-resources/13/",
	Subarches:    "generic",
	Title:        "Alma 9 Custom",
	ID:           13,
}

var sampleBootResources []BootResource = []BootResource{
	{
		Type:         "Synced",
		Name:         "ubuntu/bionic",
		Architecture: "amd64/ga-18.04",
		ResourceURI:  "/MAAS/api/2.0/boot-resources/95/",
		Subarches:    "generic,hwe-p,hwe-q,hwe-r,hwe-s,hwe-t,hwe-u,hwe-v,hwe-w,ga-16.04,ga-16.10,ga-17.04,ga-17.10,ga-18.04",
		ID:           95,
	},
	{
		Type:         "Uploaded",
		Name:         "alma9",
		Architecture: "amd64/generic",
		BaseImage:    "rhel/9",
		ResourceURI:  "/MAAS/api/2.0/boot-resources/13/",
		Subarches:    "generic",
		Title:        "Alma 9 Custom",
		ID:           13,
	},
}

func TestBootResource(t *testing.T) {
	br := new(BootResource)
	brs := new([]BootResource)

	// Unmarshal sample data into the types
	if err := helper.TestdataFromJSON("maas/boot_resource.json", br); err != nil {
		t.Fatal(err)
	}

	if err := helper.TestdataFromJSON("maas/boot_resources.json", brs); err != nil {
		t.Fatal(err)
	}

	// Verify the values are correct
	if diff := cmp.Diff(&sampleBootResource, br); diff != "" {
		t.Fatalf("json.Decode(NetworkInterface) mismatch (-want +got):\n%s", diff)
	}

	if diff := cmp.Diff(&sampleBootResources, brs); diff != "" {
		t.Fatalf("json.Decode([]NetworkInterface) mismatch (-want +got):\n%s", diff)
	}
}
