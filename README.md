# Golang MAAS Client

This repository contains the following  packages:

* `api` - defines an interface to each MAAS API endpoint.

* `gmaw` (**G**o**M**aas**A**pi**W**rapper) - is a wrapper (read: adapter) for [gomaasapi](github.com/juju/gomaasapi) to make it compatible with the client interfaces defined in and expected by the adjacent
`maas` package.

* `maas` - encapsulates a MAAS API client with thread safety, state management, Go types for resource data, and enhanced functionality for managing API calls.

## Credit

This work was initially started by [Brian Hazeltine (@onwsk8r)](https://github.com/onwsk8r) as part of his [Terraform MAAS provider](https://github.com/Roblox/terraform-provider-maas) implementation.
