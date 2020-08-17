# databricks_cluster Resource


## Example Usage
my_usage

## Argument Reference

The following arguments are required:

* `ssh_public_keys` - (Optional) (List) 

* `init_scripts` - (Optional) (List)  This field is a block and is documented below.

* `idempotency_token` - (Optional) (String) 

* `library_jar` - (Optional) (Set)  This field is a block and is documented below.

* `enable_elastic_disk` - (Computed) (Bool) 

* `single_user_name` - (Optional) (String) 

* `libraries` - (Optional) (Set)  This field is a block and is documented below.

* `library_maven` - (Optional) (Set)  This field is a block and is documented below.

* `cluster_id` - (Computed) (String) 

* `spark_version` - (Required) (String) 

* `instance_pool_id` - (Optional) (String) 

* `policy_id` - (Optional) (String) 

* `spark_env_vars` - (Optional) (Map) 

* `num_workers` - (Optional) (Integer) 

* `driver_node_type_id` - (Computed) (String) 

* `custom_tags` - (Optional) (Map) 

* `cluster_log_conf` - (Optional) (List)  This field is a block and is documented below.

* `library_egg` - (Optional) (Set)  This field is a block and is documented below.

* `library_whl` - (Optional) (Set)  This field is a block and is documented below.

* `aws_attributes` - (Optional) (List)  This field is a block and is documented below.

* `autotermination_minutes` - (Optional) (Integer) 

* `spark_conf` - (Optional) (Map) 

* `docker_image` - (Optional) (List)  This field is a block and is documented below.

* `library_pypi` - (Optional) (Set)  This field is a block and is documented below.

* `library_cran` - (Optional) (Set)  This field is a block and is documented below.

* `cluster_name` - (Optional) (String) 

* `autoscale` - (Optional) (List)  This field is a block and is documented below.

* `node_type_id` - (Computed) (String) 



### autoscale Configuration Block


* `min_workers` - (Optional) (Integer) 

* `max_workers` - (Optional) (Integer) 


### library_cran Configuration Block


* `messages` - (Computed) (List) 

* `status` - (Computed) (String) 

* `package` - (Optional) (String) 

* `repo` - (Optional) (String) 


### library_pypi Configuration Block


* `repo` - (Optional) (String) 

* `messages` - (Computed) (List) 

* `status` - (Computed) (String) 

* `package` - (Optional) (String) 


### docker_image Configuration Block


* `basic_auth` - (Optional) (List)  This field is a block and is documented below.

* `url` - (Required) (String) 


### basic_auth for docker_image Configuration Block


* `password` - (Required) (String) 

* `username` - (Required) (String) 


### aws_attributes Configuration Block


* `first_on_demand` - (Computed) (Integer) 

* `availability` - (Computed) (String) 

* `zone_id` - (Computed) (String) 

* `instance_profile_arn` - (Optional) (String) 

* `spot_bid_price_percent` - (Computed) (Integer) 

* `ebs_volume_type` - (Computed) (String) 

* `ebs_volume_count` - (Computed) (Integer) 

* `ebs_volume_size` - (Computed) (Integer) 


### library_whl Configuration Block


* `messages` - (Computed) (List) 

* `status` - (Computed) (String) 

* `path` - (Optional) (String) 


### library_egg Configuration Block


* `status` - (Computed) (String) 

* `path` - (Optional) (String) 

* `messages` - (Computed) (List) 


### cluster_log_conf Configuration Block


* `dbfs` - (Optional) (List)  This field is a block and is documented below.

* `s3` - (Optional) (List)  This field is a block and is documented below.


### s3 for cluster_log_conf Configuration Block


* `encryption_type` - (Optional) (String) 

* `kms_key` - (Optional) (String) 

* `canned_acl` - (Optional) (String) 

* `destination` - (Required) (String) 

* `region` - (Optional) (String) 

* `endpoint` - (Optional) (String) 

* `enable_encryption` - (Optional) (Bool) 


### dbfs for cluster_log_conf Configuration Block


* `destination` - (Required) (String) 


### library_maven Configuration Block


* `status` - (Computed) (String) 

* `coordinates` - (Optional) (String) 

* `repo` - (Optional) (String) 

* `exclusions` - (Optional) (List) 

* `messages` - (Computed) (List) 


### libraries Configuration Block


* `pypi` - (Optional) (List)  This field is a block and is documented below.

* `maven` - (Optional) (List)  This field is a block and is documented below.

* `cran` - (Optional) (List)  This field is a block and is documented below.

* `jar` - (Optional) (String) 

* `egg` - (Optional) (String) 

* `whl` - (Optional) (String) 


### cran for libraries Configuration Block


* `repo` - (Optional) (String) 

* `package` - (Required) (String) 


### maven for libraries Configuration Block


* `coordinates` - (Required) (String) 

* `repo` - (Optional) (String) 

* `exclusions` - (Optional) (List) 


### pypi for libraries Configuration Block


* `package` - (Required) (String) 

* `repo` - (Optional) (String) 


### library_jar Configuration Block


* `messages` - (Computed) (List) 

* `status` - (Computed) (String) 

* `path` - (Optional) (String) 


### init_scripts Configuration Block


* `s3` - (Optional) (List)  This field is a block and is documented below.

* `dbfs` - (Optional) (List)  This field is a block and is documented below.


### dbfs for init_scripts Configuration Block


* `destination` - (Required) (String) 


### s3 for init_scripts Configuration Block


* `canned_acl` - (Optional) (String) 

* `destination` - (Required) (String) 

* `region` - (Optional) (String) 

* `endpoint` - (Optional) (String) 

* `enable_encryption` - (Optional) (Bool) 

* `encryption_type` - (Optional) (String) 

* `kms_key` - (Optional) (String) 




## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the cluster.

* `state` - (String) 

* `default_tags` - (Map) 


## Import

The resource cluster can be imported using the `object`, e.g.

```bash
$ terraform import databricks_cluster.object <cluster id>
```