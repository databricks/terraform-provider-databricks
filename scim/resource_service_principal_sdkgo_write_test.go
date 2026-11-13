package scim

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	sdkcommon "github.com/databricks/databricks-sdk-go/common"
	sdkapierr "github.com/databricks/sdk-go/core/apierr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func servicePrincipalWithEveryField() User {
	return User{
		ApplicationID: "00000000-0000-4000-8000-000000000001",
		DisplayName:   "SPN",
		Active:        false,
		ExternalID:    "external",
		Schemas:       []URN{ServicePrincipalSchema},
		Entitlements:  entitlements{{Value: "allow-cluster-create"}},
		Groups:        []ComplexValue{{Value: "301", Display: "data-engineers", Ref: "Groups/301", Type: "direct"}},
		Roles:         []ComplexValue{{Value: "arn:aws:iam::123456789012:instance-profile/sp", Type: "direct"}},
	}
}

func jsonOf(t *testing.T, v any) map[string]any {
	t.Helper()
	data, err := json.Marshal(v)
	require.NoError(t, err)
	var out map[string]any
	require.NoError(t, json.Unmarshal(data, &out))
	return out
}

func TestCreateServicePrincipalRequestMatchesExistingBody(t *testing.T) {
	for name, u := range map[string]User{
		"every_field": servicePrincipalWithEveryField(),
		"no_entitlements": {
			DisplayName:  "SPN",
			Active:       true,
			Schemas:      []URN{ServicePrincipalSchema},
			Entitlements: entitlements{{}},
		},
	} {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, jsonOf(t, u), jsonOf(t, createServicePrincipalRequest(u)))
		})
	}
}

func TestUpdateServicePrincipalRequestAddsOnlyID(t *testing.T) {
	u := servicePrincipalWithEveryField()
	expected := jsonOf(t, u)
	expected["id"] = "201"
	assert.Equal(t, expected, jsonOf(t, updateServicePrincipalRequest("201", u)))
}

func TestAccountServicePrincipalRequestsDropFieldsOutsideTheAccountModel(t *testing.T) {
	u := servicePrincipalWithEveryField()
	expected := jsonOf(t, u)
	delete(expected, "schemas")
	delete(expected, "entitlements")
	delete(expected, "groups")
	assert.Equal(t, expected, jsonOf(t, createAccountServicePrincipalRequest(u)))
	expected["id"] = "201"
	assert.Equal(t, expected, jsonOf(t, updateAccountServicePrincipalRequest("201", u)))
}

func TestPatchServicePrincipalRequestsAddOnlyID(t *testing.T) {
	r := PatchRequestWithValue("replace", "active", "false")
	r.Operations = append(r.Operations, patchOperation{Op: "remove", Path: "roles"})
	expected := jsonOf(t, r)
	expected["id"] = "201"
	workspaceRequest, err := patchServicePrincipalRequest("201", r)
	require.NoError(t, err)
	assert.Equal(t, expected, jsonOf(t, workspaceRequest))
	accountRequest, err := patchAccountServicePrincipalRequest("201", r)
	require.NoError(t, err)
	assert.Equal(t, expected, jsonOf(t, accountRequest))
}

func classicAPIError(t *testing.T, body string) error {
	t.Helper()
	request := &http.Request{Method: http.MethodPost, URL: &url.URL{Path: "/scim"}}
	return apierr.GetAPIError(context.Background(), sdkcommon.ResponseWrapper{
		Response:   &http.Response{StatusCode: http.StatusConflict, Request: request},
		ReadCloser: io.NopCloser(bytes.NewReader([]byte(body))),
		DebugBytes: []byte(body),
	})
}

func sdkGoAPIError(body string) error {
	return sdkapierr.FromHTTPError(http.StatusConflict, http.Header{}, []byte(body))
}

func TestIsKnownScimError(t *testing.T) {
	known := []string{"Service principal with application ID abc already exists."}
	for name, tc := range map[string]struct {
		body     string
		expected bool
	}{
		"known_detail":                {`{"detail": "` + known[0] + `", "status": "409"}`, true},
		"known_detail_with_scim_type": {`{"detail": "` + known[0] + `", "scimType": "uniqueness"}`, false},
		"other_detail":                {`{"detail": "Service principal is locked.", "status": "409"}`, false},
		"known_message":               {`{"message": "` + known[0] + `"}`, true},
		"message_before_detail":       {`{"message": "Conflict.", "detail": "` + known[0] + `"}`, false},
		"known_api_1_2_error":         {`{"error": "` + known[0] + `", "message": "Conflict."}`, true},
		"not_json":                    {`<html>` + known[0] + `</html>`, false},
	} {
		t.Run(name, func(t *testing.T) {
			classic := classicAPIError(t, tc.body)
			assert.Equal(t, tc.expected, isKnownScimError(classic, known), "classic error %q", classic)
			sdkGo := fmt.Errorf("cannot create: %w", sdkGoAPIError(tc.body))
			assert.Equal(t, tc.expected, isKnownScimError(sdkGo, known), "sdk-go error %q", sdkGo)
		})
	}
	assert.True(t, isKnownScimError(fmt.Errorf("%s", known[0]), known))
	assert.False(t, isKnownScimError(fmt.Errorf("uniqueness %s", known[0]), known))
}
