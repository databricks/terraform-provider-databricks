default: build

build: vendor
	@echo "✓ Building source code with go build ..."
	@go build -mod vendor -v

fmt:
	@echo "✓ Formatting source code with goimports ..."
	@go run golang.org/x/tools/cmd/goimports@v0.34.0 -w $(shell find . -type f -name '*.go' -not -path "./vendor/*")
	@echo "✓ Formatting source code with gofmt ..."
	@gofmt -w $(shell find . -type f -name '*.go' -not -path "./vendor/*")

lint: vendor
	@echo "✓ Linting source code with https://staticcheck.io/ ..."
	@go run honnef.co/go/tools/cmd/staticcheck@v0.6.1 ./...

test: vendor
	@echo "✓ Running tests ..."
	@go run gotest.tools/gotestsum@v1.13.0 --format pkgname-and-test-fails \
		--no-summary=skipped --raw-command go test -v \
		-json -short -race -coverprofile=coverage.txt ./...

coverage: test
	@echo "✓ Opening coverage for unit tests ..."
	@go tool cover -html=coverage.txt

vendor:
	@echo "✓ Filling vendor folder with library code ..."
	@go mod vendor

download:
	@echo "✓ Downloading dependencies ..."
	@go mod download

doc:
	@echo "Open http://localhost:6060"
	@go run golang.org/x/tools/cmd/godoc@v0.34.0 -http=localhost:6060

.PHONY: fmt vendor fmt coverage test lint doc download
