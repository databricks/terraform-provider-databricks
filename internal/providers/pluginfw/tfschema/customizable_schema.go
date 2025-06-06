// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
package tfschema

import (
	"fmt"
	"reflect"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// CustomizableSchema is a wrapper struct on top of BaseSchemaBuilder that can be used to navigate through nested schema add customizations.
// The methods of CustomizableSchema that modify the underlying schema should return the same CustomizableSchema object to allow chaining.
type CustomizableSchema struct {
	attr BaseSchemaBuilder
}

// ConstructCustomizableSchema constructs a CustomizableSchema given a NestedBlockObject.
func ConstructCustomizableSchema(nestedObject NestedBlockObject) *CustomizableSchema {
	return &CustomizableSchema{attr: SingleNestedBlockBuilder{NestedObject: nestedObject}}
}

// ToAttributeMap converts CustomizableSchema into BaseSchemaBuilder.
func (s *CustomizableSchema) ToNestedBlockObject() NestedBlockObject {
	return attributeToNestedBlockObject(&s.attr)
}

// attributeToMap converts AttributeBuilder into a map from string to AttributeBuilder.
func attributeToNestedBlockObject(attr *BaseSchemaBuilder) NestedBlockObject {
	var res = NestedBlockObject{}
	switch attr := (*attr).(type) {
	case SingleNestedAttributeBuilder:
		res.Attributes = attr.Attributes
	case ListNestedAttributeBuilder:
		res.Attributes = attr.NestedObject.Attributes
	case MapNestedAttributeBuilder:
		res.Attributes = attr.NestedObject.Attributes
	case SingleNestedBlockBuilder:
		res.Attributes = attr.NestedObject.Attributes
		res.Blocks = attr.NestedObject.Blocks
	case ListNestedBlockBuilder:
		res.Attributes = attr.NestedObject.Attributes
		res.Blocks = attr.NestedObject.Blocks
	default:
		panic(fmt.Errorf("cannot convert to map, attribute is not nested"))
	}

	return res
}

func (s *CustomizableSchema) AddValidator(v any, path ...string) *CustomizableSchema {
	cb := func(attr BaseSchemaBuilder) BaseSchemaBuilder {
		switch a := attr.(type) {
		case BoolAttributeBuilder:
			return a.AddValidator(v.(validator.Bool))
		case Float64AttributeBuilder:
			return a.AddValidator(v.(validator.Float64))
		case Int64AttributeBuilder:
			return a.AddValidator(v.(validator.Int64))
		case ListAttributeBuilder:
			return a.AddValidator(v.(validator.List))
		case ListNestedAttributeBuilder:
			return a.AddValidator(v.(validator.List))
		case ListNestedBlockBuilder:
			return a.AddValidator(v.(validator.List))
		case MapAttributeBuilder:
			return a.AddValidator(v.(validator.Map))
		case MapNestedAttributeBuilder:
			return a.AddValidator(v.(validator.Map))
		case SingleNestedAttributeBuilder:
			return a.AddValidator(v.(validator.Object))
		case StringAttributeBuilder:
			return a.AddValidator(v.(validator.String))
		default:
			panic(fmt.Errorf("cannot add validator, attribute builder type is invalid: %s. %s", reflect.TypeOf(attr).String(), common.TerraformBugErrorMessage))
		}
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)

	return s
}

func (s *CustomizableSchema) AddPlanModifier(v any, path ...string) *CustomizableSchema {
	cb := func(attr BaseSchemaBuilder) BaseSchemaBuilder {
		switch a := attr.(type) {
		case BoolAttributeBuilder:
			return a.AddPlanModifier(v.(planmodifier.Bool))
		case Float64AttributeBuilder:
			return a.AddPlanModifier(v.(planmodifier.Float64))
		case Int64AttributeBuilder:
			return a.AddPlanModifier(v.(planmodifier.Int64))
		case ListAttributeBuilder:
			return a.AddPlanModifier(v.(planmodifier.List))
		case ListNestedAttributeBuilder:
			return a.AddPlanModifier(v.(planmodifier.List))
		case MapAttributeBuilder:
			return a.AddPlanModifier(v.(planmodifier.Map))
		case MapNestedAttributeBuilder:
			return a.AddPlanModifier(v.(planmodifier.Map))
		case SingleNestedAttributeBuilder:
			return a.AddPlanModifier(v.(planmodifier.Object))
		case StringAttributeBuilder:
			return a.AddPlanModifier(v.(planmodifier.String))
		case ListNestedBlockBuilder:
			return a.AddPlanModifier(v.(planmodifier.List))
		case SingleNestedBlockBuilder:
			return a.AddPlanModifier(v.(planmodifier.Object))
		default:
			panic(fmt.Errorf("cannot add planmodifier, attribute builder type is invalid: %s. %s", reflect.TypeOf(attr).String(), common.TerraformBugErrorMessage))
		}
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)

	return s
}

func (s *CustomizableSchema) SetOptional(path ...string) *CustomizableSchema {
	cb := func(attr BaseSchemaBuilder) BaseSchemaBuilder {
		switch a := attr.(type) {
		case AttributeBuilder:
			return a.SetOptional()
		default:
			panic(fmt.Errorf("SetOptional called on invalid attribute type: %s. %s", reflect.TypeOf(attr).String(), common.TerraformBugErrorMessage))
		}
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)

	return s
}

func (s *CustomizableSchema) SetRequired(path ...string) *CustomizableSchema {
	cb := func(attr BaseSchemaBuilder) BaseSchemaBuilder {
		switch a := attr.(type) {
		case AttributeBuilder:
			return a.SetRequired()
		default:
			panic(fmt.Errorf("SetRequired called on invalid attribute type: %s. %s", reflect.TypeOf(attr).String(), common.TerraformBugErrorMessage))
		}
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)

	return s
}

func (s *CustomizableSchema) SetSensitive(path ...string) *CustomizableSchema {
	cb := func(attr BaseSchemaBuilder) BaseSchemaBuilder {
		switch a := attr.(type) {
		case AttributeBuilder:
			return a.SetSensitive()
		default:
			panic(fmt.Errorf("SetSensitive called on invalid attribute type: %s. %s", reflect.TypeOf(attr).String(), common.TerraformBugErrorMessage))
		}
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)
	return s
}

func (s *CustomizableSchema) SetDeprecated(msg string, path ...string) *CustomizableSchema {
	cb := func(attr BaseSchemaBuilder) BaseSchemaBuilder {
		return attr.SetDeprecated(msg)
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)

	return s
}

func (s *CustomizableSchema) SetComputed(path ...string) *CustomizableSchema {
	cb := func(attr BaseSchemaBuilder) BaseSchemaBuilder {
		switch a := attr.(type) {
		case AttributeBuilder:
			return a.SetComputed()
		default:
			panic(fmt.Errorf("SetComputed called on invalid attribute type: %s. %s", reflect.TypeOf(attr).String(), common.TerraformBugErrorMessage))
		}
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)
	return s
}

// SetReadOnly sets the schema to be read-only (i.e. computed, non-optional).
// This should be used for fields that are not user-configurable but are returned
// by the platform.
func (s *CustomizableSchema) SetReadOnly(path ...string) *CustomizableSchema {
	cb := func(attr BaseSchemaBuilder) BaseSchemaBuilder {
		switch a := attr.(type) {
		case AttributeBuilder:
			return a.SetReadOnly()
		default:
			panic(fmt.Errorf("SetReadOnly called on invalid attribute type: %s. %s", reflect.TypeOf(attr).String(), common.TerraformBugErrorMessage))
		}
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)

	return s
}

// ConfigureAsSdkV2Compatible modifies the underlying schema to be compatible with SDKv2. This method must
// be called on all resources that were originally implemented using the SDKv2 and are migrated to the plugin
// framework.
func (s *CustomizableSchema) ConfigureAsSdkV2Compatible() *CustomizableSchema {
	nbo := s.attr.(SingleNestedBlockBuilder).NestedObject
	s.attr = SingleNestedBlockBuilder{NestedObject: convertAttributesToBlocks(nbo.Attributes, nbo.Blocks)}
	return s
}

// navigateSchemaWithCallback navigates through schema attributes and executes callback on the target, panics if path does not exist or invalid.
func navigateSchemaWithCallback(s *BaseSchemaBuilder, cb func(BaseSchemaBuilder) BaseSchemaBuilder, path ...string) {
	currentScm := s
	for i, p := range path {
		m := attributeToNestedBlockObject(currentScm)
		mAttr := m.Attributes
		mBlock := m.Blocks

		if v, ok := mAttr[p]; ok {
			if i == len(path)-1 {
				newV := cb(v).(AttributeBuilder)
				mAttr[p] = newV
				return
			}
			castedV := v.(BaseSchemaBuilder)
			currentScm = &castedV
		} else if v, ok := mBlock[p]; ok {
			if i == len(path)-1 {
				newV := cb(v).(BlockBuilder)
				mBlock[p] = newV
				return
			}
			castedV := v.(BaseSchemaBuilder)
			currentScm = &castedV
		} else {
			panic(fmt.Errorf("missing key %s", p))
		}

	}
	panic(fmt.Errorf("path %v is incomplete", path))
}
