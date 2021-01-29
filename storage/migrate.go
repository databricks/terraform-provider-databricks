package storage

import (
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DbfsFileV0 contains v0.2.x schema
func DbfsFileV0() cty.Type {
	return (&schema.Resource{
		Schema: map[string]*schema.Schema{
			"content": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"source": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"content"},
			},
			"content_b64_md5": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"path": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"overwrite": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"mkdirs": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"validate_remote_file": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"file_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}).CoreConfigSchema().ImpliedType()
}
