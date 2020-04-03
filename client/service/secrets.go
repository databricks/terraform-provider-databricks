package service

import (
	"encoding/json"
	"errors"
	"github.com/stikkireddy/databricks-tf-provider/client/model"
	"net/http"
)

// SecretsAPI exposes the Secrets API
type SecretsAPI struct {
	Client DBApiClient
}

func (a SecretsAPI) init(client DBApiClient) SecretsAPI {
	a.Client = client
	return a
}

// PutSecretString creates or modifies a string secret depends on the type of scope backend
func (a SecretsAPI) Create(stringValue, scope, key string) error {
	data := struct {
		StringValue string `json:"string_value,omitempty"`
		Scope       string `json:"scope,omitempty"`
		Key         string `json:"key,omitempty"`
	}{
		stringValue,
		scope,
		key,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/secrets/put", "2.0", nil, data)
	return err
}

// DeleteSecret deletes a secret depends on the type of scope backend
func (a SecretsAPI) Delete(scope, key string) error {
	data := struct {
		Scope string `json:"scope,omitempty"`
		Key   string `json:"key,omitempty"`
	}{
		scope,
		key,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/secrets/delete", "2.0", nil, data)
	return err
}

// ListSecrets lists the secret keys that are stored at this scope
func (a SecretsAPI) List(scope string) ([]model.SecretMetadata, error) {
	var secretsList struct {
		Secrets []model.SecretMetadata `json:"secrets,omitempty"`
	}

	data := struct {
		Scope string `json:"scope,omitempty" url:"scope,omitempty"`
	}{
		scope,
	}

	resp, err := a.Client.performQuery(http.MethodGet, "/secrets/list", "2.0", nil, data)
	if err != nil {
		return secretsList.Secrets, err
	}

	err = json.Unmarshal(resp, &secretsList)
	return secretsList.Secrets, err
}

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
	return secretMeta, errors.New("No Secret Scope found with secret metadata scope name " + scope + " and key " + key)
}
