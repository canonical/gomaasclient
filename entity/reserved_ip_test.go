package entity

import (
	"testing"

	"github.com/canonical/gomaasclient/test/helper"
)

func TestReservedIP(t *testing.T) {
	reservedIP := new(ReservedIP)
	reservedIPs := new([]ReservedIP)

	// Unmarshal sample data into the types
	if err := helper.TestdataFromJSON("maas/reserved_ip.json", reservedIP); err != nil {
		t.Fatal(err)
	}

	if err := helper.TestdataFromJSON("maas/reserved_ips.json", reservedIPs); err != nil {
		t.Fatal(err)
	}
}
