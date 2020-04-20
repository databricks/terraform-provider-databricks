# Resource: databricks_scim_group

This resource allows you to create groups in Databricks. You can also associate Databricks users to the following groups. 

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

    resource "databricks_scim_group" "my-group" {
      display_name = "Sri Test Group"
      members = ["${databricks_scim_user.my-user.id}"]
    }

## Argument Reference

The following arguments are supported:

.. _r_scim_group_display_name:
* :ref:`display_name <r_scim_group_display_name>`- **(Required)** This is the display name for the given group.

.. _r_scim_group_members:
* :ref:`members <r_scim_group_members>` - **(Optional)** This is a list of users associated to the given group.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

.. _r_scim_group_id:
* :ref:`id <r_scim_group_id>` - The id for the scim group object.

.. _r_scim_group_display_name1:
* :ref:`display_name <r_scim_group_display_name1>`- This is the display name for the given group.

.. _r_scim_group_members1:
* :ref:`members <r_scim_group_members1>` - This is a list of users associated to the given group.

## Import

.. Note:: Importing this resource is not currently supported.