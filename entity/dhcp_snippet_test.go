package entity

import (
	"testing"

	"github.com/canonical/gomaasclient/test/helper"
)

func TestDHCPSnippett(t *testing.T) {
	snippets := new([]DHCPSnippet)
	if err := helper.TestdataFromJSON("maas/dhcp_snippets.json", snippets); err != nil {
		t.Fatal(err)
	}
}
