package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestConnectionReadInlineSecretReferences(t *testing.T) {
	const oldReference = "{{uc_secrets/main.default.old_password}}"
	const newReference = "{{uc_secrets/main.default.new_password}}"
	testCases := []struct {
		name     string
		previous map[string]any
		returned map[string]string
		expected map[string]any
	}{
		{
			name:     "returned reference replaces a stored literal",
			previous: map[string]any{"password": "literal"},
			returned: map[string]string{"password": newReference},
			expected: map[string]any{"password": newReference},
		},
		{
			name:     "remote reference change is retained",
			previous: map[string]any{"password": oldReference},
			returned: map[string]string{"password": newReference},
			expected: map[string]any{"password": newReference},
		},
		{
			name:     "omitted literal credential is preserved",
			previous: map[string]any{"password": "literal"},
			returned: nil,
			expected: map[string]any{"password": "literal"},
		},
		{
			name:     "returned literal replaces a stored reference",
			previous: map[string]any{"password": oldReference},
			returned: map[string]string{"password": "returned-literal"},
			expected: map[string]any{"password": "returned-literal"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			qa.ResourceFixture{
				MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
					w.GetMockConnectionsAPI().EXPECT().GetByName(mock.Anything, "my_conn").Return(&catalog.ConnectionInfo{
						Name:           "my_conn",
						FullName:       "my_conn",
						ConnectionType: catalog.ConnectionTypeMysql,
						MetastoreId:    "abc",
						Options:        tc.returned,
					}, nil)
				},
				Resource: ResourceConnection(),
				Read:     true,
				ID:       "abc|my_conn",
				State: map[string]any{
					"name":            "my_conn",
					"connection_type": "MYSQL",
					"options":         tc.previous,
				},
			}.ApplyAndExpectData(t, map[string]any{"options": tc.expected})
		})
	}
}

func TestConnectionInlineSecretReferenceDiff(t *testing.T) {
	suppress := ResourceConnection().Schema["options"].DiffSuppressFunc
	testCases := []struct {
		name     string
		key      string
		old      string
		new      string
		expected bool
	}{
		{
			name:     "reference name case is ignored",
			key:      "options.password",
			old:      "{{uc_secrets/main.default.password}}",
			new:      "{{uc_secrets/MAIN.DEFAULT.Password}}",
			expected: true,
		},
		{
			name: "different reference is detected",
			key:  "options.password",
			old:  "{{uc_secrets/main.default.password}}",
			new:  "{{uc_secrets/main.default.other}}",
		},
		{
			name: "literal case remains significant",
			key:  "options.password",
			old:  "Password",
			new:  "password",
		},
		{
			name: "literal to reference remains significant",
			key:  "options.password",
			old:  "main.default.password",
			new:  "{{uc_secrets/main.default.password}}",
		},
		{
			name: "reference envelope is case sensitive",
			key:  "options.password",
			old:  "{{uc_secrets/main.default.password}}",
			new:  "{{UC_SECRETS/main.default.password}}",
		},
		{
			name: "embedded text remains significant",
			key:  "options.password",
			old:  "prefix{{uc_secrets/main.default.password}}",
			new:  "prefix{{uc_secrets/MAIN.DEFAULT.Password}}",
		},
		{
			name: "unknown map size is preserved",
			key:  "options.%",
			old:  "1",
			new:  unknownValue,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, suppress(tc.key, tc.old, tc.new, nil))
		})
	}
}
