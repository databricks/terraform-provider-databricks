---
subcategory: "Compute"
---
# databricks_library resource

Installs a [library](https://docs.databricks.com/libraries/index.html) on [databricks_cluster](cluster.md). Each different type of library has a slightly different syntax. It's possible to set only one type of library within one resource. Otherwise, the plan will fail with an error.

-> This resource can only be used with a workspace-level provider!

-> `databricks_library` resource would always start the associated cluster if it's not running, so make sure to have auto-termination configured. It's not possible to atomically change the version of the same library without cluster restart. Libraries are fully removed from the cluster only after restart.

## Plugin Framework Migration

The library resource has been migrated from sdkv2 to plugin framework。 If you encounter any problem with this resource and suspect it is due to the migration, you can fallback to sdkv2 by setting the environment variable in the following way `export USE_SDK_V2_RESOURCES="databricks_library"`.

## Installing library on all clusters

You can install libraries on all clusters with the help of [databricks_clusters](../data-sources/clusters.md) data resource:

```hcl
data "databricks_clusters" "all" {
}

resource "databricks_library" "cli" {
  for_each   = data.databricks_clusters.all.ids
  cluster_id = each.key
  pypi {
    package = "databricks-cli"
  }
}
```

## Java/Scala JAR

```hcl
resource "databricks_file" "app" {
  source = "${path.module}/app-0.0.1.jar"
  path   = "/Volumes/catalog/schema/volume/app-0.0.1.jar"
}

resource "databricks_library" "app" {
  cluster_id = databricks_cluster.this.id
  jar        = databricks_file.app.path
}
```

## Java/Scala Maven

Installing artifacts from Maven repository. You can also optionally specify a `repo` parameter for a custom Maven-style repository, that should be accessible without any authentication. Maven libraries are resolved in Databricks Control Plane, so repo should be accessible from it. It can even be properly configured [maven s3 wagon](https://github.com/seahen/maven-s3-wagon), [AWS CodeArtifact](https://aws.amazon.com/codeartifact/) or [Azure Artifacts](https://azure.microsoft.com/en-us/services/devops/artifacts/).

```hcl
resource "databricks_library" "deequ" {
  cluster_id = databricks_cluster.this.id
  maven {
    coordinates = "com.amazon.deequ:deequ:1.0.4"
    // exclusions block is optional
    exclusions = ["org.apache.avro:avro"]
  }
}
```

## Python Wheel

```hcl
resource "databricks_file" "app" {
  source = "${path.module}/baz.whl"
  path   = "/Volumes/catalog/schema/volume/baz.whl"
}

resource "databricks_library" "app" {
  cluster_id = databricks_cluster.this.id
  whl        = databricks_file.app.path
}
```

## Python PyPI

Installing Python PyPI artifacts. You can optionally also specify the `repo` parameter for a custom PyPI mirror, which should be accessible without any authentication for the network that cluster runs in.

-> `repo` host should be accessible from the Internet by Databricks control plane. If connectivity to custom PyPI repositories is required, please modify cluster-node `/etc/pip.conf` through [databricks_global_init_script](global_init_script.md).

```hcl
resource "databricks_library" "fbprophet" {
  cluster_id = databricks_cluster.this.id
  pypi {
    package = "fbprophet==0.6"
    // repo can also be specified here
  }
}
```

## Python requirements files

Installing Python libraries listed in the `requirements.txt` file.  Only Workspace paths and Unity Catalog Volumes paths are supported.  Requires a cluster with DBR 15.0+.

```hcl
resource "databricks_library" "libraries" {
  cluster_id   = databricks_cluster.this.id
  requirements = "/Workspace/path/to/requirements.txt"
}
```

## Python EGG (Deprecated)

```hcl
resource "databricks_dbfs_file" "app" {
  source = "${path.module}/foo.egg"
  path   = "/FileStore/foo.egg"
}

resource "databricks_library" "app" {
  cluster_id = databricks_cluster.this.id
  egg        = databricks_dbfs_file.app.dbfs_path
}
```

## R CRan

Installing artifacts from CRan. You can also optionally specify a `repo` parameter for a custom cran mirror.

```hcl
resource "databricks_library" "rkeops" {
  cluster_id = databricks_cluster.this.id
  cran {
    package = "rkeops"
  }
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required) ID of the [databricks_cluster](cluster.md) to install the library on.

You must specify exactly **one** of the following library types:

* `jar` - (Optional) Path to the JAR library. Supported URIs include Workspace paths, Unity Catalog Volumes paths, and S3 URIs. For example: `/Workspace/path/to/library.jar`, `/Volumes/path/to/library.jar` or `s3://my-bucket/library.jar`. If S3 is used, make sure the cluster has read access to the library. You may need to launch the cluster with an IAM role to access the S3 URI.

* `egg` - (Optional, Deprecated) Path to the EGG library. Installing Python egg files is deprecated and is not supported in Databricks Runtime 14.0 and above. Use `whl` or `pypi` instead.

* `whl` - (Optional) Path to the wheel library. Supported URIs include Workspace paths, Unity Catalog Volumes paths, and S3 URIs. For example: `/Workspace/path/to/library.whl`, `/Volumes/path/to/library.whl` or `s3://my-bucket/library.whl`. If S3 is used, make sure the cluster has read access to the library. You may need to launch the cluster with an IAM role to access the S3 URI.

* `requirements` - (Optional) Path to the requirements.txt file. Only Workspace paths and Unity Catalog Volumes paths are supported. For example: `/Workspace/path/to/requirements.txt` or `/Volumes/path/to/requirements.txt`. Requires a cluster with DBR 15.0+.

* `maven` - (Optional) Configuration block for a Maven library. The block consists of the following fields:
  * `coordinates` - (Required) Gradle-style Maven coordinates. For example: `org.jsoup:jsoup:1.7.2`.
  * `repo` - (Optional) Maven repository to install the Maven package from. If omitted, both Maven Central Repository and Spark Packages are searched.
  * `exclusions` - (Optional) List of dependencies to exclude. For example: `["slf4j:slf4j", "*:hadoop-client"]`. See [Maven dependency exclusions](https://maven.apache.org/guides/introduction/introduction-to-optional-and-excludes-dependencies.html) for more information.

* `pypi` - (Optional) Configuration block for a PyPI library. The block consists of the following fields:
  * `package` - (Required) The name of the PyPI package to install. An optional exact version specification is also supported. For example: `simplejson` or `simplejson==3.8.0`.
  * `repo` - (Optional) The repository where the package can be found. If not specified, the default pip index is used.

* `cran` - (Optional) Configuration block for a CRAN library. The block consists of the following fields:
  * `package` - (Required) The name of the CRAN package to install.
  * `repo` - (Optional) The repository where the package can be found. If not specified, the default CRAN repo is used.

* `provider_config` - (Optional) Configuration block for management through the account provider. This block consists of the following fields:
  * `workspace_id` - (Required) Workspace ID that the resource belongs to. This workspace must be part of the account that the provider is configured with.

## Import

!> Importing this resource is not currently supported.

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_clusters](../data-sources/clusters.md) data to retrieve a list of [databricks_cluster](cluster.md) ids.
* [databricks_cluster](cluster.md) to create [Databricks Clusters](https://docs.databricks.com/clusters/index.html).
* [databricks_cluster_policy](cluster_policy.md) to create a [databricks_cluster](cluster.md) policy, which limits the ability to create clusters based on a set of rules.
* [databricks_global_init_script](global_init_script.md) to manage [global init scripts](https://docs.databricks.com/clusters/init-scripts.html#global-init-scripts), which are run on all [databricks_cluster](cluster.md#init_scripts) and [databricks_job](job.md#new_cluster).
* [databricks_job](job.md) to manage [Databricks Jobs](https://docs.databricks.com/jobs.html) to run non-interactive code in a [databricks_cluster](cluster.md).
* [databricks_pipeline](pipeline.md) to deploy [Lakeflow Declarative Pipelines](https://docs.databricks.com/aws/en/dlt).
* [databricks_repo](repo.md) to manage [Databricks Repos](https://docs.databricks.com/repos.html).
