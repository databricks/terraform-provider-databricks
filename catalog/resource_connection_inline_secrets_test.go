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
		name            string
		previous        map[string]any
		returned        map[string]string
		expected        map[string]any
		previousSecrets map[string]any
		returnedSecrets map[string]string
		expectedSecrets map[string]any
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
			name:     "omitted reference is removed",
			previous: map[string]any{"password": oldReference},
			returned: nil,
			expected: map[string]any{},
		},
		{
			name:     "omitted malformed reference is also removed",
			previous: map[string]any{"password": "{{uc_secrets/invalid-name}}"},
			returned: nil,
			expected: map[string]any{},
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
		{
			name:     "import retains a returned reference without prior options",
			previous: map[string]any{},
			returned: map[string]string{"password": newReference},
			expected: map[string]any{"password": newReference},
		},
		{
			name:            "legacy configuration retains the secret map with inline API responses",
			previous:        map[string]any{"host": "test.com"},
			previousSecrets: map[string]any{"password": "main.default.old_password"},
			returned:        map[string]string{"host": "test.com", "password": oldReference},
			expected:        map[string]any{"host": "test.com"},
			expectedSecrets: map[string]any{"password": "main.default.old_password"},
		},
		{
			name:            "legacy configuration detects a changed inline reference",
			previous:        map[string]any{"host": "test.com"},
			previousSecrets: map[string]any{"password": "main.default.old_password"},
			returned:        map[string]string{"host": "test.com", "password": newReference},
			expected:        map[string]any{"host": "test.com"},
			expectedSecrets: map[string]any{"password": "main.default.new_password"},
		},
		{
			name:            "missing legacy reference is removed",
			previous:        map[string]any{"host": "test.com"},
			previousSecrets: map[string]any{"password": "main.default.old_password"},
			returned:        map[string]string{"host": "test.com"},
			expected:        map[string]any{"host": "test.com"},
		},
		{
			name:            "returned literal remains in sensitive options for a legacy configuration",
			previous:        map[string]any{"host": "test.com"},
			previousSecrets: map[string]any{"password": "main.default.old_password"},
			returned:        map[string]string{"host": "test.com", "password": "returned-literal"},
			expected:        map[string]any{"host": "test.com", "password": "returned-literal"},
		},
		{
			name:            "legacy API response retains the secret map",
			previous:        map[string]any{"host": "test.com"},
			previousSecrets: map[string]any{"password": "main.default.old_password"},
			returned:        map[string]string{"host": "test.com"},
			returnedSecrets: map[string]string{"password": "main.default.new_password"},
			expected:        map[string]any{"host": "test.com"},
			expectedSecrets: map[string]any{"password": "main.default.new_password"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expectedSecrets == nil {
				tc.expectedSecrets = map[string]any{}
			}
			qa.ResourceFixture{
				MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
					w.GetMockConnectionsAPI().EXPECT().GetByName(mock.Anything, "my_conn").Return(&catalog.ConnectionInfo{
						Name:           "my_conn",
						FullName:       "my_conn",
						ConnectionType: catalog.ConnectionTypeMysql,
						MetastoreId:    "abc",
						Options:        tc.returned,
						Secrets:        tc.returnedSecrets,
					}, nil)
				},
				Resource: ResourceConnection(),
				Read:     true,
				ID:       "abc|my_conn",
				State: map[string]any{
					"name":            "my_conn",
					"connection_type": "MYSQL",
					"options":         tc.previous,
					"secrets":         tc.previousSecrets,
				},
			}.ApplyAndExpectData(t, map[string]any{"options": tc.expected, "secrets": tc.expectedSecrets})
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

func TestConnectionsCreateInlineSecretReference(t *testing.T) {
	const reference = "{{uc_secrets/main.default.password}}"
	options := map[string]string{"host": "test.com", "password": reference}
	connection := &catalog.ConnectionInfo{
		Name:           "my_conn",
		FullName:       "my_conn",
		ConnectionType: catalog.ConnectionTypeMysql,
		MetastoreId:    "abc",
		Options:        options,
	}
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockConnectionsAPI().EXPECT()
			e.Create(mock.Anything, catalog.CreateConnection{
				Name:           "my_conn",
				ConnectionType: catalog.ConnectionTypeMysql,
				Options:        options,
			}).Return(connection, nil)
			e.GetByName(mock.Anything, "my_conn").Return(connection, nil)
		},
		Resource: ResourceConnection(),
		Create:   true,
		State: map[string]any{
			"name":            "my_conn",
			"connection_type": "MYSQL",
			"options":         map[string]any{"host": "test.com", "password": reference},
		},
	}.ApplyAndExpectData(t, map[string]any{
		"options": map[string]any{"host": "test.com", "password": reference},
		"secrets": map[string]any{},
	})
}

func TestConnectionsUpdateInlineSecretReference(t *testing.T) {
	const reference = "{{uc_secrets/main.default.new_password}}"
	options := map[string]string{"host": "test.com", "password": reference}
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockConnectionsAPI().EXPECT()
			e.Update(mock.Anything, catalog.UpdateConnection{
				Name:    "my_conn",
				Options: options,
			}).Return(&catalog.ConnectionInfo{}, nil)
			e.GetByName(mock.Anything, "my_conn").Return(&catalog.ConnectionInfo{
				Name:           "my_conn",
				FullName:       "my_conn",
				ConnectionType: catalog.ConnectionTypeMysql,
				MetastoreId:    "abc",
				Options:        options,
			}, nil)
		},
		Resource: ResourceConnection(),
		Update:   true,
		ID:       "abc|my_conn",
		InstanceState: map[string]string{
			"name":             "my_conn",
			"connection_type":  "MYSQL",
			"metastore_id":     "abc",
			"options.%":        "2",
			"options.host":     "test.com",
			"options.password": "{{uc_secrets/main.default.old_password}}",
		},
		State: map[string]any{
			"name":            "my_conn",
			"connection_type": "MYSQL",
			"options":         map[string]any{"host": "test.com", "password": reference},
		},
	}.ApplyAndExpectData(t, map[string]any{
		"options": map[string]any{"host": "test.com", "password": reference},
		"secrets": map[string]any{},
	})
}
