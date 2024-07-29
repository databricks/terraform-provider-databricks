package catalog

import (
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestArtifactAllowlistCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceArtifactAllowlist(), qa.CornerCaseID("a|b"))
}

var id = "abc|INIT_SCRIPT"

var setArtifact = catalog.SetArtifactAllowlist{
	ArtifactType: catalog.ArtifactTypeInitScript,
	ArtifactMatchers: []catalog.ArtifactMatcher{
		{
			Artifact:  "/Volumes/inits",
			MatchType: catalog.MatchTypePrefixMatch,
		},
	},
}

var updateArtifact = catalog.SetArtifactAllowlist{
	ArtifactType: catalog.ArtifactTypeInitScript,
	ArtifactMatchers: []catalog.ArtifactMatcher{
		{
			Artifact:  "/Volumes/inits",
			MatchType: catalog.MatchTypePrefixMatch,
		},
		{
			Artifact:  "/Volumes/new_inits",
			MatchType: catalog.MatchTypePrefixMatch,
		},
	},
}

var artifactInfo = catalog.ArtifactAllowlistInfo{
	ArtifactMatchers: []catalog.ArtifactMatcher{
		{
			Artifact:  "/Volumes/inits",
			MatchType: catalog.MatchTypePrefixMatch,
		},
	},
	CreatedAt:   12345,
	CreatedBy:   "admin",
	MetastoreId: "abc",
}

var updatedArtifactInfo = catalog.ArtifactAllowlistInfo{
	ArtifactMatchers: []catalog.ArtifactMatcher{
		{
			Artifact:  "/Volumes/inits",
			MatchType: catalog.MatchTypePrefixMatch,
		},
		{
			Artifact:  "/Volumes/new_inits",
			MatchType: catalog.MatchTypePrefixMatch,
		},
	},
	CreatedAt:   12345,
	CreatedBy:   "admin",
	MetastoreId: "abc",
}

func TestArtifactAllowlistCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          http.MethodPut,
				Resource:        "/api/2.1/unity-catalog/artifact-allowlists/INIT_SCRIPT",
				ExpectedRequest: setArtifact,
				Response:        artifactInfo,
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/artifact-allowlists/INIT_SCRIPT?",
				Response: artifactInfo,
			},
		},
		Resource: ResourceArtifactAllowlist(),
		Create:   true,
		HCL: `
		artifact_type = "INIT_SCRIPT"
		artifact_matcher {
			artifact = "/Volumes/inits"
			match_type = "PREFIX_MATCH"
		}
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, id, d.Id())
	assert.Equal(t, 12345, d.Get("created_at"))
}

func TestArtifactAllowlistCreate_Error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          http.MethodPut,
				Resource:        "/api/2.1/unity-catalog/artifact-allowlists/INIT_SCRIPT",
				ExpectedRequest: setArtifact,
				Response: apierr.APIError{
					ErrorCode: "SERVER_ERROR",
					Message:   "Something unexpected happened",
				},
				Status: 500,
			},
		},
		Resource: ResourceArtifactAllowlist(),
		Create:   true,
		HCL: `
		artifact_type = "INIT_SCRIPT"
		artifact_matcher {
			artifact = "/Volumes/inits"
			match_type = "PREFIX_MATCH"
		}
		`,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Something unexpected")
}

func TestArtifactAllowlistRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/artifact-allowlists/INIT_SCRIPT?",
				Response: artifactInfo,
			},
		},
		Resource: ResourceArtifactAllowlist(),
		Read:     true,
		ID:       id,
		HCL: `
		artifact_type = "INIT_SCRIPT"
		artifact_matcher {
			artifact = "/Volumes/inits"
			match_type = "PREFIX_MATCH"
		}
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "INIT_SCRIPT", d.Get("artifact_type"))
	assert.Equal(t, 1, d.Get("artifact_matcher.#"))
}

func TestResourceArtifactAllowlistRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/artifact-allowlists/INIT_SCRIPT?",
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceArtifactAllowlist(),
		Read:     true,
		ID:       id,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, id, d.Id(), "Id should not be empty for error reads")
}

func TestArtifactAllowlistUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          http.MethodPut,
				Resource:        "/api/2.1/unity-catalog/artifact-allowlists/INIT_SCRIPT",
				ExpectedRequest: updateArtifact,
				Response:        artifactInfo,
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/artifact-allowlists/INIT_SCRIPT?",
				Response: updatedArtifactInfo,
			},
		},
		Resource: ResourceArtifactAllowlist(),
		Update:   true,
		InstanceState: map[string]string{
			"artifact_type":                 "INIT_SCRIPT",
			"artifact_matcher.#":            "1",
			"artifact_matcher.0.artifact":   "/Volumes/new_inits",
			"artifact_matcher.0.match_type": "PREFIX_MATCH",
		},
		ID: id,
		HCL: `
		artifact_type = "INIT_SCRIPT"
		artifact_matcher {
			artifact = "/Volumes/new_inits"
			match_type = "PREFIX_MATCH"
		}
		artifact_matcher {
			artifact = "/Volumes/inits"
			match_type = "PREFIX_MATCH"
		}		
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "INIT_SCRIPT", d.Get("artifact_type"))
	assert.Equal(t, 2, d.Get("artifact_matcher.#"))
}

func TestArtifactAllowlistUpdate_Error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          http.MethodPut,
				Resource:        "/api/2.1/unity-catalog/artifact-allowlists/INIT_SCRIPT",
				ExpectedRequest: updateArtifact,
				Response: apierr.APIError{
					ErrorCode: "SERVER_ERROR",
					Message:   "Something unexpected happened",
				},
				Status: 500,
			},
		},
		Resource: ResourceArtifactAllowlist(),
		Update:   true,
		InstanceState: map[string]string{
			"artifact_type":                 "INIT_SCRIPT",
			"artifact_matcher.#":            "1",
			"artifact_matcher.0.artifact":   "/Volumes/new_inits",
			"artifact_matcher.0.match_type": "PREFIX_MATCH",
		},
		ID: id,
		HCL: `
		artifact_type = "INIT_SCRIPT"
		artifact_matcher {
			artifact = "/Volumes/new_inits"
			match_type = "PREFIX_MATCH"
		}
		artifact_matcher {
			artifact = "/Volumes/inits"
			match_type = "PREFIX_MATCH"
		}	
		`,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Something unexpected")
}

func TestArtifactAllowlistDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPut,
				Resource: "/api/2.1/unity-catalog/artifact-allowlists/INIT_SCRIPT",
				ExpectedRequest: catalog.SetArtifactAllowlist{
					ArtifactType:     catalog.ArtifactTypeInitScript,
					ArtifactMatchers: []catalog.ArtifactMatcher{},
				},
			},
		},
		Resource: ResourceArtifactAllowlist(),
		Delete:   true,
		ID:       id,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, id, d.Id())
}

func TestArtifactAllowlistDelete_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPut,
				Resource: "/api/2.1/unity-catalog/artifact-allowlists/INIT_SCRIPT",
				ExpectedRequest: catalog.SetArtifactAllowlist{
					ArtifactType:     catalog.ArtifactTypeInitScript,
					ArtifactMatchers: []catalog.ArtifactMatcher{},
				},
				Response: apierr.APIError{
					ErrorCode: "INVALID_STATE",
					Message:   "Something went wrong",
				},
				Status: 400,
			},
		},
		Resource: ResourceArtifactAllowlist(),
		Delete:   true,
		Removed:  true,
		ID:       id,
	}.ExpectError(t, "Something went wrong")
}
