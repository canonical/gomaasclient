# Golang MAAS Client

This repository contains the following  packages:

* `api` - defines an interface to each MAAS API endpoint.
* `entity` - defines types for the MAAS API endpoints' return types.
* `client` - contains the MAAS client source code.

## Usage

```Go
import (
    gomaasclient "github.com/ionutbalutoiu/gomaasclient/client"
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
