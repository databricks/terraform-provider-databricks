locals {
  azureit_go  = "${path.module}/../../internal/azureit/azureit.go"
  azureit_sha = sha1(filebase64sha256(local.azureit_go))
  build       = "${path.module}/.terraform"
  target      = "${local.build}/azureit"
}

resource "azurerm_virtual_network" "this" {
  name                = "${local.prefix}-vnet"
  location            = azurerm_resource_group.this.location
  resource_group_name = azurerm_resource_group.this.name
  tags                = azurerm_resource_group.this.tags
  address_space       = ["10.1.0.0/16"]
}

resource "azurerm_subnet" "this" {
  name                 = "${local.prefix}-subnet"
  resource_group_name  = azurerm_resource_group.this.name
  virtual_network_name = azurerm_virtual_network.this.name
  address_prefixes     = ["10.1.0.0/24"]

  delegation {
    name = "delegation"

    service_delegation {
      name    = "Microsoft.ContainerInstance/containerGroups"
      actions = ["Microsoft.Network/virtualNetworks/subnets/action"]
    }
  }
}

resource "azurerm_network_profile" "this" {
  name                = "${local.prefix}-networkprofile"
  location            = azurerm_resource_group.this.location
  resource_group_name = azurerm_resource_group.this.name
  tags                = azurerm_resource_group.this.tags

  container_network_interface {
    name = "beep"
    ip_configuration {
      name      = "beep"
      subnet_id = azurerm_subnet.this.id
    }
  }
}

resource "azurerm_storage_container" "abfss" {
  name                  = "${local.prefix}-abfss"
  storage_account_name  = azurerm_storage_account.this.name
  container_access_type = "private"
}

resource "azurerm_storage_blob" "example" {
  name                   = "main.tf"
  storage_account_name   = azurerm_storage_account.this.name
  storage_container_name = azurerm_storage_container.abfss.name
  type                   = "Block"
  source                 = "${path.module}/main.tf"
}

resource "azurerm_storage_container" "wasbs" {
  name                  = "${local.prefix}-wasbs"
  storage_account_name  = azurerm_storage_account.this.name
  container_access_type = "private"
}

data "azurerm_storage_account_blob_container_sas" "this" {
  connection_string = azurerm_storage_account.this.primary_connection_string
  container_name    = azurerm_storage_container.wasbs.name
  https_only        = true
  start             = "2020-02-01"
  expiry            = "2021-12-31"
  permissions {
    read   = true
    add    = true
    create = true
    write  = true
    delete = true
    list   = true
  }
}

resource "null_resource" "azureit_binary" {
  triggers = {
    file_changed = local.azureit_sha
  }
  provisioner "local-exec" {
    command = "GOOS=linux go build -ldflags '-s -w' -o ${local.target}/azureit ${local.azureit_go}"
  }
}

resource "local_file" "function" {
  filename = "${local.target}/TriggerStart/function.json"
  content = jsonencode({
    "bindings" : [
      # {
      #   "type" : "httpTrigger",
      #   "name" : "req",
      #   "direction" : "in",
      #   "methods" : ["post"]
      # },
      # {
      #   "type" : "http",
      #   "name" : "res",
      #   "direction" : "out"
      # },
      {
        "name" : "nightly",
        "type" : "timerTrigger",
        "direction" : "in",
        "schedule" : "35 11-17 * * 1-5"
      }
    ]
  })
}

resource "local_file" "host" {
  filename = "${local.target}/host.json"
  content = jsonencode({
    "version" : "2.0",
    "customHandler" : {
      "description" : {
        "defaultExecutablePath" : "azureit"
      },
      "enableForwardingHttpRequest" : true
    },
    "extensionBundle" : {
      "id" : "Microsoft.Azure.Functions.ExtensionBundle",
      "version" : "[2.*, 3.0.0)"
    },
    "functionTimeout" : "00:00:10",
    "healthMonitor" : {
      "enabled" : false
    },
    "logging" : {
      "logLevel" : {
        "default" : "Information"
      }
    }
  })
}

data "archive_file" "this" {
  depends_on  = [null_resource.azureit_binary, local_file.host, local_file.function]
  type        = "zip"
  source_dir  = local.target
  output_path = "${local.build}/azureit.zip"
}

resource "azurerm_storage_container" "azureit" {
  storage_account_name  = azurerm_storage_account.this.name
  name                  = "azureit"
  container_access_type = "private"
}

resource "azurerm_storage_blob" "azureit" {
  type                   = "Block"
  name                   = "azureit-${data.archive_file.this.output_sha}.zip"
  storage_account_name   = azurerm_storage_account.this.name
  storage_container_name = azurerm_storage_container.azureit.name
  source                 = data.archive_file.this.output_path
}

resource "azurerm_app_service_plan" "azureit" {
  name                = "${local.prefix}-splan"
  resource_group_name = azurerm_resource_group.this.name
  location            = azurerm_resource_group.this.location
  tags                = azurerm_resource_group.this.tags
  kind                = "FunctionApp"
  // If Linux app service plan true, false otherwise.
  // https://docs.microsoft.com/en-us/rest/api/appservice/app-service-plans/create-or-update#request-body
  reserved = true
  sku {
    tier = "Dynamic"
    size = "Y1"
  }
}

data "azurerm_storage_account_blob_container_sas" "read" {
  connection_string = azurerm_storage_account.this.primary_connection_string
  container_name    = azurerm_storage_container.azureit.name

  start  = "2021-12-01T00:00:00Z"
  expiry = "2022-01-01T00:00:00Z" // TODO: make time-rotation

  permissions {
    read   = true
    add    = false
    create = false
    write  = false
    delete = false
    list   = false
  }
}

resource "azurerm_function_app" "azureit" {
  name                       = "${local.prefix}-azureit"
  resource_group_name        = azurerm_resource_group.this.name
  location                   = azurerm_resource_group.this.location
  tags                       = azurerm_resource_group.this.tags
  storage_account_name       = azurerm_storage_account.this.name
  storage_account_access_key = azurerm_storage_account.this.primary_access_key
  app_service_plan_id        = azurerm_app_service_plan.azureit.id
  version                    = "~4"
  os_type                    = "linux"

  identity {
    type         = "UserAssigned"
    identity_ids = [azurerm_user_assigned_identity.this.id]
  }

  app_settings = {
    AZURE_CLIENT_ID     = azurerm_user_assigned_identity.this.client_id
    ACI_CONTAINER_GROUP = azurerm_container_group.this.id

    // TODO: try getting this to work with MSI instead of SAS
    WEBSITE_RUN_FROM_PACKAGE    = "${azurerm_storage_blob.azureit.id}${data.azurerm_storage_account_blob_container_sas.read.sas}"
    FUNCTIONS_WORKER_RUNTIME    = "custom"
    AzureWebJobsDisableHomepage = "true"
  }

  site_config {
    ftps_state                  = "Disabled"
    use_32_bit_worker_process   = false
    scm_use_main_ip_restriction = true
    ip_restriction {
      action                    = "Allow"
      virtual_network_subnet_id = azurerm_subnet.this.id
    }
  }
}

resource "azurerm_container_group" "this" {
  name                = "${local.prefix}-run"
  location            = azurerm_resource_group.this.location
  resource_group_name = azurerm_resource_group.this.name
  tags                = azurerm_resource_group.this.tags

  os_type            = "Linux"
  restart_policy     = "Never"
  ip_address_type    = "Private"
  network_profile_id = azurerm_network_profile.this.id

  identity {
    type         = "UserAssigned"
    identity_ids = [azurerm_user_assigned_identity.this.id]
  }

  container {
    name   = "acceptance"
    image  = "ghcr.io/databrickslabs/terraform-provider-it:master"
    cpu    = "2"
    memory = "2"
    environment_variables = {
      TEST_FILTER                  = "TestAcc"
      CLOUD_ENV                    = "azure"
      ARM_USE_MSI                  = "true"
      DATABRICKS_AZURE_RESOURCE_ID = azurerm_databricks_workspace.this.id
      TEST_STORAGE_V2_ABFSS        = azurerm_storage_container.abfss.name
      TEST_STORAGE_V2_ACCOUNT      = azurerm_storage_account.this.name
      TEST_STORAGE_V2_WASBS        = azurerm_storage_container.wasbs.name
    }

    secure_environment_variables = {
      TEST_STORAGE_V2_KEY       = azurerm_storage_account.this.primary_access_key
      TEST_STORAGE_V2_WASBS_SAS = data.azurerm_storage_account_blob_container_sas.this.sas
    }

    ports {
      port     = 443
      protocol = "TCP"
    }
  }
}
