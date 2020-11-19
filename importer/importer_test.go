package importer

import (
	"log"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestAccImporter(t *testing.T) {
	log.SetOutput(&levelWriter{"[INFO]", "[ERROR]", "[WARN]"})
	c := common.NewClientFromEnvironment()
	err := newImportContext(c).Run()
	assert.NoError(t, err)
}

func TestAccImportIdentity(t *testing.T) {
	err := Run("-directory", "/tmp/data-group", "-services", "groups,users")
	assert.NoError(t, err)
}

func TestAccImportSecrets(t *testing.T) {
	err := Run("-directory", "/tmp/data-group", "-services", "secrets,users,groups")
	assert.NoError(t, err)
}

func TestAccImportJobs(t *testing.T) {
	err := Run("-directory", "/tmp/data-group", "-services", "jobs")
	assert.NoError(t, err)
}

func TestResourceName(t *testing.T) {
	ic := newImportContext(&common.DatabricksClient{})
	norm := ic.ResourceName(&resource{
		Name: "9721431b_bcd3_4526_b90f_f5de2befec8c-dbutils_extensions_2_11_0_0_1-18dc8.jar",
	}, &schema.ResourceData{})
	assert.Equal(t, "dbutils_extensions_jar", norm)

	norm = ic.ResourceName(&resource{
		Name: "9721431b_bcd3_4526_b90f_f5de2befec8c|8737798193",
	}, &schema.ResourceData{})
	assert.Equal(t, "r7322b058678", norm)

	norm = ic.ResourceName(&resource{
		Name: "General Policy - All Users",
	}, &schema.ResourceData{})
	assert.Equal(t, "general_policy_all_users", norm)
}
