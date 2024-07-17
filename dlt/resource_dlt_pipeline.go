package pipelines

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"

	// "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type createPipelineRequestStruct struct {
	pipelines.CreatePipeline
}

type updatePipelineRequestStruct struct {
	pipelines.EditPipeline
}

// *******************
// Original one only maintains Pipeline spec as the schema whereas this one maintains the entire pipeline struct
// *******************

type Pipeline struct {
	// not setting group:node_type for fields in pipelineCluster
	// Cluster->Autoscale->Mode (difference in type)
	// Cluster->AwsAttributes->Availability (difference in type)
	// Cluster->enable_local_disk_encryption (not in new)
	// In cluster, ID in old Id in new
	// Cluster->init_scripts->Abfss (difference in type)
	// Libraries->Whl does not exist in new
	// Why is serverless Optional in Serverless
	// Why is every field omitempty by default
	pipelines.PipelineSpec
	AllowDuplicateNames bool `json:"allow_duplicate_names,omitempty"`
	DryRun              bool `json:"dry_run,omitempty"`
	// An optional message detailing the cause of the pipeline state.
	Cause string `json:"cause,omitempty"`
	// The ID of the cluster that the pipeline is running on.
	ClusterId string `json:"cluster_id,omitempty"`
	// The username of the pipeline creator.
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// The health of a pipeline.
	Health pipelines.GetPipelineResponseHealth `json:"health,omitempty"`
	// The last time the pipeline settings were modified or created.
	LastModified int64 `json:"last_modified,omitempty"`
	// Status of the latest updates for the pipeline. Ordered with the newest
	// update first.
	LatestUpdates []pipelines.UpdateStateInfo `json:"latest_updates,omitempty"`
	// Username of the user that the pipeline will run on behalf of.
	RunAsUserName string `json:"run_as_user_name,omitempty"`
	// The pipeline state.
	State pipelines.PipelineState `json:"state,omitempty"`
	// URL field
	URL string `json:"url,omitempty"`
}

// Constants for PipelineStates
const (
	StateDeploying  pipelines.PipelineState = "DEPLOYING"
	StateStarting   pipelines.PipelineState = "STARTING"
	StateRunning    pipelines.PipelineState = "RUNNING"
	StateStopping   pipelines.PipelineState = "STOPPPING"
	StateDeleted    pipelines.PipelineState = "DELETED"
	StateRecovering pipelines.PipelineState = "RECOVERING"
	StateFailed     pipelines.PipelineState = "FAILED"
	StateResetting  pipelines.PipelineState = "RESETTING"
	StateIdle       pipelines.PipelineState = "IDLE"
)

type workspaceClient struct {
	*databricks.WorkspaceClient
}

func (w workspaceClient) Create(ctx context.Context, d *schema.ResourceData, timeout time.Duration) error {
	var createPipelineRequest createPipelineRequestStruct
	common.DataToStructPointer(d, pipelineSchema, &createPipelineRequest)

	for i := range createPipelineRequest.Clusters {
		cluster := &createPipelineRequest.Clusters[i]
		// TF Go SDK doesn't differentiate between the default and not set values.
		// If nothing is specified, DLT creates a cluster with enhanced autoscaling
		// from 1 to 5 nodes, which is different than sending a request for zero workers.
		// The solution here is to look for the Spark configuration to determine
		// if the user only wants a single node cluster (only master, no workers).
		if cluster.SparkConf["spark.databricks.cluster.profile"] == "singleNode" {
			cluster.ForceSendFields = append(cluster.ForceSendFields, "NumWorkers")
		}
	}

	createdPipeline, err := w.Pipelines.Create(ctx, createPipelineRequest.CreatePipeline)
	if err != nil {
		return err
	}
	id := createdPipeline.PipelineId
	err = w.waitForState(ctx, id, timeout, StateRunning)
	if err != nil {
		log.Printf("[INFO] Pipeline creation failed, attempting to clean up pipeline %s", id)
		err2 := w.Delete(ctx, id, timeout)
		if err2 != nil {
			log.Printf("[WARN] Unable to delete pipeline %s; this resource needs to be manually cleaned up", id)
			return fmt.Errorf("multiple errors occurred when creating pipeline. Error while waiting for creation: \"%v\"; error while attempting to clean up failed pipeline: \"%v\"", err, err2)
		}
		log.Printf("[INFO] Successfully cleaned up pipeline %s", id)
		return err
	}
	d.SetId(createdPipeline.PipelineId)
	return nil
}

func (w workspaceClient) Update(ctx context.Context, d *schema.ResourceData, timeout time.Duration) error {
	var updatePipelineRequest updatePipelineRequestStruct
	common.DataToStructPointer(d, pipelineSchema, &updatePipelineRequest)

	for i := range updatePipelineRequest.Clusters {
		cluster := &updatePipelineRequest.Clusters[i]
		// TF Go SDK doesn't differentiate between the default and not set values.
		// If nothing is specified, DLT creates a cluster with enhanced autoscaling
		// from 1 to 5 nodes, which is different than sending a request for zero workers.
		// The solution here is to look for the Spark configuration to determine
		// if the user only wants a single node cluster (only master, no workers).
		if cluster.SparkConf["spark.databricks.cluster.profile"] == "singleNode" {
			cluster.ForceSendFields = append(cluster.ForceSendFields, "NumWorkers")
		}
	}

	err := w.Pipelines.Update(ctx, updatePipelineRequest.EditPipeline)
	if err != nil {
		return err
	}
	return w.waitForState(ctx, d.Id(), timeout, StateRunning)
}

func (w workspaceClient) Read(ctx context.Context, id string) (*pipelines.GetPipelineResponse, error) {
	return w.Pipelines.Get(ctx, pipelines.GetPipelineRequest{
		PipelineId: id,
	})
}

func (w workspaceClient) Delete(ctx context.Context, id string, timeout time.Duration) error {
	err := w.Pipelines.Delete(ctx, pipelines.DeletePipelineRequest{
		PipelineId: id,
	})
	if err != nil {
		return err
	}
	return retry.RetryContext(ctx, timeout,
		func() *retry.RetryError {
			i, err := w.Read(ctx, id)
			if err != nil {
				if apierr.IsMissing(err) {
					return nil
				}
				return retry.NonRetryableError(err)
			}
			message := fmt.Sprintf("Pipeline %s is in state %s, not yet deleted", id, i.State)
			log.Printf("[DEBUG] %s", message)
			return retry.RetryableError(fmt.Errorf(message))
		})
}

func (w workspaceClient) waitForState(ctx context.Context, id string, timeout time.Duration, desiredState pipelines.PipelineState) error {
	return retry.RetryContext(ctx, timeout,
		func() *retry.RetryError {
			i, err := w.Read(ctx, id)
			// i, err := w.Pipelines.Get(ctx, pipelines.GetPipelineRequest{
			// 	PipelineId: id,
			// })
			if err != nil {
				return retry.NonRetryableError(err)
			}
			state := i.State
			if state == desiredState {
				return nil
			}
			if state == StateFailed {
				return retry.NonRetryableError(fmt.Errorf("pipeline %s has failed", id))
			}
			if !i.Spec.Continuous {
				// continuous pipelines just need a non-FAILED check
				return nil
			}
			message := fmt.Sprintf("Pipeline %s is in state %s, not yet in state %s", id, state, desiredState)
			log.Printf("[DEBUG] %s", message)
			return retry.RetryableError(fmt.Errorf(message))
		})
}

// func adjustForceSendFields(s *pipelines.CreatePipeline) {
// 	for i := range s.Clusters {
// 		cluster := &s.Clusters[i]
// 		// TF Go SDK doesn't differentiate between the default and not set values.
// 		// If nothing is specified, DLT creates a cluster with enhanced autoscaling
// 		// from 1 to 5 nodes, which is different than sending a request for zero workers.
// 		// The solution here is to look for the Spark configuration to determine
// 		// if the user only wants a single node cluster (only master, no workers).
// 		if cluster.SparkConf["spark.databricks.cluster.profile"] == "singleNode" {
// 			cluster.ForceSendFields = append(cluster.ForceSendFields, "NumWorkers")
// 		}
// 	}
// }

func suppressStorageDiff(k, old, new string, d *schema.ResourceData) bool {
	defaultStorageRegex := regexp.MustCompile(
		`^dbfs:/pipelines/[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)
	res := defaultStorageRegex.MatchString(old)
	if new == "" && res {
		log.Printf("[DEBUG] Suppressing diff for %v: platform=%#v config=%#v", k, old, new)
		return true
	}
	return false
}

func (Pipeline) Aliases() map[string]map[string]string {
	return map[string]map[string]string{
		"pipelines.Pipeline": {
			"clusters":      "cluster",
			"libraries":     "library",
			"notifications": "notification",
		},
	}

}

func (createPipelineRequestStruct) Aliases() map[string]map[string]string {
	return map[string]map[string]string{
		"pipelines.createPipelineRequestStruct": {
			"clusters":      "cluster",
			"libraries":     "library",
			"notifications": "notification",
		},
	}
}

func (updatePipelineRequestStruct) Aliases() map[string]map[string]string {
	return map[string]map[string]string{
		"pipelines.updatePipelineRequestStruct": {
			"clusters":      "cluster",
			"libraries":     "library",
			"notifications": "notification",
		},
	}
}

func (createPipelineRequestStruct) CustomizeSchema(s *common.CustomizableSchema) *common.CustomizableSchema {
	return s
}

func (updatePipelineRequestStruct) CustomizeSchema(s *common.CustomizableSchema) *common.CustomizableSchema {
	return s
}

func (Pipeline) CustomizeSchema(s *common.CustomizableSchema) *common.CustomizableSchema {
	// Required fields
	// s.SchemaPath("name").SetRequired()
	// s.SchemaPath("storage").SetRequired()
	// s.SchemaPath("cluster").SetRequired()
	// s.SchemaPath("continuous").SetRequired()

	// ForceNew fields
	s.SchemaPath("storage").SetForceNew()
	s.SchemaPath("catalog").SetForceNew()

	// Computed fields
	s.SchemaPath("id").SetComputed()
	s.SchemaPath("cluster", "node_type_id").SetComputed()
	s.SchemaPath("cluster", "driver_node_type_id").SetComputed()
	// s.SchemaPath("cluster", "enable_local_disk_encryption").SetComputed()
	s.SchemaPath("url").SetComputed()

	// SuppressDiff fields
	s.SchemaPath("edition").SetSuppressDiff()
	s.SchemaPath("channel").SetSuppressDiff()
	s.SchemaPath("cluster", "spark_conf").SetCustomSuppressDiff(clusters.SparkConfDiffSuppressFunc)
	s.SchemaPath("cluster", "aws_attributes", "zone_id").SetCustomSuppressDiff(clusters.ZoneDiffSuppress)
	s.SchemaPath("cluster", "autoscale", "mode").SetCustomSuppressDiff(common.EqualFoldDiffSuppress)
	s.SchemaPath("edition").SetCustomSuppressDiff(common.EqualFoldDiffSuppress)
	s.SchemaPath("storage").SetCustomSuppressDiff(suppressStorageDiff)

	// Deprecated fields
	s.SchemaPath("cluster", "init_scripts", "dbfs").SetDeprecated(clusters.DbfsDeprecationWarning)

	// Delete fields
	s.SchemaPath("cluster", "gcp_attributes").RemoveField("use_preemptible_executors")
	s.SchemaPath("cluster", "gcp_attributes").RemoveField("boot_disk_size")

	// Default values
	s.SchemaPath("edition").SetDefault("ADVANCED")
	s.SchemaPath("channel").SetDefault("CURRENT")

	// ConflictsWith fields
	s.SchemaPath("storage").SetConflictsWith([]string{"catalog"})
	s.SchemaPath("catalog").SetConflictsWith([]string{"storage"})

	// MinItems fields
	s.SchemaPath("library").SetMinItems(1)
	s.SchemaPath("notification", "email_recipients").SetMinItems(1)
	s.SchemaPath("notification", "alerts").SetMinItems(1)

	// MaxItems fields
	s.SchemaPath("cluster", "ssh_public_keys").SetMaxItems(10)
	s.SchemaPath("cluster", "init_scripts").SetMaxItems(10)

	// ValidateFunc fields
	s.SchemaPath("channel").SetValidateFunc(validation.StringInSlice([]string{"current", "preview"}, true))
	s.SchemaPath("edition").SetValidateFunc(validation.StringInSlice([]string{"pro", "core", "advanced"}, true))

	return s
}

var pipelineSchema = common.StructToSchema(Pipeline{}, nil)

func ResourcePipeline() common.Resource {
	return common.Resource{
		Schema: pipelineSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			wsClient := workspaceClient{w}
			err = wsClient.Create(ctx, d, d.Timeout(schema.TimeoutCreate))
			if err != nil {
				return err
			}
			// var createPipelineRequest pipelines.CreatePipeline

			// adjustForceSendFields(&createPipelineRequest)

			// common.DataToStructPointer(d, pipelineSchema, &createPipelineRequest)
			// createdPipeline, err := wsClient.Pipelines.Create(ctx, createPipelineRequest)
			// if err != nil {
			// 	return err
			// }
			// id := createdPipeline.PipelineId
			// timeout := d.Timeout(schema.TimeoutCreate)
			// err = waitForState(w, ctx, id, timeout, StateRunning)
			// if err != nil {
			// 	log.Printf("[INFO] Pipeline creation failed, attempting to clean up pipeline %s", id)
			// 	err2 := a.Delete(id, timeout)
			// 	if err2 != nil {
			// 		log.Printf("[WARN] Unable to delete pipeline %s; this resource needs to be manually cleaned up", id)
			// 		return "", fmt.Errorf("multiple errors occurred when creating pipeline. Error while waiting for creation: \"%v\"; error while attempting to clean up failed pipeline: \"%v\"", err, err2)
			// 	}
			// 	log.Printf("[INFO] Successfully cleaned up pipeline %s", id)
			// 	return "", err
			// }
			// return id, nil

			// d.SetId(createdPipeline.PipelineId)
			// why is it required
			d.Set("url", c.FormatURL("#joblist/pipelines/", d.Id()))
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			wsClient := workspaceClient{w}
			readPipeline, err := wsClient.Read(ctx, d.Id())
			// readPipeline, err := w.Pipelines.Get(ctx, pipelines.GetPipelineRequest{
			// 	PipelineId: d.Id(),
			// })
			if err != nil {
				return err
			}
			if readPipeline.Spec == nil {
				return fmt.Errorf("pipeline spec is nil for '%v'", readPipeline.PipelineId)
			}
			fmt.Println("readPipeline.Spec", readPipeline.Spec)
			// err = common.StructToData(readPipeline, pipelineSchema, d)
			// if err != nil {
			// 	return err
			// }
			err = common.StructToData(readPipeline.Spec, pipelineSchema, d)
			if err != nil {
				return err
			}
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			wsClient := workspaceClient{w}
			return wsClient.Update(ctx, d, d.Timeout(schema.TimeoutUpdate))
			// var updatePipelineRequest pipelines.EditPipeline
			// common.DataToStructPointer(d, pipelineSchema, &updatePipelineRequest)
			// return w.Pipelines.Update(ctx, updatePipelineRequest)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			wsClient := workspaceClient{w}
			return wsClient.Delete(ctx, d.Id(), d.Timeout(schema.TimeoutDelete))
			// return w.Pipelines.Delete(ctx, pipelines.DeletePipelineRequest{
			// 	PipelineId: d.Id(),
			// })
		},
	}
}
