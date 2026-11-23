package exporter

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/lakewatch"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetLakewatchBuiltinParsersIsCached(t *testing.T) {
	qa.MockWorkspaceApply(t, func(mw *mocks.MockWorkspaceClient) {
		mw.GetMockLakewatchAPI().EXPECT().
			ListBuiltinParsers(mock.Anything, lakewatch.ListBuiltinParsersRequest{}).
			Return(createIteratorFromSlice([]lakewatch.BuiltinParser{
				{Name: "builtin-parsers/b1", ParserId: "b1", ParserName: "azure_signin"},
				{Name: "builtin-parsers/b2", ParserName: "okta"},
			})).Once()
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		parsers := ic.getLakewatchBuiltinParsers()
		assert.Len(t, parsers, 2)
		assert.Equal(t, "azure_signin", parsers["b1"].ParserName)
		// ID is extracted from the name when it's not set
		assert.Equal(t, "okta", parsers["b2"].ParserName)
		// second call must not trigger the API
		assert.Len(t, ic.getLakewatchBuiltinParsers(), 2)
	})
}

func TestGetLakewatchBuiltinParsersError(t *testing.T) {
	qa.MockWorkspaceApply(t, func(mw *mocks.MockWorkspaceClient) {
		mw.GetMockLakewatchAPI().EXPECT().
			ListBuiltinParsers(mock.Anything, mock.Anything).
			Return(createErrorIterator[lakewatch.BuiltinParser](
				fmt.Errorf("not enabled: %w", apierr.ErrNotImplemented))).Once()
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		parsers := ic.getLakewatchBuiltinParsers()
		assert.NotNil(t, parsers)
		assert.Empty(t, parsers)
		assert.Empty(t, ic.getLakewatchBuiltinParsers())
	})
}

// assertLakewatchSpecFile checks that the parser spec is loaded from a file with expected content.
func assertLakewatchSpecFile(t *testing.T, dir, hcl, resourceName, expectedSpec string) {
	fileName := lakewatchParsersDir + "/" + resourceName + ".yaml"
	assert.Contains(t, hcl, `spec = file("${path.module}/`+fileName+`")`)
	content, err := os.ReadFile(filepath.Join(dir, fileName))
	require.NoError(t, err)
	assert.Equal(t, expectedSpec, string(content))
}

func newLakewatchParserTestWrapper(t *testing.T, ic *importContext, latest int64) ResourceDataWrapper {
	return newLakewatchTestWrapper(t, ic, "databricks_lakewatch_parser", "parsers/p1", map[string]any{
		"name":           types.StringValue("parsers/p1"),
		"parser_id":      types.StringValue("p1"),
		"catalog_name":   types.StringValue("sec"),
		"schema_name":    types.StringValue("raw"),
		"parser_name":    types.StringValue("trivy"),
		"latest_version": types.Int64Value(latest),
	})
}

func TestImportLakewatchParser(t *testing.T) {
	qa.MockWorkspaceApply(t, func(mw *mocks.MockWorkspaceClient) {
		mw.GetMockLakewatchAPI().EXPECT().
			GetParserVersion(mock.Anything, lakewatch.GetParserVersionRequest{Name: "parsers/p1/versions/1"}).
			Return(&lakewatch.ParserVersion{Name: "parsers/p1/versions/1", Spec: "fields: v1",
				Description: "initial", Version: 1}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("lakewatch,uc-schemas")
		wrapper := newLakewatchParserTestWrapper(t, ic, 3)
		assert.Equal(t, "sec_raw_trivy",
			ic.Importables["databricks_lakewatch_parser"].NameUnified(ic, wrapper))

		ic.Directory = t.TempDir()
		err := importLakewatchParser(ic, &resource{Resource: "databricks_lakewatch_parser",
			ID: "parsers/p1", Name: "sec_raw_trivy", DataWrapper: wrapper})
		require.NoError(t, err)
		assert.Equal(t, "lakewatch_parsers/sec_raw_trivy.yaml", wrapper.Get("spec"))
		content, err := os.ReadFile(filepath.Join(ic.Directory, "lakewatch_parsers/sec_raw_trivy.yaml"))
		require.NoError(t, err)
		assert.Equal(t, "fields: v1", string(content))
		assert.Equal(t, "initial", wrapper.Get("description"))
		assert.Equal(t, map[string]bool{
			"databricks_lakewatch_parser_version[<unknown>] (id: parsers/p1/versions/3)": true,
			"databricks_schema[<unknown>] (id: sec.raw)":                                 true,
		}, ic.testEmits)
	})
}

func TestImportLakewatchParserWithoutFirstVersion(t *testing.T) {
	qa.MockWorkspaceApply(t, func(mw *mocks.MockWorkspaceClient) {
		api := mw.GetMockLakewatchAPI().EXPECT()
		api.GetParserVersion(mock.Anything, lakewatch.GetParserVersionRequest{Name: "parsers/p1/versions/1"}).
			Return(nil, fmt.Errorf("deleted: %w", apierr.ErrNotFound))
		api.ListParserVersionsAll(mock.Anything, lakewatch.ListParserVersionsRequest{Parent: "parsers/p1"}).
			Return([]lakewatch.ParserVersion{{Name: "parsers/p1/versions/2", Version: 2},
				{Name: "parsers/p1/versions/3", Version: 3}}, nil)
		api.GetParserVersion(mock.Anything, lakewatch.GetParserVersionRequest{Name: "parsers/p1/versions/2"}).
			Return(&lakewatch.ParserVersion{Name: "parsers/p1/versions/2", Spec: "fields: v2", Version: 2}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("lakewatch")
		wrapper := newLakewatchParserTestWrapper(t, ic, 2)

		ic.Directory = t.TempDir()
		err := importLakewatchParser(ic, &resource{Resource: "databricks_lakewatch_parser",
			ID: "parsers/p1", Name: "sec_raw_trivy", DataWrapper: wrapper})
		require.NoError(t, err)
		content, err := os.ReadFile(filepath.Join(ic.Directory, "lakewatch_parsers/sec_raw_trivy.yaml"))
		require.NoError(t, err)
		assert.Equal(t, "fields: v2", string(content))
		// the latest version is used as the parser's spec, so it isn't emitted separately
		assert.Empty(t, ic.testEmits)
	})
}

func TestImportLakewatchParserNoVersions(t *testing.T) {
	qa.MockWorkspaceApply(t, func(mw *mocks.MockWorkspaceClient) {
		api := mw.GetMockLakewatchAPI().EXPECT()
		api.GetParserVersion(mock.Anything, mock.Anything).
			Return(nil, fmt.Errorf("deleted: %w", apierr.ErrNotFound))
		api.ListParserVersionsAll(mock.Anything, mock.Anything).Return(nil, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		err := importLakewatchParser(ic, &resource{Resource: "databricks_lakewatch_parser",
			ID: "parsers/p1", DataWrapper: newLakewatchParserTestWrapper(t, ic, 0)})
		assert.ErrorContains(t, err, "doesn't have versions")
	})
}

func TestImportLakewatchParserVersion(t *testing.T) {
	ic := importContextForTest()
	ic.Context = context.Background()
	ic.enableServices("lakewatch")
	wrapper := newLakewatchTestWrapper(t, ic, "databricks_lakewatch_parser_version", "parsers/p1/versions/2",
		map[string]any{
			"name":        types.StringValue("parsers/p1/versions/2"),
			"parser_name": types.StringValue("trivy"),
			"version":     types.Int64Value(2),
			"spec":        types.StringValue("fields: v2"),
		})
	assert.Equal(t, "trivy_v2_"+generateUniqueID("parsers/p1/versions/2")[:6],
		ic.Importables["databricks_lakewatch_parser_version"].NameUnified(ic, wrapper))

	ic.Directory = t.TempDir()
	err := importLakewatchParserVersion(ic, &resource{Resource: "databricks_lakewatch_parser_version",
		ID: "parsers/p1/versions/2", Name: "trivy_v2", DataWrapper: wrapper})
	require.NoError(t, err)
	assert.Equal(t, "parsers/p1", wrapper.Get("parent"))
	assert.Equal(t, "lakewatch_parsers/trivy_v2.yaml", wrapper.Get("spec"))
	content, err := os.ReadFile(filepath.Join(ic.Directory, "lakewatch_parsers/trivy_v2.yaml"))
	require.NoError(t, err)
	assert.Equal(t, "fields: v2", string(content))
	assert.Contains(t, ic.testEmits, "databricks_lakewatch_parser[<unknown>] (id: parsers/p1)")

	err = importLakewatchParserVersion(ic, &resource{Resource: "databricks_lakewatch_parser_version",
		ID: "parsers/p1", DataWrapper: wrapper})
	assert.ErrorContains(t, err, "invalid name")
}

func TestSaveLakewatchParserSpecError(t *testing.T) {
	ic := importContextForTest()
	ic.Directory = filepath.Join(t.TempDir(), "file")
	require.NoError(t, os.WriteFile(ic.Directory, []byte{}, 0644)) // a file, so a directory can't be created
	r := &resource{ID: "parsers/p1", Name: "p1", DataWrapper: newDataSourceData("parsers/p1", nil)}
	assert.ErrorContains(t, ic.saveLakewatchParserSpec(r, "fields: v1"), "can't save spec of parsers/p1")
	// nothing is saved for an empty spec
	assert.NoError(t, ic.saveLakewatchParserSpec(r, ""))
}

func TestLakewatchParserVersionParent(t *testing.T) {
	assert.Equal(t, "parsers/p1", lakewatchParserVersionParent("parsers/p1/versions/2"))
	assert.Equal(t, "", lakewatchParserVersionParent("parsers/p1"))
	assert.Equal(t, "", lakewatchParserVersionParent("parsers//versions/2"))
	assert.Equal(t, "", lakewatchParserVersionParent("builtin-parsers/b1/versions/2"))
}

func TestLakewatchBuiltinParserNames(t *testing.T) {
	ic := importContextForTest()
	parser := newDataSourceData("builtin-parsers/b1", map[string]any{"parser_name": "azure_signin"})
	assert.Equal(t, "azure_signin", lakewatchBuiltinParserNameUnified(ic, parser))
	assert.Equal(t, generateUniqueID("builtin-parsers/b2"),
		lakewatchBuiltinParserNameUnified(ic, newDataSourceData("builtin-parsers/b2", nil)))

	version := newDataSourceData("builtin-parsers/b1/versions/3",
		map[string]any{"parser_name": "azure_signin", "version": "3"})
	assert.Equal(t, "azure_signin_v3", lakewatchParserVersionNameUnified(ic, version))
	assert.Equal(t, generateUniqueID("builtin-parsers/b1/versions/4"),
		lakewatchParserVersionNameUnified(ic, newDataSourceData("builtin-parsers/b1/versions/4", nil)))
}

func TestGenerateLakewatchBuiltinBody(t *testing.T) {
	f := hclwrite.NewEmptyFile()
	err := generateLakewatchBuiltinBody(importContextForTest(), f.Body(), &resource{
		Resource: "databricks_lakewatch_builtin_parser", Name: "azure_signin", ID: "builtin-parsers/b1"})
	require.NoError(t, err)
	assert.Equal(t, "data \"databricks_lakewatch_builtin_parser\" \"azure_signin\" {\n"+
		"  name = \"builtin-parsers/b1\"\n}\n", string(hclwrite.Format(f.Bytes())))
}

func TestIsMatchingLakewatchParserVersion(t *testing.T) {
	ic := importContextForTest()
	res := &resource{DataWrapper: newDataSourceData("datasources/d1", map[string]any{
		"parsers.0.id":      "p1",
		"parsers.0.version": "2",
	})}
	newRa := func(name string) *resourceApproximation {
		return &resourceApproximation{Instances: []instanceApproximation{
			{Attributes: map[string]any{"name": name}}}}
	}
	assert.True(t, isMatchingLakewatchParserVersion(ic, res, newRa("parsers/p1/versions/2"), "parsers.0.version"))
	assert.True(t, isMatchingLakewatchParserVersion(ic, res,
		newRa("builtin-parsers/p1/versions/2"), "parsers.0.version"))
	assert.False(t, isMatchingLakewatchParserVersion(ic, res, newRa("parsers/p2/versions/2"), "parsers.0.version"))
	assert.False(t, isMatchingLakewatchParserVersion(ic, res, newRa("parsers/p1/versions/3"), "parsers.0.version"))
	assert.False(t, isMatchingLakewatchParserVersion(ic, res, newRa("parsers/p1/versions/2"), "parsers.0.id"))
	assert.False(t, isMatchingLakewatchParserVersion(ic, res, newRa("parsers/p1/versions/2"), "parsers.1.version"))
}

func TestDataSourceData(t *testing.T) {
	d := newDataSourceData("builtin-parsers/b1", map[string]any{"parser_id": "b1", "name": "builtin-parsers/b1"})
	assert.Equal(t, "builtin-parsers/b1", d.Id())
	assert.True(t, d.IsPluginFramework())
	v, ok := d.GetOk("parser_id")
	assert.True(t, ok)
	assert.Equal(t, "b1", v)
	_, ok = d.GetOk("missing")
	assert.False(t, ok)
	require.NoError(t, d.Set("version", "3"))
	assert.Equal(t, "3", d.Get("version"))
	d.SetId("builtin-parsers/b2")
	assert.Equal(t, "builtin-parsers/b2", d.Id())
	schema := d.GetSchema()
	assert.Equal(t, []string{"name", "parser_id", "version"}, schema.GetFields())
	assert.Nil(t, schema.GetField("name"))
	assert.True(t, schema.IsPluginFramework())
	assert.Error(t, d.GetTypedStruct(context.Background(), &struct{}{}))
}

func TestImportDataSourceOnlyResource(t *testing.T) {
	ic := importContextForTest()
	ic.enableServices("lakewatch")
	ic.importing = map[string]bool{}
	ic.waitGroup.Add(2)
	(&resource{Resource: "databricks_lakewatch_builtin_parser", ID: "builtin-parsers/b1",
		DataWrapper: newDataSourceData("builtin-parsers/b1", map[string]any{
			"parser_id": "b1", "parser_name": "azure_signin"})}).ImportResource(ic)
	// data must be provided for data sources
	(&resource{Resource: "databricks_lakewatch_builtin_parser", ID: "builtin-parsers/b2"}).ImportResource(ic)

	ra := ic.State.Get("databricks_lakewatch_builtin_parser", "parser_id", "b1")
	require.NotNil(t, ra)
	assert.Equal(t, "data", ra.Mode)
	assert.Equal(t, "azure_signin", ra.Name)
	assert.Nil(t, ic.State.Get("databricks_lakewatch_builtin_parser", "id", "builtin-parsers/b2"))
}
