package databricks

import (
	"encoding/base64"
	"fmt"
	"hash/crc32"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/pkg/errors"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceNotebook() *schema.Resource {
	return &schema.Resource{
		Create: resourceNotebookCreate,
		Read:   resourceNotebookRead,
		Delete: resourceNotebookDelete,

		Schema: map[string]*schema.Schema{
			"content": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				StateFunc: func(i interface{}) string {
					base64String := i.(string)
					base64, err := convertBase64ToCheckSum(base64String)
					if err != nil {
						return ""
					}
					return base64
				},
				ValidateFunc: validation.StringIsBase64,
			},
			"path": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: ValidateNotebookPath,
			},
			"language": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(model.Scala),
					string(model.Python),
					string(model.R),
					string(model.SQL),
				}, false),
			},
			"overwrite": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},
			"mkdirs": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
				ForceNew: true,
			},
			"format": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  string(model.Source),
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					// Only supports source format as it is easiest to identify diff/delta
					string(model.Source),
				}, false),
			},
			"object_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"object_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceNotebookCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DatabricksClient)
	path := d.Get("path").(string)
	content := d.Get("content").(string)
	language := d.Get("language").(string)
	format := d.Get("format").(string)
	overwrite := d.Get("overwrite").(bool)
	mkdirs := d.Get("mkdirs").(bool)

	if mkdirs {
		parentDir, err := GetParentDirPath(path)
		switch err {
		// Notebook path is root directory so no need to make directory and there is no parent
		case DirPathRootDirError:
			break
		// Notebook path is empty thus a valid error
		case PathEmptyError:
			return err
		//	Notebook path is valid and has a parent directory
		case nil:
			workspaceObj, err := client.Notebooks().Read(parentDir)
			// Notebook parent path is not a directory and it could be a notebook
			if err == nil && workspaceObj.ObjectType != model.Directory {
				return fmt.Errorf("parent path %s should be a directory and not a %s",
					workspaceObj.Path,
					workspaceObj.ObjectType,
				)
			}
			// Parent path is missing thus needs to be created as a directory
			if e, ok := err.(service.APIError); ok && e.IsMissing() {
				err := client.Notebooks().Mkdirs(parentDir)
				if err != nil {
					return errors.Wrapf(err, "failed to create directory %s", parentDir)
				}
			}
		}
	}

	err := client.Notebooks().Create(path, content, model.Language(language), model.ExportFormat(format), overwrite)
	if err != nil {
		return err
	}
	d.SetId(path)

	err = d.Set("format", format)
	if err != nil {
		return err
	}
	err = d.Set("overwrite", overwrite)
	if err != nil {
		return err
	}
	err = d.Set("mkdirs", mkdirs)
	if err != nil {
		return err
	}
	return resourceNotebookRead(d, m)
}

func resourceNotebookRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DatabricksClient)
	format := d.Get("format").(string)
	notebookData, err := client.Notebooks().Export(id, model.ExportFormat(format))
	if err != nil {
		if e, ok := err.(service.APIError); ok && e.IsMissing() {
			log.Printf("missing resource due to error: %v\n", e)
			d.SetId("")
			return nil
		}
		return err
	}
	notebookInfo, err := client.Notebooks().Read(id)
	if err != nil {
		return err
	}
	err = d.Set("path", id)
	if err != nil {
		return err
	}

	crc, err := convertBase64ToCheckSum(notebookData)
	if err != nil {
		return err
	}
	err = d.Set("content", crc)
	if err != nil {
		return err
	}
	err = d.Set("language", string(notebookInfo.Language))
	if err != nil {
		return err
	}
	err = d.Set("object_id", int(notebookInfo.ObjectID))
	if err != nil {
		return err
	}
	err = d.Set("object_type", string(notebookInfo.ObjectType))

	return err
}

func resourceNotebookDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DatabricksClient)

	return resource.Retry(5*time.Minute, func() *resource.RetryError {
		err := client.Notebooks().Delete(id, true)
		if err == nil {
			return nil
		}
		var e service.APIError
		if errors.As(err, &e) && e.IsTooManyRequests() {
			// Wait for requests to clear up
			baseDuration := 250 * time.Millisecond
			jitter := time.Duration(rand.Intn(1000)) * time.Millisecond
			// So all threads dont wake up at once
			time.Sleep(baseDuration + jitter)
			return resource.RetryableError(fmt.Errorf("request rate limit hit %w", err))
		}
		return resource.NonRetryableError(err)
	})
}

func convertBase64ToCheckSum(b64 string) (string, error) {
	dataArr, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return "error", errors.Wrap(err, "error while trying to decode base64 content")
	}
	return strconv.Itoa(int(crc32.ChecksumIEEE(normalizeNotebookSource(dataArr)))), nil
}

// This will normalize the notebooks to only validate the contents of the notebook that is editable
// and not be impacted by autogenerated text by databricks import and export
func normalizeNotebookSource(dataArr []byte) []byte {
	scalaMagic := "// MAGIC "
	pythonRMagic := "# MAGIC "
	sqlMagic := "-- MAGIC "
	filteredDataArr := filter(strings.Split(string(dataArr), "\n"), func(s string) bool {
		trimmedS := strings.TrimRight(s, " ")
		scalaMagicStatements := trimmedS != strings.Trim(scalaMagic, " ")
		pythonRMagicStatements := trimmedS != strings.Trim(pythonRMagic, " ")
		sqlMagicStatements := trimmedS != strings.Trim(sqlMagic, " ")
		ignoreScalaCmd := trimmedS != "// COMMAND ----------"
		ignorePythonRCmd := trimmedS != "# COMMAND ----------"
		ignoreSqlCmd := trimmedS != "-- COMMAND ----------"
		ignoreNotebookHeader := !strings.HasSuffix(trimmedS, "Databricks notebook source")
		return scalaMagicStatements && pythonRMagicStatements && sqlMagicStatements &&
			ignoreScalaCmd && ignorePythonRCmd && ignoreSqlCmd && ignoreNotebookHeader && trimmedS != ""
	})
	transformedDataArr := transform(filteredDataArr, func(s string) string {
		if strings.HasPrefix(s, scalaMagic) {
			return strings.TrimPrefix(s, scalaMagic)
		}
		if strings.HasPrefix(s, pythonRMagic) {
			return strings.TrimPrefix(s, pythonRMagic)
		}
		if strings.HasPrefix(s, sqlMagic) {
			return strings.TrimPrefix(s, sqlMagic)
		}
		return s
	})
	return []byte(strings.Join(transformedDataArr, "\n"))
}

func filter(ss []string, test func(string) bool) (ret []string) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func transform(ss []string, t func(string) string) (ret []string) {
	for _, s := range ss {
		ret = append(ret, t(s))
	}
	return
}

func ValidateNotebookPath(val interface{}, key string) (warns []string, errs []error) {
	v := val.(string)
	switch {
	case v == "":
		errs = append(errs, fmt.Errorf("%s is empty must have a value", key))
		fallthrough
	case !strings.HasPrefix(v, "/"):
		errs = append(errs, fmt.Errorf("%s must start with /, got: %s", key, v))
	}

	return nil, errs
}
