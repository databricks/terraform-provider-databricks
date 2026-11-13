package scim

import (
	"testing"

	sdkscim "github.com/databricks/sdk-go/scim/v1"
	"github.com/stretchr/testify/assert"
)

func ptr[T any](v T) *T {
	return &v
}

const testInstanceProfile = "arn:aws:iam::123456789012:instance-profile/sp"

func TestUserFromSdkGoServicePrincipal(t *testing.T) {
	sp := &sdkscim.ServicePrincipal{
		Id:            ptr("201"),
		ApplicationId: ptr("app-201"),
		DisplayName:   ptr("SPN"),
		Active:        ptr(true),
		ExternalId:    ptr("external-201"),
		Entitlements:  []sdkscim.ComplexValue{{Value: ptr("allow-cluster-create")}},
		Groups: []sdkscim.ComplexValue{{
			Display: ptr("data-engineers"), Value: ptr("301"), Ref: ptr("Groups/301"),
			Type: ptr("direct"), Primary: ptr(true),
		}},
		Roles: []sdkscim.ComplexValue{{Value: ptr(testInstanceProfile)}},
		Schemas: []sdkscim.ServicePrincipalSchema{
			sdkscim.ServicePrincipalSchema_UrnIetfParamsScimSchemasCore20ServicePrincipal,
		},
	}
	assert.Equal(t, User{
		ID:            "201",
		ApplicationID: "app-201",
		DisplayName:   "SPN",
		Active:        true,
		ExternalID:    "external-201",
		Entitlements:  entitlements{{Value: "allow-cluster-create"}},
		Groups: []ComplexValue{{
			Display: "data-engineers", Value: "301", Ref: "Groups/301", Type: "direct",
		}},
		Roles:   []ComplexValue{{Value: testInstanceProfile}},
		Schemas: []URN{ServicePrincipalSchema},
	}, userFromSdkGoServicePrincipal(sp))
}

func TestUserFromSdkGoServicePrincipalDefaultsMissingFields(t *testing.T) {
	assert.Equal(t, User{}, userFromSdkGoServicePrincipal(&sdkscim.ServicePrincipal{}))
}

func TestUserFromSdkGoAccountServicePrincipal(t *testing.T) {
	sp := &sdkscim.AccountServicePrincipal{
		Id:            ptr("201"),
		ApplicationId: ptr("app-201"),
		DisplayName:   ptr("SPN"),
		Active:        ptr(false),
		ExternalId:    ptr("external-201"),
		Roles: []sdkscim.AccountComplexValue{{
			Value: ptr("account_admin"), Type: ptr("direct"),
		}},
		AccountId: ptr("account"),
	}
	assert.Equal(t, User{
		ID:            "201",
		ApplicationID: "app-201",
		DisplayName:   "SPN",
		ExternalID:    "external-201",
		Roles:         []ComplexValue{{Value: "account_admin", Type: "direct"}},
	}, userFromSdkGoAccountServicePrincipal(sp))
}

func TestComplexValuesFromSdkGoKeepsEmptyLists(t *testing.T) {
	assert.Nil(t, complexValuesFromSdkGo([]sdkscim.ComplexValue(nil)))
	assert.Equal(t, []ComplexValue{}, complexValuesFromSdkGo([]sdkscim.ComplexValue{}))
}
