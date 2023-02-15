provider "databricks" {
    host = "https://adb-5524270988312315.15.azuredatabricks.net/"
    token = "dapi56e0ef3490d901e82cdd34f60e90a425-3"
}

terraform {
    required_providers {
        databricks = {
            source = "databricks/databricks"
            version = "1.9.2"
        }
    }
}

variable "username" {
    type = string
    default = "test1@example.com"
}

resource "databricks_user" "repos_user" {
  user_name = var.username
  delete_home_dir = true
  delete_repos = true
}
