package docs_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/docs"
	"github.com/stretchr/testify/assert"
)

func expectedBaseUrl() string {
	return "https://registry.terraform.io/providers/databricks/databricks/" + common.Version() + "/"
}

func TestProviderRegistryUrl(t *testing.T) {
	assert.Equal(t, expectedBaseUrl(), docs.ProviderRegistryUrl())
}

func TestProviderDocumentationRootUrl(t *testing.T) {
	assert.Equal(t, expectedBaseUrl()+"docs/", docs.ProviderDocumentationRootUrl())
}

func TestDocumentationUrl(t *testing.T) {
	type testCase struct {
		name        string
		opts        docs.DocOptions
		expectedUrl string
	}
	cases := []testCase{
		{
			name: "default section",
			opts: docs.DocOptions{
				Slug: "resource",
			},
			expectedUrl: expectedBaseUrl() + "docs/resources/resource",
		},
		{
			name: "explicit resource section",
			opts: docs.DocOptions{
				Section: docs.Resources,
				Slug:    "resource",
			},
			expectedUrl: expectedBaseUrl() + "docs/resources/resource",
		},
		{
			name: "data source",
			opts: docs.DocOptions{
				Section: docs.DataSources,
				Slug:    "resource",
			},
			expectedUrl: expectedBaseUrl() + "docs/data-sources/resource",
		},
		{
			name: "guides",
			opts: docs.DocOptions{
				Section: docs.Guides,
				Slug:    "resource",
			},
			expectedUrl: expectedBaseUrl() + "docs/guides/resource",
		},
		{
			name: "guide with fragment",
			opts: docs.DocOptions{
				Section:  docs.Guides,
				Slug:     "resource",
				Fragment: "important-section",
			},
			expectedUrl: expectedBaseUrl() + "docs/guides/resource#important-section",
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			actualUrl := docs.DocumentationUrl(testCase.opts)
			if actualUrl != testCase.expectedUrl {
				t.Errorf("want %s, got %s", testCase.expectedUrl, actualUrl)
			}
		})
	}
}
