package common

import (
	"context"
	"log"
	"regexp"
	"strconv"
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

func removeEmptyLines(text string) string {
	lines := strings.Split(text, "\n")
	var nonEmptyLines []string
	for _, line := range lines {
		if line != "" {
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}
	return strings.Join(nonEmptyLines, "\n")
}

func SuppressDiffWhitespaceChange(k, old, new string, d *schema.ResourceData) bool {
	log.Printf("[DEBUG] Suppressing diff for %v: old=%#v new=%#v", k, old, new)
	return removeEmptyLines(strings.TrimSpace(old)) == removeEmptyLines(strings.TrimSpace(new))
}

func MustInt64(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}
