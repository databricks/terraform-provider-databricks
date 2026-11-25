# NEXT CHANGELOG

## Release v1.145.0

### Important Changes

### Breaking Changes

* Replace the connection `secrets` map with inline Unity Catalog secret references in `options` ([#6211](https://github.com/databricks/terraform-provider-databricks/pull/6211)).

  The `secrets` argument has been removed. Move each entry into the existing `options` map and wrap its secret name in `{{secrets/...}}`:

  ```hcl
  # Before
  secrets = {
    password = "main.default.mysql_password"
  }

  # After
  options = {
    # Keep the other connection options here.
    password = "{{secrets/main.default.mysql_password}}"
  }
  ```

### New Features and Improvements

### Bug Fixes

### Documentation

### Exporter

### Internal Changes
