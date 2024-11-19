package notificationdestinations

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/settings_tf"
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
	DisplayNameContains      types.String                                     `tfsdk:"display_name_contains" tf:"optional"`
	Type                     types.String                                     `tfsdk:"type" tf:"optional"`
	NotificationDestinations []settings_tf.ListNotificationDestinationsResult `tfsdk:"notification_destinations" tf:"computed"`
}

func (d *NotificationDestinationsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(dataSourceName)
}

func (d *NotificationDestinationsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(NotificationDestinationsInfo{}, nil)
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
	w, diags := d.Client.GetWorkspaceClient()
	if AppendDiagAndCheckErrors(resp, diags) {
		return
	}

	var notificationInfo NotificationDestinationsInfo
	if AppendDiagAndCheckErrors(resp, req.Config.Get(ctx, &notificationInfo)) {
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

	var notificationsTfSdk []settings_tf.ListNotificationDestinationsResult
	for _, notification := range notificationsGoSdk {
		if (notificationType != "" && notification.DestinationType.String() != notificationType) ||
			(notificationDisplayName != "" && !strings.Contains(strings.ToLower(notification.DisplayName), notificationDisplayName)) {
			continue
		}

		var notificationDestination settings_tf.ListNotificationDestinationsResult
		if AppendDiagAndCheckErrors(resp, converters.GoSdkToTfSdkStruct(ctx, notification, &notificationDestination)) {
			return
		}
		notificationsTfSdk = append(notificationsTfSdk, notificationDestination)
	}

	notificationInfo.NotificationDestinations = notificationsTfSdk
	resp.Diagnostics.Append(resp.State.Set(ctx, notificationInfo)...)

}
