package databricks

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pkg/errors"
)

func resourceFolder() *schema.Resource {
	return &schema.Resource{
		Create: resourceFolderCreate,
		Read:   resourceFolderRead,
		Update: resourceFolderUpdate,
		Delete: resourceFolderDelete,

		Schema: map[string]*schema.Schema{
			"path": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"recursive_delete": {
				Type:     schema.TypeBool,
				Required: true,
			},
		},
	}
}

func resourceFolderCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	path := d.Get("path").(string)
	recursiveDelete := d.Get("recursive_delete").(bool)

	_, err := client.Notebooks().Read(path)
	if err == nil {
		return fmt.Errorf("object already exists in path: %s; please delete before creating the resource", path)
	}

	err = client.Notebooks().Mkdirs(path)
	if err != nil {
		return errors.Wrapf(err, "failed to create directory: %s", path)
	}

	d.SetId(path)

	err = d.Set("recursive_delete", recursiveDelete)
	if err != nil {
		return err
	}
	return resourceFolderRead(d, m)
}

func resourceFolderRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)

	_, err := client.Notebooks().Read(id)
	if err != nil {
		if isFolderMissing(err.Error()) {
			log.Printf("Missing folder with id: %s.", id)
			d.SetId("")
			return nil
		}
		return errors.Wrapf(err, "error reading folder with id %s", err)
	}

	err = d.Set("path", id)
	if err != nil {
		return err
	}

	return nil
}

func resourceFolderUpdate(d *schema.ResourceData, m interface{}) error {
	recursiveDelete := d.Get("recursive_delete").(bool)
	err := d.Set("recursive_delete", recursiveDelete)
	if err != nil {
		return err
	}
	return resourceFolderRead(d, m)
}

func resourceFolderDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	recursiveDelete := d.Get("recursive_delete").(bool)
	client := m.(*service.DBApiClient)

	return resource.Retry(5*time.Minute, func() *resource.RetryError {
		err := client.Notebooks().Delete(id, recursiveDelete)
		if err == nil {
			return nil
		}
		var e *service.DBApiError
		if errors.As(err, &e) && e.StatusCode == 429 {
			// Wait for requests to clear up
			baseDuration := 250 * time.Millisecond
			jitter := time.Duration(rand.Intn(1000)) * time.Millisecond
			// So all threads dont wake up at once
			time.Sleep(baseDuration + jitter)
			return resource.RetryableError(fmt.Errorf("request rate limit hit %w", err))
		}
		return nil
	})
}

func isFolderMissing(errorMsg string) bool {
	return strings.Contains(errorMsg, "RESOURCE_DOES_NOT_EXIST")
}
