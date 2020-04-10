# Resource: db_secret

With this resource you can insert a secret under the provided scope with the given name. If a secret already exists 
with the same name, this command overwrites the existing secret’s value. The server encrypts the secret using the 
secret scope’s encryption settings before storing it. You must have WRITE or MANAGE permission on the secret scope.

The secret key must consist of alphanumeric characters, dashes, underscores, and periods, and cannot exceed 
128 characters. The maximum allowed secret value size is 128 KB. The maximum number of secrets in a given scope is 1000.

You can read a secret value only from within a command on a cluster (for example, through a notebook); there is no API 
to read a secret value outside of a cluster. The permission applied is based on who is invoking the command and you must 
have at least READ permission.

## Example Usage

.. code-block:: tf

    resource "db_secret_scope" "my-scope" {
      name = "terraform-demo-scope"
      initial_manage_principal = "users"
    }
    
    resource "db_secret" "my_secret" {
      key = "test-secret-1"
      string_value = "hello world 123"
      scope = "${db_secret_scope.my-scope.name}"
    }

## Argument Reference

The following arguments are supported:

.. _r_secret_string_value:
* :ref:`string_value <r_secret_string_value>`- **(Required)** If string_value, if specified, the value will be stored in UTF-8 
(MB4) form.

.. _r_secret_scope:
* :ref:`scope <r_secret_scope>` - **(Required)** The name of the scope to which the secret will be associated with. 
This field is required.

.. _r_secret_key:
* :ref:`key <r_secret_key>` - **(Required)** A unique name to identify the secret. This field is required.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

.. _r_secret_id:
* :ref:`id <r_secret_id>` - The id for the secret object.

.. _r_secret_string_value-1:
* :ref:`string_value <r_secret_string_value-1>`- If string_value, if specified, the value will be stored in UTF-8 
(MB4) form.

.. _r_secret_scope-1:
* :ref:`scope <r_secret_scope-1>` - The name of the scope to which the secret will be associated with. 
This field is required.

.. _r_secret_key-1:
* :ref:`key <r_secret_key-1>` - A unique name to identify the secret.

## Import

.. Note:: Importing this resource is not currently supported.