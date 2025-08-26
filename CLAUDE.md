# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Common Development Commands

### Building and Testing
- `make build` - Build the provider binary
- `make install` - Build and install the provider locally for Terraform development
- `make test` - Run unit tests with linting
- `make lint` - Run linting with staticcheck (always use this instead of running staticcheck directly)

### Code Quality
- `make fmt` - Format Go code with goimports and gofmt
- `make fmt-docs` - Format code samples in documentation (requires terrafmt)
- `make ws` - Validate whitespace in files

### Development Helpers
- `make vendor` - Populate vendor directory with dependencies
- `make schema` - Print provider schema
- `make diff-schema` - Compare current schema with previous version (useful for migration verification)

### Single Test Execution
```bash
# Run a specific test
go test -v -run TestSpecificTest ./path/to/package

# Run acceptance tests for a specific resource
TF_ACC=1 go test -v -run TestAccResourceName ./path/to/package
```

## Code Architecture

### Dual Provider Implementation
This codebase implements two Terraform provider architectures that are muxed together:
- **SDKv2 Provider**: Legacy implementation in root directories (e.g., `catalog/`, `clusters/`, `jobs/`)
- **Plugin Framework Provider**: New implementation in `internal/providers/pluginfw/products/`

### Key Architecture Components

#### Provider Structure (`internal/providers/`)
- `providers.go` - Main provider muxing logic combining SDKv2 and Plugin Framework
- `sdkv2/` - SDKv2-specific provider implementation
- `pluginfw/` - Plugin Framework provider implementation with auto-generated schemas
- `common/` - Shared utilities between both providers

#### Service Models (`internal/service/`)
Auto-generated Go structs from Databricks SDK:
- `*_tf/model.go` - Current Plugin Framework compatible structs
- `*_tf/legacy_model.go` - SDKv2 compatible structs with `_SdkV2` suffix

#### Resource Organization
- **Root directories** (e.g., `catalog/`, `jobs/`, `clusters/`): SDKv2 resources and data sources
- **`internal/providers/pluginfw/products/`**: Plugin Framework resources organized by service

### Migration Pattern
Resources are being migrated from SDKv2 to Plugin Framework. When migrating:
1. Use `_SdkV2` suffixed structs from `internal/service/` for schema compatibility
2. Call `cs.ConfigureAsSdkV2Compatible()` in schema definition
3. Ensure no schema breaking changes with `make diff-schema`

### Resource Development Patterns

#### Adding SDKv2 Resources
1. Create resource file in appropriate root directory (e.g., `catalog/resource_new_thing.go`)
2. Use `common.Resource{}` helper with auto-generated schema from struct tags
3. Add to provider in `providers/sdkv2/sdkv2.go`

#### Adding Plugin Framework Resources
1. Create in `internal/providers/pluginfw/products/{service}/`
2. Use `ResourceStructToSchema()` with structs from `internal/service/{service}_tf/`
3. Implement required interfaces (`ResourceWithConfigure`, etc.)
4. Add to `internal/providers/pluginfw/pluginfw.go`

### Client Architecture
- `common.DatabricksClient` - Core client wrapper
- Access workspace client via `client.GetWorkspaceClient()`
- Access account client via `client.GetAccountClient()`
- Client automatically handles authentication and retries

### Testing Structure
- Unit tests: `*_test.go` files using `qa.ResourceFixture` for HTTP mocking
- Integration tests: `*_acc_test.go` files with live API testing
- Test naming conventions determine environment:
  - `TestAcc*` - Workspace-level tests across all clouds
  - `TestMwsAcc*` - Account-level tests across all clouds
  - `TestUcAcc*` - Unity Catalog tests across all clouds

## Development Guidelines

### Code Organization
- Files should not exceed 600 lines
- Functions should fit on a 13" screen (max 40 lines, except tests)
- No unnecessary package exports (avoid public structs/types unless needed outside package)
- Use `qa.EnvironmentTemplate()` instead of complex `fmt.Sprintf` with >4 placeholders

### Import Conventions
Order imports as: Go standard library, vendor packages, current provider packages.
Within each section, maintain alphabetical order.

### Documentation
- All resources require Terraform Registry compatible documentation in `docs/`
- Code samples must be formatted with `make fmt-docs`
- Cross-link integrity between markdown files is required
- Use Terraform Registry Doc Preview Tool for validation

### Changelog Requirements
All user-facing changes must be documented in `NEXT_CHANGELOG.md` with format:
```
* <Summary of change> ([#<PR number>](<PR link>)).

  <Optional additional information>
```

### Migration Verification
When migrating resources to Plugin Framework, always run `make diff-schema` to ensure no breaking changes to the Terraform schema.
