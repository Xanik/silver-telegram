lint:
	@hash golangci-lint > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		export BINARY="golangci-lint"; \
		curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(GOPATH)/bin v1.31.0; \
	fi
	golangci-lint run --timeout 5m

.PHONY: test
test:
	go test -v ./... -cover

gen-mocks:
	go generate ./...