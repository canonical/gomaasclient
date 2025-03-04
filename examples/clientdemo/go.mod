module example.com/clientdemo

go 1.21.1

require github.com/canonical/gomaasclient v0.0.0-20230912053103-eb0ecc7ab134

require (
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/juju/collections v1.0.4 // indirect
	github.com/juju/errors v1.0.0 // indirect
	github.com/juju/gomaasapi/v2 v2.3.0 // indirect
	github.com/juju/loggo v1.0.0 // indirect
	github.com/juju/mgo/v2 v2.0.2 // indirect
	github.com/juju/schema v1.0.1 // indirect
	github.com/juju/version v0.0.0-20210303051006-2015802527a8 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Replace the path after => with the path to your local copy of the repository, if required.
// e.g. /home/username/repos/gomaasclient
replace github.com/canonical/gomaasclient v0.0.0-20230912053103-eb0ecc7ab134 => ../../../gomaasclient
