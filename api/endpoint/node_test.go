package endpoint_test

import (
	"testing"

	. "github.com/ionutbalutoiu/gomaasclient/api/endpoint"
	"github.com/ionutbalutoiu/gomaasclient/test/helper"
)

func TestNodet(t *testing.T) {
	node := new(Node)
	nodes := new([]Node)

	// Unmarshal sample data into the types
	if err := helper.TestdataFromJSON("maas/node.json", node); err != nil {
		t.Fatal(err)
	}
	if err := helper.TestdataFromJSON("maas/nodes.json", nodes); err != nil {
		t.Fatal(err)
	}
}
