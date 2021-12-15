---
subcategory: "Compute"
---
# databricks_library resource

Installs a library on [databricks_cluster](cluster.md). Each different type of library has a slightly different syntax. It's possible to set only one type of library within one resource. Otherwise, the plan will fail with an error. 

-> **Note** `databricks_library` resource would always start the associated cluster if it's not running, so make sure to have auto-termination configured. It's not possible to atomically change the version of the same library without cluster restart. Libraries are fully removed from the cluster only after restart.

## Installing library on all clusters

You can install libraries on all clusters with the help of [databricks_clusters](../data-sources/clusters.md) data resource:

```hcl
data "databricks_clusters" "all" {
}

resource "databricks_library" "cli" {
  for_each = data.databricks_clusters.all.ids
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
  path = "/FileStore/app-0.0.1.jar"
}

resource "databricks_library" "app" {
  cluster_id = databricks_cluster.this.id
  jar = databricks_dbfs_file.app.dbfs_path
}
```

## Java/Scala Maven

Installing artifacts from Maven repository. You can also optionally specify a `repo` parameter for custom Maven-style repository, that should be accessible without any authentication. Maven libraries are resolved in Databricks Control Plane, so repo should be accessible from it. It can even be properly configured [maven s3 wagon](https://github.com/seahen/maven-s3-wagon), [AWS CodeArtifact](https://aws.amazon.com/codeartifact/) or [Azure Artifacts](https://azure.microsoft.com/en-us/services/devops/artifacts/).

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
  path = "/FileStore/baz.whl"
}

resource "databricks_library" "app" {
  cluster_id = databricks_cluster.this.id
  whl = databricks_dbfs_file.app.dbfs_path
}
```

## Python PyPI

Installing Python PyPI artifacts. You can optionally also specify the `repo` parameter for custom PyPI mirror, which should be accessible without any authentication for the network that cluster runs in.

-> **Note** `repo` host should be accessible from Internet by Databricks control plane. If connectivity to custom PyPI repositories is required, please modify cluster-node `/etc/pip.conf` through [databricks_global_init_script](global_init_script.md).

```hcl
resource "databricks_library" "fbprophet" {
  cluster_id = databricks_cluster.this.id
  pypi {
    package = "fbprophet==0.6"
    // repo can also be specified here
  }
}
```

## Python EGG

```hcl
resource "databricks_dbfs_file" "app" {
  source = "${path.module}/foo.egg"
  path = "/FileStore/foo.egg"
}

resource "databricks_library" "app" {
  cluster_id = databricks_cluster.this.id
  egg = databricks_dbfs_file.app.dbfs_path
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