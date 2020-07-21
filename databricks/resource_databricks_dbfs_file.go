package databricks

import (
	"bufio"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"

	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pkg/errors"
)

func resourceDBFSFile() *schema.Resource {
	return &schema.Resource{
		Create: resourceDBFSFileCreate,
		Read:   resourceDBFSFileRead,
		Delete: resourceDBFSFileDelete,
		Update: resourceDBFSFileUpdate,

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
	}
}

func resourceDBFSFileCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DatabricksClient)
	path := d.Get("path").(string)
	overwrite := d.Get("overwrite").(bool)
	//checksum := d.Get("content_b64_md5").(string)
	mkdirs := d.Get("mkdirs").(bool)

	if mkdirs {
		err := handleDBFSParentDirs(client, path)
		if err != nil {
			return err
		}
	}

	content := d.Get("content").(string)
	source := d.Get("source").(string)

	var base64Content string
	if !reflect.ValueOf(source).IsZero() {
		b64, err := getLocalFileB64(source)
		if err != nil {
			return err
		}
		base64Content = b64
	} else {
		base64Content = content
	}

	err := client.DBFS().Create(path, overwrite, base64Content)
	if err != nil {
		return err
	}

	d.SetId(path)

	return resourceDBFSFileRead(d, m)
}

func resourceDBFSFileRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DatabricksClient)

	fileInfo, err := client.DBFS().Status(id)
	if err != nil {
		if e, ok := err.(service.APIError); ok && e.IsMissing() {
			log.Printf("missing resource due to error: %v\n", e)
			d.SetId("")
			return nil
		}
		return err
	}
	err = d.Set("path", fileInfo.Path)
	if err != nil {
		return err
	}
	err = d.Set("file_size", fileInfo.FileSize)

	if validateRemoteFile, ok := d.GetOk("validate_remote_file"); ok {
		validateFile := validateRemoteFile.(bool)
		if validateFile {
			log.Println("Validating remote file!")
			data, err := client.DBFS().Read(id)
			if err != nil {
				return err
			}
			// Both source & content ways of providing data will validate the checksum
			contentCheckSum, err := getMD5(data)
			if err != nil {
				return err
			}
			err = d.Set("content_b64_md5", contentCheckSum)
			if err != nil {
				return err
			}
		}
	}

	return err
}

func resourceDBFSFileUpdate(d *schema.ResourceData, m interface{}) error {
	overwrite := d.Get("overwrite").(bool)
	mkdirs := d.Get("mkdirs").(bool)
	validateRemoteFile := d.Get("validate_remote_file").(bool)

	err := d.Set("overwrite", overwrite)
	if err != nil {
		return err
	}
	err = d.Set("mkdirs", mkdirs)
	if err != nil {
		return err
	}
	err = d.Set("validate_remote_file", validateRemoteFile)
	if err != nil {
		return err
	}

	return resourceDBFSFileRead(d, m)
}

func getLocalFileB64(absPath string) (base64Content string, err error) {
	f, err := os.Open(absPath)
	if err != nil {
		return base64Content, err
	}
	defer f.Close()

	// Read entire file into byte slice.
	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return base64Content, err
	}

	// Encode as base64.
	base64Content = base64.StdEncoding.EncodeToString(content)
	return
}

func getMD5(text string) (string, error) {
	hasher := md5.New()
	_, err := hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil)), err
}

func resourceDBFSFileDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DatabricksClient)
	err := client.DBFS().Delete(id, false)
	return err
}

// handleDBFSParentDirs handles the different branch paths to create dbfs parent directories
func handleDBFSParentDirs(client *service.DatabricksClient, path string) error {
	parentDir, err := GetParentDirPath(path)
	switch err {
	// Notebook path is root directory so no need to make directory and there is no parent
	case DirPathRootDirError:
		return nil
	// Notebook path is empty thus a valid error
	case PathEmptyError:
		return err
	//	Notebook path is valid and has a parent directory
	case nil:
		dbfsObj, err := client.DBFS().Status(parentDir)
		// Notebook parent path is not a directory and it could be a notebook
		if err == nil && !dbfsObj.IsDir {
			return fmt.Errorf("parent path: %s caused error: %w", parentDir, ParentPathIsFileError)
		}
		// Parent path is missing thus needs to be created as a directory
		if e, ok := err.(service.APIError); ok && e.IsMissing() {
			err := client.DBFS().Mkdirs(parentDir)
			if err != nil {
				return errors.Wrapf(err, "failed to create directory %s", parentDir)
			}
		}
	}
	return nil
}

var ParentPathIsFileError = errors.New("parent path should be a directory and not a file")
