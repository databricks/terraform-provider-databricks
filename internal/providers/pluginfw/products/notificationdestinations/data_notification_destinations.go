package notificationdestinations

import (
	"context"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/settings_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dataSourceName = "notification_destinations"

func DataSourceNotificationDestinations() datasource.DataSource {
	return &NotificationDestinationsDataSource{}
}

var _ datasource.DataSourceWithConfigure = &NotificationDestinationsDataSource{}

type NotificationDestinationsDataSource struct {
	Client *common.DatabricksClient
}

type NotificationDestinationsInfo struct {
	DisplayNameContains      types.String `tfsdk:"display_name_contains"`
	Type                     types.String `tfsdk:"type"`
	NotificationDestinations types.List   `tfsdk:"notification_destinations"`
	tfschema.Namespace
}

func (NotificationDestinationsInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["display_name_contains"] = attrs["display_name_contains"].SetOptional()
	attrs["type"] = attrs["type"].SetOptional()
	attrs["notification_destinations"] = attrs["notification_destinations"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	return attrs
}

func (NotificationDestinationsInfo) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"notification_destinations": reflect.TypeOf(settings_tf.ListNotificationDestinationsResult{}),
		"provider_config":           reflect.TypeOf(tfschema.ProviderConfigData{}),
	}
}

func (d *NotificationDestinationsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(dataSourceName)
}

func (d *NotificationDestinationsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, NotificationDestinationsInfo{}, nil)
	resp.Schema = schema.Schema{
		Attributes: attrs,
		Blocks:     blocks,
	}
}

func (d *NotificationDestinationsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func validateType(notificationType string) diag.Diagnostics {
	validTypes := []string{
		string(settings.DestinationTypeEmail),
		string(settings.DestinationTypeMicrosoftTeams),
		string(settings.DestinationTypePagerduty),
		string(settings.DestinationTypeSlack),
		string(settings.DestinationTypeWebhook),
	}

	if !slices.Contains(validTypes, notificationType) {
		return diag.Diagnostics{diag.NewErrorDiagnostic(fmt.Sprintf("Invalid type '%s'; valid types are %s.", notificationType, strings.Join(validTypes, ", ")), "")}
	}
	return nil
}

func AppendDiagAndCheckErrors(resp *datasource.ReadResponse, diags diag.Diagnostics) bool {
	resp.Diagnostics.Append(diags...)
	return resp.Diagnostics.HasError()
}

func (d *NotificationDestinationsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var notificationInfo NotificationDestinationsInfo
	if AppendDiagAndCheckErrors(resp, req.Config.Get(ctx, &notificationInfo)) {
		return
	}

	workspaceID, diags := tfschema.GetWorkspaceIDDataSource(ctx, notificationInfo.ProviderConfig)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	w, diags := d.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, workspaceID)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	notificationType := notificationInfo.Type.ValueString()
	notificationDisplayName := strings.ToLower(notificationInfo.DisplayNameContains.ValueString())

	if notificationType != "" && AppendDiagAndCheckErrors(resp, validateType(notificationType)) {
		return
	}

	notificationsGoSdk, err := w.NotificationDestinations.ListAll(ctx, settings.ListNotificationDestinationsRequest{})
	if err != nil {
		resp.Diagnostics.AddError("Failed to fetch Notification Destinations", err.Error())
		return
	}

	var notificationsTfSdk []attr.Value
	for _, notification := range notificationsGoSdk {
		if (notificationType != "" && notification.DestinationType.String() != notificationType) ||
			(notificationDisplayName != "" && !strings.Contains(strings.ToLower(notification.DisplayName), notificationDisplayName)) {
			continue
		}

		var notificationDestination settings_tf.ListNotificationDestinationsResult
		if AppendDiagAndCheckErrors(resp, converters.GoSdkToTfSdkStruct(ctx, notification, &notificationDestination)) {
			return
		}
		notificationsTfSdk = append(notificationsTfSdk, notificationDestination.ToObjectValue(ctx))
	}

	notificationInfo.NotificationDestinations = types.ListValueMust(settings_tf.ListNotificationDestinationsResult{}.Type(ctx), notificationsTfSdk)
	resp.Diagnostics.Append(resp.State.Set(ctx, notificationInfo)...)

}
