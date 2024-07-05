# Golang MAAS Client

## :warning: Repository ownership and module name change

The GoMAASClient repository now lives under the [Canonical GitHub organisation](https://github.com/canonical) with a new module name `github.com/canonical/gomaasclient`.

If you are not using `GOPROXY=https://proxy.golang.org/cached-only` or not caching
modules, your development environment or CI might encounter failures when trying
to retrieve the module by its old name.


If you encounter an error to the likes of:
```go
module declares its path as: github.com/canonical/gomaasclient
        but was required as: github.com/maas/gomaasclient
```

Ensure you are pointing at the new module URL, which is `github.com/canonical/gomaasclient`.
You can use the [replace](https://go.dev/ref/mod#go-mod-file-replace) directive, so
you don't have to change old imports everywhere.

---

This repository contains the following  packages:

* `api` - defines an interface to each MAAS API endpoint.
* `entity` - defines types for the MAAS API endpoints' return types.
* `client` - contains the MAAS client source code.

## Usage

```Go
import (
    gomaasclient "github.com/canonical/gomaasclient/client"
    "github.com/canonical/gomaasclient/entity"
)

c, _ := gomaasclient.GetClient("<MAAS_URL>", "<API_KEY>", "2.0")

// List MAAS machines
machines, _ := c.Machines.Get(&entity.MachinesParams{})

// Get MAAS machine details
machine, _ := c.Machine.Get(machines[0].SystemID)

// List MAAS VM hosts
vmHosts, _ := c.VMHosts.Get()
```

## Credit

This work was initially started by [Brian Hazeltine (@onwsk8r)](https://github.com/onwsk8r) as part of his [Terraform MAAS provider](https://github.com/Roblox/terraform-provider-maas) implementation.
