package common

import (
	"bufio"
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	uuidRegex                = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	TerraformBugErrorMessage = "This is a bug. Please report this to the maintainers at github.com/databricks/terraform-provider-databricks."
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

func MustInt64(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}

// GetInt64 returns the data for the given key and cast it to int64
func GetInt64(d *schema.ResourceData, key string) int64 {
	return int64(d.Get(key).(int))
}

// Reads the file content from a given path
func ReadFileContent(source string) ([]byte, error) {
	log.Printf("[INFO] Reading %s", source)
	f, err := os.Open(source)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	return io.ReadAll(reader)
}

// Calculates MD5 hash of the given content
func CalculateMd5Hash(content []byte) string {
	return fmt.Sprintf("%x", md5.Sum(content))
}

// Reads content from a JSON string or a file path and returns the content and its MD5 hash
func ReadSerializedJsonContent(jsonStr, filePath string) (serJSON string, md5Hash string, err error) {
	var content []byte
	if filePath != "" {
		content, err = ReadFileContent(filePath)
		if err != nil {
			return "", "", err
		}
	} else {
		log.Printf("[INFO] Reading `serialized_json` of %d bytes", len(jsonStr))
		content = []byte(jsonStr)
	}
	md5Hash = CalculateMd5Hash(content)
	return string(content), md5Hash, nil
}

// Suppress the error if it is 404
func IgnoreNotFoundError(err error) error {
	var apiErr *apierr.APIError
	if !errors.As(err, &apiErr) {
		return err
	}
	if apiErr.StatusCode == 404 {
		return nil
	}
	if strings.Contains(apiErr.Message, "does not exist") {
		return nil
	}
	return err
}
