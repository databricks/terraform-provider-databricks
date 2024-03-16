package acceptance

import (
	"os"
	"testing"
)

var onlineTableHcl = `
locals {
	STICKY_RANDOM = replace(var.STICKY_RANDOM, "/[^a-zA-Z0-9_]/", "_")
}

resource "databricks_catalog" "sandbox" {
	name    = "sandbox_${var.STICKY_RANDOM}"
	comment = "this catalog is managed by terraform"
	properties = {
	  purpose = "testing"
	}
	force_destroy = true
  }
  
  resource "databricks_schema" "things" {
	catalog_name = databricks_catalog.sandbox.id
	name         = "things_${local.STICKY_RANDOM}"
	comment      = "this database is managed by terraform"
	properties = {
	  kind = "various"
	}
	force_destroy = true
  }
  
  resource "databricks_sql_endpoint" "this" {
	name             = "tf-${local.STICKY_RANDOM}"
	cluster_size     = "2X-Small"
	max_num_clusters = 1
	warehouse_type   = "PRO"
	enable_serverless_compute = true
  }
  
  resource "databricks_sql_table" "table" {
	catalog_name       = databricks_catalog.sandbox.id
	schema_name        = databricks_schema.things.name
	name               = "ot_src_${local.STICKY_RANDOM}"
	table_type         = "MANAGED"
	data_source_format = "DELTA"
	warehouse_id       = databricks_sql_endpoint.this.id
  
	column {
	  name = "id"
	  type = "int"
	}
	column {
	  name = "timestamp"
	  type = "int"
	}
  
	properties = {
	  "delta.enableChangeDataFeed": true
	}
  }
  
  resource "databricks_online_table" "this" {
	name = "${databricks_sql_table.table.id}_ot"
	spec {
	  source_table_full_name = databricks_sql_table.table.id
	  primary_key_columns = [
		"id"
	  ]
	  run_triggered {
	  }
	}
  }
`

func TestUcAccOnlineTable(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		t.Skipf("databricks_online_table resource is not available on GCP")
	}
	unityWorkspaceLevel(t, step{Template: onlineTableHcl})
}
