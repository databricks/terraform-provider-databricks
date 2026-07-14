# NEXT CHANGELOG

## Release v1.122.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Exporter

### Internal Changes

* Migrate `databricks_repo` resource off the hand-rolled `ReposAPI`/`reposCreateRequest` HTTP client onto the Go SDK `workspace.Repos` service, generating the schema directly from `workspace.RepoInfo` via field aliases.
