default: build

fmt:
	@echo "✓ Formatting source code with gofmt..."
	@goimports -w $(shell find . -type f -name '*.go' -not -path "./vendor/*")
	@gofmt -w $(shell find . -type f -name '*.go' -not -path "./vendor/*")
	@go fmt ./...

lint: 
	@echo "✓ Linting source code with golangci-lint make sure you run make fmt ..."
	@golangci-lint run --skip-dirs-use-default --timeout 5m

test:
	@echo "✓ Running tests..."
	@gotestsum --format pkgname-and-test-fails --no-summary=skipped --raw-command go test -v -json -short -coverprofile=coverage.txt ./...

coverage: test
	@echo "✓ Opening coverage for unit tests..."
	@go tool cover -html=coverage.txt

build:
	@echo "✓ Building source code with go build..."
	@go build -mod vendor -v -ldflags="-X 'common.version=$(git describe --long --always | sed 's/v//')'" -o terraform-provider-databricks

install: build
	@echo "✓ Installing provider into ~/.terraform.d/plugins ..."
	@test -d $(HOME)/.terraform.d/plugins && rm $(HOME)/.terraform.d/plugins/terraform-provider-databricks* || mkdir -p $(HOME)/.terraform.d/plugins
	@cp terraform-provider-databricks $(HOME)/.terraform.d/plugins
	@mkdir -p '$(HOME)/.terraform.d/plugins/registry.terraform.io/databrickslabs/databricks/$(shell git describe --long --always | sed 's/v//')/$(shell go version | awk '{print $$4}' | sed 's#/#_#')'
	@cp terraform-provider-databricks '$(HOME)/.terraform.d/plugins/registry.terraform.io/databrickslabs/databricks/$(shell git describe --long --always | sed 's/v//')/$(shell go version | awk '{print $$4}' | sed 's#/#_#')'
	@echo "✓ Use the following configuration to enable the version you've built"
	@echo 
	@echo "terraform {"
	@echo "  required_providers {"
	@echo "    databricks = {"
	@echo '      source = "databrickslabs/databricks"'
	@echo '      version = "$(shell git describe --long --always | sed 's/v//')"'
	@echo "    }"
	@echo "  }"
	@echo "}"
	
vendor:
	@echo "✓ Filling vendor folder with library code..."
	@go mod vendor

test-azcli:
	@echo "✓ Running Terraform Acceptance Tests for Azure..."
	@/bin/bash scripts/run.sh azcli '^(TestAcc|TestAzureAcc)' --debug --tee

test-azsp:
	@echo "✓ Running Terraform Acceptance Tests for Azure..."
	@/bin/bash scripts/run.sh azsp '^(TestAcc|TestAzureAcc)' --debug --tee

test-mws:
	@echo "✓ Running acceptance Tests for Multiple Workspace APIs on AWS..."
	@/bin/bash scripts/run.sh mws '^TestMwsAcc' --debug --tee

test-awsst:
	@echo "✓ Running Terraform Acceptance Tests for AWS ST..."
	@/bin/bash scripts/run.sh awsst '^(TestAcc|TestAwsAcc)' --debug --tee

test-awsmt:
	@echo "✓ Running Terraform Acceptance Tests for AWS MT..."
	@/bin/bash scripts/run.sh awsmt '^(TestAcc|TestAwsAcc)' --debug --tee

terraform-setup: build
	@echo "✓ Initializing Terraform..."
	@terraform init

terraform-apply: terraform-setup
	@echo "✓ Initializing Terraform plan..."
	@TF_LOG_PATH=log.out TF_LOG=debug terraform apply

snapshot:
	@echo "✓ Making Snapshot..."
	@goreleaser release --rm-dist --snapshot

hugo:
	@echo "✓ Making Docs..."
	@cd website && hugo -d ../docs/

.PHONY: build fmt python-setup docs vendor terraform-local build fmt coverage test lint