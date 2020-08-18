# databricks_instance_pool Resource


## Example Usage
my_usage

## Argument Reference

The following arguments are required:

* `enable_elastic_disk` - (Optional) (Bool) 

* `disk_spec` - (Optional) (List)  This field is a block and is documented below.

* `preloaded_spark_versions` - (Optional) (List) 

* `state` - (Computed) (String) 

* `aws_attributes` - (Optional) (List)  This field is a block and is documented below.

* `custom_tags` - (Optional) (Map) 

* `max_capacity` - (Optional) (Integer) 

* `idle_instance_autotermination_minutes` - (Required) (Integer) 

* `node_type_id` - (Required) (String) 

* `instance_pool_name` - (Required) (String) 

* `min_idle_instances` - (Required) (Integer) 



### aws_attributes Configuration Block


* `spot_bid_price_percent` - (Optional) (Integer) 

* `availability` - (Optional) (String) 

* `zone_id` - (Required) (String) 


### disk_spec Configuration Block


* `ebs_volume_type` - (Optional) (String) 

* `azure_disk_volume_type` - (Optional) (String) 

* `disk_count` - (Optional) (Integer) 

* `disk_size` - (Optional) (Integer) 




## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the instance pool.

* `default_tags` - (Map) 


## Import

The resource instance pool can be imported using the `object`, e.g.

```bash
$ terraform import databricks_instance_pool.object <instance pool id>
```