# Resource: databricks_instance_profile

This resource allows you to create, get, and delete instance profiles that users can launch clusters with.

## Example Usage

.. code-block:: tf

    resource "databricks_instance_profile" "db-instance-profile" {
      instance_profile_arn = "arn:aws:iam::999999999999:instance-profile/custom-s3-access-instance-profile"
      skip_validation = true
    }
    
## Argument Reference

The following arguments are supported:

.. _r_instance_profile_instance_profile_arn:
* :ref:`instance_profile_arn <r_instance_profile_instance_profile_arn>` - **(Required)** 

.. _r_instance_profile_skip_validation:
* :ref:`skip_validation <r_instance_profile_skip_validation>` - **(Required)** 


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

.. _r_instance_profile_id:
* :ref:`id <r_instance_profile_id>` - The id for the instance profile object.


## Import

.. Note:: Importing this resource is not currently supported.