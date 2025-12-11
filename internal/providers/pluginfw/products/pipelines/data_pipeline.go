package pipelines

import (
	"context"
	"fmt"
	"reflect"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/pipelines_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const pipelineDataSource = "pipeline"

func DataSourcePipeline() datasource.DataSource {
	return &PipelineDataSource{}
}

var _ datasource.DataSourceWithConfigure = &PipelineDataSource{}

type PipelineDataSource struct {
	Client *common.DatabricksClient
}

type PipelineInfo struct {
	Id           types.String `tfsdk:"id"`
	PipelineInfo types.List   `tfsdk:"pipeline_info"`
	PipelineName types.String `tfsdk:"pipeline_name"`
	tfschema.Namespace
}

func (PipelineInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetOptional().SetComputed()
	attrs["pipeline_name"] = attrs["pipeline_name"].SetOptional().SetComputed()
	attrs["pipeline_info"] = attrs["pipeline_info"].SetOptional().SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	return attrs
}

func (PipelineInfo) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"pipeline_info":   reflect.TypeOf(pipelines_tf.PipelineStateInfo_SdkV2{}),
		"provider_config": reflect.TypeOf(tfschema.ProviderConfigData{}),
	}
}

func (d *PipelineDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksStagingName(pipelineDataSource)
}

func (d *PipelineDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, PipelineInfo{}, nil)
	resp.Schema = schema.Schema{
		Attributes: attrs,
		Blocks:     blocks,
	}
}

func (d *PipelineDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func (d *PipelineDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, pipelineDataSource)

	var pipelineInfo PipelineInfo
	resp.Diagnostics.Append(req.Config.Get(ctx, &pipelineInfo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	workspaceID, diags := tfschema.GetWorkspaceIDDataSource(ctx, pipelineInfo.ProviderConfig)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	w, diags := d.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, workspaceID)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	pipelineName := pipelineInfo.PipelineName.ValueString()
	pipelineId := pipelineInfo.Id.ValueString()
	pipeline, diag := d.getPipelineDetails(ctx, w, pipelineName, pipelineId)
	resp.Diagnostics.Append(diag...)
	if resp.Diagnostics.HasError() {
		return
	}

	var tfPipeline pipelines_tf.PipelineStateInfo_SdkV2
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, pipeline, &tfPipeline)...)
	if resp.Diagnostics.HasError() {
		return
	}

	pipelineInfo.Id = tfPipeline.PipelineId
	pipelineInfo.PipelineName = tfPipeline.Name
	pipelineInfo.PipelineInfo = types.ListValueMust(tfPipeline.Type(ctx), []attr.Value{tfPipeline.ToObjectValue(ctx)})
	resp.Diagnostics.Append(resp.State.Set(ctx, pipelineInfo)...)
}

func (d *PipelineDataSource) getPipelineDetails(ctx context.Context, w *databricks.WorkspaceClient, pipelineName, pipelineId string) (p pipelines.PipelineStateInfo, dd diag.Diagnostics) {
	if pipelineId == "" && pipelineName == "" {
		dd.AddError("either 'Id' or 'PipelineName' is required", "")
		return
	}

	if pipelineId == "" {
		pipelineSearch := pipelines.ListPipelinesRequest{Filter: fmt.Sprintf("name LIKE '%s'", pipelineName), MaxResults: 100}
		pipelineList, err := w.Pipelines.ListPipelinesAll(ctx, pipelineSearch)
		if err != nil {
			dd.AddError("failed to list pipelines", err.Error())
		}

		for _, pp := range pipelineList {
			if pp.Name == pipelineName {
				pipelineId = pp.PipelineId
				break
			}
		}
	}

	pipelineInfo, err := w.Pipelines.Get(ctx, pipelines.GetPipelineRequest{PipelineId: pipelineId})
	if err != nil {
		dd.AddError(fmt.Sprintf("failed to get pipeline with pipeline name '%s'", pipelineName), err.Error())
	}

	return pipelines.PipelineStateInfo{
		ClusterId:       pipelineInfo.ClusterId,
		CreatorUserName: pipelineInfo.CreatorUserName,
		Health:          pipelines.PipelineStateInfoHealth(pipelineInfo.Health),
		LatestUpdates:   pipelineInfo.LatestUpdates,
		Name:            pipelineInfo.Name,
		PipelineId:      pipelineInfo.PipelineId,
		RunAsUserName:   pipelineInfo.RunAsUserName,
		State:           pipelineInfo.State,
		ForceSendFields: pipelineInfo.ForceSendFields,
	}, dd
}
