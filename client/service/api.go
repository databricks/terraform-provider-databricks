package service

import (
	db "github.com/stikkireddy/databricks-tf-provider/client"
)

var scimHeaders = map[string]string{
	"Content-Type": "application/scim+json",
}

type DBApiClient struct {
	Option db.DBClientOption
}

// Init initializes the client
func (c *DBApiClient) Init(option db.DBClientOption) DBApiClient {
	c.Option = option
	option.Init()
	return *c
}

// Clusters returns an instance of ClustersAPI
func (c DBApiClient) Clusters() ClustersAPI {
	var clustersAPI ClustersAPI
	return clustersAPI.init(c)
}

func (c DBApiClient) Secrets() SecretsAPI {
	var secretsAPI SecretsAPI
	return secretsAPI.init(c)
}

func (c DBApiClient) SecretScopes() SecretScopesAPI {
	var secretScopesAPI SecretScopesAPI
	return secretScopesAPI.init(c)
}

func (c DBApiClient) SecretAcls() SecretAclsAPI {
	var secretAclsAPI SecretAclsAPI
	return secretAclsAPI.init(c)
}

func (c DBApiClient) Tokens() TokensAPI {
	var tokensAPI TokensAPI
	return tokensAPI.init(c)
}

func (c DBApiClient) Users() UsersAPI {
	var usersAPI UsersAPI
	return usersAPI.init(c)
}

func (c DBApiClient) Groups() GroupsAPI {
	var groupsAPI GroupsAPI
	return groupsAPI.init(c)
}

func (c DBApiClient) InstancePools() InstancePoolsAPI {
	var instancePoolsAPI InstancePoolsAPI
	return instancePoolsAPI.init(c)
}

func (c *DBApiClient) performQuery(method, path string, apiVersion string, headers map[string]string, data interface{}) ([]byte, error) {
	return db.PerformQuery(c.Option, method, path, "2.0", headers, true, false, data)
}
