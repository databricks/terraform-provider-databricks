# Resource: db_secret_scope

This resource creates a Databricks-backed secret scope in which secrets are stored in Databricks-managed storage and 
encrypted with a cloud-based specific encryption key. 

The scope name:

* Must be unique within a workspace.
* Must consist of alphanumeric characters, dashes, underscores, and periods, and may not exceed 128 characters.

## Example Usage

.. code-block:: tf

    resource "db_secret_scope" "my-scope" {
      name = "terraform-demo-scope"
      initial_manage_principal = "users"
    }

## Argument Reference

The following arguments are supported:

.. _r_secret_scope_name:
* :ref:`name <r_secret_scope_name>`- **(Required)** Scope name requested by the user. Scope names are unique. This field is required.

.. _r_secret_scope_initial_manage_principal:
* :ref:`initial_manage_principal <r_secret_scope_initial_manage_principal>` - **(Optional)** The principal that is initially granted 
MANAGE permission to the created scope.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

.. _r_secret_scope_id:
* :ref:`id <r_secret_scope_id>` - The id for the secret scope object.

.. _r_secret_scope_name-1:
* :ref:`name <r_secret_scope_name-1>`- Scope name requested by the user. Scope names are unique. This field is required.

.. _r_secret_scope_initial_manage_principal-1:
* :ref:`initial_manage_principal <r_secret_scope_initial_manage_principal-1>` - The principal that is initially granted 
MANAGE permission to the created scope.

.. _r_secret_scope_backend_type:
* :ref:`backend_type <r_secret_scope_backend_type>` - The type of secret scope backend.

## Import

.. Note:: Importing this resource is not currently supported.