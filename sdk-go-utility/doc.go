// Package sdkgoutility lets SDKv2 resources work with the structs of the
// modular Databricks SDK, github.com/databricks/sdk-go.
//
// The reflection helpers in package common (StructToSchema, DataToStructPointer
// and StructToData) read the name and the optionality of a field from its json
// tag and treat every pointer as a nested block. Modular SDK structs carry no
// json tag, use pointers for optional scalars and named string types for
// enums, so those helpers can't work on them. The helpers in this package do
// the same jobs for modular SDK structs.
//
// Each helper lives in its own file, next to its tests.
package sdkgoutility
