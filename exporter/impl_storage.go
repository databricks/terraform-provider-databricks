package exporter

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/databricks/terraform-provider-databricks/storage"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/zclconf/go-cty/cty"
)

var (
	adlsGen2Regex = regexp.MustCompile(`^(abfss?)://([^@]+)@([^.]+)\.(?:[^/]+)(/.*)?$`)
	adlsGen1Regex = regexp.MustCompile(`^(adls?)://([^.]+)\.(?:[^/]+)(/.*)?$`)
	wasbsRegex    = regexp.MustCompile(`^(wasbs?)://([^@]+)@([^.]+)\.(?:[^/]+)(/.*)?$`)
	s3Regex       = regexp.MustCompile(`^(s3a?)://([^/]+)(/.*)?$`)
	gsRegex       = regexp.MustCompile(`^gs://([^/]+)(/.*)?$`)
)

func generateMountBody(ic *importContext, body *hclwrite.Body, r *resource) error {
	mount := ic.mountMap[r.ID]

	b := body.AppendNewBlock("resource", []string{r.Resource, r.Name}).Body()
	b.SetAttributeValue("name", cty.StringVal(strings.Replace(r.ID, "/mnt/", "", 1)))
	if res := s3Regex.FindStringSubmatch(mount.URL); res != nil {
		block := b.AppendNewBlock("s3", nil).Body()
		block.SetAttributeValue("bucket_name", cty.StringVal(res[2]))
		if mount.InstanceProfile != "" {
			block.SetAttributeValue("instance_profile", cty.StringVal(mount.InstanceProfile))
		} else if mount.ClusterID != "" {
			b.SetAttributeValue("cluster_id", cty.StringVal(mount.ClusterID))
		}
	} else if res := gsRegex.FindStringSubmatch(mount.URL); res != nil {
		block := b.AppendNewBlock("gs", nil).Body()
		block.SetAttributeValue("bucket_name", cty.StringVal(res[1]))
		if mount.ClusterID != "" {
			b.SetAttributeValue("cluster_id", cty.StringVal(mount.ClusterID))
		}
	} else if res := adlsGen2Regex.FindStringSubmatch(mount.URL); res != nil {
		containerName := res[2]
		storageAccountName := res[3]
		block := b.AppendNewBlock("abfs", nil).Body()
		block.SetAttributeValue("container_name", cty.StringVal(containerName))
		block.SetAttributeValue("storage_account_name", cty.StringVal(storageAccountName))
		if res[4] != "" && res[4] != "/" {
			block.SetAttributeValue("directory", cty.StringVal(res[4]))
		}

		varName := ic.regexFix("_"+storageAccountName+"_"+containerName+"_abfs", ic.nameFixes)
		textStr := fmt.Sprintf(" for mounting ADLSv2 resource %s://%s@%s",
			res[1], containerName, storageAccountName)

		block.SetAttributeRaw("client_id", ic.variable(
			"client_id"+varName, "Client ID"+textStr))
		block.SetAttributeRaw("tenant_id", ic.variable(
			"tenant_id"+varName, "Tenant ID"+textStr))
		block.SetAttributeRaw("client_secret_scope", ic.variable(
			"client_secret_scope"+varName,
			"Secret scope name that stores app client secret"+textStr))
		block.SetAttributeRaw("client_secret_key", ic.variable(
			"client_secret_key"+varName,
			"Key in secret scope that stores app client secret"+textStr))
		block.SetAttributeValue("initialize_file_system", cty.BoolVal(false))
	} else if res := adlsGen1Regex.FindStringSubmatch(mount.URL); res != nil {
		block := b.AppendNewBlock("adl", nil).Body()
		storageResourceName := res[2]
		block.SetAttributeValue("storage_resource_name", cty.StringVal(storageResourceName))
		if res[3] != "" && res[3] != "/" {
			block.SetAttributeValue("directory", cty.StringVal(res[3]))
		}
		varName := ic.regexFix("_"+storageResourceName+"_adl", ic.nameFixes)
		textStr := fmt.Sprintf(" for mounting ADLSv1 resource %s://%s", res[1], storageResourceName)

		block.SetAttributeRaw("client_id", ic.variable("client_id"+varName, "Client ID"+textStr))
		block.SetAttributeRaw("tenant_id", ic.variable("tenant_id"+varName, "Tenant IDs"+textStr))
		block.SetAttributeRaw("client_secret_scope", ic.variable(
			"client_secret_scope"+varName, "Secret scope name that stores app client secret"+textStr))
		block.SetAttributeRaw("client_secret_key", ic.variable(
			"client_secret_key"+varName, "Key in secret scope that stores app client secret"+textStr))
	} else if res := wasbsRegex.FindStringSubmatch(mount.URL); res != nil {
		containerName := res[2]
		storageAccountName := res[3]
		block := b.AppendNewBlock("wasb", nil).Body()
		block.SetAttributeValue("container_name", cty.StringVal(containerName))
		block.SetAttributeValue("storage_account_name", cty.StringVal(storageAccountName))
		if res[4] != "" && res[4] != "/" {
			block.SetAttributeValue("directory", cty.StringVal(res[4]))
		}
		block.SetAttributeValue("auth_type", cty.StringVal("ACCESS_KEY"))

		varName := ic.regexFix("_"+storageAccountName+"_"+containerName+"_wasb", ic.nameFixes)
		textStr := fmt.Sprintf(" for mounting WASB resource %s://%s@%s",
			res[1], containerName, storageAccountName)

		block.SetAttributeRaw("token_secret_scope", ic.variable(
			"client_secret_scope"+varName,
			"Secret scope name that stores app client secret"+textStr))
		block.SetAttributeRaw("token_secret_key", ic.variable(
			"client_secret_key"+varName,
			"Key in secret scope that stores app client secret"+textStr))
	} else {
		return fmt.Errorf("no matching handler for: %s", mount.URL)
	}
	body.AppendNewline()

	return nil
}

func importDbfsFile(ic *importContext, r *resource) error {
	dbfsAPI := storage.NewDbfsAPI(ic.Context, ic.Client)
	content, err := dbfsAPI.Read(r.ID)
	if err != nil {
		return err
	}
	name := ic.Importables["databricks_dbfs_file"].Name(ic, r.Data)
	fileName, err := ic.saveFileIn("dbfs_files", name, content)
	log.Printf("Creating %s for %s", fileName, r)
	if err != nil {
		return err
	}
	r.Data.Set("source", fileName)
	return nil
}

func listMounts(ic *importContext) error {
	if !ic.mounts {
		return nil
	}
	if err := ic.refreshMounts(); err != nil {
		return err
	}
	for mountName, source := range ic.mountMap {
		if !ic.MatchesName(mountName) {
			continue
		}
		if strings.HasPrefix(source.URL, "s3a://") {
			log.Printf("[INFO] Emitting databricks_mount: %s", source.URL)
			if source.InstanceProfile != "" {
				ic.Emit(&resource{
					Resource: "databricks_instance_profile",
					ID:       source.InstanceProfile,
				})
			} else if source.ClusterID != "" {
				ic.Emit(&resource{
					Resource: "databricks_cluster",
					ID:       source.ClusterID,
				})
			}
		} else if strings.HasPrefix(source.URL, "gs://") {
			if source.ClusterID != "" {
				ic.Emit(&resource{
					Resource: "databricks_cluster",
					ID:       source.ClusterID,
				})
			}
		} else if res := adlsGen2Regex.FindStringSubmatch(source.URL); res != nil {
		} else if res := adlsGen1Regex.FindStringSubmatch(source.URL); res != nil {
		} else if res := wasbsRegex.FindStringSubmatch(source.URL); res != nil {
		} else {
			log.Printf("[INFO] No matching handler for: %s", source.URL)
			continue
		}
		log.Printf("[INFO] Emitting databricks_mount: %s", source.URL)
		ic.Emit(&resource{
			Resource: "databricks_mount",
			ID:       mountName,
			Data: ic.Resources["databricks_mount"].Data(
				&terraform.InstanceState{
					ID:         mountName,
					Attributes: map[string]string{},
				}),
		})

	}
	return nil
}
