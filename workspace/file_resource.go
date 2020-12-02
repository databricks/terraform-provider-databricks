package workspace

import (
	"bufio"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ReadContent to work with `content_base64` and `source` properties accordingly
func ReadContent(d *schema.ResourceData) (content []byte, err error) {
	b64 := d.Get("content_base64").(string)
	source := d.Get("source").(string)
	if b64 == "" {
		log.Printf("[INFO] Reading %s", source)
		f, rre := os.Open(source)
		if rre != nil {
			err = rre
			return
		}
		// TODO: size error
		defer f.Close()
		reader := bufio.NewReader(f)
		content, err = ioutil.ReadAll(reader)
	} else {
		log.Printf("[INFO] Reading `content_base64` of %d bytes", len(b64))
		content, err = base64.StdEncoding.DecodeString(b64)
	}
	if err != nil {
		return
	}
	// TODO: file size
	d.Set("md5", fmt.Sprintf("%x", md5.Sum(content)))
	log.Printf("[INFO] Setting file content hash to %s", d.Get("md5"))
	return
}

// FileContentSchema returns common schema for file resources
func FileContentSchema(extra map[string]*schema.Schema) map[string]*schema.Schema {
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
			ValidateDiagFunc: func(i interface{}, p cty.Path) diag.Diagnostics {
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
		"path": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
			ValidateDiagFunc: func(i interface{}, p cty.Path) diag.Diagnostics {
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
	for k, v := range extra {
		s[k] = v
	}
	return s
}

// PathListHash ...
func PathListHash(v interface{}) int {
	h := fnv.New32a()
	m := v.(map[string]interface{})
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
