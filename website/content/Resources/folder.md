+++
title = "folder"
date = 2020-04-20T23:34:03-04:00
weight = 15
chapter = false
pre = ""
+++

## Resource: `databricks_folder`

This resource allows you to manage the create, read, and delete folders in your Databricks workspace.

{{% notice note %}}
Creating this resource with a path will generate all the parent folders but they will not be deleted during a delete. 
Meaning that if you have a path `/parent_folder/sub_folder` parent_folder will be created automatically for you but will 
not be deleted when the resource is destroyed. After destroyed you will have residual folders in your workspace. To 
avoid this please make one level folders and make them dependent on each other using `depends_on` so that way 
the destroy will cleanly delete all the appropriate folders.
{{% /notice %}}

## Example Usage

```hcl
resource "databricks_folder" "folder_a" {
  path = "/my/folder/path"
  recursive_delete = "false"
}
```
    
## Argument Reference

The following arguments are supported:

#### - `path`:
> **(Required)** The absolute path of the directory, beginning with "/", e.g. "/myfolder". Changing this will force 
>the creation of a new folder resource

#### - `recursive_delete`:
> **(Required)** During the deletion of this resource this field will determine whether to delete the folder recursively.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

#### - `id`:
> The id is the path for the folder object in the workspace.

## Import

{{% notice note %}}
Importing this resource is not currently supported.
{{% /notice %}}
