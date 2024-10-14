package notificationdestinations

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/settings_tf"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func DataSourceNotificationDestinations() datasource.DataSource {
	return &NotificationDestinationsDataSource{}
}

var _ datasource.DataSourceWithConfigure = &NotificationDestinationsDataSource{}

type NotificationDestinationsDataSource struct {
	Client *common.DatabricksClient
}

type NotificationDestinationsInfo struct {
	Id                       types.String                                     `tfsdk:"id" tf:"computed"`
	DisplayNameContains      types.String                                     `tfsdk:"display_name_contains" tf:"optional"`
	Type                     types.String                                     `tfsdk:"type" tf:"optional"`
	NotificationDestinations []settings_tf.ListNotificationDestinationsResult `tfsdk:"notification_destinations" tf:"computed"`
}

func (d *NotificationDestinationsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "databricks_notification_destinations_pluginframework"
}

func (d *NotificationDestinationsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: tfschema.DataSourceStructToSchemaMap(NotificationDestinationsInfo{}, nil),
	}
}

func (d *NotificationDestinationsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func validateType(notificationType string) diag.Diagnostics {
	validTypes := map[string]struct{}{
		string(settings.DestinationTypeEmail):          {},
		string(settings.DestinationTypeMicrosoftTeams): {},
		string(settings.DestinationTypePagerduty):      {},
		string(settings.DestinationTypeSlack):          {},
		string(settings.DestinationTypeWebhook):        {},
	}

	if _, valid := validTypes[notificationType]; !valid {
		return diag.Diagnostics{diag.NewErrorDiagnostic(fmt.Sprintf("Invalid type '%s'; valid types are EMAIL, MICROSOFT_TEAMS, PAGERDUTY, SLACK, WEBHOOK.", notificationType), "")}
	}
	return nil
}

func validateLength(destinations []settings_tf.ListNotificationDestinationsResult) diag.Diagnostics {
	if len(destinations) == 0 {
		return diag.Diagnostics{diag.NewErrorDiagnostic("Could not find any notification destinations with the specified criteria.", "")}
	}
	return nil
}

func appendToResponse(resp *datasource.ReadResponse, diags diag.Diagnostics) bool {
	resp.Diagnostics.Append(diags...)
	return resp.Diagnostics.HasError()
}

func (d *NotificationDestinationsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	w, diags := d.Client.GetWorkspaceClient()
	if appendToResponse(resp, diags) {
		return
	}

	var notificationInfo NotificationDestinationsInfo
	if appendToResponse(resp, req.Config.Get(ctx, &notificationInfo)) {
		return
	}

	notificationType := notificationInfo.Type.ValueString()
	notificationDisplayName := notificationInfo.DisplayNameContains.ValueString()

	if notificationType != "" {
		if appendToResponse(resp, validateType(notificationType)) {
			return
		}
	}

	notificationsGoSdk, err := w.NotificationDestinations.ListAll(ctx, settings.ListNotificationDestinationsRequest{})
	if err != nil {
		resp.Diagnostics.AddError("Failed to fetch Notification Destinations", err.Error())
		return
	}

	var notificationsTfSdk []settings_tf.ListNotificationDestinationsResult
	for _, notification := range notificationsGoSdk {
		if (notification.DestinationType.String() != notificationType) || (notificationDisplayName != "" && !strings.Contains(notification.DisplayName, notificationDisplayName)) {
			continue
		}

		var notificationDestination settings_tf.ListNotificationDestinationsResult
		if appendToResponse(resp, converters.GoSdkToTfSdkStruct(ctx, notification, &notificationDestination)) {
			return
		}
		notificationsTfSdk = append(notificationsTfSdk, notificationDestination)
	}

	if appendToResponse(resp, validateLength(notificationsTfSdk)) {
		return
	}

	notificationInfo.NotificationDestinations = notificationsTfSdk
	resp.Diagnostics.Append(resp.State.Set(ctx, notificationInfo)...)

}
