# Automatic Effective Field Mapping Pattern

## Overview

This document describes the generalized pattern for automatically copying values from `effective_*` fields to their corresponding input-only fields during exporter operations.

## Background

Many Databricks resources have a pattern where:
- Users configure input fields (e.g., `node_count`, `usage_policy_id`, `custom_tags`)
- The API returns `effective_*` versions of these fields (e.g., `effective_node_count`, `effective_usage_policy_id`, `effective_custom_tags`)
- The input fields are marked as "input-only" and are not returned by the API's Get operation

When exporting resources, this creates a problem: the exporter reads the resource state via the API (which only returns `effective_*` fields), but needs to generate Terraform configuration with the input fields.

## Solution

The solution consists of two parts implemented in `exporter/impl_lakebase.go`:

### 1. Automatic Field Copying (`copyEffectiveFieldsToInputFields`)

This function automatically discovers and copies all `effective_*` fields to their input counterparts:

```go
func copyEffectiveFieldsToInputFields(r *resource) {
    // Iterates through all fields in the resource schema
    // Finds fields starting with "effective_"
    // Checks if corresponding input field exists (without "effective_" prefix)
    // Copies value from effective field to input field
    // Note: Zero values are copied too, but the HCL generator skips them for optional fields
}
```

**Supported Types:**
- ✅ Simple types: `string`, `bool`, `int`, `int64`, `float32`, `float64`
- ✅ Zero values are handled by the HCL code generator (not filtered here)
- ❌ Complex types: Lists and objects require additional type conversion logic

**Examples of fields automatically handled:**
- `effective_node_count` → `node_count`
- `effective_usage_policy_id` → `usage_policy_id`
- `effective_enable_readable_secondaries` → `enable_readable_secondaries`
- `effective_capacity` → `capacity`
- `effective_enable_pg_native_login` → `enable_pg_native_login`
- `effective_retention_window_in_days` → `retention_window_in_days`

### 2. Schema-Based Omission Logic

In `exporter/importables.go`, the `ShouldOmitFieldUnified` callback checks if input-only fields with `effective_*` counterparts have non-zero values:

```go
ShouldOmitFieldUnified: func(ic *importContext, pathString string, fieldSchema FieldSchema, wrapper ResourceDataWrapper, r *resource) bool {
    effectiveFieldName := "effective_" + pathString
    effectiveFieldSchema := wrapper.GetSchema().GetField(effectiveFieldName)
    if effectiveFieldSchema != nil {
        // This is an input field that has an effective_* counterpart
        // Check if the value is actually a zero value for its type
        v, ok := wrapper.GetOk(pathString)
        if !ok {
            return true // Field not set, omit it
        }

        // Required fields should never be omitted, even if zero
        if fieldSchema.IsRequired() {
            return false
        }

        // Check if it's a zero value using reflection
        if v == nil {
            return true
        }
        rv := reflect.ValueOf(v)
        if rv.IsZero() {
            return true // Zero value, omit it (e.g., false for bool, 0 for int)
        }

        // Check against default value if one is defined
        if def := fieldSchema.GetDefault(); def != nil && reflect.DeepEqual(v, def) {
            return true
        }

        // Non-zero value, don't omit it
        return false
    }
    // Use default omission logic for other fields (e.g., omit computed-only fields)
    return DefaultShouldOmitFieldFuncWithAbstraction(ic, pathString, fieldSchema, wrapper, r)
},
```

**Why Use Reflection for Zero-Value Checking?**

The key insight is that `wrapper.GetOk()` returns `nonZero=true` even when a boolean field is set to `false`, because `GetOk()` indicates whether a field is **set**, not whether it's a zero value. For proper zero-value detection:

- Use `reflect.ValueOf(v).IsZero()` which correctly identifies zero values for all types:
  - `false` for booleans
  - `0` for integers
  - `""` for strings
  - `nil` for pointers/slices/maps

This ensures that fields like `enable_pg_native_login = false` are correctly omitted from the generated HCL, keeping it clean and minimal.

**Important: Required Fields**

The logic includes a critical check for required fields: `if fieldSchema.IsRequired() { return false }`. This ensures that required fields are **never omitted**, even if they have zero values. This is essential because Terraform requires these fields to be present in the configuration, regardless of their value.

## Usage

To use this pattern for a resource:

1. In the resource's `Import` function, call `copyEffectiveFieldsToInputFields(r)`:

```go
func importDatabaseInstance(ic *importContext, r *resource) error {
    // Automatically copy all effective_* fields to input fields
    copyEffectiveFieldsToInputFields(r)

    // ... rest of import logic (e.g., emit permissions)
    return nil
}
```

2. In the resource's `importable` definition, use the schema-based omission logic with reflection:

```go
"databricks_database_instance": {
    // ... other fields ...
    ShouldOmitFieldUnified: func(ic *importContext, pathString string, fieldSchema FieldSchema, wrapper ResourceDataWrapper, r *resource) bool {
        effectiveFieldName := "effective_" + pathString
        effectiveFieldSchema := wrapper.GetSchema().GetField(effectiveFieldName)
        if effectiveFieldSchema != nil {
            v, ok := wrapper.GetOk(pathString)
            if !ok {
                return true
            }
            // Required fields should never be omitted, even if zero
            if fieldSchema.IsRequired() {
                return false
            }
            if v == nil {
                return true
            }
            rv := reflect.ValueOf(v)
            if rv.IsZero() {
                return true // Omit zero values
            }
            if def := fieldSchema.GetDefault(); def != nil && reflect.DeepEqual(v, def) {
                return true
            }
            return false // Don't omit non-zero values
        }
        return DefaultShouldOmitFieldFuncWithAbstraction(ic, pathString, fieldSchema, wrapper, r)
    },
},
```

## Benefits

1. **Automatic Discovery**: No need to hardcode field names - the solution automatically finds all `effective_*` → input field mappings
2. **Maintainable**: Works for any future fields that follow the `effective_*` pattern
3. **Type-Safe**: Handles type conversion for common types (int, int64, bool, string, etc.)
4. **Relies on Code Generator**: Zero/default values are handled by the HCL code generator (see `codegen.go:394-400`), keeping the copying logic simple

## Limitations

- **Complex Types**: Lists of objects (like `custom_tags`) require additional type conversion logic and are currently not supported. See [COMPLEX_TYPES_HANDLING.md](./COMPLEX_TYPES_HANDLING.md) for a detailed explanation of why complex types need special handling and what would be required to support them.
- **Custom Logic**: If a field requires special handling beyond simple copying, you'll need to add custom logic

## Testing

The `TestDatabaseInstanceExport` test verifies that:
- Input fields are correctly copied from their effective counterparts
- Fields are included in the generated HCL
- All simple-type fields work correctly

## Future Improvements

To support complex types like `custom_tags`, we would need to:
1. Detect the field type (list, object, etc.)
2. Convert the effective value to the expected input type
3. Handle nested structures appropriately

This could be added to `copyEffectiveFieldsToInputFields` if needed in the future.
