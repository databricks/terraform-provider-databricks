// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
package converters

import (
	"strings"
)

var reservedNames = map[string]struct{}{
	"Type": {},
}

func toGoSdkName(tfSdkName string) string {
	return strings.TrimSuffix(tfSdkName, "_")
}

func toTfSdkName(goSdkName string) string {
	if _, ok := reservedNames[goSdkName]; ok {
		return goSdkName + "_"
	}
	return goSdkName
}
