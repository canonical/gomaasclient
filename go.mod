module github.com/maas/gomaasclient

go 1.17

require (
	github.com/google/go-cmp v0.5.6
	github.com/google/go-querystring v1.1.0
	github.com/juju/gomaasapi/v2 v2.1.0
	github.com/stretchr/testify v1.8.4
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/juju/collections v0.0.0-20220203020748-febd7cad8a7a // indirect
	github.com/juju/errors v0.0.0-20220203013757-bd733f3c86b9 // indirect
	github.com/juju/loggo v0.0.0-20210728185423-eebad3a902c4 // indirect
	github.com/juju/mgo/v2 v2.0.0-20220111072304-f200228f1090 // indirect
	github.com/juju/schema v1.0.1-0.20190814234152-1f8aaeef0989 // indirect
	github.com/juju/version v0.0.0-20191219164919-81c1be00b9a6 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/juju/gomaasapi/v2 v2.1.0 => github.com/skatsaounis/gomaasapi/v2 v2.2.0
