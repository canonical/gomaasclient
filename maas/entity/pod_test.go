package entity_test

import (
	"testing"

	. "github.com/ionutbalutoiu/gomaasclient/maas/entity"

	"github.com/ionutbalutoiu/gomaasclient/test/helper"
)

func TestPodt(t *testing.T) {
	pod := new(Pod)
	pods := new([]Pod)

	// Unmarshal sample data into the types
	if err := helper.TestdataFromJSON("maas/pod.json", pod); err != nil {
		t.Fatal(err)
	}
	if err := helper.TestdataFromJSON("maas/pods.json", pods); err != nil {
		t.Fatal(err)
	}
}
