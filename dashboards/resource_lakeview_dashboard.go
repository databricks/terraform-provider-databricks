package dashboards

import (
	"bufio"
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"

	"time"

	"github.com/databricks/databricks-sdk-go/service/dashboards"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Reads the file content from a given path
func readFileContent(source string) ([]byte, error) {
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
func calculateMd5Hash(content []byte) string {
	return fmt.Sprintf("%x", md5.Sum(content))
}

// Reads content from a JSON string or a file path and returns the content and its MD5 hash
func readSerializedJsonContent(jsonStr, filePath string) (serJSON string, md5Hash string, err error) {
	var content []byte
	if filePath != "" {
		content, err = readFileContent(filePath)
		if err != nil {
			return "", "", err
		}
	} else {
		log.Printf("[INFO] Reading `serialized_json` of %d bytes", len(jsonStr))
		content = []byte(jsonStr)
	}
	md5Hash = calculateMd5Hash(content)
	return string(content), md5Hash, nil
}

// ResourceLakeviewDashboard manages lakeview dashboards
func ResourceLakeviewDashboard() common.Resource {
	s := common.StructToSchema(dashboards.Dashboard{},
		func(s map[string]*schema.Schema) map[string]*schema.Schema {
			s["create_time"].Computed = true
			s["dashboard_id"].Computed = true
			s["display_name"].Optional = false
			s["display_name"].Required = true
			s["etag"].Computed = true
			s["lifecycle_state"].Computed = true
			s["parent_path"].Optional = false
			s["parent_path"].Required = true
			s["parent_path"].ForceNew = true
			s["path"].Computed = true
			s["serialized_dashboard"].Optional = true
			s["serialized_dashboard"].ConflictsWith = []string{"file_path"}
			s["serialized_dashboard"].DiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
				_, ok := d.GetOk("file_path")
				if ok {
					return true
				}
				_, new_json_hash, _ := readSerializedJsonContent(new, "")
				// fmt.Println(new_json_hash, d.Get("md5").(string))
				// fmt.Println(new, d.Get("serialized_dashboard").(string), old)
				return d.Get("md5").(string) == new_json_hash && !d.Get("dashboard_change_detected").(bool)
				// return !d.Get("dashboard_change_detected").(bool)
			}
			s["update_time"].Computed = true
			s["warehouse_id"].Optional = false
			s["warehouse_id"].Required = true

			s["file_path"] = &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"serialized_dashboard"},
			}
			s["md5"] = &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			}
			s["dashboard_change_detected"] = &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			}
			s["embed_credentials"] = &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			}
			return s
		})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var new_dashboard dashboards.CreateDashboardRequest
			common.DataToStructPointer(d, s, &new_dashboard)
			content, md5Hash, err := readSerializedJsonContent(d.Get("serialized_dashboard").(string), d.Get("file_path").(string))
			if err != nil {
				return err
			}
			d.Set("md5", md5Hash)
			new_dashboard.SerializedDashboard = content
			created_dashboard, err := w.Lakeview.Create(ctx, new_dashboard)
			if err != nil {
				return err
			}

			time.Sleep(2 * time.Second)
			dashboard := dashboards.GetDashboardRequest{
				DashboardId: created_dashboard.DashboardId,
			}
			resp, err := w.Lakeview.Get(ctx, dashboard)
			if err != nil {
				return err
			}
			d.Set("etag", resp.Etag)

			// Publish the dashboard
			// fmt.Println("Publishing the dashboard", d.Get("embed_credentials").(bool))
			_, err = w.Lakeview.Publish(ctx, dashboards.PublishRequest{
				DashboardId:      created_dashboard.DashboardId,
				WarehouseId:      d.Get("warehouse_id").(string),
				EmbedCredentials: d.Get("embed_credentials").(bool),
				ForceSendFields:  []string{"EmbedCredentials"},
			})
			if err != nil {
				return err
			}

			d.SetId(created_dashboard.DashboardId)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			dashboard := dashboards.GetDashboardRequest{
				DashboardId: d.Id(),
			}
			resp, err := w.Lakeview.Get(ctx, dashboard)
			if err != nil {
				return err
			}
			// fmt.Println(d.Get("md5").(string))
			d.Set("dashboard_change_detected", (resp.Etag != d.Get("etag").(string)) || func() bool {
				if v, ok := d.GetOk("file_path"); ok {
					_, new_file_hash, _ := readSerializedJsonContent("", v.(string))
					return (new_file_hash != d.Get("md5").(string))
				}
				// if v, ok := d.GetOk("serialized_dashboard"); ok {
				// 	_, new_json_hash, _ := readSerializedJsonContent(v.(string), "")
				// 	fmt.Println(new_json_hash, d.Get("md5").(string))
				// 	fmt.Println(v.(string), d.Get("serialized_dashboard").(string))
				// 	return (new_json_hash != d.Get("md5").(string))
				// }
				return false
			}())

			// read published dashboard
			// published_dashboard, err := w.Lakeview.GetPublished(ctx, dashboards.GetPublishedDashboardRequest{
			// 	DashboardId: d.Id(),
			// })
			// if err != nil {
			// 	return err
			// }
			// fmt.Println(published_dashboard)

			return common.StructToData(resp, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var dashboard dashboards.UpdateDashboardRequest
			common.DataToStructPointer(d, s, &dashboard)
			dashboard.DashboardId = d.Get("dashboard_id").(string)
			content, md5Hash, err := readSerializedJsonContent(d.Get("serialized_dashboard").(string), d.Get("file_path").(string))
			if err != nil {
				return err
			}
			d.Set("md5", md5Hash)
			dashboard.SerializedDashboard = content
			resp, err := w.Lakeview.Update(ctx, dashboard)
			if err != nil {
				return err
			}
			d.Set("etag", resp.Etag)

			// Publish the dashboard
			_, err = w.Lakeview.Publish(ctx, dashboards.PublishRequest{
				DashboardId:      d.Id(),
				WarehouseId:      d.Get("warehouse_id").(string),
				EmbedCredentials: d.Get("embed_credentials").(bool),
				ForceSendFields:  []string{"EmbedCredentials"},
			})
			if err != nil {
				return err
			}

			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			err = w.Lakeview.Trash(ctx, dashboards.TrashDashboardRequest{
				DashboardId: d.Id(),
			})
			if err != nil {
				return err
			}
			return nil
		},
	}
}
