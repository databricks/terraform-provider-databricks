package common

import (
	"context"
	"log"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	uuidRegex = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
)

func StringIsUUID(s string) bool {
	return uuidRegex.MatchString(s)
}

func GetTerraformVersionFromContext(ctx context.Context) string {
	tfVersion := "unknown"
	p, ok := ctx.Value(Provider).(*schema.Provider)
	if ok {
		tfVersion = p.TerraformVersion
	}
	return tfVersion
}

func IsExporter(ctx context.Context) bool {
	return GetTerraformVersionFromContext(ctx) == "exporter"
}

func SuppressDiffWhitespaceChange(k, old, new string, d *schema.ResourceData) bool {
	log.Printf("[DEBUG] Suppressing diff for %v: old=%#v new=%#v", k, old, new)
	return strings.TrimSpace(old) == strings.TrimSpace(new)
}

func CopyMap[K comparable, V any](originalMap map[K]V) map[K]V {
	newMap := make(map[K]V)
	for key, value := range originalMap {
		newMap[key] = value
	}
	return newMap
}
