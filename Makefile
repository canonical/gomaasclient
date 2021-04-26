TEST?=$$(go list ./... | grep -v 'vendor')

.PHONY: test
test:
	go test $(TEST) -v $(TESTARGS) -timeout=5m -parallel=4
