package common

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var kindMap = map[reflect.Kind]string{
	reflect.Bool:          "Bool",
	reflect.Int:           "Int",
	reflect.Int8:          "Int8",
	reflect.Int16:         "Int16",
	reflect.Int32:         "Int32",
	reflect.Int64:         "Int64",
	reflect.Uint:          "Uint",
	reflect.Uint8:         "Uint8",
	reflect.Uint16:        "Uint16",
	reflect.Uint32:        "Uint32",
	reflect.Uint64:        "Uint64",
	reflect.Uintptr:       "Uintptr",
	reflect.Float32:       "Float32",
	reflect.Float64:       "Float64",
	reflect.Complex64:     "Complex64",
	reflect.Complex128:    "Complex128",
	reflect.Array:         "Array",
	reflect.Chan:          "Chan",
	reflect.Func:          "Func",
	reflect.Interface:     "Interface",
	reflect.Ptr:           "Ptr",
	reflect.Slice:         "Slice",
	reflect.String:        "String",
	reflect.Struct:        "Struct",
	reflect.UnsafePointer: "UnsafePointer",
}

// Generic interface for ResourceProvider. Using CustomizeSchema function to keep track of additional information
// on top of the generated go-sdk struct.
type ResourceProvider interface {
	CustomizeSchema(map[string]*schema.Schema) map[string]*schema.Schema
}

// Interface for ResourceProvider instances that need aliases for fields.
type ResourceProviderWithAlias interface {
	ResourceProvider
	// Aliases() returns a two dimensional map where the top level key is the name of the struct, the second level key is the name of the field,
	// the values are the alias for the corresponding field under the specified struct.
	// Example:
	//
	//	{
	//	    "compute.ClusterSpec": {
	//	        "libraries": "library"
	//	    }
	//	}
	Aliases() map[string]map[string]string
}

// Interface for ResourceProvider instances that have recursive references in its schema.
// The function MaxDepthForTypes allows us to specify the max number of recursive depth for a specific field
//
// Example:
//
//	func (JobSettings) MaxDepthForTypes map[string]int {
//	    return map[string]int{"for_each_task": 2}
//	}
type RecursiveResourceProvider interface {
	ResourceProvider
	MaxDepthForTypes() map[string]int
}

// Takes in a ResourceProvider and converts that into a map from string to schema.
func resourceProviderStructToSchema(v ResourceProvider) map[string]*schema.Schema {
	rv := reflect.ValueOf(v)
	var scm map[string]*schema.Schema
	aliases := map[string]map[string]string{}
	if rpwa, ok := v.(ResourceProviderWithAlias); ok {
		aliases = rpwa.Aliases()
	}
	if rrp, ok := v.(RecursiveResourceProvider); ok {
		scm = typeToSchema(rv, aliases, getRecursionTrackingContext(rrp))
	} else {
		scm = typeToSchema(rv, aliases, getEmptyRecursionTrackingContext())
	}
	scm = v.CustomizeSchema(scm)
	return scm
}

func reflectKind(k reflect.Kind) string {
	n, ok := kindMap[k]
	if !ok {
		return "other"
	}
	return n
}

func chooseFieldNameWithAliases(typeField reflect.StructField, parentType reflect.Type, aliases map[string]map[string]string) string {
	parentTypeName := getNameForType(parentType)
	fieldNameWithAliasTag := chooseFieldName(typeField)
	// If nothing in the aliases map, return the field name from plain chooseFieldName method.
	if len(aliases) == 0 {
		return fieldNameWithAliasTag
	}

	jsonFieldName := getJsonFieldName(typeField)
	if jsonFieldName == "-" {
		return "-"
	}

	if parentMap, ok := aliases[parentTypeName]; ok {
		if value, ok := parentMap[jsonFieldName]; ok {
			return value
		}
	}
	return fieldNameWithAliasTag
}

func getJsonFieldName(typeField reflect.StructField) string {
	jsonTag := typeField.Tag.Get("json")
	// fields without JSON tags would be treated as if ignored,
	// but keeping linters happy
	if jsonTag == "" {
		return "-"
	}
	return strings.Split(jsonTag, ",")[0]
}

// SchemaPath helps to navigate
func SchemaPath(s map[string]*schema.Schema, path ...string) (*schema.Schema, error) {
	cs := s
	for i, p := range path {
		v, ok := cs[p]
		if !ok {
			return nil, fmt.Errorf("missing key %s", p)
		}
		if i == len(path)-1 {
			return v, nil
		}
		cv, ok := v.Elem.(*schema.Resource)
		if !ok {
			return nil, fmt.Errorf("%s is not nested resource", p)
		}
		cs = cv.Schema
	}
	return nil, fmt.Errorf("%v does not compute", path)
}

func MustSchemaPath(s map[string]*schema.Schema, path ...string) *schema.Schema {
	sch, err := SchemaPath(s, path...)
	if err != nil {
		panic(err)
	}
	return sch
}

// StructToSchema makes schema from a struct type & applies customizations from callback given
func StructToSchema(v any, customize func(map[string]*schema.Schema) map[string]*schema.Schema) map[string]*schema.Schema {
	// If the input 'v' is an instance of ResourceProvider, call resourceProviderStructToSchema instead.
	if rp, ok := v.(ResourceProvider); ok {
		if customize != nil {
			panic("customize should be nil if the input implements the ResourceProvider interface; use CustomizeSchema of ResourceProvider instead")
		}
		return resourceProviderStructToSchema(rp)
	}
	rv := reflect.ValueOf(v)
	scm := typeToSchema(rv, map[string]map[string]string{}, getEmptyRecursionTrackingContext())
	if customize != nil {
		scm = customize(scm)
	}
	return scm
}

// SetDefault sets the default value for a schema.
func SetDefault(v *schema.Schema, value any) {
	v.Default = value
	v.Optional = true
	v.Required = false
}

// SetReadOnly sets the schema to be read-only (i.e. computed, non-optional).
// This should be used for fields that are not user-configurable but are returned
// by the platform.
func SetReadOnly(v *schema.Schema) {
	v.Optional = false
	v.Required = false
	v.MaxItems = 0
	v.Computed = true
}

// SetRequired sets the schema to be required.
func SetRequired(v *schema.Schema) {
	v.Optional = false
	v.Required = true
	v.Computed = false
}

func isOptional(typeField reflect.StructField) bool {
	if strings.Contains(typeField.Tag.Get("json"), "omitempty") {
		return true
	}
	tfTags := strings.Split(typeField.Tag.Get("tf"), ",")
	for _, tag := range tfTags {
		if tag == "optional" {
			return true
		}
	}
	return false
}

func handleOptional(typeField reflect.StructField, schema *schema.Schema) {
	if isOptional(typeField) {
		schema.Optional = true
	} else {
		schema.Required = true
	}
}

func handleComputed(typeField reflect.StructField, schema *schema.Schema) {
	tfTags := strings.Split(typeField.Tag.Get("tf"), ",")
	for _, tag := range tfTags {
		if tag == "computed" {
			schema.Computed = true
			schema.Required = false
			break
		}
	}
}

func handleForceNew(typeField reflect.StructField, schema *schema.Schema) {
	tfTags := strings.Split(typeField.Tag.Get("tf"), ",")
	for _, tag := range tfTags {
		if tag == "force_new" {
			schema.ForceNew = true
			break
		}
	}
}

func handleSensitive(typeField reflect.StructField, schema *schema.Schema) {
	tfTags := strings.Split(typeField.Tag.Get("tf"), ",")
	for _, tag := range tfTags {
		if tag == "sensitive" {
			schema.Sensitive = true
			break
		}
	}
}

func handleSuppressDiff(typeField reflect.StructField, fieldName string, v *schema.Schema) {
	tfTags := strings.Split(typeField.Tag.Get("tf"), ",")
	for _, tag := range tfTags {
		if tag == "suppress_diff" {
			v.DiffSuppressFunc = diffSuppressor(fieldName, v)
			break
		}
	}
}

func getAlias(typeField reflect.StructField) string {
	tfTags := strings.Split(typeField.Tag.Get("tf"), ",")
	for _, tag := range tfTags {
		if strings.HasPrefix(tag, "alias:") {
			return strings.TrimPrefix(tag, "alias:")
		}
	}
	return ""
}

func chooseFieldName(typeField reflect.StructField) string {
	alias := getAlias(typeField)
	if alias != "" {
		return alias
	}
	return getJsonFieldName(typeField)
}

func diffSuppressor(fieldName string, v *schema.Schema) func(k, old, new string, d *schema.ResourceData) bool {
	zero := fmt.Sprintf("%v", v.Type.Zero())
	return func(k, old, new string, d *schema.ResourceData) bool {
		if new == zero && old != zero {
			log.Printf("[DEBUG] Suppressing diff for %v: platform=%#v config=%#v", k, old, new)
			return true
		}
		// When suppressing diffs for a list of attributes, SuppressDiffFunc is called for each diff
		// recursively. To verify that the list is empty, we need to ensure that the key being diffed
		// is exactly the list's attribute itself and not one of its elements.
		if strings.HasSuffix(k, fmt.Sprintf("%s.#", fieldName)) && new == "0" && old != "0" {
			log.Printf("[DEBUG] Suppressing diff for list or set %v: no value configured but platform returned some value (likely {})", fieldName)
			return true
		}
		return false
	}
}

type field struct {
	sf reflect.StructField
	v  reflect.Value
}

func listAllFields(v reflect.Value) []field {
	t := v.Type()
	fields := make([]field, 0, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)
		if f.Anonymous {
			fields = append(fields, listAllFields(v.Field(i))...)
		} else {
			fields = append(fields, field{
				sf: f,
				v:  v.Field(i),
			})
		}
	}
	return fields
}

func typeToSchema(v reflect.Value, aliases map[string]map[string]string, rt recursionTrackingContext) map[string]*schema.Schema {
	scm := map[string]*schema.Schema{}
	rk := v.Kind()
	if rk == reflect.Ptr {
		v = v.Elem()
		rk = v.Kind()
	}
	if rk != reflect.Struct {
		panic(fmt.Errorf("Schema value of Struct is expected, but got %s: %#v", reflectKind(rk), v))
	}
	rt = rt.copy()
	rt.visit(v)
	fields := listAllFields(v)
	for _, field := range fields {
		typeField := field.sf
		if rt.depthExceeded(typeField) {
			// Skip the field if recursion depth is over the limit.
			log.Printf("[TRACE] over recursion limit, skipping field: %s, max depth: %d", getNameForType(typeField.Type), rt.getMaxDepthForTypeField(typeField))
			continue
		}
		tfTag := typeField.Tag.Get("tf")

		fieldName := chooseFieldNameWithAliases(typeField, v.Type(), aliases)
		if fieldName == "-" {
			continue
		}
		scm[fieldName] = &schema.Schema{}
		for _, token := range strings.Split(tfTag, ",") {
			colonSplit := strings.Split(token, ":")
			if len(colonSplit) == 2 {
				tfKey := colonSplit[0]
				tfValue := colonSplit[1]
				switch tfKey {
				case "default":
					scm[fieldName].Default = tfValue
				case "max_items":
					maxItems, err := strconv.Atoi(tfValue)
					if err != nil {
						continue
					}
					scm[fieldName].MaxItems = maxItems
				case "min_items":
					minItems, err := strconv.Atoi(tfValue)
					if err != nil {
						continue
					}
					scm[fieldName].MinItems = minItems
				}
			}
		}
		handleOptional(typeField, scm[fieldName])
		handleComputed(typeField, scm[fieldName])
		handleForceNew(typeField, scm[fieldName])
		handleSensitive(typeField, scm[fieldName])
		switch typeField.Type.Kind() {
		case reflect.Int, reflect.Int32, reflect.Int64:
			scm[fieldName].Type = schema.TypeInt
			// diff suppression needs type for zero value
			handleSuppressDiff(typeField, fieldName, scm[fieldName])
		case reflect.Float64:
			scm[fieldName].Type = schema.TypeFloat
			// diff suppression needs type for zero value
			handleSuppressDiff(typeField, fieldName, scm[fieldName])
		case reflect.Bool:
			scm[fieldName].Type = schema.TypeBool
		case reflect.String:
			scm[fieldName].Type = schema.TypeString
			// diff suppression needs type for zero value
			handleSuppressDiff(typeField, fieldName, scm[fieldName])
		case reflect.Map:
			scm[fieldName].Type = schema.TypeMap
			elem := typeField.Type.Elem()
			switch elem.Kind() {
			case reflect.String:
				scm[fieldName].Elem = schema.TypeString
			case reflect.Int64:
				scm[fieldName].Elem = schema.TypeInt
			default:
				panic(fmt.Errorf("unsupported map value for %s: %s", fieldName, reflectKind(elem.Kind())))
			}
		case reflect.Ptr:
			scm[fieldName].MaxItems = 1
			scm[fieldName].Type = schema.TypeList
			elem := typeField.Type.Elem()
			sv := reflect.New(elem).Elem()
			nestedSchema := typeToSchema(sv, aliases, rt)
			if strings.Contains(tfTag, "suppress_diff") {
				scm[fieldName].DiffSuppressFunc = diffSuppressor(fieldName, scm[fieldName])
				for k, v := range nestedSchema {
					v.DiffSuppressFunc = diffSuppressor(k, v)
				}
			}
			scm[fieldName].Elem = &schema.Resource{
				Schema: nestedSchema,
			}
		case reflect.Struct:
			scm[fieldName].MaxItems = 1
			scm[fieldName].Type = schema.TypeList

			elem := typeField.Type  // changed from ptr
			sv := reflect.New(elem) // changed from ptr

			nestedSchema := typeToSchema(sv, aliases, rt)
			if strings.Contains(tfTag, "suppress_diff") {
				scm[fieldName].DiffSuppressFunc = diffSuppressor(fieldName, scm[fieldName])
				for k, v := range nestedSchema {
					v.DiffSuppressFunc = diffSuppressor(k, v)
				}
			}
			scm[fieldName].Elem = &schema.Resource{
				Schema: nestedSchema,
			}
		case reflect.Slice:
			ft := schema.TypeList
			if strings.Contains(tfTag, "slice_set") {
				ft = schema.TypeSet
			}
			scm[fieldName].Type = ft
			elem := typeField.Type.Elem()
			switch elem.Kind() {
			case reflect.Int, reflect.Int32, reflect.Int64:
				scm[fieldName].Elem = &schema.Schema{Type: schema.TypeInt}
			case reflect.Float64:
				scm[fieldName].Elem = &schema.Schema{Type: schema.TypeFloat}
			case reflect.Bool:
				scm[fieldName].Elem = &schema.Schema{Type: schema.TypeBool}
			case reflect.String:
				scm[fieldName].Elem = &schema.Schema{Type: schema.TypeString}
			case reflect.Struct:
				sv := reflect.New(elem).Elem()
				scm[fieldName].Elem = &schema.Resource{
					Schema: typeToSchema(sv, aliases, rt),
				}
			}
		default:
			panic(fmt.Errorf("unknown type for %s: %s", fieldName, reflectKind(typeField.Type.Kind())))
		}
	}
	return scm
}

func IsRequestEmpty(v any) (bool, error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return false, fmt.Errorf("value of Struct is expected, but got %s: %#v", reflectKind(rv.Kind()), rv)
	}
	var isNotEmpty bool
	err := iterFields(rv, []string{}, StructToSchema(v, nil), map[string]map[string]string{}, func(fieldSchema *schema.Schema, path []string, valueField *reflect.Value) error {
		if isNotEmpty {
			return nil
		}
		if !valueField.IsZero() {
			isNotEmpty = true
		}
		return nil
	})
	return !isNotEmpty, err
}

// isGoSdk returns true if the struct is from databricks-sdk-go or embeds a struct from databricks-sdk-go.
func isGoSdk(v reflect.Value) bool {
	if strings.Contains(v.Type().PkgPath(), "databricks-sdk-go") {
		return true
	}
	for i := 0; i < v.NumField(); i++ {
		f := v.Type().Field(i)
		if f.Anonymous && isGoSdk(v.Field(i)) {
			return true
		}
	}
	return false
}

// Iterate through each field of the given reflect.Value object and execute a callback function with the corresponding
// terraform schema object as the input.
func iterFields(rv reflect.Value, path []string, s map[string]*schema.Schema, aliases map[string]map[string]string,
	cb func(fieldSchema *schema.Schema, path []string, valueField *reflect.Value) error) error {
	rk := rv.Kind()
	if rk != reflect.Struct {
		return fmt.Errorf("value of Struct is expected, but got %s: %#v", reflectKind(rk), rv)
	}
	if !rv.IsValid() {
		return fmt.Errorf("%s: got invalid reflect value %#v", path, rv)
	}
	isGoSDK := isGoSdk(rv)
	fields := listAllFields(rv)
	for _, field := range fields {
		typeField := field.sf
		fieldName := chooseFieldNameWithAliases(typeField, rv.Type(), aliases)
		if fieldName == "-" {
			continue
		}
		fieldSchema, ok := s[fieldName]
		if !ok {
			continue
		}
		omitEmpty := isOptional(typeField)
		// TODO: fix in https://github.com/databricks/databricks-sdk-go/issues/268
		if !isGoSDK && omitEmpty && !fieldSchema.Optional {
			return fmt.Errorf("inconsistency: %s has omitempty, but is not optional", fieldName)
		}
		defaultEmpty := reflect.ValueOf(fieldSchema.Default).Kind() == reflect.Invalid
		if fieldSchema.Optional && defaultEmpty && !omitEmpty {
			return fmt.Errorf("inconsistency: %s is optional, default is empty, but has no omitempty", fieldName)
		}
		valueField := field.v
		err := cb(fieldSchema, append(path, fieldName), &valueField)
		if err != nil {
			return fmt.Errorf("%s: %s", fieldName, err)
		}
	}
	return nil
}

func collectionToMaps(v any, s *schema.Schema, aliases map[string]map[string]string) ([]any, error) {
	resultList := []any{}
	if sl, ok := v.([]string); ok {
		// most likely list of parameters to job task
		for _, str := range sl {
			resultList = append(resultList, str)
		}
		return resultList, nil
	}
	r, ok := s.Elem.(*schema.Resource)
	if !ok {
		return nil, fmt.Errorf("not resource")
	}
	var allItems []reflect.Value
	rv := reflect.ValueOf(v)
	rvType := rv.Type().Kind()
	isList := rvType == reflect.Array || rvType == reflect.Slice
	if isList {
		for i := 0; i < rv.Len(); i++ {
			allItems = append(allItems, rv.Index(i))
		}
	} else {
		allItems = append(allItems, rv)
	}
	for _, v := range allItems {
		data := map[string]any{}
		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				continue
			}
			v = v.Elem()
		}
		err := iterFields(v, []string{}, r.Schema, aliases, func(fieldSchema *schema.Schema,
			path []string, valueField *reflect.Value) error {
			fieldName := path[len(path)-1]
			fieldValue := valueField.Interface()
			fieldPath := strings.Join(path, ".")
			switch fieldSchema.Type {
			case schema.TypeList, schema.TypeSet:
				nv, err := collectionToMaps(fieldValue, fieldSchema, aliases)
				if err != nil {
					return fmt.Errorf("%s: %v", path, err)
				}
				data[fieldName] = nv
			default:
				if fieldSchema.Optional && isValueNilOrEmpty(valueField, fieldPath) {
					return nil
				}
				data[fieldName] = fieldValue
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
		if len(data) == 0 {
			continue
		}
		resultList = append(resultList, data)
	}
	return resultList, nil
}

func isValueNilOrEmpty(valueField *reflect.Value, fieldPath string) bool {
	switch valueField.Kind() {
	case reflect.Ptr:
		if valueField.IsNil() {
			log.Printf("[TRACE] skipping empty %s %#v", fieldPath, valueField)
			return true
		}
	case reflect.Array, reflect.Map, reflect.String, reflect.Slice:
		if valueField.Len() == 0 {
			log.Printf("[TRACE] skipping empty %s %#v", fieldPath, valueField)
			return true
		}
	}
	return false
}

// StructToData reads result using schema onto resource data
func StructToData(result any, s map[string]*schema.Schema, d *schema.ResourceData) error {
	aliases := getAliasesMapFromStruct(result)
	v := reflect.ValueOf(result)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return iterFields(v, []string{}, s, aliases, func(
		fieldSchema *schema.Schema, path []string, valueField *reflect.Value) error {
		fieldValue := valueField.Interface()
		if fieldValue == nil {
			return nil
		}
		fieldPath := strings.Join(path, ".")
		if fieldSchema.Optional && isValueNilOrEmpty(valueField, fieldPath) {
			return nil
		}
		_, configured := d.GetOk(fieldPath)
		if !d.IsNewResource() && !fieldSchema.Computed && !configured {
			log.Printf("[TRACE] Removing default fields sent back by server: %s - %#v",
				fieldPath, fieldValue)
			return nil
		}
		switch fieldSchema.Type {
		case schema.TypeList, schema.TypeSet:
			es, ok := fieldSchema.Elem.(*schema.Schema)
			if ok {
				log.Printf("[TRACE] Set %s %s %v", es.Type, fieldPath, fieldValue)
				// here we rely on Terraform SDK to perform
				// validation, so we don't to it twice
				return d.Set(fieldPath, fieldValue)
			}
			nv, err := collectionToMaps(fieldValue, fieldSchema, aliases)
			if err != nil {
				return fmt.Errorf("%s: %v", fieldPath, err)
			}
			if len(nv) == 0 {
				return nil
			}
			log.Printf("[TRACE] set %s %#v", fieldPath, nv)
			return d.Set(fieldPath, nv)
		default:
			log.Printf("[TRACE] set %s %#v", fieldPath, fieldValue)
			return d.Set(fieldPath, fieldValue)
		}
	})
}

// attributeGetter is a generalization between schema.ResourceDiff & schema.ResourceData
// to those who'll be reading this code and would know public equivalent interface from
// TF SDK - feel free to replace the usages of this interface in a PR.
type attributeGetter interface {
	GetOk(key string) (any, bool)
	// This method should only be used for optional boolean parameters according to the Terraform docs:
	// https://pkg.go.dev/github.com/hashicorp/terraform-plugin-sdk/helper/schema#ResourceData.GetOkExists
	GetOkExists(key string) (any, bool)
}

// DiffToStructPointer reads resource diff with given schema onto result pointer. Panics.
func DiffToStructPointer(d attributeGetter, scm map[string]*schema.Schema, result any) {
	rv := reflect.ValueOf(result)
	rk := rv.Kind()
	if rk != reflect.Ptr {
		panic(fmt.Errorf("pointer is expected, but got %s: %#v", reflectKind(rk), result))
	}
	rv = rv.Elem()
	aliases := getAliasesMapFromStruct(result)
	err := readReflectValueFromData([]string{}, d, rv, scm, aliases)
	if err != nil {
		panic(err)
	}
}

// DataToStructPointer reads resource data with given schema onto result pointer. Panics.
func DataToStructPointer(d *schema.ResourceData, scm map[string]*schema.Schema, result any) {
	aliases := getAliasesMapFromStruct(result)
	rv := reflect.ValueOf(result)
	rk := rv.Kind()
	if rk != reflect.Ptr {
		panic(fmt.Errorf("pointer is expected, but got %s: %#v", reflectKind(rk), result))
	}
	rv = rv.Elem()
	err := readReflectValueFromData([]string{}, d, rv, scm, aliases)
	if err != nil {
		panic(err)
	}
}

// DataToReflectValue reads reflect value from data
func DataToReflectValue(d *schema.ResourceData, s map[string]*schema.Schema, rv reflect.Value) error {
	// TODO: Pass in the right aliases map.
	return readReflectValueFromData([]string{}, d, rv, s, map[string]map[string]string{})
}

// Get the aliases map from the given struct if it is an instance of ResourceProvider.
// NOTE: This does not return aliases defined on `tf` tags.
func getAliasesMapFromStruct(s any) map[string]map[string]string {
	if v, ok := s.(ResourceProviderWithAlias); ok {
		return v.Aliases()
	}
	return map[string]map[string]string{}
}

func readReflectValueFromData(path []string, d attributeGetter,
	rv reflect.Value, s map[string]*schema.Schema, aliases map[string]map[string]string) error {
	return iterFields(rv, path, s, aliases, func(fieldSchema *schema.Schema,
		path []string, valueField *reflect.Value) error {
		fieldPath := strings.Join(path, ".")
		raw, ok := d.GetOk(fieldPath)
		if !ok {
			return nil
		}
		switch fieldSchema.Type {
		case schema.TypeInt:
			if v, ok := raw.(int); ok {
				valueField.SetInt(int64(v))
			}
		case schema.TypeString:
			if v, ok := raw.(string); ok {
				valueField.SetString(v)
			}
		case schema.TypeBool:
			if v, ok := raw.(bool); ok {
				valueField.SetBool(v)
			}
		case schema.TypeFloat:
			if v, ok := raw.(float64); ok {
				valueField.SetFloat(v)
			}
		case schema.TypeMap:
			mapValueKind := valueField.Type().Elem().Kind()
			valueField.Set(reflect.MakeMap(valueField.Type()))
			for key, ivalue := range raw.(map[string]any) {
				vrv, err := primitiveReflectValueFromInterface(mapValueKind, ivalue, fieldPath, key)
				if err != nil {
					return err
				}
				valueField.SetMapIndex(reflect.ValueOf(key), vrv)
			}
		case schema.TypeSet:
			// here we rely on Terraform SDK to perform validation, so we don't to it twice
			rawSet := raw.(*schema.Set)
			rawList := rawSet.List()
			return readListFromData(path, d, rawList, valueField,
				fieldSchema, aliases, func(i int) string {
					return strconv.Itoa(rawSet.F(rawList[i]))
				})
		case schema.TypeList:
			// here we rely on Terraform SDK to perform validation, so we don't to it twice
			rawList := raw.([]any)
			return readListFromData(path, d, rawList, valueField, fieldSchema, aliases, strconv.Itoa)
		default:
			return fmt.Errorf("%s[%v] unsupported field type", fieldPath, raw)
		}
		return nil
	})
}

func primitiveReflectValueFromInterface(rk reflect.Kind,
	ivalue any, fieldPath, key string) (rv reflect.Value, err error) {
	switch rk {
	case reflect.String:
		return reflect.ValueOf(fmt.Sprintf("%v", ivalue)), nil
	case reflect.Float32:
		v, ok := ivalue.(float32)
		if !ok {
			err = fmt.Errorf("%s[%s] '%v' is not %s",
				fieldPath, key, ivalue, reflectKind(rk))
			return
		}
		rv = reflect.ValueOf(v)
	case reflect.Float64:
		v, ok := ivalue.(float64)
		if !ok {
			err = fmt.Errorf("%s[%s] '%v' is not %s",
				fieldPath, key, ivalue, reflectKind(rk))
			return
		}
		rv = reflect.ValueOf(v)
	case reflect.Int:
		v, ok := ivalue.(int)
		if !ok {
			err = fmt.Errorf("%s[%s] '%v' is not %s",
				fieldPath, key, ivalue, reflectKind(rk))
			return
		}
		rv = reflect.ValueOf(v)
	case reflect.Bool:
		v, ok := ivalue.(bool)
		if !ok {
			err = fmt.Errorf("%s[%s] '%v' is not %s",
				fieldPath, key, ivalue, reflectKind(rk))
			return
		}
		rv = reflect.ValueOf(v)
	default:
		err = fmt.Errorf("%s[%s] '%v' is not valid primitive",
			fieldPath, key, ivalue)
	}
	return rv, err
}

func readListFromData(path []string, d attributeGetter,
	rawList []any, valueField *reflect.Value, fieldSchema *schema.Schema, aliases map[string]map[string]string,
	offsetConverter func(i int) string) error {
	if len(rawList) == 0 {
		return nil
	}
	fieldPath := strings.Join(path, ".")
	switch valueField.Type().Kind() {
	case reflect.Ptr:
		vpointer := reflect.New(valueField.Type().Elem())
		valueField.Set(vpointer)
		ve := vpointer.Elem()
		// here we rely on Terraform SDK to perform validation, so we don't to it twice
		nestedResource := fieldSchema.Elem.(*schema.Resource)
		nestedPath := append(path, offsetConverter(0))
		return readReflectValueFromData(nestedPath, d, ve, nestedResource.Schema, aliases)
	case reflect.Struct:
		// code path for setting the struct value is different from pointer value
		// in a single way: we set the field only after readReflectValueFromData
		// traversed the graph.
		vstruct := reflect.New(valueField.Type())
		ve := vstruct.Elem()
		nestedResource := fieldSchema.Elem.(*schema.Resource)
		nestedPath := append(path, offsetConverter(0))
		err := readReflectValueFromData(nestedPath, d, ve, nestedResource.Schema, aliases)
		if err != nil {
			return err
		}
		valueField.Set(ve)
		return nil
	case reflect.Slice:
		k := valueField.Type().Elem().Kind()
		newSlice := reflect.MakeSlice(valueField.Type(), len(rawList), len(rawList))
		valueField.Set(newSlice)
		for i, elem := range rawList {
			item := newSlice.Index(i)
			switch k {
			case reflect.Struct:
				// here we rely on Terraform SDK to perform validation, so we don't to it twice
				nestedResource := fieldSchema.Elem.(*schema.Resource)
				nestedPath := append(path, offsetConverter(i))
				vpointer := reflect.New(valueField.Type().Elem())
				ve := vpointer.Elem()
				err := readReflectValueFromData(nestedPath, d, ve, nestedResource.Schema, aliases)
				if err != nil {
					return err
				}
				item.Set(ve)
			default:
				err := setPrimitiveValueOfKind(fieldPath, k, item, elem)
				if err != nil {
					return err
				}
			}
		}
	default:
		return fmt.Errorf("%s[%v] unknown collection field", fieldPath, rawList)
	}
	return nil
}

func setPrimitiveValueOfKind(
	fieldPath string, k reflect.Kind, item reflect.Value, elem any) error {
	switch k {
	case reflect.String:
		v, ok := elem.(string)
		if !ok {
			return fmt.Errorf("%s[%v] is not a string", fieldPath, elem)
		}
		item.SetString(v)
	case reflect.Int, reflect.Int64:
		v, ok := elem.(int)
		if !ok {
			return fmt.Errorf("%s[%v] is not an int", fieldPath, elem)
		}
		item.SetInt(int64(v))
	case reflect.Float64:
		v, ok := elem.(float64)
		if !ok {
			return fmt.Errorf("%s[%v] is not a float64", fieldPath, elem)
		}
		item.SetFloat(v)
	case reflect.Bool:
		v, ok := elem.(bool)
		if !ok {
			return fmt.Errorf("%s[%v] is not a bool", fieldPath, elem)
		}
		item.SetBool(v)
	default:
		return fmt.Errorf("%s[%v] is not a valid primitive", fieldPath, elem)
	}
	return nil
}
