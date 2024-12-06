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

func DataSourceJobs() common.Resource {
	return common.WorkspaceData(func(ctx context.Context, data *struct {
		Ids        map[string]string `json:"ids,omitempty" tf:"computed"`
		NameFilter string            `json:"job_name_contains,omitempty"`
	}, w *databricks.WorkspaceClient) error {
		iter := w.Jobs.List(ctx, jobs.ListJobsRequest{ExpandTasks: false, Limit: 100})
		data.Ids = map[string]string{}
		nameFilter := strings.ToLower(data.NameFilter)
		for iter.HasNext(ctx) {
			job, err := iter.Next(ctx)
			if err != nil {
				return err
			}
			name := job.Settings.Name
			if nameFilter != "" && !strings.Contains(strings.ToLower(name), nameFilter) {
				continue
			}
			_, duplicateName := data.Ids[name]
			if duplicateName {
				return fmt.Errorf("duplicate job name detected: %s", name)
			}
			data.Ids[name] = strconv.FormatInt(job.JobId, 10)
		}
		return nil
	})
}
