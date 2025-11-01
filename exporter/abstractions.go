package exporter

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	frameworkschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	sdkv2schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// FieldType represents unified field types across SDKv2 and Plugin Framework
type FieldType int

const (
	FieldTypeString FieldType = iota
	FieldTypeInt
	FieldTypeBool
	FieldTypeFloat
	FieldTypeList
	FieldTypeSet
	FieldTypeMap
	FieldTypeObject
)

// ResourceDataWrapper provides unified access to resource data across SDKv2 and Plugin Framework
type ResourceDataWrapper interface {
	// Read operations
	GetOk(key string) (interface{}, bool)
	Get(key string) interface{}
	Id() string

	// Write operations (for directly generated data)
	SetId(id string)
	Set(key string, value interface{}) error

	// Schema access
	GetSchema() SchemaWrapper

	// Type checking
	IsPluginFramework() bool
}

// SchemaWrapper abstracts schema access across SDKv2 and Plugin Framework
type SchemaWrapper interface {
	GetFields() []string
	GetFieldSchema(field string) FieldSchema
	IsPluginFramework() bool
}

// FieldSchema abstracts individual field schema properties
type FieldSchema interface {
	GetType() FieldType
	IsRequired() bool
	IsComputed() bool
	IsOptional() bool
	IsSensitive() bool
	IsDeprecated() bool
	GetDeprecationMessage() string
	GetDefault() interface{}

	// For complex types
	GetElementSchema() FieldSchema  // for list, set
	GetNestedSchema() SchemaWrapper // for object/block
	GetKeySchema() FieldSchema      // for map keys
	GetValueSchema() FieldSchema    // for map values
	GetDescription() string

	// Backward compatibility - returns nil for Plugin Framework
	GetSDKv2Schema() *sdkv2schema.Schema
}

// ========== SDKv2 Implementations ==========

// SDKv2ResourceData wraps schema.ResourceData
type SDKv2ResourceData struct {
	data   *sdkv2schema.ResourceData
	schema *sdkv2schema.Resource
}

func (s *SDKv2ResourceData) GetOk(key string) (interface{}, bool) {
	return s.data.GetOk(key)
}

func (s *SDKv2ResourceData) Get(key string) interface{} {
	return s.data.Get(key)
}

func (s *SDKv2ResourceData) Id() string {
	return s.data.Id()
}

func (s *SDKv2ResourceData) SetId(id string) {
	s.data.SetId(id)
}

func (s *SDKv2ResourceData) Set(key string, value interface{}) error {
	return s.data.Set(key, value)
}

func (s *SDKv2ResourceData) GetSchema() SchemaWrapper {
	return &SDKv2SchemaWrapper{resource: s.schema}
}

func (s *SDKv2ResourceData) IsPluginFramework() bool {
	return false
}

// SDKv2SchemaWrapper wraps *schema.Resource
type SDKv2SchemaWrapper struct {
	resource *sdkv2schema.Resource
}

func (s *SDKv2SchemaWrapper) GetFields() []string {
	fields := make([]string, 0, len(s.resource.Schema))
	for field := range s.resource.Schema {
		fields = append(fields, field)
	}
	return fields
}

func (s *SDKv2SchemaWrapper) GetFieldSchema(field string) FieldSchema {
	if sch, ok := s.resource.Schema[field]; ok {
		return &SDKv2FieldSchema{schema: sch}
	}
	return nil
}

func (s *SDKv2SchemaWrapper) IsPluginFramework() bool {
	return false
}

// SDKv2FieldSchema wraps *schema.Schema
type SDKv2FieldSchema struct {
	schema *sdkv2schema.Schema
}

func (s *SDKv2FieldSchema) GetType() FieldType {
	switch s.schema.Type {
	case sdkv2schema.TypeString:
		return FieldTypeString
	case sdkv2schema.TypeInt:
		return FieldTypeInt
	case sdkv2schema.TypeBool:
		return FieldTypeBool
	case sdkv2schema.TypeFloat:
		return FieldTypeFloat
	case sdkv2schema.TypeList:
		return FieldTypeList
	case sdkv2schema.TypeSet:
		return FieldTypeSet
	case sdkv2schema.TypeMap:
		return FieldTypeMap
	default:
		return FieldTypeString
	}
}

func (s *SDKv2FieldSchema) IsRequired() bool {
	return s.schema.Required
}

func (s *SDKv2FieldSchema) IsComputed() bool {
	return s.schema.Computed
}

func (s *SDKv2FieldSchema) IsOptional() bool {
	return s.schema.Optional
}

func (s *SDKv2FieldSchema) IsSensitive() bool {
	return s.schema.Sensitive
}

func (s *SDKv2FieldSchema) IsDeprecated() bool {
	return s.schema.Deprecated != ""
}

func (s *SDKv2FieldSchema) GetDeprecationMessage() string {
	return s.schema.Deprecated
}

func (s *SDKv2FieldSchema) GetDefault() interface{} {
	return s.schema.Default
}

func (s *SDKv2FieldSchema) GetElementSchema() FieldSchema {
	if elem, ok := s.schema.Elem.(*sdkv2schema.Schema); ok {
		return &SDKv2FieldSchema{schema: elem}
	}
	return nil
}

func (s *SDKv2FieldSchema) GetNestedSchema() SchemaWrapper {
	if res, ok := s.schema.Elem.(*sdkv2schema.Resource); ok {
		return &SDKv2SchemaWrapper{resource: res}
	}
	return nil
}

func (s *SDKv2FieldSchema) GetKeySchema() FieldSchema {
	// SDKv2 maps don't have separate key schema
	return nil
}

func (s *SDKv2FieldSchema) GetValueSchema() FieldSchema {
	if s.schema.Type == sdkv2schema.TypeMap {
		if elem, ok := s.schema.Elem.(*sdkv2schema.Schema); ok {
			return &SDKv2FieldSchema{schema: elem}
		}
	}
	return nil
}

func (s *SDKv2FieldSchema) GetDescription() string {
	return s.schema.Description
}

func (s *SDKv2FieldSchema) GetSDKv2Schema() *sdkv2schema.Schema {
	return s.schema
}

// ========== Plugin Framework Implementations ==========

// PluginFrameworkResourceData wraps tfsdk.State
type PluginFrameworkResourceData struct {
	state  *tfsdk.State
	schema frameworkschema.Schema
}

func (p *PluginFrameworkResourceData) GetOk(key string) (interface{}, bool) {
	// Parse path: "field.subfield.0.name"
	pathParts := strings.Split(key, ".")
	attrPath := path.Empty()

	for _, part := range pathParts {
		if idx, err := strconv.Atoi(part); err == nil {
			attrPath = attrPath.AtListIndex(idx)
		} else {
			attrPath = attrPath.AtName(part)
		}
	}

	var value attr.Value
	diags := p.state.GetAttribute(context.Background(), attrPath, &value)
	if diags.HasError() || value.IsNull() || value.IsUnknown() {
		return nil, false
	}

	// Convert Plugin Framework types to Go types
	return convertPluginFrameworkToGoType(value), true
}

func (p *PluginFrameworkResourceData) Get(key string) interface{} {
	val, _ := p.GetOk(key)
	return val
}

func (p *PluginFrameworkResourceData) Id() string {
	val, ok := p.GetOk("id")
	if !ok {
		return ""
	}
	if str, ok := val.(string); ok {
		return str
	}
	return ""
}

func (p *PluginFrameworkResourceData) SetId(id string) {
	p.state.SetAttribute(context.Background(), path.Root("id"), types.StringValue(id))
}

func (p *PluginFrameworkResourceData) Set(key string, value interface{}) error {
	pathParts := strings.Split(key, ".")
	attrPath := path.Empty()
	for _, part := range pathParts {
		if idx, err := strconv.Atoi(part); err == nil {
			attrPath = attrPath.AtListIndex(idx)
		} else {
			attrPath = attrPath.AtName(part)
		}
	}

	// Convert Go type to Plugin Framework type
	pfValue := convertGoToPluginFrameworkType(value)
	diags := p.state.SetAttribute(context.Background(), attrPath, pfValue)
	if diags.HasError() {
		return fmt.Errorf("failed to set attribute: %v", diags)
	}
	return nil
}

func (p *PluginFrameworkResourceData) GetSchema() SchemaWrapper {
	return &PluginFrameworkSchemaWrapper{schema: p.schema}
}

func (p *PluginFrameworkResourceData) IsPluginFramework() bool {
	return true
}

// PluginFrameworkSchemaWrapper wraps Plugin Framework frameworkschema.Schema
type PluginFrameworkSchemaWrapper struct {
	schema frameworkschema.Schema
}

func (p *PluginFrameworkSchemaWrapper) GetFields() []string {
	fields := make([]string, 0, len(p.schema.Attributes)+len(p.schema.Blocks))
	for field := range p.schema.Attributes {
		fields = append(fields, field)
	}
	for field := range p.schema.Blocks {
		fields = append(fields, field)
	}
	return fields
}

func (p *PluginFrameworkSchemaWrapper) GetFieldSchema(field string) FieldSchema {
	if attr, ok := p.schema.Attributes[field]; ok {
		return &PluginFrameworkFieldSchema{attribute: attr, isBlock: false}
	}
	if block, ok := p.schema.Blocks[field]; ok {
		return &PluginFrameworkFieldSchema{block: block, isBlock: true}
	}
	return nil
}

func (p *PluginFrameworkSchemaWrapper) IsPluginFramework() bool {
	return true
}

// PluginFrameworkFieldSchema wraps Plugin Framework attribute or block
type PluginFrameworkFieldSchema struct {
	attribute frameworkschema.Attribute
	block     frameworkschema.Block
	isBlock   bool
}

func (p *PluginFrameworkFieldSchema) GetType() FieldType {
	if p.isBlock {
		return FieldTypeObject
	}

	// Check attribute type
	switch p.attribute.(type) {
	case frameworkschema.StringAttribute:
		return FieldTypeString
	case frameworkschema.Int64Attribute:
		return FieldTypeInt
	case frameworkschema.BoolAttribute:
		return FieldTypeBool
	case frameworkschema.Float64Attribute:
		return FieldTypeFloat
	case frameworkschema.ListAttribute:
		return FieldTypeList
	case frameworkschema.SetAttribute:
		return FieldTypeSet
	case frameworkschema.MapAttribute:
		return FieldTypeMap
	case frameworkschema.SingleNestedAttribute, frameworkschema.ListNestedAttribute, frameworkschema.SetNestedAttribute, frameworkschema.MapNestedAttribute:
		return FieldTypeObject
	default:
		return FieldTypeString
	}
}

func (p *PluginFrameworkFieldSchema) IsRequired() bool {
	if p.isBlock {
		return false
	}
	return p.attribute.IsRequired()
}

func (p *PluginFrameworkFieldSchema) IsComputed() bool {
	if p.isBlock {
		return false
	}
	return p.attribute.IsComputed()
}

func (p *PluginFrameworkFieldSchema) IsOptional() bool {
	if p.isBlock {
		return true
	}
	return p.attribute.IsOptional()
}

func (p *PluginFrameworkFieldSchema) IsSensitive() bool {
	if p.isBlock {
		return false
	}
	return p.attribute.IsSensitive()
}

func (p *PluginFrameworkFieldSchema) IsDeprecated() bool {
	if p.isBlock {
		return false
	}
	return p.attribute.GetDeprecationMessage() != ""
}

func (p *PluginFrameworkFieldSchema) GetDeprecationMessage() string {
	if p.isBlock {
		return ""
	}
	return p.attribute.GetDeprecationMessage()
}

func (p *PluginFrameworkFieldSchema) GetDefault() interface{} {
	// Plugin Framework doesn't have direct default values in the same way
	return nil
}

func (p *PluginFrameworkFieldSchema) GetElementSchema() FieldSchema {
	if p.isBlock {
		return nil
	}

	// Handle list/set element types
	// Plugin Framework doesn't expose element schema in the same way as SDKv2
	// Return nil for now - element types are handled during value conversion
	return nil
}

func (p *PluginFrameworkFieldSchema) GetNestedSchema() SchemaWrapper {
	if p.isBlock {
		switch block := p.block.(type) {
		case frameworkschema.ListNestedBlock:
			return &PluginFrameworkNestedSchemaWrapper{attrs: block.NestedObject.Attributes, blocks: block.NestedObject.Blocks}
		case frameworkschema.SetNestedBlock:
			return &PluginFrameworkNestedSchemaWrapper{attrs: block.NestedObject.Attributes, blocks: block.NestedObject.Blocks}
		case frameworkschema.SingleNestedBlock:
			return &PluginFrameworkNestedSchemaWrapper{attrs: block.Attributes, blocks: block.Blocks}
		}
	}

	switch attr := p.attribute.(type) {
	case frameworkschema.SingleNestedAttribute:
		return &PluginFrameworkNestedSchemaWrapper{attrs: attr.Attributes}
	case frameworkschema.ListNestedAttribute:
		return &PluginFrameworkNestedSchemaWrapper{attrs: attr.NestedObject.Attributes}
	case frameworkschema.SetNestedAttribute:
		return &PluginFrameworkNestedSchemaWrapper{attrs: attr.NestedObject.Attributes}
	case frameworkschema.MapNestedAttribute:
		return &PluginFrameworkNestedSchemaWrapper{attrs: attr.NestedObject.Attributes}
	}
	return nil
}

func (p *PluginFrameworkFieldSchema) GetKeySchema() FieldSchema {
	return nil
}

func (p *PluginFrameworkFieldSchema) GetValueSchema() FieldSchema {
	if p.isBlock {
		return nil
	}

	// Plugin Framework doesn't expose map element schema in the same way as SDKv2
	// Return nil for now - element types are handled during value conversion
	return nil
}

func (p *PluginFrameworkFieldSchema) GetDescription() string {
	if p.isBlock {
		return p.block.GetDescription()
	}
	return p.attribute.GetDescription()
}

func (p *PluginFrameworkFieldSchema) GetSDKv2Schema() *sdkv2schema.Schema {
	return nil
}

// PluginFrameworkNestedSchemaWrapper wraps nested schema attributes
type PluginFrameworkNestedSchemaWrapper struct {
	attrs  map[string]frameworkschema.Attribute
	blocks map[string]frameworkschema.Block
}

func (p *PluginFrameworkNestedSchemaWrapper) GetFields() []string {
	fields := make([]string, 0, len(p.attrs)+len(p.blocks))
	for field := range p.attrs {
		fields = append(fields, field)
	}
	for field := range p.blocks {
		fields = append(fields, field)
	}
	return fields
}

func (p *PluginFrameworkNestedSchemaWrapper) GetFieldSchema(field string) FieldSchema {
	if attr, ok := p.attrs[field]; ok {
		return &PluginFrameworkFieldSchema{attribute: attr, isBlock: false}
	}
	if block, ok := p.blocks[field]; ok {
		return &PluginFrameworkFieldSchema{block: block, isBlock: true}
	}
	return nil
}

func (p *PluginFrameworkNestedSchemaWrapper) IsPluginFramework() bool {
	return true
}

// ========== Type Conversion Helpers ==========

// convertPluginFrameworkToGoType converts Plugin Framework types to Go types
func convertPluginFrameworkToGoType(value attr.Value) interface{} {
	if value.IsNull() || value.IsUnknown() {
		return nil
	}

	switch v := value.(type) {
	case basetypes.StringValue:
		return v.ValueString()
	case basetypes.Int64Value:
		return int(v.ValueInt64())
	case basetypes.BoolValue:
		return v.ValueBool()
	case basetypes.Float64Value:
		return v.ValueFloat64()
	case basetypes.ListValue:
		elements := v.Elements()
		result := make([]interface{}, len(elements))
		for i, elem := range elements {
			result[i] = convertPluginFrameworkToGoType(elem)
		}
		return result
	case basetypes.SetValue:
		elements := v.Elements()
		result := make([]interface{}, len(elements))
		for i, elem := range elements {
			result[i] = convertPluginFrameworkToGoType(elem)
		}
		return result
	case basetypes.MapValue:
		elements := v.Elements()
		result := make(map[string]interface{}, len(elements))
		for k, elem := range elements {
			result[k] = convertPluginFrameworkToGoType(elem)
		}
		return result
	case basetypes.ObjectValue:
		attrs := v.Attributes()
		result := make(map[string]interface{}, len(attrs))
		for k, attr := range attrs {
			result[k] = convertPluginFrameworkToGoType(attr)
		}
		return result
	default:
		return fmt.Sprintf("%v", value)
	}
}

// convertGoToPluginFrameworkType converts Go types to Plugin Framework types
func convertGoToPluginFrameworkType(value interface{}) attr.Value {
	if value == nil {
		return types.StringNull()
	}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String:
		return types.StringValue(v.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return types.Int64Value(v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return types.Int64Value(int64(v.Uint()))
	case reflect.Bool:
		return types.BoolValue(v.Bool())
	case reflect.Float32, reflect.Float64:
		return types.Float64Value(v.Float())
	default:
		// For complex types, convert to string as fallback
		return types.StringValue(fmt.Sprintf("%v", value))
	}
}
