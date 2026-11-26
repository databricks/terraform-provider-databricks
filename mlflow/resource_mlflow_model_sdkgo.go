package mlflow

import (
	"context"
	"errors"

	modelregistry "github.com/databricks/sdk-go/modelregistry/v1"
	"github.com/databricks/sdk-go/options/call"
	"github.com/databricks/terraform-provider-databricks/common"
	sdkgoutility "github.com/databricks/terraform-provider-databricks/sdk-go-utility"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// mlflowModelResourceType is the name by which USE_SDK_GO_RESOURCES opts
// databricks_mlflow_model in to the modular SDK.
const mlflowModelResourceType = "databricks_mlflow_model"

type mlflowModelOperation = func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error

// mlflowModelWithSdkGo adds the modular SDK path to classic, the resource on
// the classic client. Every operation asks sdkgoutility.Enabled when it runs,
// so the opt-in can change between two operations of one provider. The schema
// is built once, from the modular request when the resource is opted in by
// then. Both schemas are equal in every property, so an operation of either
// path works on either of them.
func mlflowModelWithSdkGo(classic common.Resource) common.Resource {
	r := classic
	if sdkgoutility.Enabled(mlflowModelResourceType) {
		r.Schema = mlflowModelSdkGoSchema()
	}
	sdkGo := mlflowModelSdkGo{schema: r.Schema}
	r.Create = mlflowModelPickPath(classic.Create, sdkGo.create)
	r.Read = mlflowModelPickPath(classic.Read, sdkGo.read)
	r.Update = mlflowModelPickPath(classic.Update, sdkGo.update)
	r.Delete = mlflowModelPickPath(classic.Delete, sdkGo.delete)
	return r
}

func mlflowModelPickPath(classic, sdkGo mlflowModelOperation) mlflowModelOperation {
	return func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
		if sdkgoutility.Enabled(mlflowModelResourceType) {
			return sdkGo(ctx, d, c)
		}
		return classic(ctx, d, c)
	}
}

// mlflowModelSdkGoSchema is the schema of ResourceMlflowModel, built from the
// create request of the modular SDK instead of the classic one. StructToSchema
// makes every attribute optional and takes no embedded common.Namespace, so
// the customization adds what the classic struct tags say.
func mlflowModelSdkGoSchema() map[string]*schema.Schema {
	return sdkgoutility.StructToSchema(modelregistry.CreateRegisteredModelRequest{},
		func(s map[string]*schema.Schema) map[string]*schema.Schema {
			common.SetRequired(s["name"])
			s["name"].ForceNew = true
			s["registered_model_id"] = &schema.Schema{
				Computed: true,
				Type:     schema.TypeString,
			}
			return common.NamespaceCustomizeSchemaMap(common.AddNamespaceInSchema(s))
		})
}

// mlflowModelSdkGo is databricks_mlflow_model on the modular SDK client
// modelregistry/v1. Each operation sends the request of its classic twin and
// leaves the same state.
type mlflowModelSdkGo struct {
	schema map[string]*schema.Schema
}

// mlflowModelSdkGoCall runs method, a method of the modular client, with req
// against the workspace that d is routed to. The workspace client comes from
// the provider as on the classic path, so provider_config routes the call the
// same way. Its config builds the modular client, and ClassicCall retries,
// limits and reports the call by it as the classic client does. The config
// has authenticated, as NewClient asks, and is the same pointer for every call
// to a workspace, as ClassicCall asks: the provider builds each workspace
// client once, with databricks.NewWorkspaceClient.
func mlflowModelSdkGoCall[Req, Res any](
	ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient,
	method func(*modelregistry.Client, context.Context, Req, ...call.Option) (Res, error), req Req,
) (Res, error) {
	var zero Res
	w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
	if err != nil {
		return zero, err
	}
	registry, err := sdkgoutility.NewClient(ctx, w.Config, modelregistry.NewClient)
	if err != nil {
		return zero, err
	}
	return sdkgoutility.ClassicCall(ctx, w.Config, func(ctx context.Context, opts ...call.Option) (Res, error) {
		return method(registry, ctx, req, opts...)
	})
}

// mlflowModelSdkGoName returns the name for a request body. DataToStruct
// leaves an empty name nil, and the classic requests send name even when it is
// empty.
func mlflowModelSdkGoName(d *schema.ResourceData) *string {
	name := d.Get("name").(string)
	return &name
}

func (r mlflowModelSdkGo) create(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
	var req modelregistry.CreateRegisteredModelRequest
	if err := sdkgoutility.DataToStruct(d, r.schema, &req); err != nil {
		return err
	}
	req.Name = mlflowModelSdkGoName(d)
	res, err := mlflowModelSdkGoCall(ctx, d, c, (*modelregistry.Client).CreateRegisteredModel, req)
	if err != nil {
		return err
	}
	if res.RegisteredModel == nil {
		// The classic path panics on the nil pointer here.
		return errors.New("the response has no registered_model")
	}
	d.SetId(mlflowModelSdkGoString(res.RegisteredModel.Name))
	return nil
}

func (r mlflowModelSdkGo) read(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
	id := d.Id()
	req := modelregistry.GetRegisteredModelDatabricksRequest{Name: &id}
	res, err := mlflowModelSdkGoCall(ctx, d, c, (*modelregistry.Client).GetRegisteredModelDatabricks, req)
	if err != nil {
		return err
	}
	model := res.RegisteredModelDatabricks
	if model == nil {
		// The classic path reports the nil pointer as an error of its reflection helper.
		return errors.New("the response has no registered_model_databricks")
	}
	if err := sdkgoutility.StructToData(model, r.schema, d); err != nil {
		return err
	}
	d.Set("registered_model_id", mlflowModelSdkGoString(model.Id)) // alias
	return nil
}

func (r mlflowModelSdkGo) update(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
	var req modelregistry.UpdateRegisteredModelRequest
	if err := sdkgoutility.DataToStruct(d, r.schema, &req); err != nil {
		return err
	}
	req.Name = mlflowModelSdkGoName(d)
	_, err := mlflowModelSdkGoCall(ctx, d, c, (*modelregistry.Client).UpdateRegisteredModel, req)
	return err
}

func (r mlflowModelSdkGo) delete(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
	id := d.Id()
	req := modelregistry.DeleteRegisteredModelRequest{Name: &id}
	_, err := mlflowModelSdkGoCall(ctx, d, c, (*modelregistry.Client).DeleteRegisteredModel, req)
	return err
}

// mlflowModelSdkGoString returns the string that the classic struct holds for
// a field of a response: empty when the response leaves the field out.
func mlflowModelSdkGoString(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}
