We happily welcome contributions to databricks-terraform. We use GitHub Issues to track community reported issues and GitHub Pull Requests for accepting changes.

## Unit testing resources

In order to unit test a resource, which runs fast and could be included in code coverage, one should use `ResourceTester`, that launches embedded HTTP server with `HTTPFixture`'s containing all calls that should have been made in given scenario:

```go
func TestPermissionsCreate(t *testing.T) {
	_, err := ResourceTester(t, []HTTPFixture{
		{
            Method:   http.MethodPatch,
            // requires full URI
            Resource: "/api/2.0/preview/permissions/clusters/abc",
            // works with entities, not JSON. Diff is displayed in case of missmatch 
			ExpectedRequest: model.AccessControlChangeList{
				AccessControlList: []*model.AccessControlChange{
					{
						UserName:        &TestingUser,
						PermissionLevel: "CAN_USE",
					},
				},
			},
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/preview/permissions/clusters/abc?",
			Response: model.AccessControlChangeList{
				AccessControlList: []*model.AccessControlChange{
					{
						UserName:        &TestingUser,
						PermissionLevel: "CAN_MANAGE",
					},
				},
			},
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/preview/scim/v2/Me?",
			Response: model.User{
				UserName: "chuck.norris",
			},
		},
    }, 
    // next argument is function, that creates resource (to make schema for ResourceData)
    resourcePermissions, 
    // state represented as native structure (though a bit clunky)
    map[string]interface{}{
		"cluster_id": "abc",
		"access_control": []interface{}{
			map[string]interface{}{
				"user_name":        TestingUser,
				"permission_level": "CAN_USE",
			},
		},
    },
    // the last argument is a function, that performs a stage on resource (Create/update/delete/read)
    resourcePermissionsCreate)
	assert.NoError(t, err, err)
}
```