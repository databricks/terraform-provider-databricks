package docs

import "github.com/databricks/terraform-provider-databricks/common"

// ProviderRegistryUrl is the URL of the registry page for the Terraform provider for Databricks.
func ProviderRegistryUrl() string {
	return "https://registry.terraform.io/providers/databricks/databricks/" + common.Version() + "/"
}

// ProviderDocumentationRootUrl is the URL of the documentation for the Terraform provider for Databricks.
func ProviderDocumentationRootUrl() string {
	return ProviderRegistryUrl() + "docs/"
}

// Section is the section of the documentation for the provider in the Terraform registry.
type Section int

const (
	// Resources is the section containing documentation for an individual resource.
	Resources Section = iota

	// DataSources is the section containing documentation for an individual data source.
	DataSources

	// Guides is the section containing guides or walkthroughs for customers to get started using the provider.
	Guides
)

// sectionFragment returns the URI segment for the given section.
func (s Section) sectionFragment() string {
	switch s {
	case Resources:
		return "resources"
	case DataSources:
		return "data-sources"
	case Guides:
		return "guides"
	default:
		panic("section unspecified, must be one of: Resources, Guides")
	}
}

// DocOptions describes a single documentation reference for the provider in the Terraform registry. At a minimum,
// the Slug fields must be specified. If unspecified, the default section is Resources. If desired, a fragment can
// be provided to link to a specific anchor in the documentation page.
type DocOptions struct {
	// Section is the section of the doc that should be linked to.
	Section Section

	// Slug is the URI slug for the page. This should be the name of the markdown file containing the documentation
	// for the resource, data source, or guide, not including the `.md` extension.
	Slug string

	// Fragment is an optional fragment to link to in the documentation page. If unspecified, the link will
	// refer to the page itself.
	Fragment string
}

// DocumentationUrl is the URL for a specific guide in the provider documentation. See the documentation for
// DocOptions for more information.
func DocumentationUrl(opts DocOptions) string {
	url := ProviderDocumentationRootUrl() + opts.Section.sectionFragment() + "/" + opts.Slug
	if opts.Fragment != "" {
		url += "#" + opts.Fragment
	}
	return url
}
