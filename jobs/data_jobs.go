package jobs

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/terraform-provider-databricks/common"
)

const (
	JobsGroupByName = "name"
	JobsGroupByID   = "id"
)

func DataSourceJobs() common.Resource {
	return common.WorkspaceDataWithUnifiedProvider(func(ctx context.Context, data *struct {
		Ids                map[string]string         `json:"ids,omitempty" tf:"computed"`
		NameFilter         string                    `json:"job_name_contains,omitempty"`
		Key                string                    `json:"key,omitempty" tf:"default:name"`
		ProviderConfigData common.ProviderConfigData `json:"provider_config,omitempty"`
	}, w *databricks.WorkspaceClient) error {
		iter := w.Jobs.List(ctx, jobs.ListJobsRequest{ExpandTasks: false, Limit: 100})
		data.Ids = map[string]string{}
		nameFilter := strings.ToLower(data.NameFilter)
		keyAttribute := strings.ToLower(data.Key)
		for iter.HasNext(ctx) {
			job, err := iter.Next(ctx)
			if err != nil {
				return err
			}
			name := job.Settings.Name
			if nameFilter != "" && !strings.Contains(strings.ToLower(name), nameFilter) {
				continue
			}
			jobId := strconv.FormatInt(job.JobId, 10)

			key := name
			if strings.EqualFold(keyAttribute, JobsGroupByName) {
				key = name
			} else if strings.EqualFold(keyAttribute, JobsGroupByID) {
				key = jobId
			} else {
				return fmt.Errorf("unsupported key %s, must be one of %s or %s", keyAttribute, JobsGroupByName, JobsGroupByID)
			}

			_, duplicateKey := data.Ids[key]
			if duplicateKey {
				return fmt.Errorf("duplicate job %s detected: %s", keyAttribute, key)
			}
			data.Ids[key] = jobId
		}
		return nil
	})
}
