variable "tenant_id" {
  type = string
}

variable "subscription_id" {
  type = string
}

variable "client_id" {
  type = string
}

variable "client_secret" {
  type = string
}

variable "region" {
  type = string
  default = "uswest2"
}

variable "dbws_rg_name" {
  type = string
}

variable "managed_rg_name" {
  type = string
}

variable "dbws_name" {
  type = string
}

variable "breakglass_user" {
  type = string
}

variable "vnet_id" {
  type = string
}

variable "subnet_public" {
  type = string
}

variable "subnet_private" {
  type = string
}

variable "workspace_tags" {
  type = map
}
