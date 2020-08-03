default: build

int:
	@echo "✓ Running tests..."
	@gotestsum --raw-command go test -v -json -coverprofile=coverage.txt ./...

coverage-int: int
	@echo "✓ Opening coverage for unit tests..."
	@go tool cover -html=coverage.txt

int-build: int build

fmt:
	@echo "✓ Formatting source code with gofmt..."
	@goimports -w client
	@goimports -w databricks
	@goimports -w main.go
	@gofmt -s -w client
	@gofmt -s -w databricks
	@gofmt -s -w main.go
	@go fmt ./...

lint: 
	@echo "✓ Linting source code with golangci-lint make sure you run make fmt ..."
	@golangci-lint run --skip-dirs-use-default --timeout 5m

test: lint
	@echo "✓ Running tests..."
	@gotestsum --format pkgname-and-test-fails --no-summary=skipped --raw-command go test -v -json -short -coverprofile=coverage.txt ./...

coverage: test
	@echo "✓ Opening coverage for unit tests..."
	@go tool cover -html=coverage.txt

build:
	@echo "✓ Building source code with go build..."
	@go build -mod vendor -v -o terraform-provider-databricks

install: build
	@echo "✓ Installing provider into ~/.terraform.d/plugins ..."
	@test -d $(HOME)/.terraform.d/plugins && rm $(HOME)/.terraform.d/plugins/terraform-provider-databricks* || mkdir -p $(HOME)/.terraform.d/plugins
	@mv terraform-provider-databricks $(HOME)/.terraform.d/plugins

vendor:
	@echo "✓ Filling vendor folder with library code..."
	@go mod vendor

test-azure:
	@echo "✓ Running Terraform Acceptance Tests for Azure..."
	@/bin/bash scripts/run.sh mws '^(TestAcc|TestAzureAcc)'

test-mws:
	@echo "✓ Running acceptance Tests for Multiple Workspace APIs on AWS..."
	@/bin/bash scripts/run.sh mws TestAccMWSWorkspaces

# INTEGRATION TESTING WITH AZURE
terraform-acc-azure: lint
	@echo "✓ Running Terraform Acceptance Tests for Azure..."
	@/bin/bash integration-environment-azure/run.sh

# INTEGRATION TESTING WITH AWS
terraform-acc-aws: lint
	@echo "✓ Running Terraform Acceptance Tests for AWS..."
	@CLOUD_ENV="aws" TF_ACC=1 gotestsum --format short-verbose --raw-command go test -v -json -short -coverprofile=coverage.out -run 'TestAccAws' ./...

# INTEGRATION TESTING WITH AWS
terraform-acc-mws: lint
	@echo "✓ Running Terraform Acceptance Tests for Multiple Workspace APIs on AWS..."
	@/bin/bash integration-environment-mws/run.sh

# Launch VSCode with Azure integration test ENV variables
code-azure:
	export $(scripts/run.sh azure --export)
	code .

code-aws:
	export $(scripts/run.sh aws --export)
	code .

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