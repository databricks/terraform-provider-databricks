package exporter

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"maps"
	"os"
	"reflect"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/databricks/terraform-provider-databricks/workspace"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/zclconf/go-cty/cty"
)

// TODO: move to IC
var dependsRe = regexp.MustCompile(`(\.[\d]+)`)

func (ic *importContext) generateVariableName(attrName, name string) string {
	return fmt.Sprintf("%s_%s", attrName, name)
}

func maybeAddQuoteCharacter(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "\"", "\\\"")
	return s
}

func genTraversalTokens(sr *resourceApproximation, pick string) hcl.Traversal {
	if sr.Mode == "data" {
		return hcl.Traversal{
			hcl.TraverseRoot{Name: "data"},
			hcl.TraverseAttr{Name: sr.Type},
			hcl.TraverseAttr{Name: sr.Name},
			hcl.TraverseAttr{Name: pick},
		}
	}
	return hcl.Traversal{
		hcl.TraverseRoot{Name: sr.Type},
		hcl.TraverseAttr{Name: sr.Name},
		hcl.TraverseAttr{Name: pick},
	}
}

func (ic *importContext) isIgnoredResourceApproximation(ra *resourceApproximation) bool {
	var ignored bool
	if ra != nil && ra.Resource != nil {
		ignoreFunc := ic.Importables[ra.Type].Ignore
		if ignoreFunc != nil && ignoreFunc(ic, ra.Resource) {
			log.Printf("[WARN] Found reference to the ignored resource %s: %s", ra.Type, ra.Name)
			return true
		}
	}
	return ignored
}

func (ic *importContext) Find(value, attr string, ref reference, origResource *resource, origPath string) (string, hcl.Traversal, bool) {
	log.Printf("[DEBUG] Starting searching for reference for resource %s, attr='%s', value='%s', ref=%v",
		ref.Resource, attr, value, ref)
	// optimize performance by avoiding doing regexp matching multiple times
	matchValue := ""
	switch ref.MatchType {
	case MatchRegexp:
		if ref.Regexp == nil {
			log.Printf("[WARN] you must provide regular expression for 'regexp' match type")
			return "", nil, false
		}
		res := ref.Regexp.FindStringSubmatch(value)
		if res == nil {
			return "", nil, false
		}
		if len(res) < 2 {
			log.Printf("[WARN] no match for regexp: %v in string %s", ref.Regexp, value)
			return "", nil, false
		}
		matchValue = res[1]
	case MatchCaseInsensitive:
		matchValue = strings.ToLower(value) // performance optimization to avoid doing it in the loop
	case MatchExact, MatchDefault:
		matchValue = value
	case MatchPrefix, MatchLongestPrefix:
		if ref.MatchValueTransformFunc != nil {
			matchValue = ref.MatchValueTransformFunc(value)
		} else {
			matchValue = value
		}
	}
	// doing explicit lookup in the state.  For case insensitive matches, first attempt to lookup for the value,
	// and do iteration if it's not found
	if (ref.MatchType == MatchExact || ref.MatchType == MatchDefault || ref.MatchType == MatchRegexp ||
		ref.MatchType == MatchCaseInsensitive) && !ref.SkipDirectLookup {
		sr := ic.State.Get(ref.Resource, attr, matchValue)
		if sr != nil && (ref.IsValidApproximation == nil || ref.IsValidApproximation(ic, origResource, sr, origPath)) &&
			!ic.isIgnoredResourceApproximation(sr) {
			log.Printf("[DEBUG] Finished direct lookup for reference for resource %s, attr='%s', value='%s', ref=%v. Found: type=%s name=%s",
				ref.Resource, attr, value, ref, sr.Type, sr.Name)
			return matchValue, genTraversalTokens(sr, attr), sr.Mode == "data"
		}
		if ref.MatchType != MatchCaseInsensitive { // for case-insensitive matching we'll try iteration
			log.Printf("[DEBUG] Finished direct lookup for reference for resource %s, attr='%s', value='%s', ref=%v. Not found",
				ref.Resource, attr, value, ref)
			return "", nil, false
		}
	} else if ref.MatchType == MatchLongestPrefix && ref.ExtraLookupKey != "" {
		extraKeyValue, exists := origResource.GetExtraData(ref.ExtraLookupKey)
		if exists && extraKeyValue.(string) != "" {
			sr := ic.State.Get(ref.Resource, attr, extraKeyValue.(string))
			if sr != nil && (ref.IsValidApproximation == nil || ref.IsValidApproximation(ic, origResource, sr, origPath)) &&
				!ic.isIgnoredResourceApproximation(sr) {
				log.Printf("[DEBUG] Finished direct lookup by key %s for reference for resource %s, attr='%s', value='%s', ref=%v. Found: type=%s name=%s",
					ref.ExtraLookupKey, ref.Resource, attr, value, ref, sr.Type, sr.Name)
				return extraKeyValue.(string), genTraversalTokens(sr, attr), sr.Mode == "data"
			}
		}
	}

	maxPrefixLen := 0
	maxPrefixOrigValue := ""
	var maxPrefixResource *resourceApproximation
	srs := *ic.State.Resources(ref.Resource)
	for _, sr := range srs {
		for _, i := range sr.Instances {
			v := i.Attributes[attr]
			if v == nil {
				log.Printf("[WARN] Can't find instance attribute '%v' in resource: '%v'", attr, ref.Resource)
				continue
			}
			strValue := v.(string)
			origValue := strValue
			if ref.SearchValueTransformFunc != nil {
				strValue = ref.SearchValueTransformFunc(strValue)
				log.Printf("[TRACE] Resource %s. Transformed value from '%s' to '%s'", ref.Resource, origValue, strValue)
			}
			matched := false
			switch ref.MatchType {
			case MatchCaseInsensitive:
				matched = (strings.ToLower(strValue) == matchValue)
			case MatchPrefix:
				matched = strings.HasPrefix(matchValue, strValue)
			case MatchLongestPrefix:
				if strings.HasPrefix(matchValue, strValue) && len(origValue) > maxPrefixLen && !ic.isIgnoredResourceApproximation(sr) {
					maxPrefixLen = len(origValue)
					maxPrefixOrigValue = origValue
					maxPrefixResource = sr
				}
			case MatchExact, MatchDefault:
				matched = (strValue == matchValue)
			default:
				log.Printf("[WARN] Unsupported match type: %s", ref.MatchType)
			}
			if !matched || (ref.IsValidApproximation != nil && !ref.IsValidApproximation(ic, origResource, sr, origPath)) ||
				ic.isIgnoredResourceApproximation(sr) {
				continue
			}
			log.Printf("[DEBUG] Finished searching for reference for resource %s, attr='%s', value='%s', ref=%v. Found: type=%s name=%s",
				ref.Resource, attr, value, ref, sr.Type, sr.Name)
			return origValue, genTraversalTokens(sr, attr), sr.Mode == "data"
		}
	}
	if ref.MatchType == MatchLongestPrefix && maxPrefixResource != nil &&
		(ref.IsValidApproximation == nil || ref.IsValidApproximation(ic, origResource, maxPrefixResource, origPath)) &&
		!ic.isIgnoredResourceApproximation(maxPrefixResource) {
		log.Printf("[DEBUG] Finished searching longest prefix for reference for resource %s, attr='%s', value='%s', ref=%v. Found: type=%s name=%s",
			ref.Resource, attr, value, ref, maxPrefixResource.Type, maxPrefixResource.Name)
		return maxPrefixOrigValue, genTraversalTokens(maxPrefixResource, attr), maxPrefixResource.Mode == "data"
	}
	log.Printf("[DEBUG] Finished searching for reference for resource %s, pick=%s, ref=%v. Not found", ref.Resource, attr, ref)
	return "", nil, false
}

func (ic *importContext) getTraversalTokens(ref reference, value string, origResource *resource, origPath string) (hclwrite.Tokens, bool) {
	matchType := ref.MatchTypeValue()
	attr := ref.MatchAttribute()
	attrValue, traversal, isData := ic.Find(value, attr, ref, origResource, origPath)
	// at least one invocation of ic.Find will assign Nil to traversal if resource with value is not found
	if traversal == nil {
		return nil, isData
	}
	// capture if it's data?
	switch matchType {
	case MatchExact, MatchDefault, MatchCaseInsensitive:
		return hclwrite.TokensForTraversal(traversal), isData
	case MatchPrefix, MatchLongestPrefix:
		rest := value[len(attrValue):]
		tokens := hclwrite.Tokens{&hclwrite.Token{Type: hclsyntax.TokenOQuote, Bytes: []byte{'"', '$', '{'}}}
		tokens = append(tokens, hclwrite.TokensForTraversal(traversal)...)
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCQuote, Bytes: []byte{'}'}})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenQuotedLit, Bytes: []byte(maybeAddQuoteCharacter(rest))})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCQuote, Bytes: []byte{'"'}})
		return tokens, isData
	case MatchRegexp:
		indices := ref.Regexp.FindStringSubmatchIndex(value)
		if len(indices) == 4 {
			tokens := hclwrite.Tokens{&hclwrite.Token{Type: hclsyntax.TokenOQuote, Bytes: []byte{'"'}}}
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenQuotedLit, Bytes: []byte(maybeAddQuoteCharacter(value[0:indices[2]]))})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenOQuote, Bytes: []byte{'$', '{'}})
			tokens = append(tokens, hclwrite.TokensForTraversal(traversal)...)
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCQuote, Bytes: []byte{'}'}})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenQuotedLit, Bytes: []byte(maybeAddQuoteCharacter(value[indices[3]:]))})
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCQuote, Bytes: []byte{'"'}})
			return tokens, isData
		}
		log.Printf("[WARN] Can't match found data in '%s'. Indices: %v", value, indices)
	default:
		log.Printf("[WARN] Unsupported match type: %s", ref.MatchType)
	}
	return nil, false
}

func (ic *importContext) reference(i importable, path []string, value string, ctyValue cty.Value, origResource *resource) hclwrite.Tokens {
	pathString := strings.Join(path, ".")
	match := dependsRe.ReplaceAllString(pathString, "")
	// get reference candidate, but if it's a `data`, then look for another non-data reference if possible..
	var dataTokens hclwrite.Tokens
	for _, d := range i.Depends {
		if d.Path != match {
			continue
		}
		if d.File {
			relativeFile := fmt.Sprintf("${path.module}/%s", value)
			return hclwrite.Tokens{
				&hclwrite.Token{Type: hclsyntax.TokenOQuote, Bytes: []byte{'"'}},
				&hclwrite.Token{Type: hclsyntax.TokenQuotedLit, Bytes: []byte(relativeFile)},
				&hclwrite.Token{Type: hclsyntax.TokenCQuote, Bytes: []byte{'"'}},
			}
		}
		if d.Variable {
			varName := ic.generateVariableName(path[0], value)
			return ic.variable(varName, "")
		}

		tokens, isData := ic.getTraversalTokens(d, value, origResource, pathString)
		if tokens != nil {
			if isData {
				dataTokens = tokens
				log.Printf("[DEBUG] Got reference to data for dependency %v", d)
			} else {
				return tokens
			}
		}
	}
	if len(dataTokens) > 0 {
		return dataTokens
	}
	return hclwrite.TokensForValue(ctyValue)
}

func (ic *importContext) variable(name, desc string) hclwrite.Tokens {
	ic.variablesLock.Lock()
	ic.variables[name] = desc
	ic.variablesLock.Unlock()
	return hclwrite.TokensForTraversal(hcl.Traversal{
		hcl.TraverseRoot{Name: "var"},
		hcl.TraverseAttr{Name: name},
	})
}

type fieldTuple struct {
	Field  string
	Schema *schema.Schema
}

// fieldGenerationInfo holds metadata for a single field that needs to be generated
type fieldGenerationInfo struct {
	Name        string
	PathString  string
	RawValue    interface{}
	FieldSchema FieldSchema
	ShouldSkip  bool
}

// unifiedDataToHcl is the unified entry point for HCL generation that works with both SDKv2 and Plugin Framework
func (ic *importContext) unifiedDataToHcl(imp importable, path []string,
	res *resource, body *hclwrite.Body) error {

	wrapper := res.DataWrapper
	if wrapper == nil {
		// For backward compatibility, create wrapper on-the-fly for SDKv2 resources
		if res.Data != nil && !imp.PluginFramework {
			// Get the schema resource for SDKv2
			pr, ok := ic.Resources[res.Resource]
			if !ok {
				return fmt.Errorf("resource %s not found in provider", res.Resource)
			}
			wrapper = &SDKv2ResourceData{
				data:   res.Data,
				schema: pr,
			}
			res.DataWrapper = wrapper
		} else {
			return fmt.Errorf("DataWrapper is nil for resource %s", res.Resource)
		}
	}

	// Extract fields with common logic (omission, variable references, skip evaluation)
	fields, varCnt, err := ic.extractFieldsForGeneration(imp, path, wrapper, res)
	if err != nil {
		return err
	}

	// Generate HCL for each field (type-specific logic)
	for _, field := range fields {
		if field.ShouldSkip {
			continue
		}

		var genErr error
		if wrapper.IsPluginFramework() {
			genErr = ic.generatePluginFrameworkField(imp, path, field, res, body)
		} else {
			genErr = ic.generateSdkv2Field(imp, path, field, res, body, varCnt)
		}

		if genErr != nil {
			log.Printf("[WARN] Error generating HCL for field %s: %v", field.PathString, genErr)
		}
	}

	// Generate depends_on (shared logic)
	if len(path) == 0 && len(res.DependsOn) > 0 {
		ic.generateDependsOnAttribute(res, body, wrapper)
	}

	return nil
}

// extractFieldsForGeneration extracts all fields that need to be generated with common omission and skip logic
func (ic *importContext) extractFieldsForGeneration(imp importable, path []string,
	wrapper ResourceDataWrapper, res *resource) ([]fieldGenerationInfo, int, error) {

	schema := wrapper.GetSchema()
	fieldNames := schema.GetFields()

	// Sort fields in reverse order for consistent output
	sort.Slice(fieldNames, func(i, j int) bool {
		return fieldNames[i] > fieldNames[j]
	})

	fields := []fieldGenerationInfo{}
	varCnt := 0

	for _, fieldName := range fieldNames {
		pathString := strings.Join(append(path, fieldName), ".")
		raw, nonZero := wrapper.GetOk(fieldName)

		fieldSchema := schema.GetField(fieldName)
		if fieldSchema == nil {
			continue
		}

		// Apply field omission logic with unified callbacks
		shouldOmit := false
		if imp.ShouldOmitFieldUnified != nil {
			shouldOmit = imp.ShouldOmitFieldUnified(ic, pathString, fieldSchema, wrapper, res)
		} else if imp.ShouldOmitField != nil && !wrapper.IsPluginFramework() {
			// Fall back to legacy callback for SDKv2 only
			if sdkSchema := fieldSchema.GetSDKv2Schema(); sdkSchema != nil {
				shouldOmit = imp.ShouldOmitField(ic, pathString, sdkSchema, res.Data, res)
			}
		} else {
			// Use default omission logic (match original SDKv2 behavior)
			shouldOmit = DefaultShouldOmitFieldFuncWithAbstraction(ic, pathString, fieldSchema, wrapper, res)
		}

		if shouldOmit {
			continue
		}

		// Handle variable references for sensitive fields
		mpath := dependsRe.ReplaceAllString(pathString, "")
		for _, ref := range imp.Depends {
			if ref.Path == mpath && ref.Variable {
				raw = ic.regexFix(ic.ResourceName(res), simpleNameFixes)
				if varCnt > 0 {
					raw = fmt.Sprintf("%s_%d", raw, varCnt)
				}
				nonZero = true
				varCnt++
			}
		}

		// Determine if we should skip this field
		shouldSkip := !nonZero
		if fieldSchema.IsRequired() {
			shouldSkip = false
		} else if def := fieldSchema.GetDefault(); def != nil && !reflect.DeepEqual(raw, def) {
			shouldSkip = false
		}

		// For Plugin Framework, also check for zero values in primitives
		if !shouldSkip && wrapper.IsPluginFramework() && nonZero && fieldSchema.IsOptional() {
			rv := reflect.ValueOf(raw)
			if rv.IsValid() && rv.IsZero() {
				shouldSkip = true
			}
		}

		// Check if ShouldGenerateField forces generation
		if shouldSkip {
			forceGenerate := false
			if imp.ShouldGenerateFieldUnified != nil {
				forceGenerate = imp.ShouldGenerateFieldUnified(ic, pathString, fieldSchema, wrapper, res)
			} else if imp.ShouldGenerateField != nil && !wrapper.IsPluginFramework() {
				if sdkSchema := fieldSchema.GetSDKv2Schema(); sdkSchema != nil {
					forceGenerate = imp.ShouldGenerateField(ic, pathString, sdkSchema, res.Data, res)
				}
			}

			if !forceGenerate {
				// Don't add to list at all
				continue
			}
			// Force generation by not skipping
			shouldSkip = false
		}

		fields = append(fields, fieldGenerationInfo{
			Name:        fieldName,
			PathString:  pathString,
			RawValue:    raw,
			FieldSchema: fieldSchema,
			ShouldSkip:  shouldSkip,
		})
	}

	return fields, varCnt, nil
}

// generateDependsOnAttribute generates the depends_on attribute (shared between SDKv2 and Plugin Framework)
func (ic *importContext) generateDependsOnAttribute(res *resource, body *hclwrite.Body, wrapper ResourceDataWrapper) {
	notIgnoredResources := []*resource{}
	for _, dr := range res.DependsOn {
		dr := dr
		if dr.Data == nil && dr.DataWrapper == nil {
			tdr := ic.Scope.FindById(dr.Resource, dr.ID)
			if tdr == nil {
				log.Printf("[WARN] can't find resource %s in scope", dr)
				continue
			}
			dr = tdr
		}
		if ic.Importables[dr.Resource].Ignore == nil || !ic.Importables[dr.Resource].Ignore(ic, dr) {
			found := false
			for _, v := range notIgnoredResources {
				if v.ID == dr.ID && v.Resource == dr.Resource {
					found = true
					break
				}
			}
			if !found {
				notIgnoredResources = append(notIgnoredResources, dr)
			}
		}
	}

	if len(notIgnoredResources) > 0 {
		toks := hclwrite.Tokens{}
		toks = append(toks, &hclwrite.Token{
			Type:  hclsyntax.TokenOBrack,
			Bytes: []byte{'['},
		})
		for i, dr := range notIgnoredResources {
			if i > 0 {
				toks = append(toks, &hclwrite.Token{
					Type:  hclsyntax.TokenComma,
					Bytes: []byte{','},
				})
			}
			toks = append(toks, hclwrite.TokensForTraversal(hcl.Traversal{
				hcl.TraverseRoot{Name: dr.Resource},
				hcl.TraverseAttr{Name: ic.ResourceName(dr)},
			})...)
		}
		toks = append(toks, &hclwrite.Token{
			Type:  hclsyntax.TokenCBrack,
			Bytes: []byte{']'},
		})
		body.SetAttributeRaw("depends_on", toks)
	}
}

// generateSdkv2Field generates HCL for a single SDKv2 field
func (ic *importContext) generateSdkv2Field(imp importable, path []string,
	field fieldGenerationInfo, res *resource, body *hclwrite.Body, varCnt int) error {

	sdkSchema := field.FieldSchema.GetSDKv2Schema()
	if sdkSchema == nil {
		return fmt.Errorf("SDKv2 schema is nil for field %s", field.Name)
	}

	switch sdkSchema.Type {
	case schema.TypeString:
		value := field.RawValue.(string)
		tokens := ic.reference(imp, append(path, field.Name), value, cty.StringVal(value), res)
		body.SetAttributeRaw(field.Name, tokens)
	case schema.TypeBool:
		body.SetAttributeValue(field.Name, cty.BoolVal(field.RawValue.(bool)))
	case schema.TypeInt:
		var num int64
		switch iv := field.RawValue.(type) {
		case int:
			num = int64(iv)
		case int32:
			num = int64(iv)
		case int64:
			num = iv
		}
		body.SetAttributeRaw(field.Name, ic.reference(imp, append(path, field.Name),
			strconv.FormatInt(num, 10), cty.NumberIntVal(num), res))
	case schema.TypeFloat:
		body.SetAttributeValue(field.Name, cty.NumberFloatVal(field.RawValue.(float64)))
	case schema.TypeMap:
		attrs := []hclwrite.ObjectAttrTokens{}
		m := field.RawValue.(map[string]any)
		keys := slices.Sorted(maps.Keys(m))
		for _, key := range keys {
			iv := m[key]
			var nameTokens hclwrite.Tokens
			if hclsyntax.ValidIdentifier(key) {
				nameTokens = hclwrite.TokensForIdentifier(key)
			} else {
				nameTokens = hclwrite.TokensForValue(cty.StringVal(key))
			}
			v := cty.StringVal(fmt.Sprintf("%v", iv))
			attrs = append(attrs, hclwrite.ObjectAttrTokens{
				Name:  nameTokens,
				Value: ic.reference(imp, append(path, field.Name), v.AsString(), v, res),
			})
		}
		body.SetAttributeRaw(field.Name, hclwrite.TokensForObject(attrs))
	case schema.TypeSet:
		if rawSet, ok := field.RawValue.(*schema.Set); ok {
			rawList := rawSet.List()
			err := ic.readListFromData(imp, append(path, field.Name), res, rawList, body, sdkSchema, func(i int) string {
				return strconv.Itoa(rawSet.F(rawList[i]))
			})
			if err != nil {
				return err
			}
		}
	case schema.TypeList:
		if rawList, ok := field.RawValue.([]any); ok {
			err := ic.readListFromData(imp, append(path, field.Name), res, rawList, body, sdkSchema, strconv.Itoa)
			if err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported schema type: %v", path)
	}

	return nil
}

// generatePluginFrameworkField generates HCL for a single Plugin Framework field
func (ic *importContext) generatePluginFrameworkField(imp importable, path []string,
	field fieldGenerationInfo, res *resource, body *hclwrite.Body) error {

	return ic.pluginFrameworkFieldToHcl(imp, path, field.Name, field.FieldSchema, field.RawValue, res, body)
}

func (ic *importContext) dataToHcl(imp importable, path []string,
	pr *schema.Resource, res *resource, body *hclwrite.Body) error {
	d := res.Data
	ss := []fieldTuple{}
	for a, as := range pr.Schema {
		ss = append(ss, fieldTuple{a, as})
	}
	sort.Slice(ss, func(i, j int) bool {
		// it just happens that reverse field order
		// makes the most beautiful configs
		return ss[i].Field > ss[j].Field
	})
	var_cnt := 0
	for _, tuple := range ss {
		a, as := tuple.Field, tuple.Schema
		pathString := strings.Join(append(path, a), ".")
		raw, nonZero := d.GetOk(pathString)
		// log.Printf("[DEBUG] path=%s, raw='%v'", pathString, raw)
		if imp.ShouldOmitField == nil { // we don't have custom function, so skip computed & default fields
			if defaultShouldOmitFieldFunc(ic, pathString, as, d, res) {
				continue
			}
		} else if imp.ShouldOmitField(ic, pathString, as, d, res) {
			continue
		}
		mpath := dependsRe.ReplaceAllString(pathString, "")
		for _, ref := range imp.Depends {
			if ref.Path == mpath && ref.Variable {
				// sensitive fields are moved to variable depends, variable name is normalized
				// TODO: handle a case when we have multiple blocks, so names won't be unique
				raw = ic.regexFix(ic.ResourceName(res), simpleNameFixes)
				if var_cnt > 0 {
					raw = fmt.Sprintf("%s_%d", raw, var_cnt)
				}
				nonZero = true
				var_cnt++
			}
		}
		shouldSkip := !nonZero
		if as.Required { // for required fields we must produce a value, even empty...
			shouldSkip = false
		} else if as.Default != nil && !reflect.DeepEqual(raw, as.Default) {
			// In case when have zero value, but there is non-zero default, we also need to produce it
			shouldSkip = false
		}
		if shouldSkip && (imp.ShouldGenerateField == nil || !imp.ShouldGenerateField(ic, pathString, as, d, res)) {
			continue
		}
		switch as.Type {
		case schema.TypeString:
			value := raw.(string)
			tokens := ic.reference(imp, append(path, a), value, cty.StringVal(value), res)
			body.SetAttributeRaw(a, tokens)
		case schema.TypeBool:
			body.SetAttributeValue(a, cty.BoolVal(raw.(bool)))
		case schema.TypeInt:
			var num int64
			switch iv := raw.(type) {
			case int:
				num = int64(iv)
			case int32:
				num = int64(iv)
			case int64:
				num = iv
			}
			body.SetAttributeRaw(a, ic.reference(imp, append(path, a),
				strconv.FormatInt(num, 10), cty.NumberIntVal(num), res))
		case schema.TypeFloat:
			body.SetAttributeValue(a, cty.NumberFloatVal(raw.(float64)))
		case schema.TypeMap:
			// Resolve references in maps as well, and also support different types inside map...
			attrs := []hclwrite.ObjectAttrTokens{}
			m := raw.(map[string]any)
			keys := slices.Sorted(maps.Keys(m))
			for _, key := range keys {
				iv := m[key]
				var nameTokens hclwrite.Tokens
				if hclsyntax.ValidIdentifier(key) {
					nameTokens = hclwrite.TokensForIdentifier(key)
				} else {
					nameTokens = hclwrite.TokensForValue(cty.StringVal(key))
				}
				v := cty.StringVal(fmt.Sprintf("%v", iv)) // TODO: support different types inside map...
				attrs = append(attrs, hclwrite.ObjectAttrTokens{
					Name:  nameTokens,
					Value: ic.reference(imp, append(path, a), v.AsString(), v, res),
				})
			}
			body.SetAttributeRaw(a, hclwrite.TokensForObject(attrs))
		case schema.TypeSet:
			if rawSet, ok := raw.(*schema.Set); ok {
				rawList := rawSet.List()
				err := ic.readListFromData(imp, append(path, a), res, rawList, body, as, func(i int) string {
					return strconv.Itoa(rawSet.F(rawList[i]))
				})
				if err != nil {
					return err
				}
			}
		case schema.TypeList:
			if rawList, ok := raw.([]any); ok {
				err := ic.readListFromData(imp, append(path, a), res, rawList, body, as, strconv.Itoa)
				if err != nil {
					return err
				}
			}
		default:
			return fmt.Errorf("unsupported schema type: %v", path)
		}
	}
	// Generate `depends_on` only for top-level resource because `dataToHcl` is called recursively
	if len(path) == 0 && len(res.DependsOn) > 0 {
		notIgnoredResources := []*resource{}
		for _, dr := range res.DependsOn {
			dr := dr
			if dr.Data == nil {
				tdr := ic.Scope.FindById(dr.Resource, dr.ID)
				if tdr == nil {
					log.Printf("[WARN] can't find resource %s in scope", dr)
					continue
				}
				dr = tdr
			}
			if ic.Importables[dr.Resource].Ignore == nil || !ic.Importables[dr.Resource].Ignore(ic, dr) {
				found := false
				for _, v := range notIgnoredResources {
					if v.ID == dr.ID && v.Resource == dr.Resource {
						found = true
						break
					}
				}
				if !found {
					notIgnoredResources = append(notIgnoredResources, dr)
				}
			}
		}
		if len(notIgnoredResources) > 0 {
			toks := hclwrite.Tokens{}
			toks = append(toks, &hclwrite.Token{
				Type:  hclsyntax.TokenOBrack,
				Bytes: []byte{'['},
			})
			for i, dr := range notIgnoredResources {
				if i > 0 {
					toks = append(toks, &hclwrite.Token{
						Type:  hclsyntax.TokenComma,
						Bytes: []byte{','},
					})
				}
				toks = append(toks, hclwrite.TokensForTraversal(hcl.Traversal{
					hcl.TraverseRoot{Name: dr.Resource},
					hcl.TraverseAttr{Name: ic.ResourceName(dr)},
				})...)
			}
			toks = append(toks, &hclwrite.Token{
				Type:  hclsyntax.TokenCBrack,
				Bytes: []byte{']'},
			})
			body.SetAttributeRaw("depends_on", toks)
		}
	}
	return nil
}

// pluginFrameworkFieldToHcl generates HCL for a single Plugin Framework field
func (ic *importContext) pluginFrameworkFieldToHcl(imp importable, path []string, fieldName string,
	fieldSchema FieldSchema, raw interface{}, res *resource, body *hclwrite.Body) error {

	// Check for nested types first (before checking primitive types)
	// This handles ListNestedAttribute, SetNestedAttribute, SingleNestedAttribute, etc.
	if fieldSchema.IsNested() {
		// Check if this is a list/set of nested objects
		if rawList, ok := raw.([]interface{}); ok {
			if len(rawList) == 0 {
				return nil
			}

			nestedSchema := fieldSchema.GetNestedSchema()
			if nestedSchema != nil {
				// Generate list/set of objects as attribute: field = [{ ... }, { ... }]
				listTokens := hclwrite.Tokens{}
				listTokens = append(listTokens, &hclwrite.Token{
					Type:  hclsyntax.TokenOBrack,
					Bytes: []byte{'['},
				})

				for i, item := range rawList {
					if i > 0 {
						listTokens = append(listTokens, &hclwrite.Token{
							Type:  hclsyntax.TokenComma,
							Bytes: []byte{','},
						})
						listTokens = append(listTokens, &hclwrite.Token{
							Type:  hclsyntax.TokenNewline,
							Bytes: []byte{'\n'},
						})
					}

					if nestedData, ok := item.(map[string]interface{}); ok {
						// Include the index in the path for proper reference resolution
						nestedPath := append(path, fieldName, strconv.Itoa(i))
						objTokens := ic.pluginFrameworkNestedObjectToTokens(imp, nestedPath, nestedSchema, nestedData, res)
						listTokens = append(listTokens, objTokens...)
					}
				}

				listTokens = append(listTokens, &hclwrite.Token{
					Type:  hclsyntax.TokenCBrack,
					Bytes: []byte{']'},
				})
				body.SetAttributeRaw(fieldName, listTokens)
				return nil
			}
		}

		// Single nested object (not in a list)
		if nestedData, ok := raw.(map[string]interface{}); ok {
			nestedSchema := fieldSchema.GetNestedSchema()
			if nestedSchema != nil {
				// Generate single nested object as attribute: field = { ... }
				objTokens := ic.pluginFrameworkNestedObjectToTokens(imp, append(path, fieldName), nestedSchema, nestedData, res)
				body.SetAttributeRaw(fieldName, objTokens)
				return nil
			}
		}

		log.Printf("[DEBUG] Nested field %s could not be processed, raw type: %T", fieldName, raw)
		return nil
	}

	// Now handle primitive types
	switch {
	case fieldSchema.IsString():
		value, ok := raw.(string)
		if !ok {
			return nil
		}
		tokens := ic.reference(imp, append(path, fieldName), value, cty.StringVal(value), res)
		body.SetAttributeRaw(fieldName, tokens)

	case fieldSchema.IsBool():
		value, ok := raw.(bool)
		if !ok {
			return nil
		}
		body.SetAttributeValue(fieldName, cty.BoolVal(value))

	case fieldSchema.IsInt():
		var num int64
		switch iv := raw.(type) {
		case int:
			num = int64(iv)
		case int32:
			num = int64(iv)
		case int64:
			num = iv
		default:
			return nil
		}
		body.SetAttributeRaw(fieldName, ic.reference(imp, append(path, fieldName),
			strconv.FormatInt(num, 10), cty.NumberIntVal(num), res))

	case fieldSchema.IsFloat():
		value, ok := raw.(float64)
		if !ok {
			return nil
		}
		body.SetAttributeValue(fieldName, cty.NumberFloatVal(value))

	case fieldSchema.IsMap():
		// Handle maps
		m, ok := raw.(map[string]interface{})
		if !ok {
			return nil
		}
		attrs := []hclwrite.ObjectAttrTokens{}
		keys := slices.Sorted(maps.Keys(m))
		for _, key := range keys {
			iv := m[key]
			var nameTokens hclwrite.Tokens
			if hclsyntax.ValidIdentifier(key) {
				nameTokens = hclwrite.TokensForIdentifier(key)
			} else {
				nameTokens = hclwrite.TokensForValue(cty.StringVal(key))
			}
			v := cty.StringVal(fmt.Sprintf("%v", iv))
			attrs = append(attrs, hclwrite.ObjectAttrTokens{
				Name:  nameTokens,
				Value: ic.reference(imp, append(path, fieldName), v.AsString(), v, res),
			})
		}
		body.SetAttributeRaw(fieldName, hclwrite.TokensForObject(attrs))

	case fieldSchema.IsList() || fieldSchema.IsSet():
		// Handle lists and sets of primitives (nested types are handled above)
		rawList, ok := raw.([]interface{})
		if !ok {
			return nil
		}
		if len(rawList) == 0 {
			return nil
		}

		// Generate list/set of primitives
		toks := hclwrite.Tokens{}
		toks = append(toks, &hclwrite.Token{
			Type:  hclsyntax.TokenOBrack,
			Bytes: []byte{'['},
		})
		for _, item := range rawList {
			if len(toks) != 1 {
				toks = append(toks, &hclwrite.Token{
					Type:  hclsyntax.TokenComma,
					Bytes: []byte{','},
				})
			}
			switch x := item.(type) {
			case string:
				toks = append(toks, ic.reference(imp, append(path, fieldName), x, cty.StringVal(x), res)...)
			case int, int32, int64:
				var num int64
				switch v := x.(type) {
				case int:
					num = int64(v)
				case int32:
					num = int64(v)
				case int64:
					num = v
				}
				toks = append(toks, hclwrite.TokensForValue(cty.NumberIntVal(num))...)
			default:
				// For complex types, convert to string
				toks = append(toks, hclwrite.TokensForValue(cty.StringVal(fmt.Sprintf("%v", x)))...)
			}
		}
		toks = append(toks, &hclwrite.Token{
			Type:  hclsyntax.TokenCBrack,
			Bytes: []byte{']'},
		})
		body.SetAttributeRaw(fieldName, toks)

	default:
		log.Printf("[DEBUG] Unsupported field type for %s in Plugin Framework resource (type: %T, fieldType: %v)", fieldName, raw, fieldSchema.GetType())
	}

	return nil
}

// pluginFrameworkNestedObjectToTokens generates HCL tokens for a nested object
func (ic *importContext) pluginFrameworkNestedObjectToTokens(imp importable, path []string,
	nestedSchema SchemaWrapper, nestedData map[string]interface{}, res *resource) hclwrite.Tokens {

	if nestedSchema == nil || nestedData == nil {
		return hclwrite.Tokens{}
	}

	// Create a temporary body to generate the nested attributes
	tempFile := hclwrite.NewEmptyFile()
	tempBody := tempFile.Body()

	fields := nestedSchema.GetFields()
	// Sort fields for consistent output (reverse order like SDKv2)
	sort.Slice(fields, func(i, j int) bool {
		return fields[i] > fields[j]
	})

	for _, fieldName := range fields {
		raw, ok := nestedData[fieldName]
		nonZero := ok && raw != nil

		fieldSchema := nestedSchema.GetField(fieldName)
		if fieldSchema == nil {
			log.Printf("[DEBUG] No schema found for field %s in path %v", fieldName, path)
			continue
		}

		// Apply field omission logic - check computed-only fields
		if fieldSchema.IsComputed() && !fieldSchema.IsOptional() && !fieldSchema.IsRequired() {
			log.Printf("[DEBUG] Skipping computed-only field %s in path %v", fieldName, path)
			continue
		}

		// Skip fields with zero values unless required or different from default
		shouldSkip := !nonZero
		if fieldSchema.IsRequired() {
			// For required fields we must produce a value, even empty
			shouldSkip = false
		} else if def := fieldSchema.GetDefault(); def != nil && !reflect.DeepEqual(raw, def) {
			// If field has a default value and current value is different, we need to generate it
			shouldSkip = false
		}

		// Check for zero values in primitives
		if !shouldSkip && nonZero {
			switch v := raw.(type) {
			case bool:
				if !v && !fieldSchema.IsRequired() {
					// Skip false for optional boolean fields
					shouldSkip = true
				}
			case int, int32, int64:
				var num int64
				switch n := v.(type) {
				case int:
					num = int64(n)
				case int32:
					num = int64(n)
				case int64:
					num = n
				}
				if num == 0 && !fieldSchema.IsRequired() {
					// Skip 0 for optional numeric fields
					shouldSkip = true
				}
			case float32, float64:
				var fnum float64
				switch n := v.(type) {
				case float32:
					fnum = float64(n)
				case float64:
					fnum = n
				}
				if fnum == 0 && !fieldSchema.IsRequired() {
					// Skip 0.0 for optional float fields
					shouldSkip = true
				}
			case string:
				if v == "" && !fieldSchema.IsRequired() {
					// Skip empty strings for optional string fields
					shouldSkip = true
				}
			}
		}

		// Check if ShouldGenerateField forces generation
		pathString := strings.Join(append(path, fieldName), ".")
		if shouldSkip && (imp.ShouldGenerateField == nil || !imp.ShouldGenerateField(ic, pathString, nil, nil, res)) {
			log.Printf("[DEBUG] Skipping field %s with zero value in path %v", fieldName, path)
			continue
		}

		log.Printf("[DEBUG] Processing field %s in path %v, value type: %T", fieldName, path, raw)

		// Generate HCL for this nested field
		if err := ic.pluginFrameworkFieldToHcl(imp, path, fieldName, fieldSchema, raw, res, tempBody); err != nil {
			log.Printf("[WARN] Error generating HCL for nested field %s: %v", fieldName, err)
		}
	}

	// Extract the attributes from the temp body and convert to object tokens
	attrs := []hclwrite.ObjectAttrTokens{}
	for name, attr := range tempBody.Attributes() {
		attrs = append(attrs, hclwrite.ObjectAttrTokens{
			Name:  hclwrite.TokensForIdentifier(name),
			Value: attr.Expr().BuildTokens(nil),
		})
	}

	// Also include any nested blocks (though for Plugin Framework these should be rare)
	// Convert nested blocks to attribute syntax
	for _, block := range tempBody.Blocks() {
		blockName := block.Type()
		blockBody := block.Body()

		// Recursively convert block body to object tokens
		blockAttrs := []hclwrite.ObjectAttrTokens{}
		for name, attr := range blockBody.Attributes() {
			blockAttrs = append(blockAttrs, hclwrite.ObjectAttrTokens{
				Name:  hclwrite.TokensForIdentifier(name),
				Value: attr.Expr().BuildTokens(nil),
			})
		}

		attrs = append(attrs, hclwrite.ObjectAttrTokens{
			Name:  hclwrite.TokensForIdentifier(blockName),
			Value: hclwrite.TokensForObject(blockAttrs),
		})
	}

	return hclwrite.TokensForObject(attrs)
}

func (ic *importContext) readListFromData(i importable, path []string, res *resource,
	rawList []any, body *hclwrite.Body, as *schema.Schema, offsetConverter func(i int) string) error {
	if len(rawList) == 0 {
		return nil
	}
	name := path[len(path)-1]
	switch elem := as.Elem.(type) {
	case *schema.Resource:
		if as.MaxItems == 1 {
			nestedPath := append(path, offsetConverter(0))
			confBlock := body.AppendNewBlock(name, []string{})
			return ic.dataToHcl(i, nestedPath, elem, res, confBlock.Body())
		}
		for offset := range rawList {
			confBlock := body.AppendNewBlock(name, []string{})
			nestedPath := append(path, offsetConverter(offset))
			err := ic.dataToHcl(i, nestedPath, elem, res, confBlock.Body())
			if err != nil {
				return err
			}
		}
	case *schema.Schema:
		toks := hclwrite.Tokens{}
		toks = append(toks, &hclwrite.Token{
			Type:  hclsyntax.TokenOBrack,
			Bytes: []byte{'['},
		})
		for _, raw := range rawList {
			if len(toks) != 1 {
				toks = append(toks, &hclwrite.Token{
					Type:  hclsyntax.TokenComma,
					Bytes: []byte{','},
				})
			}
			switch x := raw.(type) {
			case string:
				value := raw.(string)
				toks = append(toks, ic.reference(i, path, value, cty.StringVal(value), res)...)
			case int:
				// probably we don't even use integer lists?...
				toks = append(toks, hclwrite.TokensForValue(
					cty.NumberIntVal(int64(x)))...)
			default:
				return fmt.Errorf("unsupported primitive list: %#v", path)
			}
		}
		toks = append(toks, &hclwrite.Token{
			Type:  hclsyntax.TokenCBrack,
			Bytes: []byte{']'},
		})
		body.SetAttributeRaw(name, toks)
	}
	return nil
}

func (ic *importContext) generateTfvars() error {
	// TODO: make it incremental as well...
	if len(ic.tfvars) == 0 {
		return nil
	}
	f := hclwrite.NewEmptyFile()
	body := f.Body()
	fileName := fmt.Sprintf("%s/terraform.tfvars", ic.Directory)

	vf, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer vf.Close()

	for k, v := range ic.tfvars {
		body.SetAttributeValue(k, cty.StringVal(v))
	}
	// nolint
	vf.Write(f.Bytes())
	log.Printf("[INFO] Written %d tfvars", len(ic.tfvars))

	ic.generateGitIgnore()

	return nil
}

func (ic *importContext) generateVariables() error {
	if len(ic.variables) == 0 {
		return nil
	}
	f := hclwrite.NewEmptyFile()
	body := f.Body()
	fileName := fmt.Sprintf("%s/vars.tf", ic.Directory)
	if ic.incremental {
		content, err := os.ReadFile(fileName)
		if err == nil {
			ftmp, diags := hclwrite.ParseConfig(content, fileName, hcl.Pos{Line: 1, Column: 1})
			if diags.HasErrors() {
				log.Printf("[ERROR] parsing of existing file failed: %s", diags)
			} else {
				tbody := ftmp.Body()
				for _, block := range tbody.Blocks() {
					typ := block.Type()
					labels := block.Labels()
					_, present := ic.variables[labels[0]]
					if typ == "variable" && present {
						log.Printf("[DEBUG] Ignoring variable '%s' that will be re-exported", labels[0])
					} else {
						log.Printf("[DEBUG] Adding not exported object. type='%s', labels=%v", typ, labels)
						body.AppendBlock(block)
					}
				}
			}
		} else {
			log.Printf("[ERROR] opening file %s", fileName)
		}
	}
	vf, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer vf.Close()

	for k, v := range ic.variables {
		b := body.AppendNewBlock("variable", []string{k}).Body()
		b.SetAttributeValue("description", cty.StringVal(v))
	}
	// nolint
	vf.Write(f.Bytes())
	log.Printf("[INFO] Written %d variables", len(ic.variables))
	return nil
}

func (ic *importContext) generateGitIgnore() {
	fileName := fmt.Sprintf("%s/.gitignore", ic.Directory)
	vf, err := os.Create(fileName)
	if err != nil {
		log.Printf("[ERROR] can't create %s: %v", fileName, err)
		return
	}
	defer vf.Close()
	// nolint
	vf.Write([]byte("terraform.tfvars\n"))
}

func (ic *importContext) generateAndWriteResources(sh *os.File) {
	resources := ic.Scope.Sorted()
	scopeSize := ic.Scope.Len()
	t1 := time.Now()
	log.Printf("[INFO] Generating configuration for %d resources", scopeSize)

	// make configurable via environment variables
	resourceHandlersNumber := getEnvAsInt("EXPORTER_RESOURCE_GENERATORS", 50)
	resourcesChan := make(resourceChannel, defaultChannelSize)

	resourceWriters := make(map[string]dataWriteChannel, len(ic.Resources))
	for service := range ic.services {
		resourceWriters[service] = make(dataWriteChannel, defaultChannelSize)
	}
	writersWaitGroup := &sync.WaitGroup{}
	// write shell script for importing
	shellImportChan := make(importWriteChannel, defaultChannelSize)
	writersWaitGroup.Add(1)
	go func() {
		ic.writeShellImports(sh, shellImportChan)
		writersWaitGroup.Done()
	}()
	//
	nativeImportChan := make(importWriteChannel, defaultChannelSize)
	writersWaitGroup.Add(1)
	go func() {
		ic.writeNativeImports(nativeImportChan)
		writersWaitGroup.Done()
	}()
	// start resource handlers
	for i := 0; i < resourceHandlersNumber; i++ {
		i := i
		go func() {
			log.Printf("[DEBUG] Starting resource handler %d", i)
			ic.processSingleResource(resourcesChan, resourceWriters, nativeImportChan)
		}()
	}
	// start writers for specific services
	for service, ch := range resourceWriters {
		service := service
		ch := ch
		generatedFile := fmt.Sprintf("%s/%s.tf", ic.Directory, service)
		log.Printf("[DEBUG] starting writer for service %s", service)
		writersWaitGroup.Add(1)
		go func() {
			ic.handleResourceWrite(generatedFile, ch, shellImportChan)
			writersWaitGroup.Done()
		}()
	}
	// submit all extracted resources...
	for i, r := range resources {
		ic.waitGroup.Add(1)
		resourcesChan <- r
		if i%500 == 0 {
			log.Printf("[INFO] Submitted %d of %d resources", i+1, scopeSize)
		}
	}
	ic.waitGroup.Wait()
	// close all channels
	close(shellImportChan)
	close(nativeImportChan)
	close(resourcesChan)
	for service, ch := range resourceWriters {
		log.Printf("Closing writer for service %s", service)
		close(ch)
	}
	writersWaitGroup.Wait()

	log.Printf("[INFO] Finished generation of configuration for %d resources (took %v seconds)",
		scopeSize, time.Since(t1).Seconds())
}

func (ic *importContext) processSingleResource(resourcesChan resourceChannel,
	writerChannels map[string]dataWriteChannel, nativeImportChannel importWriteChannel) {
	processed := 0
	generated := 0
	ignored := 0
	for r := range resourcesChan {
		processed = processed + 1
		if r == nil {
			log.Print("[WARN] Got nil resource...")
			ic.waitGroup.Done()
			continue
		}
		ir := ic.Importables[r.Resource]
		if ir.Ignore != nil && ir.Ignore(ic, r) {
			log.Printf("[WARN] Ignoring resource %s: %s", r.Resource, r.Name)
			ignored = ignored + 1
			ic.waitGroup.Done()
			continue
		}
		var err error
		f := hclwrite.NewEmptyFile()
		log.Printf("[TRACE] Generating %s: %s", r.Resource, r.Name)
		body := f.Body()
		if ir.Body != nil {
			err = ir.Body(ic, body, r)
			if err != nil {
				log.Printf("[ERROR] error calling ir.Body for %v: %s", r, err.Error())
			}
		} else {
			blockType := "resource"
			if r.Mode == "data" {
				blockType = r.Mode
			}
			resourceBlock := body.AppendNewBlock(blockType, []string{r.Resource, r.Name})

			// Use unified HCL generation for both SDKv2 and Plugin Framework
			err = ic.unifiedDataToHcl(ic.Importables[r.Resource],
				[]string{}, r, resourceBlock.Body())
			if err != nil {
				log.Printf("[ERROR] error generating body for %v: %s", r, err.Error())
			}
		}
		if err == nil && len(body.Blocks()) > 0 {
			formatted := hclwrite.Format(f.Bytes())
			// fix some formatting in a hacky way instead of writing 100 lines of HCL AST writer code
			formatted = []byte(ic.regexFix(string(formatted), ic.hclFixes))
			writeData := &resourceWriteData{
				ResourceBody: string(formatted),
				BlockName:    generateBlockFullName(body.Blocks()[0]),
			}
			// Check if resource supports import
			// For Plugin Framework resources, import is always supported
			// For SDKv2 resources, check if Importer is defined
			supportsImport := false
			if ir.PluginFramework {
				supportsImport = true
			} else if sdkResource, ok := ic.Resources[r.Resource]; ok && sdkResource.Importer != nil {
				supportsImport = true
			}
			if r.Mode != "data" && supportsImport {
				writeData.ImportCommand = r.ImportCommand(ic)
				if ic.nativeImportSupported { // generate import block for native import
					imp := hclwrite.NewEmptyFile()
					imoBlock := imp.Body().AppendNewBlock("import", []string{})
					imoBlock.Body().SetAttributeValue("id", cty.StringVal(r.ID))
					traversal := hcl.Traversal{
						hcl.TraverseRoot{Name: r.Resource},
						hcl.TraverseAttr{Name: r.Name},
					}
					tokens := hclwrite.TokensForTraversal(traversal)
					imoBlock.Body().SetAttributeRaw("to", tokens)
					formattedImp := hclwrite.Format(imp.Bytes())
					//log.Printf("[DEBUG] Import block for %s: %s", r.ID, string(formattedImp))
					ic.waitGroup.Add(1)
					nativeImportChannel <- string(formattedImp)
				}
			}
			ch, exists := writerChannels[ir.Service]
			if exists {
				ic.waitGroup.Add(1)
				ch <- writeData
			} else {
				log.Printf("[WARN] can't find a channel for service: %s, resource: %s", ir.Service, r.Resource)
			}
			log.Printf("[TRACE] Finished generating %s: %s", r.Resource, r.Name)
			generated = generated + 1
		} else {
			log.Printf("[WARN] error generating resource body: %v, or body blocks len is 0", err)
		}
		ic.waitGroup.Done()
	}
	log.Printf("[DEBUG] processed resources: %d, generated: %d, ignored: %d", processed, generated, ignored)
}

func extractResourceIdFromImportBlock(block *hclwrite.Block) string {
	if block.Type() != "import" {
		log.Print("[WARN] it's not an import block!")
		return ""
	}
	idAttr := block.Body().GetAttribute("to")
	if idAttr == nil {
		log.Printf("[WARN] Can't find `to` attribute in the import block")
		return ""
	}
	idVal := string(idAttr.Expr().BuildTokens(nil).Bytes())
	return strings.TrimSpace(idVal)
}

func extractResourceIdFromImportBlockString(importBlock string) string {
	block, diags := hclwrite.ParseConfig([]byte(importBlock), "test.tf", hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		log.Printf("[WARN] parsing of import block %s has failed: %s", importBlock, diags.Error())
		return ""
	}
	if len(block.Body().Blocks()) == 0 {
		log.Printf("[WARN] import block %s has 0 blocks!", importBlock)
		return ""
	}
	return extractResourceIdFromImportBlock(block.Body().Blocks()[0])
}

func (ic *importContext) writeNativeImports(importChan importWriteChannel) {
	if !ic.nativeImportSupported {
		log.Print("[DEBUG] Native import is not enabled, skipping...")
		return
	}
	importsFileName := fmt.Sprintf("%s/import.tf", ic.Directory)
	// TODO: in incremental mode read existing file with imports and append them for not processed & not deleted resources
	var existingFile *hclwrite.File
	if ic.incremental {
		log.Printf("[DEBUG] Going to read existing file %s", importsFileName)
		content, err := os.ReadFile(importsFileName)
		if errors.Is(err, os.ErrNotExist) {
			log.Printf("[WARN] File %s doesn't exist when using incremental export", importsFileName)
		} else if err != nil {
			log.Printf("[ERROR] error opening %s", importsFileName)
		} else {
			log.Printf("[DEBUG] Going to parse existing file %s", importsFileName)
			var diags hcl.Diagnostics
			existingFile, diags = hclwrite.ParseConfig(content, importsFileName, hcl.Pos{Line: 1, Column: 1})
			if diags.HasErrors() {
				log.Printf("[ERROR] parsing of existing file %s failed: %s", importsFileName, diags.Error())
			} else {
				log.Printf("[DEBUG] There are %d objects in existing file %s",
					len(existingFile.Body().Blocks()), importsFileName)
			}
		}
	}
	if existingFile == nil {
		existingFile = hclwrite.NewEmptyFile()
	}

	// do actual writes
	importsFile, err := os.Create(importsFileName)
	if err != nil {
		log.Printf("[ERROR] Can't create %s: %v", importsFileName, err)
		return
	}
	defer importsFile.Close()

	newImports := make(map[string]struct{}, 100)
	log.Printf("[DEBUG] started processing new writes for %s", importsFileName)
	// write native imports
	for importBlock := range importChan {
		if importBlock != "" {
			log.Printf("[TRACE] writing import command %s", importBlock)
			importsFile.WriteString(importBlock)
			id := extractResourceIdFromImportBlockString(importBlock)
			if id != "" {
				newImports[id] = struct{}{}
			}
		} else {
			log.Print("[WARN] got empty import command...")
		}
		ic.waitGroup.Done()
	}
	// write the rest of import blocks
	numResources := len(newImports)
	log.Printf("[DEBUG] finished processing new writes for %s. Wrote %d resources", importsFileName, numResources)
	// update existing file if incremental mode
	if ic.incremental {
		log.Printf("[DEBUG] Starting to merge existing resources for %s", importsFileName)
		f := hclwrite.NewEmptyFile()
		for _, block := range existingFile.Body().Blocks() {
			blockName := extractResourceIdFromImportBlock(block)
			if blockName == "" {
				log.Printf("[WARN] can't extract resource ID from import block: %s",
					string(block.BuildTokens(nil).Bytes()))
				continue
			}
			_, exists := newImports[blockName]
			_, deleted := ic.deletedResources[blockName]
			if exists {
				log.Printf("[DEBUG] resource %s already generated, skipping...", blockName)
			} else if deleted {
				log.Printf("[DEBUG] resource %s is deleted, skipping...", blockName)
			} else {
				log.Printf("[DEBUG] resource %s doesn't exist, adding...", blockName)
				f.Body().AppendBlock(block)
				numResources = numResources + 1
			}
		}
		_, err = importsFile.WriteString(string(f.Bytes()))
		if err != nil {
			log.Printf("[ERROR] error when writing existing resources for file %s: %v", importsFileName, err)
		}
		log.Printf("[DEBUG] Finished merging existing resources for %s", importsFileName)
	}
}

func (ic *importContext) writeShellImports(sh *os.File, importChan importWriteChannel) {
	for importCommand := range importChan {
		if importCommand != "" && sh != nil {
			log.Printf("[DEBUG] writing import command %s", importCommand)
			sh.WriteString(importCommand + "\n")
			delete(ic.shImports, importCommand)
		} else {
			log.Print("[WARN] got empty import command... or file is nil")
		}
		ic.waitGroup.Done()
	}
	if sh != nil {
		log.Printf("[DEBUG] Writing the rest of import commands. len=%d", len(ic.shImports))
		for k := range ic.shImports {
			parts := strings.Split(k, " ")
			if len(parts) > 3 {
				resource := parts[2]
				_, deleted := ic.deletedResources[resource]
				if deleted {
					log.Printf("[DEBUG] Resource %s is deleted. Skipping import command for it", resource)
					continue
				}
			}
			sh.WriteString(k + "\n")
		}
	}
}

func generateResourceName(rtype, rname string) string {
	return rtype + "." + rname
}

func generateBlockFullName(block *hclwrite.Block) string {
	labels := block.Labels()
	return generateResourceName(labels[0], strings.Join(labels[1:], "_"))
}

type resourceWriteData struct {
	BlockName     string
	ResourceBody  string
	ImportCommand string
}

type dataWriteChannel chan *resourceWriteData
type importWriteChannel chan string

func (ic *importContext) handleResourceWrite(generatedFile string, ch dataWriteChannel, importChan importWriteChannel) {
	var existingFile *hclwrite.File
	if ic.incremental {
		log.Printf("[DEBUG] Going to read existing file %s", generatedFile)
		content, err := os.ReadFile(generatedFile)
		if errors.Is(err, os.ErrNotExist) {
			log.Printf("[WARN] File %s doesn't exist when using incremental export", generatedFile)
		} else if err != nil {
			log.Printf("[ERROR] error opening %s", generatedFile)
		} else {
			log.Printf("[DEBUG] Going to parse existing file %s", generatedFile)
			var diags hcl.Diagnostics
			existingFile, diags = hclwrite.ParseConfig(content, generatedFile, hcl.Pos{Line: 1, Column: 1})
			if diags.HasErrors() {
				log.Printf("[ERROR] parsing of existing file %s failed: %s", generatedFile, diags.Error())
			} else {
				log.Printf("[DEBUG] There are %d objects in existing file %s",
					len(existingFile.Body().Blocks()), generatedFile)
			}
		}
	}
	if existingFile == nil {
		existingFile = hclwrite.NewEmptyFile()
	}

	tf, err := os.Create(generatedFile)
	if err != nil {
		log.Printf("[ERROR] Can't create %s: %v", generatedFile, err)
		return
	}

	newResources := make(map[string]struct{}, 100)
	log.Printf("[DEBUG] started processing new writes for %s", generatedFile)
	for f := range ch {
		if f != nil {
			// check if we have the same blockname already written. To avoid duplicates
			_, exists := newResources[f.BlockName]
			if !exists {
				log.Printf("[DEBUG] started writing resource body for %s", f.BlockName)
				_, err = tf.WriteString(f.ResourceBody)
				if err == nil {
					newResources[f.BlockName] = struct{}{}
					if f.ImportCommand != "" {
						ic.waitGroup.Add(1)
						importChan <- f.ImportCommand
					}
					log.Printf("[DEBUG] finished writing resource body for %s", f.BlockName)
				} else {
					log.Printf("[ERROR] Error when writing to %s: %v", generatedFile, err)
				}
			} else {
				log.Printf("[WARN] Found duplicate resource: '%s'", f.BlockName)
			}
		} else {
			log.Print("[WARN] got nil as resourceWriteData!")
		}
		ic.waitGroup.Done()
	}
	numResources := len(newResources)
	log.Printf("[DEBUG] finished processing new writes for %s. Wrote %d resources", generatedFile, numResources)
	// update existing file if incremental mode
	if ic.incremental {
		log.Printf("[DEBUG] Starting to merge existing resources for %s", generatedFile)
		f := hclwrite.NewEmptyFile()
		for _, block := range existingFile.Body().Blocks() {
			blockName := generateBlockFullName(block)
			_, exists := newResources[blockName]
			_, deleted := ic.deletedResources[blockName]
			if exists {
				log.Printf("[DEBUG] resource %s already generated, skipping...", blockName)
			} else if deleted {
				log.Printf("[DEBUG] resource %s is deleted, skipping...", blockName)
			} else {
				log.Printf("[DEBUG] resource %s doesn't exist, adding...", blockName)
				f.Body().AppendBlock(block)
				numResources = numResources + 1
			}
		}
		_, err = tf.WriteString(string(f.Bytes()))
		if err != nil {
			log.Printf("[ERROR] error when writing existing resources for file %s: %v", generatedFile, err)
		}
		log.Printf("[DEBUG] Finished merging existing resources for %s", generatedFile)
	}
	tf.Close()
	if numResources == 0 {
		log.Printf("[DEBUG] removing empty file %s - no resources for a given service", generatedFile)
		os.Remove(generatedFile)
	}
}

func (ic *importContext) generateResourceIdForWorkspaceObject(obj workspace.ObjectStatus) (string, string) {
	var rtype string
	switch obj.ObjectType {
	case workspace.Directory:
		rtype = "databricks_directory"
	case workspace.File:
		rtype = "databricks_workspace_file"
	case workspace.Notebook:
		rtype = "databricks_notebook"
	default:
		log.Printf("[WARN] Unsupported WS object type: %s in obj %v", obj.ObjectType, obj)
		return "", ""
	}
	rData := ic.Resources[rtype].Data(
		&terraform.InstanceState{
			ID:         obj.Path,
			Attributes: map[string]string{},
		})
	rData.Set("object_id", obj.ObjectID)
	rData.Set("path", obj.Path)
	name := ic.ResourceName(&resource{
		ID:       obj.Path,
		Resource: rtype,
		Data:     rData,
	})
	return generateResourceName(rtype, name), rtype
}

func (ic *importContext) loadOldWorkspaceObjects(fileName string) {
	ic.oldWorkspaceObjects = []workspace.ObjectStatus{}
	// Read a list of resources from previous run
	oldDataFile, err := os.ReadFile(fileName)
	if err != nil {
		log.Printf("[WARN] Can't open the file (%s) with previous list of workspace objects: %s", fileName, err.Error())
		return
	}
	err = json.Unmarshal(oldDataFile, &ic.oldWorkspaceObjects)
	if err != nil {
		log.Printf("[WARN] Can't desereialize previous list of workspace objects: %s", err.Error())
		return
	}
	log.Printf("[DEBUG] Read previous list of workspace objects. got %d objects", len(ic.oldWorkspaceObjects))
	for _, obj := range ic.oldWorkspaceObjects {
		ic.oldWorkspaceObjectMapping[obj.ObjectID] = obj.Path
	}
}

func (ic *importContext) findDeletedResources() {
	log.Print("[INFO] Starting detection of deleted workspace objects")
	if !ic.incremental || len(ic.allWorkspaceObjects) == 0 {
		return
	}
	if len(ic.oldWorkspaceObjects) == 0 {
		log.Print("[INFO] Previous list of workspace objects is empty")
		return
	}
	// generate IDs of current objects
	currentObjs := map[string]struct{}{}
	for _, obj := range ic.allWorkspaceObjects {
		obj := obj
		if !isSupportedWorkspaceObject(obj) {
			continue
		}
		rid, _ := ic.generateResourceIdForWorkspaceObject(obj)
		currentObjs[rid] = struct{}{}
	}
	// Loop through previous objects, and if it's missing from the current list, add it to deleted, including permission
	for _, obj := range ic.oldWorkspaceObjects {
		obj := obj
		if !isSupportedWorkspaceObject(obj) {
			continue
		}
		rid, rtype := ic.generateResourceIdForWorkspaceObject(obj)
		_, exists := currentObjs[rid]
		if exists {
			log.Printf("[DEBUG] object %s still exists", rid) // change to TRACE?
			continue
		}
		log.Printf("[DEBUG] object %s is deleted!", rid)
		ic.deletedResources[rid] = struct{}{}
		// convert into permissions. This is quite fragile right now, need to think how to handle it better
		var permId string
		switch rtype {
		case "databricks_notebook":
			permId = "databricks_permissions.notebook_" + rid[len(rtype)+1:]
		case "databricks_directory":
			permId = "databricks_permissions.directory_" + rid[len(rtype)+1:]
		case "databricks_workspace_file":
			permId = "databricks_permissions.ws_file_" + rid[len(rtype)+1:]
		}
		log.Printf("[DEBUG] deleted permissions object %s", permId)
		if permId != "" {
			ic.deletedResources[permId] = struct{}{}
		}
	}
	log.Printf("[INFO] Finished detection of deleted workspace objects. Detected %d deleted objects.",
		len(ic.deletedResources))
	log.Printf("[DEBUG] Deleted objects. %v", ic.deletedResources) // change to TRACE?
}
