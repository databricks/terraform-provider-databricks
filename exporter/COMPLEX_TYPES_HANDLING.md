# Why Complex Types Need Special Handling

## The Type System Mismatch

When we try to copy `effective_custom_tags` to `custom_tags`, here's what happens:

### What We Get (Go SDK Type)
```go
// From the API response via Go SDK
[]database.CustomTag{
    {Key: "Environment", Value: "Production"},
    {Key: "Team", Value: "DataPlatform"},
}
```

### What Plugin Framework Expects (TF Type)
```go
// Plugin Framework type system
types.List[
    types.Object[
        "key":   types.String,
        "value": types.String,
    ],
]
```

### The Error We Get
```
Expected framework type: types.ListType[types.ObjectType["key":basetypes.StringType, "value":basetypes.StringType]]
Received framework type: basetypes.StringType  // Wrong! Because convertGoToPluginFrameworkType fell back to string
```

## Why Simple Types Work

For simple types, the `convertGoToPluginFrameworkType` function (in `exporter/abstractions.go:698-719`) can convert directly:

```go
func convertGoToPluginFrameworkType(value interface{}) attr.Value {
    v := reflect.ValueOf(value)
    switch v.Kind() {
    case reflect.String:
        return types.StringValue(v.String())    // ✅ Works!
    case reflect.Int64:
        return types.Int64Value(v.Int())        // ✅ Works!
    case reflect.Bool:
        return types.BoolValue(v.Bool())        // ✅ Works!
    default:
        // ❌ For complex types, falls back to string!
        return types.StringValue(fmt.Sprintf("%v", value))
    }
}
```

So when we copy:
- `effective_usage_policy_id` (string) → `usage_policy_id` ✅
- `effective_node_count` (int64) → `node_count` ✅
- `effective_enable_readable_secondaries` (bool) → `enable_readable_secondaries` ✅

## What's Needed for Complex Types

To properly handle `[]database.CustomTag`, we need to:

### 1. Detect the Schema Structure
```go
// Get the field schema to understand what type is expected
inputFieldSchema := schema.GetField("custom_tags")

// Check if it's a list
if inputFieldSchema.IsList() {
    // Get the element type (should be Object with "key" and "value" fields)
    elementType := inputFieldSchema.GetElementType()
}
```

### 2. Convert Each Element
```go
// For each CustomTag in the slice
for _, tag := range effectiveCustomTags {
    // Convert the struct to a types.Object
    obj := types.ObjectValueMust(
        map[string]attr.Type{
            "key":   types.StringType,
            "value": types.StringType,
        },
        map[string]attr.Value{
            "key":   types.StringValue(tag.Key),
            "value": types.StringValue(tag.Value),
        },
    )
}
```

### 3. Create the List
```go
// Wrap all objects in a types.List
listValue := types.ListValueMust(
    types.ObjectType{
        AttrTypes: map[string]attr.Type{
            "key":   types.StringType,
            "value": types.StringType,
        },
    },
    []attr.Value{obj1, obj2, ...},
)
```

### 4. Set the Value
```go
// Now we can set it!
wrapper.Set("custom_tags", listValue)  // ✅ Works!
```

## Implementation Approach

To implement this, we would need to enhance `copyEffectiveFieldsToInputFields`:

```go
func copyEffectiveFieldsToInputFields(r *resource) {
    // ... existing code ...

    // After detecting we have a value to copy
    if needsComplexTypeConversion(effectiveValue, inputFieldSchema) {
        // Convert complex types using schema information
        convertedValue, err := convertComplexTypeToPluginFramework(
            effectiveValue,
            inputFieldSchema,
        )
        if err != nil {
            log.Printf("[WARN] Failed to convert complex type %s: %v", fieldName, err)
            continue
        }
        wrapper.Set(inputFieldName, convertedValue)
    } else {
        // Simple types work as-is
        wrapper.Set(inputFieldName, effectiveValue)
    }
}

func convertComplexTypeToPluginFramework(value interface{}, schema FieldSchema) (interface{}, error) {
    switch {
    case schema.IsList():
        return convertSliceToPluginFrameworkList(value, schema)
    case schema.IsMap():
        return convertMapToPluginFrameworkMap(value, schema)
    case schema.IsObject():
        return convertStructToPluginFrameworkObject(value, schema)
    default:
        return value, nil
    }
}
```

## Why We Haven't Implemented This Yet

1. **Complexity**: Each nested structure needs recursive conversion with proper type information
2. **Schema Access**: Need to navigate nested schema definitions to get correct types
3. **Edge Cases**: Handling nested lists of objects, maps of lists, etc.
4. **Maintenance**: More code to maintain and test
5. **Low Priority**: Most important fields (usage_policy_id, node_count, etc.) are simple types

## Solution: Converter-Based Effective Fields Copying

The recommended approach is to use `copyEffectiveFieldsToInputFieldsWithConverters`, a generic function that automatically handles ALL types (simple and complex) by leveraging the existing converter infrastructure.

### Usage Example (Recommended)

```go
func importDatabaseInstance(ic *importContext, r *resource) error {
    // One function call handles ALL effective fields, including complex types!
    copyEffectiveFieldsToInputFieldsWithConverters[database_instance_resource.DatabaseInstance](
        ic, r, database.DatabaseInstance{})

    return nil
}
```

This single function call:
- ✅ Converts TF state → Go SDK struct
- ✅ Copies ALL `Effective*` fields to their input counterparts using reflection
- ✅ Handles both simple AND complex types automatically
- ✅ Converts back Go SDK struct → TF state
- ✅ No need for manual field-by-field handling

### How It Works Internally

The `copyEffectiveFieldsToInputFieldsWithConverters` function:

1. **Converts TF state → Go SDK struct** using `TfSdkToGoSdkStruct`
2. **Uses reflection** to find all fields starting with `Effective*` (e.g., `EffectiveNodeCount`, `EffectiveCustomTags`)
3. **Copies values** from effective fields to their input counterparts (e.g., `EffectiveNodeCount` → `NodeCount`)
4. **Converts back Go SDK struct → TF state** using `GoSdkToTfSdkStruct`
5. **Writes updated state** back to the resource

This approach leverages the battle-tested converter infrastructure that already handles:
- ✅ **Simple types**: string, int, bool, float
- ✅ **Lists of primitives**: `[]string`, `[]int`, etc.
- ✅ **Lists of structs**: `[]database.CustomTag` → proper TF types
- ✅ **Maps**: `map[string]string` and other map types
- ✅ **Nested structs**: Recursively converted with proper type information
- ✅ **Nested lists**: Lists within objects, objects within lists, etc.

### Field Name Resolution

The conversion process automatically handles field name mapping:
- **First checks** the `tfsdk` struct tag (Plugin Framework resources)
- **Then checks** the `json` struct tag (Go SDK structs)
- **Falls back** to snake_case conversion if neither tag exists

This ensures proper field name mapping between Go struct fields and Terraform schema fields.

## Summary

**Q: Why can't we simply Get the effective value and put it into the counterpart?**

**A:** Because:
1. Go SDK returns **native Go types** (`[]database.CustomTag`)
2. Plugin Framework expects **TF types** (`types.List[types.Object[...]]`)
3. There's no automatic conversion for complex types

**Solution:** Use `copyEffectiveFieldsToInputFieldsWithConverters` which:
- Works at the Go SDK struct level (native types)
- Copies all effective fields using simple reflection
- Leverages existing converters for type conversion
- Handles ALL types automatically in one function call

This is now the **recommended approach** for all Plugin Framework resources with effective_* fields. No need for manual field-by-field copying or special handling of complex types!
