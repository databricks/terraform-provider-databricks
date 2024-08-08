package workspace

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"hash/fnv"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ReadContent to work with `content_base64` and `source` properties accordingly and set MD5 checksum
func ReadContent(d *schema.ResourceData) (content []byte, err error) {
	b64 := d.Get("content_base64").(string)
	if b64 == "" {
		content, err = common.ReadFileContent(d.Get("source").(string))
	} else {
		log.Printf("[INFO] Reading `content_base64` of %d bytes", len(b64))
		content, err = base64.StdEncoding.DecodeString(b64)
	}
	if err != nil {
		return
	}
	d.Set("md5", fmt.Sprintf("%x", md5.Sum(content)))
	log.Printf("[INFO] Setting file content hash to %s", d.Get("md5"))
	return
}

// MigrateV0 migrates from version 0.2.x state
func MigrateV0(ctx context.Context,
	rawState map[string]any,
	meta any) (map[string]any, error) {
	newState := map[string]any{}
	for k, v := range rawState {
		switch k {
		case "overwrite", "mkdirs", "validate_remote_file", "content_b64_md5":
			log.Printf("[INFO] Migrated from v0.2.x and removed %s from databricks_dbfs_file", k)
			continue
		case "source":
			newState["source"] = v
			if v != nil {
				if content, err := common.ReadFileContent(v.(string)); err == nil {
					newState["md5"] = fmt.Sprintf("%x", md5.Sum(content))
					log.Printf("[INFO] State of %s file is migrated from v0.2.x", newState["md5"])
				}
			}
		case "content":
			newState["content_base64"] = v
			if v != nil {
				if content, err := base64.StdEncoding.DecodeString(v.(string)); err == nil {
					newState["md5"] = fmt.Sprintf("%x", md5.Sum(content))
					log.Printf("[INFO] State of %s direct content is migrated from v0.2.x", newState["md5"])
				}
			}
		default:
			newState[k] = v
		}
	}
	return newState, nil
}

// FileContentSchemaWithoutPath returns common schema for file resources, but without path
func FileContentSchemaWithoutPath(extra map[string]*schema.Schema) map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"md5": {
			Type:     schema.TypeString,
			Default:  "different",
			Optional: true,
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				if _, err := ReadContent(d); err != nil {
					return false
				}
				log.Printf("[INFO] Suppressing %s diff: %v", d.Id(), old == d.Get("md5"))
				return old == d.Get("md5")
			},
		},
		"content_base64": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"source"},
			ValidateFunc:  validation.StringIsBase64,
		},
		"source": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"content_base64"},
			ValidateDiagFunc: func(i any, p cty.Path) diag.Diagnostics {
				v := i.(string)
				if _, err := os.Stat(v); os.IsNotExist(err) {
					return diag.Diagnostics{
						{
							Summary:       fmt.Sprintf("File %s does not exist", v),
							Severity:      diag.Error,
							AttributePath: p,
						},
					}
				}
				return nil
			},
		},
	}
	for k, v := range extra {
		s[k] = v
	}
	return s
}

// FileContentSchema returns common schema for file resources
func FileContentSchema(extra map[string]*schema.Schema) map[string]*schema.Schema {
	pathMap := map[string]*schema.Schema{
		"path": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
			ValidateDiagFunc: func(i any, p cty.Path) diag.Diagnostics {
				v := i.(string)
				if v == "" {
					return diag.Diagnostics{
						{
							Summary:       "Path must not be empty",
							Severity:      diag.Error,
							AttributePath: p,
						},
					}
				}
				if strings.HasPrefix(v, "Volume") {
					return diag.Diagnostics{
						{
							Summary:       "Path should start with /Volumes",
							Severity:      diag.Error,
							AttributePath: p,
						},
					}
				}
				if strings.HasPrefix(v, "dbfs:") {
					return diag.Diagnostics{
						{
							Summary: "Remove `dbfs:` prefix",
							Detail: fmt.Sprintf("Even though `dbfs:` is a valid filesystem prefix, "+
								"we recommend to remove it to avoid confusion and "+
								"recreation of resource: %s", strings.ReplaceAll(v, "dbfs:", "")),
							Severity:      diag.Error,
							AttributePath: p,
						},
					}
				}
				clean := filepath.ToSlash(filepath.Clean(v))
				if v != clean {
					return diag.Diagnostics{
						{
							Summary: "Clean path required",
							Detail: fmt.Sprintf(
								"Replace value with %s to avoid resource replacement",
								clean),
							Severity:      diag.Error,
							AttributePath: p,
						},
					}
				}
				return nil
			},
		},
	}
	s := FileContentSchemaWithoutPath(pathMap)
	for k, v := range extra {
		s[k] = v
	}
	return s
}

// PathListHash ...
func PathListHash(v any) int {
	h := fnv.New32a()
	m := v.(map[string]any)
	var err error
	if v, ok := m["path"]; ok {
		_, err = h.Write([]byte(v.(string)))
		if err != nil {
			return 0
		}
	}
	c := int(h.Sum32())
	if -c >= 0 {
		return -c
	}
	return c
}
