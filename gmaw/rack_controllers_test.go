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

func TestNewRackControllers(t *testing.T) {
	NewRackControllers(client)
}

func TestRackControllers(t *testing.T) {
	// Ensure the type implements the interface
	var _ api.RackControllers = (*RackControllers)(nil)

	// Create a new rackController client to be used in the tests
	rackControllerClient := NewRackControllers(client)

	t.Run("Get", func(t *testing.T) {
		t.Parallel()
		var want []endpoint.RackController
		if err := helper.TestdataFromJSON("maas/rack_controllers.json", &want); err != nil {
			t.Fatal(err)
		}
		httpmock.RegisterResponder("GET", "/MAAS/api/2.0/rackcontrollers/",
			httpmock.NewJsonResponderOrPanic(http.StatusOK, want))
		res, err := rackControllerClient.Get(&endpoint.RackControllerSearch{})
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(want, res, cmpopts.EquateEmpty()); diff != "" {
			t.Fatalf("json.Decode() mismatch (-want +got):\n%s", diff)
		}
	})
}
