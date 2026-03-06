package app

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGitSourceAttrTypes(t *testing.T) {
	attrTypes := gitSourceAttrTypes()

	assert.Equal(t, types.StringType, attrTypes["branch"])
	assert.Equal(t, types.StringType, attrTypes["tag"])
	assert.Equal(t, types.StringType, attrTypes["commit"])
	assert.Equal(t, types.StringType, attrTypes["source_code_path"])
	assert.Equal(t, types.StringType, attrTypes["resolved_commit"])

	gitRepoType, ok := attrTypes["git_repository"].(types.ObjectType)
	require.True(t, ok)
	assert.Equal(t, types.StringType, gitRepoType.AttrTypes["provider"])
	assert.Equal(t, types.StringType, gitRepoType.AttrTypes["url"])
}

func TestGitRepositoryAttrTypes(t *testing.T) {
	attrTypes := gitRepositoryAttrTypes()

	assert.Equal(t, types.StringType, attrTypes["provider"])
	assert.Equal(t, types.StringType, attrTypes["url"])
}

func TestAppDeploymentImportState_ValidID(t *testing.T) {
	r := &resourceAppDeployment{}
	ctx := context.Background()

	// Get the schema
	schemaResp := resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, &schemaResp)

	state := tfsdk.State{
		Schema: schemaResp.Schema,
	}

	resp := resource.ImportStateResponse{
		State: state,
	}
	r.ImportState(ctx, resource.ImportStateRequest{
		ID: "my-app/deploy-123",
	}, &resp)

	assert.False(t, resp.Diagnostics.HasError(), diagnosticMessages(resp.Diagnostics))

	var model appDeploymentModel
	diags := resp.State.Get(ctx, &model)
	assert.False(t, diags.HasError())
	assert.Equal(t, "my-app", model.AppName.ValueString())
	assert.Equal(t, "deploy-123", model.DeploymentId.ValueString())
	assert.True(t, model.Triggers.IsNull())
}

func TestAppDeploymentImportState_InvalidID(t *testing.T) {
	r := &resourceAppDeployment{}
	ctx := context.Background()

	schemaResp := resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, &schemaResp)

	state := tfsdk.State{
		Schema: schemaResp.Schema,
	}

	resp := resource.ImportStateResponse{
		State: state,
	}
	r.ImportState(ctx, resource.ImportStateRequest{
		ID: "invalid-no-slash",
	}, &resp)

	assert.True(t, resp.Diagnostics.HasError())
	assert.Contains(t, resp.Diagnostics.Errors()[0].Detail(), "expected format: app_name/deployment_id")
}

func TestAppDeploymentValidateConfig_SourceCodePath(t *testing.T) {
	r := &resourceAppDeployment{}
	ctx := context.Background()

	schemaResp := resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, &schemaResp)

	config := tfsdk.Config{
		Schema: schemaResp.Schema,
	}
	config.Raw = tftypes.NewValue(config.Schema.Type().TerraformType(ctx), map[string]tftypes.Value{
		"app_name":         tftypes.NewValue(tftypes.String, "my-app"),
		"source_code_path": tftypes.NewValue(tftypes.String, "/Workspace/apps/my-app"),
		"git_source":       tftypes.NewValue(schemaResp.Schema.Attributes["git_source"].GetType().TerraformType(ctx), nil),
		"mode":             tftypes.NewValue(tftypes.String, nil),
		"triggers":         tftypes.NewValue(tftypes.Map{ElementType: tftypes.String}, nil),
		"deployment_id":    tftypes.NewValue(tftypes.String, nil),
		"status":           tftypes.NewValue(tftypes.String, nil),
		"create_time":      tftypes.NewValue(tftypes.String, nil),
	})

	resp := resource.ValidateConfigResponse{}
	r.ValidateConfig(ctx, resource.ValidateConfigRequest{Config: config}, &resp)

	assert.False(t, resp.Diagnostics.HasError(), diagnosticMessages(resp.Diagnostics))
}

func TestAppDeploymentValidateConfig_NeitherSpecified(t *testing.T) {
	r := &resourceAppDeployment{}
	ctx := context.Background()

	schemaResp := resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, &schemaResp)

	config := tfsdk.Config{
		Schema: schemaResp.Schema,
	}
	config.Raw = tftypes.NewValue(config.Schema.Type().TerraformType(ctx), map[string]tftypes.Value{
		"app_name":         tftypes.NewValue(tftypes.String, "my-app"),
		"source_code_path": tftypes.NewValue(tftypes.String, nil),
		"git_source":       tftypes.NewValue(schemaResp.Schema.Attributes["git_source"].GetType().TerraformType(ctx), nil),
		"mode":             tftypes.NewValue(tftypes.String, nil),
		"triggers":         tftypes.NewValue(tftypes.Map{ElementType: tftypes.String}, nil),
		"deployment_id":    tftypes.NewValue(tftypes.String, nil),
		"status":           tftypes.NewValue(tftypes.String, nil),
		"create_time":      tftypes.NewValue(tftypes.String, nil),
	})

	resp := resource.ValidateConfigResponse{}
	r.ValidateConfig(ctx, resource.ValidateConfigRequest{Config: config}, &resp)

	assert.True(t, resp.Diagnostics.HasError())
	assert.Contains(t, resp.Diagnostics.Errors()[0].Detail(), "Exactly one of 'source_code_path' or 'git_source' must be specified")
}

func TestAppDeploymentValidateConfig_BothSpecified(t *testing.T) {
	r := &resourceAppDeployment{}
	ctx := context.Background()

	schemaResp := resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, &schemaResp)

	gitSourceType := schemaResp.Schema.Attributes["git_source"].GetType().TerraformType(ctx)
	gitRepoType := gitSourceType.(tftypes.Object).AttributeTypes["git_repository"]

	config := tfsdk.Config{
		Schema: schemaResp.Schema,
	}
	config.Raw = tftypes.NewValue(config.Schema.Type().TerraformType(ctx), map[string]tftypes.Value{
		"app_name":         tftypes.NewValue(tftypes.String, "my-app"),
		"source_code_path": tftypes.NewValue(tftypes.String, "/Workspace/apps/my-app"),
		"git_source": tftypes.NewValue(gitSourceType, map[string]tftypes.Value{
			"branch":           tftypes.NewValue(tftypes.String, "main"),
			"tag":              tftypes.NewValue(tftypes.String, nil),
			"commit":           tftypes.NewValue(tftypes.String, nil),
			"source_code_path": tftypes.NewValue(tftypes.String, nil),
			"git_repository":   tftypes.NewValue(gitRepoType, nil),
			"resolved_commit":  tftypes.NewValue(tftypes.String, nil),
		}),
		"mode":          tftypes.NewValue(tftypes.String, nil),
		"triggers":      tftypes.NewValue(tftypes.Map{ElementType: tftypes.String}, nil),
		"deployment_id": tftypes.NewValue(tftypes.String, nil),
		"status":        tftypes.NewValue(tftypes.String, nil),
		"create_time":   tftypes.NewValue(tftypes.String, nil),
	})

	resp := resource.ValidateConfigResponse{}
	r.ValidateConfig(ctx, resource.ValidateConfigRequest{Config: config}, &resp)

	assert.True(t, resp.Diagnostics.HasError())
	assert.Contains(t, resp.Diagnostics.Errors()[0].Detail(), "Only one of 'source_code_path' or 'git_source' can be specified")
}

func TestAppDeploymentUpdate_ReturnsError(t *testing.T) {
	r := &resourceAppDeployment{}
	resp := resource.UpdateResponse{}
	r.Update(context.Background(), resource.UpdateRequest{}, &resp)

	assert.True(t, resp.Diagnostics.HasError())
	assert.Contains(t, resp.Diagnostics.Errors()[0].Summary(), "unexpected update")
}

func TestAppDeploymentDelete_NoOp(t *testing.T) {
	r := &resourceAppDeployment{}
	resp := resource.DeleteResponse{}
	r.Delete(context.Background(), resource.DeleteRequest{}, &resp)

	assert.False(t, resp.Diagnostics.HasError())
}

func TestAppDeploymentMetadata(t *testing.T) {
	r := &resourceAppDeployment{}
	resp := resource.MetadataResponse{}
	r.Metadata(context.Background(), resource.MetadataRequest{
		ProviderTypeName: "databricks",
	}, &resp)

	assert.Equal(t, "databricks_app_deployment", resp.TypeName)
}

func diagnosticMessages(diags diag.Diagnostics) string {
	var msgs string
	for _, d := range diags {
		msgs += fmt.Sprintf("%s: %s\n", d.Summary(), d.Detail())
	}
	return msgs
}
