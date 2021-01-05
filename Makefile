default: build

fmt:
	@echo "✓ Formatting source code with goimports ..."
	@goimports -w $(shell find . -type f -name '*.go' -not -path "./vendor/*")
	@echo "✓ Formatting source code with gofmt ..."
	@gofmt -w $(shell find . -type f -name '*.go' -not -path "./vendor/*")
	
lint: vendor
	@echo "✓ Linting source code with golangci-lint make sure you run make fmt ..."
	@golangci-lint run --skip-dirs-use-default --timeout 5m

test: lint
	@echo "✓ Running tests ..."
	@gotestsum --format pkgname-and-test-fails --no-summary=skipped --raw-command go test -v -json -short -coverprofile=coverage.txt ./...

coverage: test
	@echo "✓ Opening coverage for unit tests ..."
	@go tool cover -html=coverage.txt

VERSION = 0.3.0

build:
	@echo "✓ Building source code with go build ..."
	@go build -mod vendor -v -o terraform-provider-databricks

install: build
	@echo "✓ Installing provider into ~/.terraform.d/plugins ..."
	@test -d $(HOME)/.terraform.d/plugins && rm $(HOME)/.terraform.d/plugins/terraform-provider-databricks* || mkdir -p $(HOME)/.terraform.d/plugins
	@cp terraform-provider-databricks $(HOME)/.terraform.d/plugins
	@mkdir -p '$(HOME)/.terraform.d/plugins/registry.terraform.io/databrickslabs/databricks/$(shell ./terraform-provider-databricks version)/$(shell go version | awk '{print $$4}' | sed 's#/#_#')'
	@cp terraform-provider-databricks '$(HOME)/.terraform.d/plugins/registry.terraform.io/databrickslabs/databricks/$(shell ./terraform-provider-databricks version)/$(shell go version | awk '{print $$4}' | sed 's#/#_#')/terraform-provider-databricks-v$(shell ./terraform-provider-databricks version)'
	@echo "✓ Use the following configuration to enable the version you've built"
	@echo 
	@echo "terraform {"
	@echo "  required_providers {"
	@echo "    databricks = {"
	@echo '      source = "databrickslabs/databricks"'
	@echo '      version = "$(shell ./terraform-provider-databricks version)"'
	@echo "    }"
	@echo "  }"
	@echo "}"
	
vendor:
	@echo "✓ Filling vendor folder with library code ..."
	@go mod vendor

test-azcli: install
	@echo "✓ Running Terraform Acceptance Tests for Azure..."
	@/bin/bash scripts/run.sh azcli '^(TestAcc|TestAzureAcc)' --debug --tee

test-azsp: install
	@echo "✓ Running Terraform Acceptance Tests for Azure..."
	@/bin/bash scripts/run.sh azsp '^(TestAcc|TestAzureAcc)' --debug --tee

test-mws: install
	@echo "✓ Running acceptance Tests for Multiple Workspace APIs on AWS..."
	@/bin/bash scripts/run.sh mws '^TestMwsAcc' --debug --tee

test-awsst: install
	@echo "✓ Running Terraform Acceptance Tests for AWS ST..."
	@/bin/bash scripts/run.sh awsst '^(TestAcc|TestAwsAcc)' --debug --tee

test-awsmt: install
	@echo "✓ Running Terraform Acceptance Tests for AWS MT..."
	@/bin/bash scripts/run.sh awsmt '^(TestAcc|TestAwsAcc)' --debug --tee

snapshot:
	@echo "✓ Making Snapshot ..."
	@goreleaser release --rm-dist --snapshot

.PHONY: build fmt python-setup docs vendor build fmt coverage test lint