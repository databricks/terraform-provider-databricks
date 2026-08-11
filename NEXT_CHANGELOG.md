# NEXT CHANGELOG

## Release v1.126.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Honor `config.DefaultHostMetadataResolverFactory` during provider configuration ([#NNNN](https://github.com/databricks/terraform-provider-databricks/pull/NNNN)).

  The provider installs a wrapper around the SDK's host-metadata resolver to observe the host's `workspace_id`. Because the wrapper always set `HostMetadataResolver`, the SDK never consulted the `DefaultHostMetadataResolverFactory` global, so a resolver installed through that factory was silently ignored. The wrapper now replicates the SDK's resolver precedence (a pre-existing resolver, then the factory, then the built-in fetch).

### Documentation

### Exporter

### Internal Changes
