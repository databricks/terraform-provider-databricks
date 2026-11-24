package exporter

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/lakewatch"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

const (
	lakewatchParserPrefix        = "parsers/"
	lakewatchBuiltinParserPrefix = "builtin-parsers/"
	lakewatchVersionsInfix       = "/versions/"
	lakewatchParsersDir          = "lakewatch_parsers"
)

// getLakewatchBuiltinParsers returns built-in Lakewatch parsers keyed by parser ID. The list is
// fetched only once, and then reused. If it can't be fetched, an empty map is returned, so all
// parsers are treated as custom.
func (ic *importContext) getLakewatchBuiltinParsers() map[string]lakewatch.BuiltinParser {
	ic.lakewatchBuiltinParsersMutex.RLock()
	parsers := ic.lakewatchBuiltinParsers
	ic.lakewatchBuiltinParsersMutex.RUnlock()
	if parsers != nil {
		return parsers
	}

	ic.lakewatchBuiltinParsersMutex.Lock()
	defer ic.lakewatchBuiltinParsersMutex.Unlock()
	if ic.lakewatchBuiltinParsers != nil { // initialized by another goroutine
		return ic.lakewatchBuiltinParsers
	}
	parsers = map[string]lakewatch.BuiltinParser{}
	it := ic.workspaceClient.Lakewatch.ListBuiltinParsers(ic.Context, lakewatch.ListBuiltinParsersRequest{})
	for it.HasNext(ic.Context) {
		p, err := it.Next(ic.Context)
		if err != nil {
			log.Printf("[WARN] Can't list Lakewatch built-in parsers: %v", err)
			break
		}
		id := p.ParserId
		if id == "" {
			id = strings.TrimPrefix(p.Name, lakewatchBuiltinParserPrefix)
		}
		parsers[id] = p
	}
	log.Printf("[DEBUG] Found %d Lakewatch built-in parsers", len(parsers))
	ic.lakewatchBuiltinParsers = parsers
	return parsers
}

// emitLakewatchParsers emits custom parsers defined in the given catalog.
func (ic *importContext) emitLakewatchParsers(catalogName string) {
	it := ic.workspaceClient.Lakewatch.ListParsers(ic.Context, lakewatch.ListParsersRequest{CatalogName: catalogName})
	i := 0
	for it.HasNext(ic.Context) {
		p, err := it.Next(ic.Context)
		if err != nil {
			logLakewatchError("parsers", catalogName, err)
			break
		}
		fullName := fmt.Sprintf("%s.%s.%s", p.CatalogName, p.SchemaName, p.ParserName)
		if !ic.MatchesName(fullName) {
			log.Printf("[INFO] Skipping Lakewatch parser %s because it doesn't match %s", fullName, ic.match)
			continue
		}
		ic.Emit(&resource{
			Resource: "databricks_lakewatch_parser",
			ID:       p.Name,
		})
		i++
	}
	if i > 0 {
		log.Printf("[INFO] Scanned %d Lakewatch parsers in catalog %s", i, catalogName)
	}
}

// emitLakewatchParserRef emits the parser (and its version) used by a datasource. Built-in parsers
// are emitted as data sources, and custom parsers as resources. Version 1 of a custom parser is
// created together with the parser, so a separate resource is emitted only for other versions.
func (ic *importContext) emitLakewatchParserRef(ref lakewatch.ParserRef) {
	if ref.Id == "" {
		return
	}
	if bp, isBuiltin := ic.getLakewatchBuiltinParsers()[ref.Id]; isBuiltin {
		name := lakewatchBuiltinParserPrefix + ref.Id
		ic.Emit(&resource{
			Resource: "databricks_lakewatch_builtin_parser",
			ID:       name,
			DataWrapper: newDataSourceData(name, map[string]any{
				"name":        name,
				"parser_id":   ref.Id,
				"parser_name": bp.ParserName,
			}),
		})
		if ref.Version != "" {
			versionName := name + lakewatchVersionsInfix + ref.Version
			ic.Emit(&resource{
				Resource: "databricks_lakewatch_builtin_parser_version",
				ID:       versionName,
				DataWrapper: newDataSourceData(versionName, map[string]any{
					"name":        versionName,
					"parser_name": bp.ParserName,
					"version":     ref.Version,
				}),
			})
		}
		return
	}
	name := lakewatchParserPrefix + ref.Id
	ic.Emit(&resource{
		Resource: "databricks_lakewatch_parser",
		ID:       name,
	})
	if ref.Version != "" && ref.Version != "1" {
		ic.Emit(&resource{
			Resource: "databricks_lakewatch_parser_version",
			ID:       name + lakewatchVersionsInfix + ref.Version,
		})
	}
}

// importLakewatchParser populates the spec & description of the parser (they aren't returned by
// the parser's API) from its first available version, and emits its latest version.
func importLakewatchParser(ic *importContext, r *resource) error {
	w := ic.workspaceClient.Lakewatch
	pv, err := w.GetParserVersion(ic.Context, lakewatch.GetParserVersionRequest{
		Name: r.ID + lakewatchVersionsInfix + "1"})
	if apierr.IsMissing(err) {
		// version 1 was deleted, so we use the oldest available version
		var versions []lakewatch.ParserVersion
		versions, err = w.ListParserVersionsAll(ic.Context, lakewatch.ListParserVersionsRequest{Parent: r.ID})
		if err == nil && len(versions) == 0 {
			return fmt.Errorf("lakewatch parser %s doesn't have versions", r.ID)
		}
		if err == nil {
			pv, err = w.GetParserVersion(ic.Context, lakewatch.GetParserVersionRequest{Name: versions[0].Name})
		}
	}
	if err != nil {
		return err
	}
	if err = ic.saveLakewatchParserSpec(r, pv.Spec); err != nil {
		return err
	}
	if pv.Description != "" {
		r.DataWrapper.Set("description", pv.Description)
	}
	if latest, ok := r.DataWrapper.Get("latest_version").(int); ok && int64(latest) > pv.Version {
		ic.Emit(&resource{
			Resource: "databricks_lakewatch_parser_version",
			ID:       fmt.Sprintf("%s%s%d", r.ID, lakewatchVersionsInfix, latest),
		})
	}
	catalogName, _ := r.DataWrapper.Get("catalog_name").(string)
	schemaName, _ := r.DataWrapper.Get("schema_name").(string)
	ic.emitLakewatchSchema(catalogName, schemaName)
	return nil
}

// saveLakewatchParserSpec saves the YAML spec of a parser or its version into a file, and replaces
// the `spec` attribute with the relative path of that file, so it's loaded with the `file` function.
func (ic *importContext) saveLakewatchParserSpec(r *resource, spec string) error {
	if spec == "" {
		return nil
	}
	fileName, err := ic.saveContentIn(lakewatchParsersDir, r.Name+".yaml", []byte(spec))
	if err != nil {
		return fmt.Errorf("can't save spec of %s: %w", r.ID, err)
	}
	return r.DataWrapper.Set("spec", fileName)
}

// lakewatchParserVersionParent extracts the parser's name from the name of its version
// (`parsers/{parser_id}/versions/{version}`).
func lakewatchParserVersionParent(name string) string {
	parent, _, found := strings.Cut(name, lakewatchVersionsInfix)
	if !found || !strings.HasPrefix(parent, lakewatchParserPrefix) || parent == lakewatchParserPrefix {
		return ""
	}
	return parent
}

// importLakewatchParserVersion sets the input-only `parent` attribute and emits the parser.
func importLakewatchParserVersion(ic *importContext, r *resource) error {
	parent := lakewatchParserVersionParent(r.ID)
	if parent == "" {
		return fmt.Errorf("invalid name of Lakewatch parser version: %s", r.ID)
	}
	r.DataWrapper.Set("parent", parent)
	if spec, ok := r.DataWrapper.Get("spec").(string); ok {
		if err := ic.saveLakewatchParserSpec(r, spec); err != nil {
			return err
		}
	}
	ic.Emit(&resource{
		Resource: "databricks_lakewatch_parser",
		ID:       parent,
	})
	return nil
}

// lakewatchParserVersionNameUnified generates names like `{parser_name}_v{version}_{hash}`. The
// hash is included because parsers in different schemas may have the same name.
func lakewatchParserVersionNameUnified(ic *importContext, wrapper ResourceDataWrapper) string {
	parserName, _ := wrapper.Get("parser_name").(string)
	version := ""
	switch v := wrapper.Get("version").(type) {
	case int:
		version = strconv.Itoa(v)
	case string:
		version = v
	}
	if parserName == "" || version == "" {
		return generateUniqueID(wrapper.Id())
	}
	name := parserName + "_v" + version
	if strings.HasPrefix(wrapper.Id(), lakewatchParserPrefix) {
		name += "_" + generateUniqueID(wrapper.Id())[:6]
	}
	return name
}

// lakewatchBuiltinParserNameUnified uses the published name of a built-in parser.
func lakewatchBuiltinParserNameUnified(ic *importContext, wrapper ResourceDataWrapper) string {
	if parserName, _ := wrapper.Get("parser_name").(string); parserName != "" {
		return parserName
	}
	return generateUniqueID(wrapper.Id())
}

// generateLakewatchBuiltinBody generates a data source that refers to a built-in parser or its
// version by the resource name.
func generateLakewatchBuiltinBody(ic *importContext, body *hclwrite.Body, r *resource) error {
	b := body.AppendNewBlock("data", []string{r.Resource, r.Name}).Body()
	b.SetAttributeValue("name", cty.StringVal(r.ID))
	return nil
}

// isMatchingLakewatchParserVersion checks that the parser version belongs to the parser referred
// to in the same element of datasource's `parsers` (origPath is like `parsers.0.version`).
func isMatchingLakewatchParserVersion(ic *importContext, res *resource, ra *resourceApproximation,
	origPath string) bool {
	prefix, found := strings.CutSuffix(origPath, ".version")
	if !found {
		return false
	}
	parserId := resourceGetString(res, prefix+".id")
	version := resourceGetString(res, origPath)
	raName, ok := ra.Get("name")
	if parserId == "" || version == "" || !ok {
		return false
	}
	return strings.HasSuffix(raName.(string), "/"+parserId+lakewatchVersionsInfix+version)
}

// shouldOmitForLakewatchParser keeps the optional & computed `description` of the parser that
// is populated from its first version in the Import function.
func shouldOmitForLakewatchParser(ic *importContext, pathString string, fieldSchema FieldSchema,
	wrapper ResourceDataWrapper, r *resource) bool {
	if pathString == "description" {
		v, _ := wrapper.Get(pathString).(string)
		return v == ""
	}
	return DefaultShouldOmitFieldFuncWithAbstraction(ic, pathString, fieldSchema, wrapper, r)
}
