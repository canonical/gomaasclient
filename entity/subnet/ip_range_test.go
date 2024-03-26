package subnet

import (
	"net"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/canonical/gomaasclient/test/helper"
)

var sampleIPRanges []IPRange = []IPRange{
	{
		Start:        net.ParseIP("172.16.2.2"),
		End:          net.ParseIP("172.16.2.2"),
		NumAddresses: 1,
	},
	{
		Start:        net.ParseIP("172.16.2.5"),
		End:          net.ParseIP("172.16.2.10"),
		NumAddresses: 6,
	},
	{
		Start:        net.ParseIP("172.16.2.12"),
		End:          net.ParseIP("172.16.2.25"),
		NumAddresses: 14,
	},
	{
		Start:        net.ParseIP("172.16.2.27"),
		End:          net.ParseIP("172.16.2.61"),
		NumAddresses: 35,
	},
	{
		Start:        net.ParseIP("172.16.2.64"),
		End:          net.ParseIP("172.16.2.100"),
		NumAddresses: 37,
	},
	{
		Start:        net.ParseIP("172.16.2.102"),
		End:          net.ParseIP("172.16.2.108"),
		NumAddresses: 7,
	},
	{
		Start:        net.ParseIP("172.16.2.110"),
		End:          net.ParseIP("172.16.2.110"),
		NumAddresses: 1,
	},
	{
		Start:        net.ParseIP("172.16.2.112"),
		End:          net.ParseIP("172.16.2.115"),
		NumAddresses: 4,
	},
	{
		Start:        net.ParseIP("172.16.2.117"),
		End:          net.ParseIP("172.16.2.133"),
		NumAddresses: 17,
	},
	{
		Start:        net.ParseIP("172.16.2.135"),
		End:          net.ParseIP("172.16.2.173"),
		NumAddresses: 39,
	},
	{
		Start:        net.ParseIP("172.16.2.175"),
		End:          net.ParseIP("172.16.2.205"),
		NumAddresses: 31,
	},
	{
		Start:        net.ParseIP("172.16.2.207"),
		End:          net.ParseIP("172.16.2.234"),
		NumAddresses: 28,
	},
	{
		Start:        net.ParseIP("172.16.2.236"),
		End:          net.ParseIP("172.16.2.236"),
		NumAddresses: 1,
	},
	{
		Start:        net.ParseIP("172.16.2.238"),
		End:          net.ParseIP("172.16.2.251"),
		NumAddresses: 14,
	},
	{
		Start:        net.ParseIP("172.16.2.253"),
		End:          net.ParseIP("172.16.2.254"),
		NumAddresses: 2,
	},
}

func TestIPRange(t *testing.T) {
	ranges := new([]IPRange)

	// Unmarshal sample data into the types
	if err := helper.TestdataFromJSON("maas/subnets/ipranges.json", ranges); err != nil {
		t.Fatal(err)
	}

	// Verify the values are correct
	if diff := cmp.Diff(&sampleIPRanges, ranges); diff != "" {
		t.Fatalf("json.Decode([]IPRange) mismatch (-want +got):\n%s", diff)
	}
}
