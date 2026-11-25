// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package lakewatch

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just Lakewatch API methods
type lakewatchImpl struct {
	client *client.DatabricksClient
}

func (a *lakewatchImpl) CreateDatasource(ctx context.Context, request CreateDatasourceRequest) (*Datasource, error) {
	var datasource Datasource
	path := "/api/2.0/lakewatch/datasources"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Datasource, &datasource)
	return &datasource, err
}

func (a *lakewatchImpl) CreateParser(ctx context.Context, request CreateParserRequest) (*Parser, error) {
	var parser Parser
	path := "/api/2.0/lakewatch/parsers"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Parser, &parser)
	return &parser, err
}

func (a *lakewatchImpl) CreateParserVersion(ctx context.Context, request CreateParserVersionRequest) (*ParserVersion, error) {
	var parserVersion ParserVersion
	path := fmt.Sprintf("/api/2.0/lakewatch/%v/versions", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.ParserVersion, &parserVersion)
	return &parserVersion, err
}

func (a *lakewatchImpl) CreateRule(ctx context.Context, request CreateRuleRequest) (*Rule, error) {
	var rule Rule
	path := "/api/2.0/lakewatch/rules"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Rule, &rule)
	return &rule, err
}

func (a *lakewatchImpl) DeleteDatasource(ctx context.Context, request DeleteDatasourceRequest) error {
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *lakewatchImpl) DeleteParser(ctx context.Context, request DeleteParserRequest) error {
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *lakewatchImpl) DeleteParserVersion(ctx context.Context, request DeleteParserVersionRequest) error {
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *lakewatchImpl) DeleteRule(ctx context.Context, request DeleteRuleRequest) error {
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *lakewatchImpl) GetAlertConfig(ctx context.Context, request GetAlertConfigRequest) (*AlertConfig, error) {
	var alertConfig AlertConfig
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &alertConfig)
	return &alertConfig, err
}

func (a *lakewatchImpl) GetBuiltinParser(ctx context.Context, request GetBuiltinParserRequest) (*BuiltinParser, error) {
	var builtinParser BuiltinParser
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &builtinParser)
	return &builtinParser, err
}

func (a *lakewatchImpl) GetBuiltinParserVersion(ctx context.Context, request GetBuiltinParserVersionRequest) (*BuiltinParserVersion, error) {
	var builtinParserVersion BuiltinParserVersion
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &builtinParserVersion)
	return &builtinParserVersion, err
}

func (a *lakewatchImpl) GetCommonConfig(ctx context.Context, request GetCommonConfigRequest) (*CommonConfig, error) {
	var commonConfig CommonConfig
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &commonConfig)
	return &commonConfig, err
}

func (a *lakewatchImpl) GetDatasource(ctx context.Context, request GetDatasourceRequest) (*Datasource, error) {
	var datasource Datasource
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &datasource)
	return &datasource, err
}

func (a *lakewatchImpl) GetDatasourceConfig(ctx context.Context, request GetDatasourceConfigRequest) (*DatasourceConfig, error) {
	var datasourceConfig DatasourceConfig
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &datasourceConfig)
	return &datasourceConfig, err
}

func (a *lakewatchImpl) GetParser(ctx context.Context, request GetParserRequest) (*Parser, error) {
	var parser Parser
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &parser)
	return &parser, err
}

func (a *lakewatchImpl) GetParserVersion(ctx context.Context, request GetParserVersionRequest) (*ParserVersion, error) {
	var parserVersion ParserVersion
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &parserVersion)
	return &parserVersion, err
}

func (a *lakewatchImpl) GetRule(ctx context.Context, request GetRuleRequest) (*Rule, error) {
	var rule Rule
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &rule)
	return &rule, err
}

func (a *lakewatchImpl) GetRuleConfig(ctx context.Context, request GetRuleConfigRequest) (*RuleConfig, error) {
	var ruleConfig RuleConfig
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &ruleConfig)
	return &ruleConfig, err
}

// Lists a built-in parser's published versions, oldest first (ascending
// 1..resource_revision), without their specs.
func (a *lakewatchImpl) ListBuiltinParserVersions(ctx context.Context, request ListBuiltinParserVersionsRequest) listing.Iterator[BuiltinParserVersion] {

	getNextPage := func(ctx context.Context, req ListBuiltinParserVersionsRequest) (*ListBuiltinParserVersionsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListBuiltinParserVersions(ctx, req)
	}
	getItems := func(resp *ListBuiltinParserVersionsResponse) []BuiltinParserVersion {
		return resp.BuiltinParserVersions
	}
	getNextReq := func(resp *ListBuiltinParserVersionsResponse) *ListBuiltinParserVersionsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Lists a built-in parser's published versions, oldest first (ascending
// 1..resource_revision), without their specs.
func (a *lakewatchImpl) ListBuiltinParserVersionsAll(ctx context.Context, request ListBuiltinParserVersionsRequest) ([]BuiltinParserVersion, error) {
	iterator := a.ListBuiltinParserVersions(ctx, request)
	return listing.ToSlice[BuiltinParserVersion](ctx, iterator)
}

func (a *lakewatchImpl) internalListBuiltinParserVersions(ctx context.Context, request ListBuiltinParserVersionsRequest) (*ListBuiltinParserVersionsResponse, error) {
	var listBuiltinParserVersionsResponse ListBuiltinParserVersionsResponse
	path := fmt.Sprintf("/api/2.0/lakewatch/%v/versions", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listBuiltinParserVersionsResponse)
	return &listBuiltinParserVersionsResponse, err
}

// Lists the built-in parsers Databricks publishes, without their specs.
func (a *lakewatchImpl) ListBuiltinParsers(ctx context.Context, request ListBuiltinParsersRequest) listing.Iterator[BuiltinParser] {

	getNextPage := func(ctx context.Context, req ListBuiltinParsersRequest) (*ListBuiltinParsersResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListBuiltinParsers(ctx, req)
	}
	getItems := func(resp *ListBuiltinParsersResponse) []BuiltinParser {
		return resp.BuiltinParsers
	}
	getNextReq := func(resp *ListBuiltinParsersResponse) *ListBuiltinParsersRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Lists the built-in parsers Databricks publishes, without their specs.
func (a *lakewatchImpl) ListBuiltinParsersAll(ctx context.Context, request ListBuiltinParsersRequest) ([]BuiltinParser, error) {
	iterator := a.ListBuiltinParsers(ctx, request)
	return listing.ToSlice[BuiltinParser](ctx, iterator)
}

func (a *lakewatchImpl) internalListBuiltinParsers(ctx context.Context, request ListBuiltinParsersRequest) (*ListBuiltinParsersResponse, error) {
	var listBuiltinParsersResponse ListBuiltinParsersResponse
	path := "/api/2.0/lakewatch/builtin-parsers"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listBuiltinParsersResponse)
	return &listBuiltinParsersResponse, err
}

// Lists the datasources in a Unity Catalog catalog, optionally narrowed to one
// schema or filtered by name. Results are paginated. Datasources that are still
// being created or deleted are not listed.
func (a *lakewatchImpl) ListDatasources(ctx context.Context, request ListDatasourcesRequest) listing.Iterator[Datasource] {

	getNextPage := func(ctx context.Context, req ListDatasourcesRequest) (*ListDatasourcesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListDatasources(ctx, req)
	}
	getItems := func(resp *ListDatasourcesResponse) []Datasource {
		return resp.Datasources
	}
	getNextReq := func(resp *ListDatasourcesResponse) *ListDatasourcesRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Lists the datasources in a Unity Catalog catalog, optionally narrowed to one
// schema or filtered by name. Results are paginated. Datasources that are still
// being created or deleted are not listed.
func (a *lakewatchImpl) ListDatasourcesAll(ctx context.Context, request ListDatasourcesRequest) ([]Datasource, error) {
	iterator := a.ListDatasources(ctx, request)
	return listing.ToSlice[Datasource](ctx, iterator)
}

func (a *lakewatchImpl) internalListDatasources(ctx context.Context, request ListDatasourcesRequest) (*ListDatasourcesResponse, error) {
	var listDatasourcesResponse ListDatasourcesResponse
	path := "/api/2.0/lakewatch/datasources"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listDatasourcesResponse)
	return &listDatasourcesResponse, err
}

// Lists a parser's versions, oldest first, without their specs.
func (a *lakewatchImpl) ListParserVersions(ctx context.Context, request ListParserVersionsRequest) listing.Iterator[ParserVersion] {

	getNextPage := func(ctx context.Context, req ListParserVersionsRequest) (*ListParserVersionsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListParserVersions(ctx, req)
	}
	getItems := func(resp *ListParserVersionsResponse) []ParserVersion {
		return resp.ParserVersions
	}
	getNextReq := func(resp *ListParserVersionsResponse) *ListParserVersionsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Lists a parser's versions, oldest first, without their specs.
func (a *lakewatchImpl) ListParserVersionsAll(ctx context.Context, request ListParserVersionsRequest) ([]ParserVersion, error) {
	iterator := a.ListParserVersions(ctx, request)
	return listing.ToSlice[ParserVersion](ctx, iterator)
}

func (a *lakewatchImpl) internalListParserVersions(ctx context.Context, request ListParserVersionsRequest) (*ListParserVersionsResponse, error) {
	var listParserVersionsResponse ListParserVersionsResponse
	path := fmt.Sprintf("/api/2.0/lakewatch/%v/versions", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listParserVersionsResponse)
	return &listParserVersionsResponse, err
}

// Lists parsers in a catalog, without their specs.
func (a *lakewatchImpl) ListParsers(ctx context.Context, request ListParsersRequest) listing.Iterator[Parser] {

	getNextPage := func(ctx context.Context, req ListParsersRequest) (*ListParsersResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListParsers(ctx, req)
	}
	getItems := func(resp *ListParsersResponse) []Parser {
		return resp.Parsers
	}
	getNextReq := func(resp *ListParsersResponse) *ListParsersRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Lists parsers in a catalog, without their specs.
func (a *lakewatchImpl) ListParsersAll(ctx context.Context, request ListParsersRequest) ([]Parser, error) {
	iterator := a.ListParsers(ctx, request)
	return listing.ToSlice[Parser](ctx, iterator)
}

func (a *lakewatchImpl) internalListParsers(ctx context.Context, request ListParsersRequest) (*ListParsersResponse, error) {
	var listParsersResponse ListParsersResponse
	path := "/api/2.0/lakewatch/parsers"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listParsersResponse)
	return &listParsersResponse, err
}

// Lists rules, with optional filtering and pagination.
func (a *lakewatchImpl) ListRules(ctx context.Context, request ListRulesRequest) listing.Iterator[Rule] {

	getNextPage := func(ctx context.Context, req ListRulesRequest) (*ListRulesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListRules(ctx, req)
	}
	getItems := func(resp *ListRulesResponse) []Rule {
		return resp.Rules
	}
	getNextReq := func(resp *ListRulesResponse) *ListRulesRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Lists rules, with optional filtering and pagination.
func (a *lakewatchImpl) ListRulesAll(ctx context.Context, request ListRulesRequest) ([]Rule, error) {
	iterator := a.ListRules(ctx, request)
	return listing.ToSlice[Rule](ctx, iterator)
}

func (a *lakewatchImpl) internalListRules(ctx context.Context, request ListRulesRequest) (*ListRulesResponse, error) {
	var listRulesResponse ListRulesResponse
	path := "/api/2.0/lakewatch/rules"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listRulesResponse)
	return &listRulesResponse, err
}

func (a *lakewatchImpl) UpdateAlertConfig(ctx context.Context, request UpdateAlertConfigRequest) (*AlertConfig, error) {
	var alertConfig AlertConfig
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)

	if request.Etag != "" || slices.Contains(request.ForceSendFields, "Etag") {
		queryParams["etag"] = request.Etag
	}

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.AlertConfig, &alertConfig)
	return &alertConfig, err
}

func (a *lakewatchImpl) UpdateCommonConfig(ctx context.Context, request UpdateCommonConfigRequest) (*CommonConfig, error) {
	var commonConfig CommonConfig
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)

	if request.Etag != "" || slices.Contains(request.ForceSendFields, "Etag") {
		queryParams["etag"] = request.Etag
	}

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.CommonConfig, &commonConfig)
	return &commonConfig, err
}

func (a *lakewatchImpl) UpdateDatasource(ctx context.Context, request UpdateDatasourceRequest) (*Datasource, error) {
	var datasource Datasource
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)

	if request.Etag != "" || slices.Contains(request.ForceSendFields, "Etag") {
		queryParams["etag"] = request.Etag
	}

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Datasource, &datasource)
	return &datasource, err
}

func (a *lakewatchImpl) UpdateDatasourceConfig(ctx context.Context, request UpdateDatasourceConfigRequest) (*DatasourceConfig, error) {
	var datasourceConfig DatasourceConfig
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)

	if request.Etag != "" || slices.Contains(request.ForceSendFields, "Etag") {
		queryParams["etag"] = request.Etag
	}

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.DatasourceConfig, &datasourceConfig)
	return &datasourceConfig, err
}

func (a *lakewatchImpl) UpdateRule(ctx context.Context, request UpdateRuleRequest) (*Rule, error) {
	var rule Rule
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)

	if request.Etag != "" || slices.Contains(request.ForceSendFields, "Etag") {
		queryParams["etag"] = request.Etag
	}

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Rule, &rule)
	return &rule, err
}

func (a *lakewatchImpl) UpdateRuleConfig(ctx context.Context, request UpdateRuleConfigRequest) (*RuleConfig, error) {
	var ruleConfig RuleConfig
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)

	if request.Etag != "" || slices.Contains(request.ForceSendFields, "Etag") {
		queryParams["etag"] = request.Etag
	}

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.RuleConfig, &ruleConfig)
	return &ruleConfig, err
}
