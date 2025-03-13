package scim_test

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/service/iam"
)

type entitlementResource interface {
	resourceType() string
	setDisplayName(string)
	setWorkspaceClient(*databricks.WorkspaceClient)
	create(context.Context) error
	getEntitlements(context.Context) ([]iam.ComplexValue, error)
	cleanUp(context.Context) error
	dataSourceTemplate() string
	tfReference() string
}

type userResource struct {
	email string
	id    string
	w     *databricks.WorkspaceClient
}

func (u *userResource) resourceType() string {
	return "user"
}

func (u *userResource) setDisplayName(displayName string) {
	u.email = displayName + "@example.com"
}

func (u *userResource) setWorkspaceClient(w *databricks.WorkspaceClient) {
	u.w = w
}

func (u *userResource) create(ctx context.Context) error {
	user, err := u.w.Users.Create(ctx, iam.User{
		UserName: u.email,
	})
	if err != nil {
		return err
	}
	u.id = user.Id
	return nil
}

func (u *userResource) getEntitlements(ctx context.Context) ([]iam.ComplexValue, error) {
	c, err := u.w.Config.NewApiClient()
	if err != nil {
		return nil, err
	}
	res := iam.User{}
	err = c.Do(ctx, "GET", fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%s?attributes=entitlements", u.id),
		httpclient.WithResponseUnmarshal(&res))
	if err != nil {
		return nil, err
	}
	return res.Entitlements, nil
}

func (u *userResource) cleanUp(ctx context.Context) error {
	return u.w.Users.DeleteById(ctx, u.id)
}

func (u *userResource) dataSourceTemplate() string {
	return fmt.Sprintf(`
		data "databricks_user" "example" {
			user_name = "%s"
		}`, u.email)
}

func (u *userResource) tfReference() string {
	return "user_id = data.databricks_user.example.id"
}

type groupResource struct {
	displayName string
	id          string
	w           *databricks.WorkspaceClient
}

func (g *groupResource) resourceType() string {
	return "group"
}

func (g *groupResource) setDisplayName(displayName string) {
	g.displayName = displayName
}

func (g *groupResource) setWorkspaceClient(w *databricks.WorkspaceClient) {
	g.w = w
}

func (g *groupResource) create(ctx context.Context) error {
	group, err := g.w.Groups.Create(ctx, iam.Group{
		DisplayName: g.displayName,
	})
	if err != nil {
		return err
	}
	g.id = group.Id
	return nil
}

func (g *groupResource) getEntitlements(ctx context.Context) ([]iam.ComplexValue, error) {
	c, err := g.w.Config.NewApiClient()
	if err != nil {
		return nil, err
	}
	res := iam.Group{}
	err = c.Do(ctx, "GET", fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%s?attributes=entitlements", g.id),
		httpclient.WithResponseUnmarshal(&res))
	if err != nil {
		return nil, err
	}
	return res.Entitlements, nil
}

func (g *groupResource) cleanUp(ctx context.Context) error {
	return g.w.Groups.DeleteById(ctx, g.id)
}

func (g *groupResource) dataSourceTemplate() string {
	return fmt.Sprintf(`
		data "databricks_group" "example" {
			display_name = "%s"
		}`, g.displayName)
}

func (g *groupResource) tfReference() string {
	return "group_id = data.databricks_group.example.id"
}

type servicePrincipalResource struct {
	applicationId string
	cleanup       bool
	displayName   string
	id            string
	w             *databricks.WorkspaceClient
}

func (s *servicePrincipalResource) resourceType() string {
	return "service_principal"
}

func (s *servicePrincipalResource) setDisplayName(displayName string) {
	s.displayName = displayName
}

func (s *servicePrincipalResource) setWorkspaceClient(w *databricks.WorkspaceClient) {
	s.w = w
}

func (s *servicePrincipalResource) create(ctx context.Context) error {
	sp, err := s.create0(ctx)
	if err != nil {
		return err
	}
	if s.applicationId == "" {
		s.applicationId = sp.ApplicationId
	}
	s.id = sp.Id
	return nil
}

func (s *servicePrincipalResource) create0(ctx context.Context) (iam.ServicePrincipal, error) {
	if s.applicationId != "" {
		sps := s.w.ServicePrincipals.List(ctx, iam.ListServicePrincipalsRequest{
			Filter: fmt.Sprintf(`applicationId eq "%s"`, s.applicationId),
		})
		if !sps.HasNext(ctx) {
			return iam.ServicePrincipal{}, fmt.Errorf("service principal with applicationId %s not found", s.applicationId)
		}
		return sps.Next(ctx)
	}
	sp, err := s.w.ServicePrincipals.Create(ctx, iam.ServicePrincipal{
		DisplayName: s.displayName,
	})
	return *sp, err
}

func (s *servicePrincipalResource) getEntitlements(ctx context.Context) ([]iam.ComplexValue, error) {
	c, err := s.w.Config.NewApiClient()
	if err != nil {
		return nil, err
	}
	res := iam.ServicePrincipal{}
	err = c.Do(ctx, "GET", fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%s?attributes=entitlements", s.id),
		httpclient.WithResponseUnmarshal(&res))
	if err != nil {
		return nil, err
	}
	return res.Entitlements, nil
}

func (s *servicePrincipalResource) cleanUp(ctx context.Context) error {
	if !s.cleanup {
		return nil
	}
	return s.w.ServicePrincipals.DeleteById(ctx, s.applicationId)
}

func (s *servicePrincipalResource) dataSourceTemplate() string {
	fragment := fmt.Sprintf(`display_name = "%s"`, s.displayName)
	if s.applicationId != "" {
		fragment = fmt.Sprintf(`application_id = "%s"`, s.applicationId)
	}
	return fmt.Sprintf(`
		data "databricks_service_principal" "example" {
			%s
		}`, fragment)
}

func (s *servicePrincipalResource) tfReference() string {
	return "service_principal_id = data.databricks_service_principal.example.id"
}
