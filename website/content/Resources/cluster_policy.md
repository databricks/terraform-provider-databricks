+++
title = "cluster_policy"
date = 2020-06-21T23:34:03-04:00
weight = 15
chapter = false
pre = ""
+++

## Resource: `databricks_cluster_policy`

This resource creates a cluster policy, which limits the ability to create clusters based on a set of rules. The policy rules limit the attributes or attribute values available for cluster creation. Cluster policies have ACLs that limit their use to specific users and groups. Only admin users can create, edit, and delete policies. Admin users also have access to all policies.

Cluster policies let you:

* Limit users to create clusters with prescribed settings.
* Simplify the user interface and enable more users to create their own clusters (by fixing and hiding some values).
* Control cost by limiting per cluster maximum cost (by setting limits on attributes whose values contribute to hourly price).

Cluster policy permissions limit which policies a user can select in the Policy drop-down when the user creates a cluster:

* If no policies have been created in the workspace, the Policy drop-down does not display.
* A user who has cluster create permission can select the Free form policy and create fully-configurable clusters.
* A user who has both cluster create permission and access to cluster policies can select the Free form policy and policies they have access to.
* A user that has access to only cluster policies, can select the policies they have access to.

## Example Usage

The following example defines [external metastore](https://docs.databricks.com/administration-guide/clusters/policies.html#external-metastore-policy) policy:

```hcl
resource "databricks_cluster_policy" "external_metastore" {
    name = "Use Enterprise Metastore"
    definition = jsonencode({
        "spark_conf.spark.hadoop.javax.jdo.option.ConnectionURL": {
            "type": "fixed",
            "value": "jdbc:sqlserver://<jdbc-url>"
        },
        "spark_conf.spark.hadoop.javax.jdo.option.ConnectionDriverName": {
            "type": "fixed",
            "value": "com.microsoft.sqlserver.jdbc.SQLServerDriver"
        },
        "spark_conf.spark.databricks.delta.preview.enabled": {
            "type": "fixed",
            "value": true
        },
        "spark_conf.spark.hadoop.javax.jdo.option.ConnectionUserName": {
            "type": "fixed",
            "value": "<metastore-user>"
        },
        "spark_conf.spark.hadoop.javax.jdo.option.ConnectionPassword": {
            "type": "fixed",
            "value": "<metastore-password>"
        }
        })
}
```

## Argument Reference

The following arguments are required:

#### - `name`:
> **(Required)** Cluster policy name. This must be unique. Length must be between 1 and 100 characters.

#### - `definition`:
> **(Required)** Policy definition JSON document expressed in [Databricks Policy Definition Language](https://docs.databricks.com/administration-guide/clusters/policies.html#cluster-policy-definition).


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

#### - `policy_id`:
> Canonical unique identifier for the cluster policy.

## Import

{{% notice note %}}
Importing this resource is not currently supported.
{{% /notice %}}