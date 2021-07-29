provider "google" {
}

output "cloud_env" {
  value = "gcp-accounts"
}

output "databricks_host" {
  value = "https://accounts.gcp.databricks.com"
}

data "google_client_openid_userinfo" "me" {
}

output "test_prefix" {
  value = replace(split("@", data.google_client_openid_userinfo.me.email)[0], ".", "_")
}
