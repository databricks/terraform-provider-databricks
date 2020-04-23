+++
title = "instance_profile"
date = 2020-04-20T23:34:03-04:00
weight = 15
chapter = false
pre = ""
+++


## Resource: `databricks_instance_profile`

This resource allows you to create, get, and delete instance profiles that users can launch clusters with.

## Example Usage

```hcl
resource "databricks_instance_profile" "db-instance-profile" {
  instance_profile_arn = "arn:aws:iam::999999999999:instance-profile/custom-s3-access-instance-profile"
  skip_validation = true
}
```  
    
## Argument Reference

The following arguments are supported:

#### - `instance_profile_arn`:
> **(Required)** 

#### - `skip_validation`:
> **(Required)** 


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

#### - `id`:
> The id for the instance profile object.


## Import

{{% notice note %}}
Importing this resource is not currently supported.
{{% /notice %}}
