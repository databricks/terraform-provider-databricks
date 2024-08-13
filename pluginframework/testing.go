package pluginframework

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type ResourceFixturePluginFramework struct {
	Fixtures                []qa.HTTPFixture
	MockWorkspaceClientFunc func(*mocks.MockWorkspaceClient)
	MockAccountClientFunc   func(*mocks.MockAccountClient)
	Token                   string
}

func (f ResourceFixturePluginFramework) validateMocks() error {
	isMockConfigured := f.MockAccountClientFunc != nil || f.MockWorkspaceClientFunc != nil
	isFixtureConfigured := f.Fixtures != nil
	if isFixtureConfigured && isMockConfigured {
		return fmt.Errorf("either (MockWorkspaceClientFunc, MockAccountClientFunc) or Fixtures may be set, not both")
	}
	return nil
}

func (f ResourceFixturePluginFramework) Start(t *testing.T) (*common.DatabricksClient, error) {
	err := f.validateMocks()
	if err != nil {
		return nil, err
	}
	client, server, err := f.setupClient(t)
	if err != nil {
		return nil, err
	}
	defer server.Close()
	return client, nil
}

func (f ResourceFixturePluginFramework) setupClient(t *testing.T) (*common.DatabricksClient, qa.Server, error) {
	token := "..."
	if f.Token != "" {
		token = f.Token
	}
	if f.Fixtures != nil {
		client, s, err := qa.HttpFixtureClientWithToken(t, f.Fixtures, token)
		ss := qa.Server{
			Close: s.Close,
			URL:   s.URL,
		}
		return client, ss, err
	}
	mw := mocks.NewMockWorkspaceClient(t)
	ma := mocks.NewMockAccountClient(t)
	if f.MockWorkspaceClientFunc != nil {
		f.MockWorkspaceClientFunc(mw)
	}
	if f.MockAccountClientFunc != nil {
		f.MockAccountClientFunc(ma)
	}
	c := &common.DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{},
		},
	}
	c.SetWorkspaceClient(mw.WorkspaceClient)
	c.SetAccountClient(ma.AccountClient)
	c.Config.Credentials = qa.TestCredentialsProvider{Token: token}
	return c, qa.Server{
		Close: func() {},
		URL:   "does-not-matter",
	}, nil
}

func StructToTfTypesValue(s interface{}) tftypes.Value {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Struct {
		panic("provided input is not a struct")
	}

	result := make(map[string]interface{})
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := t.Field(i).Name
		if field.CanInterface() {
			result[fieldName] = field.Interface()
		}
	}
	return MapToTfTypesValue(result)
}

func MapToTfTypesValue(configMap map[string]any) tftypes.Value {
	configTypeMap := map[string]tftypes.Type{}
	rawConfigValueMap := map[string]tftypes.Value{}
	for k, v := range configMap {
		switch v := v.(type) {
		case string:
			configTypeMap[k] = tftypes.String
			rawConfigValueMap[k] = tftypes.NewValue(tftypes.String, v)
		case int:
			configTypeMap[k] = tftypes.Number
			rawConfigValueMap[k] = tftypes.NewValue(tftypes.Bool, int(v))
		case bool:
			configTypeMap[k] = tftypes.Bool
			rawConfigValueMap[k] = tftypes.NewValue(tftypes.Bool, v)
		default:
			configTypeMap[k] = tftypes.String // tfypes.Object{}?
			rawConfigValueMap[k] = tftypes.NewValue(tftypes.String, fmt.Sprintf("%v", v))
		}
	}
	rawConfigType := tftypes.Object{
		AttributeTypes: configTypeMap,
	}
	rawConfigValue := tftypes.NewValue(rawConfigType, rawConfigValueMap)
	return rawConfigValue
}
