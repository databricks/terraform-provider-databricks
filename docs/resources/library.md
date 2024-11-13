---
subcategory: "Compute"
---
# databricks_library resource

Installs a [library](https://docs.databricks.com/libraries/index.html) on [databricks_cluster](cluster.md). Each different type of library has a slightly different syntax. It's possible to set only one type of library within one resource. Otherwise, the plan will fail with an error.

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
resource "databricks_dbfs_file" "app" {
  source = "${path.module}/app-0.0.1.jar"
  path   = "/FileStore/app-0.0.1.jar"
}

resource "databricks_library" "app" {
  cluster_id = databricks_cluster.this.id
  jar        = databricks_dbfs_file.app.dbfs_path
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
resource "databricks_dbfs_file" "app" {
  source = "${path.module}/baz.whl"
  path   = "/FileStore/baz.whl"
}

resource "databricks_library" "app" {
  cluster_id = databricks_cluster.this.id
  whl        = databricks_dbfs_file.app.dbfs_path
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


## Python EGG

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


## Import

!> Importing this resource is not currently supported.

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_clusters](../data-sources/clusters.md) data to retrieve a list of [databricks_cluster](cluster.md) ids.
* [databricks_cluster](cluster.md) to create [Databricks Clusters](https://docs.databricks.com/clusters/index.html).
* [databricks_cluster_policy](cluster_policy.md) to create a [databricks_cluster](cluster.md) policy, which limits the ability to create clusters based on a set of rules.
* [databricks_dbfs_file](../data-sources/dbfs_file.md) data to get file content from [Databricks File System (DBFS)](https://docs.databricks.com/data/databricks-file-system.html).
* [databricks_dbfs_file_paths](../data-sources/dbfs_file_paths.md) data to get list of file names from get file content from [Databricks File System (DBFS)](https://docs.databricks.com/data/databricks-file-system.html).
* [databricks_dbfs_file](dbfs_file.md) to manage relatively small files on [Databricks File System (DBFS)](https://docs.databricks.com/data/databricks-file-system.html).
* [databricks_global_init_script](global_init_script.md) to manage [global init scripts](https://docs.databricks.com/clusters/init-scripts.html#global-init-scripts), which are run on all [databricks_cluster](cluster.md#init_scripts) and [databricks_job](job.md#new_cluster).
* [databricks_job](job.md) to manage [Databricks Jobs](https://docs.databricks.com/jobs.html) to run non-interactive code in a [databricks_cluster](cluster.md).
* [databricks_mount](mount.md) to [mount your cloud storage](https://docs.databricks.com/data/databricks-file-system.html#mount-object-storage-to-dbfs) on `dbfs:/mnt/name`.
* [databricks_pipeline](pipeline.md) to deploy [Delta Live Tables](https://docs.databricks.com/data-engineering/delta-live-tables/index.html).
* [databricks_repo](repo.md) to manage [Databricks Repos](https://docs.databricks.com/repos.html).
