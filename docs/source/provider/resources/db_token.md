# Resource: databricks_token

This resource creates an api token that can be used to create Databricks resources. 

.. WARNING:: This will create an API token for the user that has authenticated on the provider. So if you have used an 
    admin user to setup the provider then you will be making API tokens for that admin user. 

## Example Usage

.. code-block:: tf
    
    resource "databricks_token" "my-token" {
      lifetime_seconds = 6000
      comment = "Testing terraform v2"
    }

## Argument Reference

The following arguments are supported:

.. _r_token_lifetime_seconds:
* :ref:`lifetime_seconds <r_token_lifetime_seconds>`- **(Optional) (Numeric)** The lifetime of the token, in seconds. 
If no lifetime is specified, the token remains valid indefinitely.

.. _r_token_comment:
* :ref:`comment <r_token_comment>` - **(Optional)** Optional description to attach to the token.


## Attribute Reference
In addition to all arguments above, the following attributes are exported:

.. _r_token_id:
* :ref:`id <r_token_id>` - The id for the token object.

.. _r_token_lifetime_seconds-1:    
* :ref:`lifetime_seconds <r_token_lifetime_seconds-1>` - The lifetime of the token, in seconds. If no lifetime is specified, 
the token remains valid indefinitely.

.. _r_token_comment-1:
* :ref:`comment <r_token_comment-1>` - Optional description to attach to the token.

.. _r_token_creation_time:
* :ref:`creation_time <r_token_creation_time>` - Server time (in epoch milliseconds) when the token was created.

.. _r_token_token_value:
* :ref:`token_value <r_token_token_value>` - The value of the newly-created token.

.. _r_token_expiry_time:
* :ref:`expiry_time <r_token_expiry_time>` - Server time (in epoch milliseconds) when the token will expire, or -1 if not applicable.

## Import

.. Note:: Importing this resource is not currently supported.
  