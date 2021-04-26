package gmaw_test

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jarcoal/httpmock"

	"github.com/ionutbalutoiu/gomaasclient/api"
	"github.com/ionutbalutoiu/gomaasclient/api/endpoint"
	. "github.com/ionutbalutoiu/gomaasclient/gmaw"
	"github.com/ionutbalutoiu/gomaasclient/test/helper"
)

func TestNewVLANs(t *testing.T) {
	NewVLANs(client)
}

func TestVLANs(t *testing.T) {
	// Ensure the type implements the interface
	var _ api.VLANs = (*VLANs)(nil)

	// Create a new vlans client to be used in the tests
	vlansClient := NewVLANs(client)

	t.Run("Get", func(t *testing.T) {
		var vlans []endpoint.VLAN
		if err := helper.TestdataFromJSON("maas/vlans.json", &vlans); err != nil {
			t.Fatal(err)
		}
		t.Run("200", func(t *testing.T) {
			t.Parallel()
			httpmock.RegisterResponder("GET", "/MAAS/api/2.0/fabrics/123/vlans/",
				httpmock.NewJsonResponderOrPanic(http.StatusOK, vlans))
			res, err := vlansClient.Get(123)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(vlans, res, cmpopts.EquateEmpty()); diff != "" {
				t.Fatalf("json.Decode(VLANs) mismatch (-want +got):\n%s", diff)
			}
		})
		t.Run("404", func(t *testing.T) {
			t.Parallel()
			httpmock.RegisterResponder("GET", "/MAAS/api/2.0/fabrics/234/vlans/",
				httpmock.NewStringResponder(http.StatusNotFound, "Not Found"))

			got, err := vlansClient.Get(234)
			if diff := cmp.Diff(([]endpoint.VLAN{}), got, cmpopts.EquateEmpty()); diff != "" {
				t.Fatalf("json.Decode() mismatch (-want +got):\n%s", diff)
			}
			if err.Error() != "ServerError: 404 (Not Found)" {
				t.Fatal(err)
			}
		})
	})
	t.Run("Post", func(t *testing.T) {
		vlan := new(endpoint.VLAN)
		if err := helper.TestdataFromJSON("maas/vlan.json", vlan); err != nil {
			t.Fatal(err)
		}
		t.Run("200", func(t *testing.T) {
			t.Parallel()
			httpmock.RegisterResponder("POST", "/MAAS/api/2.0/fabrics/123/vlans/",
				httpmock.NewJsonResponderOrPanic(http.StatusOK, vlan))

			p := new(endpoint.VLANParams)
			got, err := vlansClient.Post(123, p)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(vlan, got, cmpopts.EquateEmpty()); diff != "" {
				t.Fatalf("json.Decode(VLANs) mismatch (-want +got):\n%s", diff)
			}
		})
		t.Run("404", func(t *testing.T) {
			t.Parallel()
			httpmock.RegisterResponder("POST", "/MAAS/api/2.0/fabrics/234/vlans/",
				httpmock.NewStringResponder(http.StatusNotFound, "Not Found"))

			p := new(endpoint.VLANParams)
			got, err := vlansClient.Post(234, p)
			if diff := cmp.Diff((&endpoint.VLAN{}), got, cmpopts.EquateEmpty()); diff != "" {
				t.Fatalf("json.Decode() mismatch (-want +got):\n%s", diff)
			}
			if err.Error() != "ServerError: 404 (Not Found)" {
				t.Fatal(err)
			}
		})
	})
}
