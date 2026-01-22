package entity

import (
	"testing"

	"github.com/canonical/gomaasclient/test/helper"
)

func TestPackageRepository(t *testing.T) {
	packageRepository := new(PackageRepository)
	packageRepositories := new([]PackageRepository)

	// Unmarshal sample data into the types
	if err := helper.TestdataFromJSON("maas/package_repository.json", packageRepository); err != nil {
		t.Fatal(err)
	}

	if err := helper.TestdataFromJSON("maas/package_repositories.json", packageRepositories); err != nil {
		t.Fatal(err)
	}
}
