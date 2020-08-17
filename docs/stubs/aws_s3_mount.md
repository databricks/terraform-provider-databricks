# databricks_aws_s3_mount Resource


## Example Usage
my_usage

## Argument Reference

The following arguments are required:

* `cluster_id` - (Optional) (String) 

* `mount_name` - (Required) (String) 

* `s3_bucket_name` - (Required) (String) 




## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the aws s3 mount.

* `source` - (String) 


## Import

The resource aws s3 mount can be imported using the `object`, e.g.

```bash
$ terraform import databricks_aws_s3_mount.object <aws s3 mount id>
```