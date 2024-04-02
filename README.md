# Golang MAAS Client

**Notice**

* **GoMAASClient is moving!**

The GoMAASClient codebase will be accessed through the [Canonical GitHub organisation](https://github.com/canonical/gomaasclient), but go imports will still use the original link until further notice.

ie: `go get github.com/maas/gomaasclient`.

if you encounter as error along the likes of:
```go
module declares its path as: github.com/canonical/maas
        but was required as: github.com/maas/maas
```

Ensure you are pointing at the original library URL.

**/Notice**

This repository contains the following  packages:

* `api` - defines an interface to each MAAS API endpoint.
* `entity` - defines types for the MAAS API endpoints' return types.
* `client` - contains the MAAS client source code.

## Usage

```Go
import (
    gomaasclient "github.com/maas/gomaasclient/client"
)

c, _ := gomaasclient.GetClient("<MAAS_URL>", "<API_KEY>", "2.0")

// List MAAS machines
machines, _ := c.Machines.Get()

// Get MAAS machine details
machine, _ := c.Machine.Get(machines[0].SystemID)

// List MAAS VM hosts
vmHosts, _ := c.VMHosts.Get()
```

## Credit

This work was initially started by [Brian Hazeltine (@onwsk8r)](https://github.com/onwsk8r) as part of his [Terraform MAAS provider](https://github.com/Roblox/terraform-provider-maas) implementation.
