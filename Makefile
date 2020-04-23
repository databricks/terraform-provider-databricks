default: build

test:
	@echo "==> Running tests..."
	@gotestsum --format short-verbose --raw-command go test -v -json -short -coverprofile=coverage.out ./...

int:
	@echo "==> Running tests..."
	@gotestsum --raw-command go test -v -json -coverprofile=coverage.out ./...

coverage: test
	@echo "==> Opening coverage for unit tests..."
	@go tool cover -html=coverage.out

coverage-int: int
	@echo "==> Opening coverage for unit tests..."
	@go tool cover -html=coverage.out

int-build: int build

build: lint test fmt
	@echo "==> Building source code with go build..."
	@go build -mod vendor -v -o terraform-provider-databricks

lint:
	@echo "==> Linting source code with golangci-lint..."
	@golangci-lint run --skip-dirs-use-default

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

# INTEGRATION TESTING WITH TERRAFORM EXAMPLES
terraform-acc: fmt build
	@echo "==> Running Terraform Acceptance Tests..."
	@TF_ACC=1 go test -v -short ./databricks/...

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
	@rm -rf website/content/_index.files/*
	@mkdir -p website/content/_index.files/
	@cp dist/*.gz website/content/_index.files/
	@cp dist/*.zip website/content/_index.files/
	@cd website && hugo

internal-docs-sync:
	@echo "==> Uploading Website..."
	@azcopy login --service-principal --application-id $(AZCOPY_SPA_CLIENT_ID) --tenant-id=$(AZCOPY_SPA_TENANT_ID) && azcopy sync './website/public' '$(AZCOPY_STORAGE_ACCT)' --recursive

.PHONY: build fmt python-setup docs vendor terraform-local build fmt coverage test lint