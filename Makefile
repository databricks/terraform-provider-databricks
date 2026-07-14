default: build

fmt:
	@echo "✓ Formatting source code with goimports ..."
	@go tool goimports -w $(shell find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./.git/*")
	@echo "✓ Formatting source code with gofmt ..."
	@gofmt -w $(shell find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./.git/*")

fmt-docs:
	@echo "✓ Formatting code samples in documentation"
	@terrafmt fmt -p '*.md' .

ws:
	./tools/validate_whitespace.py

lint: vendor
	@echo "✓ Linting source code with https://staticcheck.io/ ..."
	@go tool staticcheck ./...

test: lint
	@echo "✓ Running tests ..."
	@go tool gotestsum --format pkgname-and-test-fails --no-summary=skipped --raw-command go test -v -json -short -coverprofile=coverage.txt ./...

coverage: test
	@echo "✓ Opening coverage for unit tests ..."
	@go tool cover -html=coverage.txt

build: vendor
	@echo "✓ Building source code with go build ..."
	@go build -mod vendor -v -o terraform-provider-databricks

install: build
	@echo "✓ Installing provider for Terraform 1.0+ into ~/.terraform.d/plugins ..."
	@mkdir -p '$(HOME)/.terraform.d/plugins/registry.terraform.io/databricks/databricks/$(shell ./terraform-provider-databricks version)/$(shell go version | awk '{print $$4}' | sed 's#/#_#')'
	@cp terraform-provider-databricks '$(HOME)/.terraform.d/plugins/registry.terraform.io/databricks/databricks/$(shell ./terraform-provider-databricks version)/$(shell go version | awk '{print $$4}' | sed 's#/#_#')/terraform-provider-databricks'
	@echo "✓ Use the following configuration to enable the version you've built"
	@echo
	@echo "terraform {"
	@echo "  required_providers {"
	@echo "    databricks = {"
	@echo '      source = "databricks/databricks"'
	@echo '      version = "$(shell ./terraform-provider-databricks version)"'
	@echo "    }"
	@echo "  }"
	@echo "}"

vendor:
	@echo "✓ Filling vendor folder with library code ..."
	@go mod vendor

test-azcli: install
	@echo "✓ Running Terraform Acceptance Tests for Azure..."
	@/bin/bash scripts/run.sh azcli '^TestAzureAcc' --debug --tee

test-azsp: install
	@echo "✓ Running Terraform Acceptance Tests for Azure..."
	@/bin/bash scripts/run.sh azsp '^(TestAcc|TestAzureAcc)' --debug --tee

test-mws: install
	@echo "✓ Running acceptance Tests for Multiple Workspace APIs on AWS..."
	@/bin/bash scripts/run.sh mws '^TestMwsAcc' --debug --tee

test-awsmt: install
	@echo "✓ Running Terraform Acceptance Tests for AWS MT..."
	@/bin/bash scripts/run.sh awsmt '^(TestAcc|TestAwsAcc)' --debug --tee

test-gcp-accounts: install
	@echo "✓ Running acceptance Tests for Multiple Workspace APIs on GCP..."
	@/bin/bash scripts/run.sh gcp-accounts '^TestGcpaAcc' --debug --tee

test-gcp: install
	@echo "✓ Running acceptance Tests for GCP..."
	@/bin/bash scripts/run.sh gcp '^(TestAcc|TestGcpAcc)' --debug --tee

test-preview: install
	@echo "✓ Running acceptance Tests for Preview features..."
	@/bin/bash scripts/run.sh preview '^TestPreviewAcc' --debug --tee

docker-it:
	docker build -t databricks-terrafrom/test -f scripts/Dockerfile .

schema:
	@/bin/bash scripts/print-schema.sh

diff-schema:
	@/bin/bash scripts/diff-schema.sh

.PHONY: build fmt python-setup docs vendor build fmt coverage test lint ws
