package gmaw_test

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jarcoal/httpmock"

	"github.com/ionutbalutoiu/gomaasclient/api"
	"github.com/ionutbalutoiu/gomaasclient/api/params"
	. "github.com/ionutbalutoiu/gomaasclient/gmaw"
	"github.com/ionutbalutoiu/gomaasclient/maas/entity"

	"github.com/ionutbalutoiu/gomaasclient/test/helper"
)

func TestNewSubnets(t *testing.T) {
	NewSubnets(client)
}

func TestSubnets(t *testing.T) {
	// Ensure the type implements the interface
	var _ api.Subnets = (*Subnets)(nil)

	// Create a new subnets client to be used in the tests
	subnetsClient := NewSubnets(client)

	t.Run("Get", func(t *testing.T) {
		t.Parallel()
		var subnets []entity.Subnet
		if err := helper.TestdataFromJSON("maas/subnets.json", &subnets); err != nil {
			t.Fatal(err)
		}
		httpmock.RegisterResponder("GET", "/MAAS/api/2.0/subnets/",
			httpmock.NewJsonResponderOrPanic(http.StatusOK, subnets))
		res, err := subnetsClient.Get()
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(subnets, res, cmpopts.EquateEmpty()); diff != "" {
			t.Fatalf("json.Decode(Subnets) mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("Post", func(t *testing.T) {
		t.Parallel()
		subnet := new(entity.Subnet)
		if err := helper.TestdataFromJSON("maas/subnet.json", subnet); err != nil {
			t.Fatal(err)
		}
		httpmock.RegisterResponder("POST", "/MAAS/api/2.0/subnets/",
			httpmock.NewJsonResponderOrPanic(http.StatusOK, subnet))

		p := new(params.Subnet)
		res, err := subnetsClient.Post(p)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(subnet, res, cmpopts.EquateEmpty()); diff != "" {
			t.Fatalf("json.Decode(Subnets) mismatch (-want +got):\n%s", diff)
		}
	})
}
