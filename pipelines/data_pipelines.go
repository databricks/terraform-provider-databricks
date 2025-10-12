package pipelines

import (
	"context"
	"fmt"
	"sort"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourcePipelines() common.Resource {
	type pipelinesData struct {
		common.Namespace
		PipelineNameContains string   `json:"pipeline_name,omitempty"`
		Ids                  []string `json:"ids,omitempty" tf:"computed,slice_set"`
	}
	return common.WorkspaceDataWithUnifiedProvider(func(ctx context.Context, data *pipelinesData, w *databricks.WorkspaceClient) error {
		pipelineSearch := pipelines.ListPipelinesRequest{MaxResults: 100}

		if data.PipelineNameContains != "" {
			pipelineSearch = pipelines.ListPipelinesRequest{Filter: fmt.Sprintf("name LIKE '%s'", data.PipelineNameContains), MaxResults: 100}
		}

		pipelines, err := w.Pipelines.ListPipelinesAll(ctx, pipelineSearch)

		if err != nil {
			return err
		}

		for _, p := range pipelines {
			data.Ids = append(data.Ids, p.PipelineId)
		}

		sort.Strings(data.Ids)

		return nil

	})
}
