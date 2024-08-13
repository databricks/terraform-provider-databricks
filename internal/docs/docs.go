package docs

import "github.com/databricks/terraform-provider-databricks/common"

func ProviderRegistryUrl() string {
	return "https://registry.terraform.io/providers/databricks/databricks/" + common.Version() + "/"
}

func ProviderDocumentationRootUrl() string {
	return ProviderRegistryUrl() + "docs/"
}

func ResourceDocumentationUrl(resource string) string {
	return ProviderDocumentationRootUrl() + "resources/" + resource
}
