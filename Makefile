TEST?=$$(go list ./... | grep -v 'vendor')

.PHONY: test
test:
	go test $(TEST) -v $(TESTARGS) -timeout=5m -parallel=4

.PHONY: lint
lint: lint-go

.PHONY: lint-go
lint-go:
	golangci-lint run $(if $(LINT_AUTOFIX),--fix,) ./...

.PHONY: lint-go-fix
lint-go-fix: LINT_AUTOFIX=true
lint-go-fix: lint-go
