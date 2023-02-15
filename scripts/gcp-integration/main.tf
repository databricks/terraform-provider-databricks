terraform {
  required_providers {
    databricks = {
      source = "databrickslabs/databricks"
    }
  }
}

locals {
  username   = replace(split("@", data.google_client_openid_userinfo.me.email)[0], ".", "_")
  prefix     = "${local.username}-${random_string.naming.result}"
  account_id = data.external.env.result.DATABRICKS_ACCOUNT_ID
}

data "external" "env" {
  program = ["python", "-c", "import sys,os,json;json.dump(dict(os.environ), sys.stdout)"]
}

resource "random_string" "naming" {
  special = false
  upper   = false
  length  = 6
}

// configured via env
provider "google" {}

// account_id & google_service_account configured via env
provider "databricks" {
  alias = "accounts"
  host  = "https://accounts.gcp.databricks.com"
}

data "google_client_config" "current" {}

resource "databricks_mws_workspaces" "this" {
  provider       = databricks.accounts
  account_id     = local.account_id
  workspace_name = local.prefix
  location       = data.google_client_config.current.region
  cloud_resource_container {
    gcp {
      project_id = data.google_client_config.current.project
    }
  }
  token {}
}

// google_service_account configured via env
provider "databricks" {
  alias = "workspace"
  host  = databricks_mws_workspaces.this.workspace_url
}

data "google_client_openid_userinfo" "me" {}

data "databricks_group" "admins" {
  depends_on   = [databricks_mws_workspaces.this]
  provider     = databricks.workspace
  display_name = "admins"
}

resource "databricks_user" "me" {
  depends_on = [databricks_mws_workspaces.this]
  provider   = databricks.workspace
  user_name  = data.google_client_openid_userinfo.me.email
}

resource "databricks_group_member" "allow_me_to_admin" {
  depends_on = [databricks_mws_workspaces.this]
  provider   = databricks.workspace
  group_id   = data.databricks_group.admins.id
  member_id  = databricks_user.me.id
}

output "databricks_host" {
  value = databricks_mws_workspaces.this.workspace_url
}

output "token" {
  value     = databricks_mws_workspaces.this.token[0].token_value
  sensitive = true
}

output "cloud_env" {
  value = "gcp"
}
