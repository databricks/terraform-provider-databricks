default: build

test:
	@echo "==> Running tests..."
	@gotestsum --format short-verbose --raw-command go test -v -json -short -coverprofile=coverage.txt ./...

client-test:
	@echo "==> Running tests..."
	@gotestsum --format short-verbose --raw-command go test -v -json -short -coverprofile=client-coverage.txt ./client/...

provider-test:
	@echo "==> Running tests..."
	@gotestsum --format short-verbose --raw-command go test -v -json -short -coverprofile=provider-coverage.txt ./databricks/...

int:
	@echo "==> Running tests..."
	@gotestsum --raw-command go test -v -json -coverprofile=coverage.txt ./...

coverage: test
	@echo "==> Opening coverage for unit tests..."
	@go tool cover -html=coverage.txt

coverage-int: int
	@echo "==> Opening coverage for unit tests..."
	@go tool cover -html=coverage.txt

int-build: int build

build: lint test fmt
	@echo "==> Building source code with go build..."
	@go build -mod vendor -v -o terraform-provider-databricks

lint:
	@echo "==> Linting source code with golangci-lint..."
	@golangci-lint run --skip-dirs-use-default --timeout 5m --build-tags=azure
	@golangci-lint run --skip-dirs-use-default --timeout 5m --build-tags=aws

fmt: lint
	@echo "==> Formatting source code with gofmt..."
	@go fmt ./...


python-setup:
	@echo "==> Setting up virtual env and installing python libraries..."
	@python -m pip install virtualenv
	@cd docs && python -m virtualenv venv
	@cd docs && source venv/bin/activate && python -m pip install -r requirements.txt

docs: python-setup
	@echo "==> Building Docs ..."
	@cd docs && source venv/bin/activate && make clean && make html

opendocs: python-setup docs
	@echo "==> Opening Docs ..."
	@cd docs && open build/html/index.html

singlehtmldocs: python-setup
	@echo "==> Building Docs ..."
	@cd docs && source venv/bin/activate && make clean && make singlehtml

vendor:
	@echo "==> Filling vendor folder with library code..."
	@go mod vendor

# INTEGRATION TESTING WITH AZURE
terraform-acc-azure: fmt
	@echo "==> Running Terraform Acceptance Tests for Azure..."
	@CLOUD_ENV="azure" TF_ACC=1 gotestsum --format short-verbose --raw-command go test -v -json -tags=azure  -short -coverprofile=coverage.out ./...

# INTEGRATION TESTING WITH AWS
terraform-acc-aws: fmt
	@echo "==> Running Terraform Acceptance Tests for AWS..."
	@CLOUD_ENV="aws" TF_ACC=1 gotestsum --format short-verbose --raw-command go test -v -json -tags=aws  -short -coverprofile=coverage.out ./...

terraform-setup: build
	@echo "==> Initializing Terraform..."
	@terraform init

terraform-apply: terraform-setup
	@echo "==> Initializing Terraform plan..."
	@TF_LOG_PATH=log.out TF_LOG=debug terraform apply

snapshot:
	@echo "==> Making Snapshot..."
	@goreleaser release --rm-dist --snapshot

hugo:
	@echo "==> Making Docs..."
	@cd website && hugo -d ../docs/

internal-docs-sync:
	@echo "==> Uploading Website..."
	@azcopy login --service-principal --application-id $(AZCOPY_SPA_CLIENT_ID) --tenant-id=$(AZCOPY_SPA_TENANT_ID) && azcopy sync './website/public' '$(AZCOPY_STORAGE_ACCT)' --recursive

.PHONY: build fmt python-setup docs vendor terraform-local build fmt coverage test lint