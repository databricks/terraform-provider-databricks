# Resource: databricks_secret_acl

Create or overwrite the ACL associated with the given principal (user or group) on the specified scope point. 
In general, a user or group will use the most powerful permission available to them, and 
permissions are ordered as follows:

* **MANAGE** - Allowed to change ACLs, and read and write to this secret scope.
* **WRITE** - Allowed to read and write to this secret scope.
* **READ** - Allowed to read this secret scope and list what secrets are available.

## Example Usage

.. code-block:: tf

    resource "databricks_secret_scope" "my-scope" {
      name = "terraform-demo-scope"
    }
    
    resource "databricks_secret_acl" "my-acl" {
      principal = "USERS"
      permission = "READ"
      scope = "${databricks_secret_scope.my-scope.name}"
    }

## Argument Reference

The following arguments are supported:

.. _r_secret_acl_scope:
* :ref:`scope <r_secret_acl_scope>`- **(Required)** The name of the scope to remove permissions from. This field is required. 
(MB4) form.

.. _r_secret_acl_principal:
* :ref:`principal <r_secret_acl_principal>` - **(Required)** The principal to remove an existing ACL from. The principal is a user 
or group name corresponding to an existing Databricks principal to be granted or revoked access. This field is required. 

.. _r_secret_acl_permission:
* :ref:`permission <r_secret_acl_permission>` - **(Required)** The permission level applied to the principal. 
Options are: `"READ", "WRITE", "MANAGE"`. This field is required.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

.. _r_secret_acl_id:
* :ref:`id <r_secret_acl_id>` - The id for the secret scope acl object.

.. _r_secret_acl_scope-1:
* :ref:`scope <r_secret_acl_scope-1>` - The name of the scope to remove permissions from.
(MB4) form.

.. _r_secret_acl_principal-1:
* :ref:`principal <r_secret_acl_principal-1>` - The principal to remove an existing ACL from. The principal is a user 
or group name corresponding to an existing Databricks principal to be granted or revoked access.  

.. _r_secret_acl_permission-1:
* :ref:`permission <r_secret_acl_permission-1>` - The permission level applied to the principal. 

## Import

.. Note:: Importing this resource is not currently supported.