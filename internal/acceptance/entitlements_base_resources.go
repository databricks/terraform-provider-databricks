package acceptance

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/iam"
)

type entitlementResource interface {
	setDisplayName(string)
	setWorkspaceClient(*databricks.WorkspaceClient)
	create(context.Context) error
	cleanUp(context.Context) error
	dataSourceTemplate() string
	tfReference() string
}

type userResource struct {
	email string
	id    string
	w     *databricks.WorkspaceClient
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
	return fmt.Sprintf("data.databricks_user.example.id")
}

type groupResource struct {
	displayName string
	id          string
	w           *databricks.WorkspaceClient
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
	return fmt.Sprintf("data.databricks_group.example.id")
}

type servicePrincipalResource struct {
	applicationId string
	displayName   string
	w             *databricks.WorkspaceClient
}

func (s *servicePrincipalResource) setDisplayName(displayName string) {
	s.displayName = displayName
}

func (s *servicePrincipalResource) setWorkspaceClient(w *databricks.WorkspaceClient) {
	s.w = w
}

func (s *servicePrincipalResource) create(ctx context.Context) error {
	sp, err := s.w.ServicePrincipals.Create(ctx, iam.ServicePrincipal{
		ApplicationId: s.applicationId,
		DisplayName:   s.displayName,
	})
	if err != nil {
		return err
	}
	s.applicationId = sp.ApplicationId
	return nil
}

func (s *servicePrincipalResource) cleanUp(ctx context.Context) error {
	return s.w.ServicePrincipals.DeleteById(ctx, s.applicationId)
}

func (s *servicePrincipalResource) dataSourceTemplate() string {
	return fmt.Sprintf(`
		data "databricks_service_principal" "example" {
			display_name = "%s"
		}`, s.displayName)
}

func (s *servicePrincipalResource) tfReference() string {
	return fmt.Sprintf("data.databricks_service_principal.example.id")
}
