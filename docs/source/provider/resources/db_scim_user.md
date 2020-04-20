# Resource: databricks_scim_user

This resource allows you to create users in Databricks and give them the proper level of access, as well as 
remove access for users (deprovision them) when they leave your organization or no longer need access to Databricks.

.. Important:: You must be a Databricks administrator API token to use SCIM resources. 

## Example Usage

.. code-block:: tf

    resource "databricks_scim_user" "my-user" {
      user_name = "testuser@databricks.com"
      display_name = "Test User"
      entitlements = [
        "allow-cluster-create",
      ]
    }

## Argument Reference

The following arguments are supported:

.. _r_scim_user_user_name:
* :ref:`user_name <r_scim_user_user_name>`- **(Required)** This is the username of the given user and will be their form of access 
and identity.

.. _r_scim_user_display_name:
* :ref:`display_name <r_scim_user_display_name>` - **(Optional)** This is an alias for the username can be the full name of the user.

.. _r_scim_user_roles:
* :ref:`roles <r_scim_user_roles>` - **(Optional)** This is a list of roles assigned to the user, specific to the AWS environment for 
user to assume roles on clusters.

.. _r_scim_user_entitlements:
* :ref:`entitlements <r_scim_user_entitlements>` - **(Optional)** Entitlements for the user to be able to have the ability to create 
clusters and pools. Current options are: `"allow-cluster-create", "allow-instance-pool-create"`.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

.. _r_scim_user_id:
* :ref:`id <r_scim_user_id>` - The id for the scim user object.

.. _r_scim_user_user_name1:
* :ref:`user_name <r_scim_user_user_name1>`- This is the username of the given user and will be their form of access 
and identity.

.. _r_scim_user_display_name1:
* :ref:`display_name <r_scim_user_display_name1>` - This is an alias for the username can be the full name of the user.

.. _r_scim_user_roles1:
* :ref:`roles <r_scim_user_roles1>` - This is a list of role arns assigned to the user, specific to the AWS environment for 
user to assume roles on clusters.

.. _r_scim_user_entitlements1:
* :ref:`entitlements <r_scim_user_entitlements1>` - This is a list of entitlements for the user to be able to have the ability to create 
clusters and pools.

## Import

.. Note:: Importing this resource is not currently supported.