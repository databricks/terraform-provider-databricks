package docs_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/docs"
	"github.com/stretchr/testify/assert"
)

func TestProviderRegistryUrl(t *testing.T) {
	assert.Equal(t, "https://registry.terraform.io/providers/databricks/databricks/"+common.Version()+"/", docs.ProviderRegistryUrl())
}

func TestProviderDocumentationRootUrl(t *testing.T) {
	assert.Equal(t, "https://registry.terraform.io/providers/databricks/databricks/"+common.Version()+"/docs/", docs.ProviderDocumentationRootUrl())
}

func TestResourceDocumentationUrl(t *testing.T) {
	assert.Equal(t, "https://registry.terraform.io/providers/databricks/databricks/"+common.Version()+"/docs/resources/my_resource", docs.ResourceDocumentationUrl("my_resource"))
}
