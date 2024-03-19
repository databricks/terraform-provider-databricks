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

func SuppressDiffWhitespaceAndEmptyLines(k, old, new string, d *schema.ResourceData) bool {
	log.Printf("[DEBUG] Suppressing diff for %v due to potential whitespace and empty line differences: old=%#v new=%#v", k, old, new)

	cleanString := func(input string) string {
		lines := strings.Split(input, "\n")
		processedLines := []string{}
		for _, line := range lines {
			trimmedLine := strings.TrimSpace(line)
			if trimmedLine != "" {
				processedLines = append(processedLines, trimmedLine)
			}
		}
		return strings.Join(processedLines, " ")
	}

	cleanedOld := cleanString(old)
	cleanedNew := cleanString(new)

	return strings.TrimSpace(cleanedOld) == strings.TrimSpace(cleanedNew)
}
