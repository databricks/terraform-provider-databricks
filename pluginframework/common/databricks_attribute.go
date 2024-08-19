package pluginframework

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dataschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// Common interface for all attributes, we need this because in terraform plugin framework, the datasource schema and resource
// schema are in two separate packages. This common interface prevents us from keeping two copies of StructToSchema and CustomizableSchema.
type Attribute interface {
	ToDataSourceAttribute() dataschema.Attribute
	ToResourceAttribute() schema.Attribute
	SetOptional() Attribute
	SetRequired() Attribute
	SetSensitive() Attribute
	SetComputed() Attribute
	SetReadOnly() Attribute
	SetDeprecated(string) Attribute
	AddValidators(any) Attribute
}

type StringAttribute struct {
	Optional           bool
	Required           bool
	Sensitive          bool
	Computed           bool
	DeprecationMessage string
	Validators         []validator.String
}

func (a StringAttribute) ToDataSourceAttribute() dataschema.Attribute {
	return dataschema.StringAttribute{Optional: a.Optional, Required: a.Required, Sensitive: a.Sensitive, DeprecationMessage: a.DeprecationMessage, Computed: a.Computed, Validators: a.Validators}
}

func (a StringAttribute) ToResourceAttribute() schema.Attribute {
	return schema.StringAttribute{Optional: a.Optional, Required: a.Required, Sensitive: a.Sensitive, DeprecationMessage: a.DeprecationMessage, Computed: a.Computed, Validators: a.Validators}
}

func (a StringAttribute) SetOptional() Attribute {
	a.Optional = true
	a.Required = false
	return a
}

func (a StringAttribute) SetRequired() Attribute {
	a.Optional = false
	a.Required = true
	return a
}

func (a StringAttribute) SetSensitive() Attribute {
	a.Sensitive = true
	return a
}

func (a StringAttribute) SetComputed() Attribute {
	a.Computed = true
	return a
}

func (a StringAttribute) SetReadOnly() Attribute {
	a.Computed = true
	a.Optional = false
	a.Required = false
	return a
}

func (a StringAttribute) SetDeprecated(msg string) Attribute {
	a.DeprecationMessage = msg
	return a
}

func (a StringAttribute) AddValidators(v any) Attribute {
	a.Validators = append(a.Validators, v.(validator.String))
	return a
}

type Float64Attribute struct {
	Optional           bool
	Required           bool
	Sensitive          bool
	Computed           bool
	DeprecationMessage string
	Validators         []validator.Float64
}

func (a Float64Attribute) ToDataSourceAttribute() dataschema.Attribute {
	return dataschema.Float64Attribute{Optional: a.Optional, Required: a.Required, Sensitive: a.Sensitive, DeprecationMessage: a.DeprecationMessage, Computed: a.Computed, Validators: a.Validators}
}

func (a Float64Attribute) ToResourceAttribute() schema.Attribute {
	return schema.Float64Attribute{Optional: a.Optional, Required: a.Required, Sensitive: a.Sensitive, DeprecationMessage: a.DeprecationMessage, Computed: a.Computed, Validators: a.Validators}
}

func (a Float64Attribute) SetOptional() Attribute {
	a.Optional = true
	a.Required = false
	return a
}

func (a Float64Attribute) SetRequired() Attribute {
	a.Optional = false
	a.Required = true
	return a
}

func (a Float64Attribute) SetSensitive() Attribute {
	a.Sensitive = true
	return a
}

func (a Float64Attribute) SetComputed() Attribute {
	a.Computed = true
	return a
}

func (a Float64Attribute) SetReadOnly() Attribute {
	a.Computed = true
	a.Optional = false
	a.Required = false
	return a
}

func (a Float64Attribute) SetDeprecated(msg string) Attribute {
	a.DeprecationMessage = msg
	return a
}

func (a Float64Attribute) AddValidators(v any) Attribute {
	a.Validators = append(a.Validators, v.(validator.Float64))
	return a
}

type Int64Attribute struct {
	Optional           bool
	Required           bool
	Sensitive          bool
	Computed           bool
	DeprecationMessage string
	Validators         []validator.Int64
}

func (a Int64Attribute) ToDataSourceAttribute() dataschema.Attribute {
	return dataschema.Int64Attribute{Optional: a.Optional, Required: a.Required, Sensitive: a.Sensitive, DeprecationMessage: a.DeprecationMessage, Computed: a.Computed, Validators: a.Validators}
}

func (a Int64Attribute) ToResourceAttribute() schema.Attribute {
	return schema.Int64Attribute{Optional: a.Optional, Required: a.Required, Sensitive: a.Sensitive, DeprecationMessage: a.DeprecationMessage, Computed: a.Computed, Validators: a.Validators}
}

func (a Int64Attribute) SetOptional() Attribute {
	a.Optional = true
	a.Required = false
	return a
}

func (a Int64Attribute) SetRequired() Attribute {
	a.Optional = false
	a.Required = true
	return a
}

func (a Int64Attribute) SetSensitive() Attribute {
	a.Sensitive = true
	return a
}

func (a Int64Attribute) SetComputed() Attribute {
	a.Computed = true
	return a
}

func (a Int64Attribute) SetReadOnly() Attribute {
	a.Computed = true
	a.Optional = false
	a.Required = false
	return a
}

func (a Int64Attribute) SetDeprecated(msg string) Attribute {
	a.DeprecationMessage = msg
	return a
}

func (a Int64Attribute) AddValidators(v any) Attribute {
	a.Validators = append(a.Validators, v.(validator.Int64))
	return a
}

type BoolAttribute struct {
	Optional           bool
	Required           bool
	Sensitive          bool
	Computed           bool
	DeprecationMessage string
	Validators         []validator.Bool
}

func (a BoolAttribute) ToDataSourceAttribute() dataschema.Attribute {
	return dataschema.BoolAttribute{Optional: a.Optional, Required: a.Required, Sensitive: a.Sensitive, DeprecationMessage: a.DeprecationMessage, Computed: a.Computed, Validators: a.Validators}
}

func (a BoolAttribute) ToResourceAttribute() schema.Attribute {
	return schema.BoolAttribute{Optional: a.Optional, Required: a.Required, Sensitive: a.Sensitive, DeprecationMessage: a.DeprecationMessage, Computed: a.Computed, Validators: a.Validators}
}

func (a BoolAttribute) SetOptional() Attribute {
	a.Optional = true
	a.Required = false
	return a
}

func (a BoolAttribute) SetRequired() Attribute {
	a.Optional = false
	a.Required = true
	return a
}

func (a BoolAttribute) SetSensitive() Attribute {
	a.Sensitive = true
	return a
}

func (a BoolAttribute) SetComputed() Attribute {
	a.Computed = true
	return a
}

func (a BoolAttribute) SetReadOnly() Attribute {
	a.Computed = true
	a.Optional = false
	a.Required = false
	return a
}

func (a BoolAttribute) SetDeprecated(msg string) Attribute {
	a.DeprecationMessage = msg
	return a
}

func (a BoolAttribute) AddValidators(v any) Attribute {
	a.Validators = append(a.Validators, v.(validator.Bool))
	return a
}

type MapAttribute struct {
	ElementType        attr.Type
	Optional           bool
	Required           bool
	Sensitive          bool
	Computed           bool
	DeprecationMessage string
	Validators         []validator.Map
}

func (a MapAttribute) ToDataSourceAttribute() dataschema.Attribute {
	return dataschema.MapAttribute{ElementType: a.ElementType, Optional: a.Optional, Required: a.Required, Sensitive: a.Sensitive, DeprecationMessage: a.DeprecationMessage, Computed: a.Computed, Validators: a.Validators}
}

func (a MapAttribute) ToResourceAttribute() schema.Attribute {
	return schema.MapAttribute{ElementType: a.ElementType, Optional: a.Optional, Required: a.Required, Sensitive: a.Sensitive, DeprecationMessage: a.DeprecationMessage, Computed: a.Computed, Validators: a.Validators}
}

func (a MapAttribute) SetOptional() Attribute {
	a.Optional = true
	a.Required = false
	return a
}

func (a MapAttribute) SetRequired() Attribute {
	a.Optional = false
	a.Required = true
	return a
}

func (a MapAttribute) SetSensitive() Attribute {
	a.Sensitive = true
	return a
}

func (a MapAttribute) SetComputed() Attribute {
	a.Computed = true
	return a
}

func (a MapAttribute) SetReadOnly() Attribute {
	a.Computed = true
	a.Optional = false
	a.Required = false
	return a
}

func (a MapAttribute) SetDeprecated(msg string) Attribute {
	a.DeprecationMessage = msg
	return a
}

func (a MapAttribute) AddValidators(v any) Attribute {
	a.Validators = append(a.Validators, v.(validator.Map))
	return a
}

type ListAttribute struct {
	ElementType        attr.Type
	Optional           bool
	Required           bool
	Sensitive          bool
	Computed           bool
	DeprecationMessage string
	Validators         []validator.List
}

func (a ListAttribute) ToDataSourceAttribute() dataschema.Attribute {
	return dataschema.ListAttribute{ElementType: a.ElementType, Optional: a.Optional, Required: a.Required, Sensitive: a.Sensitive, DeprecationMessage: a.DeprecationMessage, Computed: a.Computed, Validators: a.Validators}
}

func (a ListAttribute) ToResourceAttribute() schema.Attribute {
	return schema.ListAttribute{ElementType: a.ElementType, Optional: a.Optional, Required: a.Required, Sensitive: a.Sensitive, DeprecationMessage: a.DeprecationMessage, Computed: a.Computed, Validators: a.Validators}
}

func (a ListAttribute) SetOptional() Attribute {
	a.Optional = true
	a.Required = false
	return a
}

func (a ListAttribute) SetRequired() Attribute {
	a.Optional = false
	a.Required = true
	return a
}

func (a ListAttribute) SetSensitive() Attribute {
	a.Sensitive = true
	return a
}

func (a ListAttribute) SetComputed() Attribute {
	a.Computed = true
	return a
}

func (a ListAttribute) SetReadOnly() Attribute {
	a.Computed = true
	a.Optional = false
	a.Required = false
	return a
}

func (a ListAttribute) SetDeprecated(msg string) Attribute {
	a.DeprecationMessage = msg
	return a
}

func (a ListAttribute) AddValidators(v any) Attribute {
	a.Validators = append(a.Validators, v.(validator.List))
	return a
}

type SingleNestedAttribute struct {
	Attributes         map[string]Attribute
	Optional           bool
	Required           bool
	Sensitive          bool
	Computed           bool
	DeprecationMessage string
	Validators         []validator.Object
}

func (a SingleNestedAttribute) ToDataSourceAttribute() dataschema.Attribute {
	return dataschema.SingleNestedAttribute{Attributes: ToDataSourceAttributeMap(a.Attributes), Optional: a.Optional, Required: a.Required, Sensitive: a.Sensitive, DeprecationMessage: a.DeprecationMessage, Computed: a.Computed, Validators: a.Validators}
}

func (a SingleNestedAttribute) ToResourceAttribute() schema.Attribute {
	return schema.SingleNestedAttribute{Attributes: ToResourceAttributeMap(a.Attributes), Optional: a.Optional, Required: a.Required, Sensitive: a.Sensitive, DeprecationMessage: a.DeprecationMessage, Computed: a.Computed, Validators: a.Validators}
}

func (a SingleNestedAttribute) SetOptional() Attribute {
	a.Optional = true
	a.Required = false
	return a
}

func (a SingleNestedAttribute) SetRequired() Attribute {
	a.Optional = false
	a.Required = true
	return a
}

func (a SingleNestedAttribute) SetSensitive() Attribute {
	a.Sensitive = true
	return a
}

func (a SingleNestedAttribute) SetComputed() Attribute {
	a.Computed = true
	return a
}

func (a SingleNestedAttribute) SetReadOnly() Attribute {
	a.Computed = true
	a.Optional = false
	a.Required = false
	return a
}

func (a SingleNestedAttribute) SetDeprecated(msg string) Attribute {
	a.DeprecationMessage = msg
	return a
}

func (a SingleNestedAttribute) AddValidators(v any) Attribute {
	a.Validators = append(a.Validators, v.(validator.Object))
	return a
}

type ListNestedAttribute struct {
	NestedObject       NestedAttributeObject
	Optional           bool
	Required           bool
	Sensitive          bool
	Computed           bool
	DeprecationMessage string
	Validators         []validator.List
}

func (a ListNestedAttribute) ToDataSourceAttribute() dataschema.Attribute {
	return dataschema.ListNestedAttribute{NestedObject: a.NestedObject.ToDataSourceAttribute(), Optional: a.Optional, Required: a.Required, Sensitive: a.Sensitive, DeprecationMessage: a.DeprecationMessage, Computed: a.Computed, Validators: a.Validators}
}

func (a ListNestedAttribute) ToResourceAttribute() schema.Attribute {
	return schema.ListNestedAttribute{NestedObject: a.NestedObject.ToResourceAttribute(), Optional: a.Optional, Required: a.Required, Sensitive: a.Sensitive, DeprecationMessage: a.DeprecationMessage, Computed: a.Computed, Validators: a.Validators}
}

func (a ListNestedAttribute) SetOptional() Attribute {
	a.Optional = true
	a.Required = false
	return a
}

func (a ListNestedAttribute) SetRequired() Attribute {
	a.Optional = false
	a.Required = true
	return a
}

func (a ListNestedAttribute) SetSensitive() Attribute {
	a.Sensitive = true
	return a
}

func (a ListNestedAttribute) SetComputed() Attribute {
	a.Computed = true
	return a
}

func (a ListNestedAttribute) SetReadOnly() Attribute {
	a.Computed = true
	a.Optional = false
	a.Required = false
	return a
}

func (a ListNestedAttribute) SetDeprecated(msg string) Attribute {
	a.DeprecationMessage = msg
	return a
}

func (a ListNestedAttribute) AddValidators(v any) Attribute {
	a.Validators = append(a.Validators, v.(validator.List))
	return a
}

type NestedAttributeObject struct {
	Attributes map[string]Attribute
}

func (a NestedAttributeObject) ToDataSourceAttribute() dataschema.NestedAttributeObject {
	dataSourceAttributes := ToDataSourceAttributeMap(a.Attributes)

	return dataschema.NestedAttributeObject{
		Attributes: dataSourceAttributes,
	}
}

func (a NestedAttributeObject) ToResourceAttribute() schema.NestedAttributeObject {
	resourceAttributes := ToResourceAttributeMap(a.Attributes)

	return schema.NestedAttributeObject{
		Attributes: resourceAttributes,
	}
}

func ToDataSourceAttributeMap(attributes map[string]Attribute) map[string]dataschema.Attribute {
	dataSourceAttributes := make(map[string]dataschema.Attribute)

	for key, attribute := range attributes {
		dataSourceAttributes[key] = attribute.ToDataSourceAttribute()
	}

	return dataSourceAttributes
}

func ToResourceAttributeMap(attributes map[string]Attribute) map[string]schema.Attribute {
	resourceAttributes := make(map[string]schema.Attribute)

	for key, attribute := range attributes {
		resourceAttributes[key] = attribute.ToResourceAttribute()
	}

	return resourceAttributes
}

type MapNestedAttribute struct {
	NestedObject       NestedAttributeObject
	Optional           bool
	Required           bool
	Sensitive          bool
	Computed           bool
	DeprecationMessage string
	Validators         []validator.Map
}

func (a MapNestedAttribute) ToDataSourceAttribute() dataschema.Attribute {
	return dataschema.MapNestedAttribute{NestedObject: a.NestedObject.ToDataSourceAttribute(), Optional: a.Optional, Required: a.Required, Sensitive: a.Sensitive, DeprecationMessage: a.DeprecationMessage, Computed: a.Computed, Validators: a.Validators}
}

func (a MapNestedAttribute) ToResourceAttribute() schema.Attribute {
	return schema.MapNestedAttribute{NestedObject: a.NestedObject.ToResourceAttribute(), Optional: a.Optional, Required: a.Required, Sensitive: a.Sensitive, DeprecationMessage: a.DeprecationMessage, Computed: a.Computed, Validators: a.Validators}
}

func (a MapNestedAttribute) SetOptional() Attribute {
	a.Optional = true
	a.Required = false
	return a
}

func (a MapNestedAttribute) SetRequired() Attribute {
	a.Optional = false
	a.Required = true
	return a
}

func (a MapNestedAttribute) SetSensitive() Attribute {
	a.Sensitive = true
	return a
}

func (a MapNestedAttribute) SetComputed() Attribute {
	a.Computed = true
	return a
}

func (a MapNestedAttribute) SetReadOnly() Attribute {
	a.Computed = true
	a.Optional = false
	a.Required = false
	return a
}

func (a MapNestedAttribute) SetDeprecated(msg string) Attribute {
	a.DeprecationMessage = msg
	return a
}

func (a MapNestedAttribute) AddValidators(v any) Attribute {
	a.Validators = append(a.Validators, v.(validator.Map))
	return a
}
