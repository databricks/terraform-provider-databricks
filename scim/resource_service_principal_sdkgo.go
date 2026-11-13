package scim

import (
	"context"
	"encoding/json"

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
		sp, err := common.CallSdkGoScim(a.context, sc,
			func(ctx context.Context, opts ...call.Option) (*sdkscim.AccountServicePrincipal, error) {
				req := sdkscim.GetAccountServicePrincipalRequest{Id: &servicePrincipalID}
				return sc.GetAccountServicePrincipal(ctx, req, opts...)
			})
		if err != nil {
			return User{}, err
		}
		return userFromSdkGoAccountServicePrincipal(sp), nil
	}
	sp, err := common.CallSdkGoScim(a.context, sc,
		func(ctx context.Context, opts ...call.Option) (*sdkscim.ServicePrincipal, error) {
			req := sdkscim.GetServicePrincipalRequest{Id: &servicePrincipalID}
			return sc.GetServicePrincipal(ctx, req, opts...)
		})
	if err != nil {
		return User{}, err
	}
	return userFromSdkGoServicePrincipal(sp), nil
}

func (a ServicePrincipalsAPI) createWithSdkGo(u User) (User, error) {
	sc, err := a.client.SdkGoScimClient(a.context, a.ApiLevel)
	if err != nil {
		return User{}, err
	}
	if sc.IsAccount {
		sp, err := common.CallSdkGoScim(a.context, sc,
			func(ctx context.Context, opts ...call.Option) (*sdkscim.AccountServicePrincipal, error) {
				return sc.CreateAccountServicePrincipal(ctx, createAccountServicePrincipalRequest(u), opts...)
			})
		if err != nil {
			return User{}, err
		}
		return userFromSdkGoAccountServicePrincipal(sp), nil
	}
	sp, err := common.CallSdkGoScim(a.context, sc,
		func(ctx context.Context, opts ...call.Option) (*sdkscim.ServicePrincipal, error) {
			return sc.CreateServicePrincipal(ctx, createServicePrincipalRequest(u), opts...)
		})
	if err != nil {
		return User{}, err
	}
	return userFromSdkGoServicePrincipal(sp), nil
}

// filterWithSdkGo returns the first page of matches, like the existing client.
func (a ServicePrincipalsAPI) filterWithSdkGo(filter string, excludeRoles bool) ([]User, error) {
	sc, err := a.client.SdkGoScimClient(a.context, a.ApiLevel)
	if err != nil {
		return nil, err
	}
	var filterParam, excludedAttributes *string
	if filter != "" {
		filterParam = &filter
	}
	if excludeRoles {
		excludedAttributes = new("roles")
	}
	if sc.IsAccount {
		list, err := common.CallSdkGoScim(a.context, sc,
			func(ctx context.Context, opts ...call.Option) (*sdkscim.ListAccountServicePrincipalsResponse, error) {
				req := sdkscim.ListAccountServicePrincipalsRequest{
					Filter: filterParam, ExcludedAttributes: excludedAttributes,
				}
				return sc.ListAccountServicePrincipals(ctx, req, opts...)
			})
		if err != nil {
			return nil, err
		}
		return usersFromSdkGo(list.Resources, userFromSdkGoAccountServicePrincipal), nil
	}
	list, err := common.CallSdkGoScim(a.context, sc,
		func(ctx context.Context, opts ...call.Option) (*sdkscim.ListServicePrincipalResponse, error) {
			req := sdkscim.ListServicePrincipalsRequest{
				Filter: filterParam, ExcludedAttributes: excludedAttributes,
			}
			return sc.ListServicePrincipals(ctx, req, opts...)
		})
	if err != nil {
		return nil, err
	}
	return usersFromSdkGo(list.Resources, userFromSdkGoServicePrincipal), nil
}

func (a ServicePrincipalsAPI) updateWithSdkGo(servicePrincipalID string, u User) error {
	sc, err := a.client.SdkGoScimClient(a.context, a.ApiLevel)
	if err != nil {
		return err
	}
	if sc.IsAccount {
		req := updateAccountServicePrincipalRequest(servicePrincipalID, u)
		return callSdkGoScimWithoutResult(a.context, sc,
			func(ctx context.Context, opts ...call.Option) error {
				return sc.UpdateAccountServicePrincipal(ctx, req, opts...)
			})
	}
	req := updateServicePrincipalRequest(servicePrincipalID, u)
	return callSdkGoScimWithoutResult(a.context, sc,
		func(ctx context.Context, opts ...call.Option) error {
			return sc.UpdateServicePrincipal(ctx, req, opts...)
		})
}

func (a ServicePrincipalsAPI) patchWithSdkGo(servicePrincipalID string, r patchRequest) error {
	sc, err := a.client.SdkGoScimClient(a.context, a.ApiLevel)
	if err != nil {
		return err
	}
	if sc.IsAccount {
		req, err := patchAccountServicePrincipalRequest(servicePrincipalID, r)
		if err != nil {
			return err
		}
		return callSdkGoScimWithoutResult(a.context, sc,
			func(ctx context.Context, opts ...call.Option) error {
				return sc.PatchAccountServicePrincipal(ctx, req, opts...)
			})
	}
	req, err := patchServicePrincipalRequest(servicePrincipalID, r)
	if err != nil {
		return err
	}
	return callSdkGoScimWithoutResult(a.context, sc,
		func(ctx context.Context, opts ...call.Option) error {
			return sc.PatchServicePrincipal(ctx, req, opts...)
		})
}

func (a ServicePrincipalsAPI) deleteWithSdkGo(servicePrincipalID string) error {
	sc, err := a.client.SdkGoScimClient(a.context, a.ApiLevel)
	if err != nil {
		return err
	}
	if sc.IsAccount {
		req := sdkscim.DeleteAccountServicePrincipalRequest{Id: &servicePrincipalID}
		return callSdkGoScimWithoutResult(a.context, sc,
			func(ctx context.Context, opts ...call.Option) error {
				return sc.DeleteAccountServicePrincipal(ctx, req, opts...)
			})
	}
	req := sdkscim.DeleteServicePrincipalRequest{Id: &servicePrincipalID}
	return callSdkGoScimWithoutResult(a.context, sc,
		func(ctx context.Context, opts ...call.Option) error {
			return sc.DeleteServicePrincipal(ctx, req, opts...)
		})
}

func callSdkGoScimWithoutResult(ctx context.Context, sc *common.SdkGoScim,
	method func(context.Context, ...call.Option) error) error {
	_, err := common.CallSdkGoScim(ctx, sc,
		func(ctx context.Context, opts ...call.Option) (struct{}, error) {
			return struct{}{}, method(ctx, opts...)
		})
	return err
}

// createServicePrincipalRequest sends the fields the existing client sends: active
// always, and empty strings and lists omitted.
func createServicePrincipalRequest(u User) sdkscim.CreateServicePrincipalRequest {
	return sdkscim.CreateServicePrincipalRequest{
		Active:        &u.Active,
		ApplicationId: nonEmpty(u.ApplicationID),
		DisplayName:   nonEmpty(u.DisplayName),
		Entitlements:  complexValuesToSdkGo[sdkscim.ComplexValue](u.Entitlements),
		ExternalId:    nonEmpty(u.ExternalID),
		Groups:        complexValuesToSdkGo[sdkscim.ComplexValue](u.Groups),
		Roles:         complexValuesToSdkGo[sdkscim.ComplexValue](u.Roles),
		Schemas:       servicePrincipalSchemasToSdkGo(u.Schemas),
	}
}

// createAccountServicePrincipalRequest drops schemas, entitlements and groups, which
// the account-level API model doesn't have.
func createAccountServicePrincipalRequest(u User) sdkscim.CreateAccountServicePrincipalRequest {
	return sdkscim.CreateAccountServicePrincipalRequest{
		Active:        &u.Active,
		ApplicationId: nonEmpty(u.ApplicationID),
		DisplayName:   nonEmpty(u.DisplayName),
		ExternalId:    nonEmpty(u.ExternalID),
		Roles:         complexValuesToSdkGo[sdkscim.AccountComplexValue](u.Roles),
	}
}

// updateServicePrincipalRequest also sends the ID in the body, because sdk-go
// serializes its path field.
func updateServicePrincipalRequest(servicePrincipalID string, u User) sdkscim.UpdateServicePrincipalRequest {
	return sdkscim.UpdateServicePrincipalRequest{
		Id:            &servicePrincipalID,
		Active:        &u.Active,
		ApplicationId: nonEmpty(u.ApplicationID),
		DisplayName:   nonEmpty(u.DisplayName),
		Entitlements:  complexValuesToSdkGo[sdkscim.ComplexValue](u.Entitlements),
		ExternalId:    nonEmpty(u.ExternalID),
		Groups:        complexValuesToSdkGo[sdkscim.ComplexValue](u.Groups),
		Roles:         complexValuesToSdkGo[sdkscim.ComplexValue](u.Roles),
		Schemas:       servicePrincipalSchemasToSdkGo(u.Schemas),
	}
}

// updateAccountServicePrincipalRequest drops schemas, entitlements and groups, which
// the account-level API model doesn't have, and sends the ID in the body too.
func updateAccountServicePrincipalRequest(servicePrincipalID string,
	u User) sdkscim.UpdateAccountServicePrincipalRequest {
	return sdkscim.UpdateAccountServicePrincipalRequest{
		Id:            &servicePrincipalID,
		Active:        &u.Active,
		ApplicationId: nonEmpty(u.ApplicationID),
		DisplayName:   nonEmpty(u.DisplayName),
		ExternalId:    nonEmpty(u.ExternalID),
		Roles:         complexValuesToSdkGo[sdkscim.AccountComplexValue](u.Roles),
	}
}

// patchServicePrincipalRequest sends each operation's value as the existing client
// encodes it, and the ID in the body too.
func patchServicePrincipalRequest(servicePrincipalID string,
	r patchRequest) (sdkscim.PatchServicePrincipalRequest, error) {
	req := sdkscim.PatchServicePrincipalRequest{Id: &servicePrincipalID}
	for _, s := range r.Schemas {
		req.Schemas = append(req.Schemas, sdkscim.PatchSchema(s))
	}
	for _, o := range r.Operations {
		value, err := patchValue(o.Value)
		if err != nil {
			return req, err
		}
		req.Operations = append(req.Operations, sdkscim.Patch{
			Op: sdkscim.PatchOp(o.Op), Path: nonEmpty(o.Path), Value: value,
		})
	}
	return req, nil
}

func patchAccountServicePrincipalRequest(servicePrincipalID string,
	r patchRequest) (sdkscim.PatchAccountServicePrincipalRequest, error) {
	req := sdkscim.PatchAccountServicePrincipalRequest{Id: &servicePrincipalID}
	for _, s := range r.Schemas {
		req.Schemas = append(req.Schemas, sdkscim.AccountPatchSchema_PatchSchema(s))
	}
	for _, o := range r.Operations {
		value, err := patchValue(o.Value)
		if err != nil {
			return req, err
		}
		req.Operations = append(req.Operations, sdkscim.AccountPatch{
			Op: sdkscim.AccountPatchOp_PatchOp(o.Op), Path: nonEmpty(o.Path), Value: value,
		})
	}
	return req, nil
}

// patchValue leaves a nil value out of the operation, like the existing client.
func patchValue(value any) (json.RawMessage, error) {
	if value == nil {
		return nil, nil
	}
	return json.Marshal(value)
}

func servicePrincipalSchemasToSdkGo(schemas []URN) []sdkscim.ServicePrincipalSchema {
	var out []sdkscim.ServicePrincipalSchema
	for _, s := range schemas {
		out = append(out, sdkscim.ServicePrincipalSchema(s))
	}
	return out
}

func complexValuesToSdkGo[V sdkscim.ComplexValue | sdkscim.AccountComplexValue](values []ComplexValue) []V {
	if values == nil {
		return nil
	}
	out := make([]V, 0, len(values))
	for _, v := range values {
		out = append(out, V(sdkscim.ComplexValue{
			Value:   nonEmpty(v.Value),
			Display: nonEmpty(v.Display),
			Ref:     nonEmpty(v.Ref),
			Type:    nonEmpty(v.Type),
		}))
	}
	return out
}

// nonEmpty returns nil for an empty string, which the existing client omits.
func nonEmpty(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func usersFromSdkGo[S any](values []S, convert func(*S) User) []User {
	var users []User
	for i := range values {
		users = append(users, convert(&values[i]))
	}
	return users
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
func complexValuesFromSdkGo[V sdkscim.ComplexValue | sdkscim.AccountComplexValue](values []V) []ComplexValue {
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
