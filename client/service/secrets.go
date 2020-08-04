package service

import (
	"fmt"
	"net/http"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// SecretsAPI exposes the Secrets API
type SecretsAPI struct {
	client *DatabricksClient
}

// Create creates or modifies a string secret depends on the type of scope backend
func (a SecretsAPI) Create(stringValue, scope, key string) error {
	return a.client.post("/secrets/put", model.SecretsRequest{
		StringValue: stringValue,
		Scope:       scope,
		Key:         key,
	}, nil)
}

// Delete deletes a secret depends on the type of scope backend
func (a SecretsAPI) Delete(scope, key string) error {
	return a.client.post("/secrets/delete", model.SecretsRequest{
		Scope: scope,
		Key:   key,
	}, nil)
}

// List lists the secret keys that are stored at this scope
func (a SecretsAPI) List(scope string) ([]model.SecretMetadata, error) {
	var secretsList model.SecretsList
	err := a.client.get("/secrets/list", map[string]string{
		"scope": scope,
	}, &secretsList)
	return secretsList.Secrets, err
}

// Read returns the metadata for the secret and not the contents of the secret
func (a SecretsAPI) Read(scope string, key string) (model.SecretMetadata, error) {
	var secretMeta model.SecretMetadata
	secrets, err := a.List(scope)
	if err != nil {
		return secretMeta, err
	}
	for _, secret := range secrets {
		if secret.Key == key {
			return secret, nil
		}
	}
	return secretMeta, APIError{
		ErrorCode:  "NOT_FOUND",
		Message:    fmt.Sprintf("no secret Scope found with secret metadata scope name: %s and key: %s", scope, key),
		Resource:   "/api/2.0/secrets/scopes/list",
		StatusCode: http.StatusNotFound,
	}
}
