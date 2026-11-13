package scim

import (
	"context"

	"github.com/databricks/sdk-go/options/call"
	sdkscim "github.com/databricks/sdk-go/scim/v1"

	"github.com/databricks/terraform-provider-databricks/common"
)

// readWithSdkGo reads a service principal through sdk-go scim/v1. Its Get methods can't
// limit attributes, so the API returns the full object.
func (a ServicePrincipalsAPI) readWithSdkGo(servicePrincipalID string) (User, error) {
	sc, err := a.client.SdkGoScimClient(a.context, a.ApiLevel)
	if err != nil {
		return User{}, err
	}
	if sc.IsAccount {
		get := func(ctx context.Context,
			opts ...call.Option) (*sdkscim.AccountServicePrincipal, error) {
			req := sdkscim.GetAccountServicePrincipalRequest{Id: &servicePrincipalID}
			return sc.GetAccountServicePrincipal(ctx, req, opts...)
		}
		sp, err := common.CallSdkGoScim(a.context, sc, get)
		if err != nil {
			return User{}, err
		}
		return userFromSdkGoAccountServicePrincipal(sp), nil
	}
	get := func(ctx context.Context,
		opts ...call.Option) (*sdkscim.ServicePrincipal, error) {
		req := sdkscim.GetServicePrincipalRequest{Id: &servicePrincipalID}
		return sc.GetServicePrincipal(ctx, req, opts...)
	}
	sp, err := common.CallSdkGoScim(a.context, sc, get)
	if err != nil {
		return User{}, err
	}
	return userFromSdkGoServicePrincipal(sp), nil
}

func userFromSdkGoServicePrincipal(sp *sdkscim.ServicePrincipal) User {
	u := User{
		ID:            deref(sp.Id),
		DisplayName:   deref(sp.DisplayName),
		Active:        deref(sp.Active),
		ApplicationID: deref(sp.ApplicationId),
		ExternalID:    deref(sp.ExternalId),
		Groups:        complexValuesFromSdkGo(sp.Groups),
		Roles:         complexValuesFromSdkGo(sp.Roles),
		Entitlements:  complexValuesFromSdkGo(sp.Entitlements),
	}
	for _, s := range sp.Schemas {
		u.Schemas = append(u.Schemas, URN(s))
	}
	return u
}

// userFromSdkGoAccountServicePrincipal converts an account-level service principal,
// whose API model has no entitlements, groups or schemas.
func userFromSdkGoAccountServicePrincipal(sp *sdkscim.AccountServicePrincipal) User {
	return User{
		ID:            deref(sp.Id),
		DisplayName:   deref(sp.DisplayName),
		Active:        deref(sp.Active),
		ApplicationID: deref(sp.ApplicationId),
		ExternalID:    deref(sp.ExternalId),
		Roles:         complexValuesFromSdkGo(sp.Roles),
	}
}

// complexValuesFromSdkGo drops Primary, which ComplexValue doesn't model.
func complexValuesFromSdkGo[V sdkscim.ComplexValue | sdkscim.AccountComplexValue](
	values []V) []ComplexValue {
	if values == nil {
		return nil
	}
	out := make([]ComplexValue, 0, len(values))
	for _, v := range values {
		cv := sdkscim.ComplexValue(v)
		out = append(out, ComplexValue{
			Value:   deref(cv.Value),
			Display: deref(cv.Display),
			Ref:     deref(cv.Ref),
			Type:    deref(cv.Type),
		})
	}
	return out
}

func deref[T any](p *T) T {
	var zero T
	if p == nil {
		return zero
	}
	return *p
}
