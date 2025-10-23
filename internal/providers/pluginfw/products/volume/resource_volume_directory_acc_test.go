package volume_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

const volumeDirectorySetup = `
resource "databricks_catalog" "sandbox" {
	name         = "sandbox{var.STICKY_RANDOM}"
	comment      = "this catalog is managed by terraform"
	properties = {
		purpose = "testing"
	}
	force_destroy = true
}

resource "databricks_schema" "things" {
	catalog_name = databricks_catalog.sandbox.id
	name         = "things{var.STICKY_RANDOM}"
	comment      = "this schema is managed by terraform"
	properties = {
		kind = "various"
	}
}

resource "databricks_volume" "this" {
	name         = "volume{var.STICKY_RANDOM}"
	catalog_name = databricks_catalog.sandbox.name
	schema_name  = databricks_schema.things.name
	volume_type  = "MANAGED"
	comment      = "this volume is managed by terraform"
}
`

func TestUcAccVolumeDirectoryCreate(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: volumeDirectorySetup + `
resource "databricks_volume_directory" "test" {
	directory_path = "${databricks_volume.this.volume_path}/test_directory"
}
`,
	})
}

func TestUcAccVolumeDirectoryCreateNested(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: volumeDirectorySetup + `
resource "databricks_volume_directory" "parent" {
	directory_path = "${databricks_volume.this.volume_path}/parent"
}

resource "databricks_volume_directory" "child" {
	directory_path = "${databricks_volume_directory.parent.id}/child"
	depends_on = [databricks_volume_directory.parent]
}
`,
	})
}

func TestUcAccVolumeDirectoryMultiple(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: volumeDirectorySetup + `
resource "databricks_volume_directory" "dir1" {
	directory_path = "${databricks_volume.this.volume_path}/dir1"
}

resource "databricks_volume_directory" "dir2" {
	directory_path = "${databricks_volume.this.volume_path}/dir2"
}

resource "databricks_volume_directory" "dir3" {
	directory_path = "${databricks_volume.this.volume_path}/dir3"
}
`,
	})
}

func TestUcAccVolumeDirectoryUpdate(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t,
		acceptance.Step{
			Template: volumeDirectorySetup + `
resource "databricks_volume_directory" "test" {
	directory_path = "${databricks_volume.this.volume_path}/original_dir"
}
`,
		},
		acceptance.Step{
			Template: volumeDirectorySetup + `
resource "databricks_volume_directory" "test" {
	directory_path = "${databricks_volume.this.volume_path}/new_dir"
}
`,
		},
	)
}

func TestUcAccVolumeDirectoryDeepNesting(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: volumeDirectorySetup + `
resource "databricks_volume_directory" "deep" {
	directory_path = "${databricks_volume.this.volume_path}/level1/level2/level3/level4"
}
`,
	})
}
