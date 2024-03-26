package entity

import (
	"testing"

	"github.com/canonical/gomaasclient/test/helper"
)

func TestVMHostt(t *testing.T) {
	vmHost := new(VMHost)
	vmHosts := new([]VMHost)

	// Unmarshal sample data into the types
	if err := helper.TestdataFromJSON("maas/vm_host.json", vmHost); err != nil {
		t.Fatal(err)
	}

	if err := helper.TestdataFromJSON("maas/vm_hosts.json", vmHosts); err != nil {
		t.Fatal(err)
	}
}
