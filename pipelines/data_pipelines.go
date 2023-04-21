package pipelines

import (
	"context"
	"fmt"
	"sort"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourcePipelines() *schema.Resource {
	type pipelinesData struct {
		PipelineNameContains string   `json:"pipeline_name,omitempty"`
		Ids                  []string `json:"ids,omitempty" tf:"computed,slice_set"`
	}
	return common.WorkspaceData(func(ctx context.Context, data *pipelinesData, w *databricks.WorkspaceClient) error {
		pipelineSearch := pipelines.ListPipelines{MaxResults: 100}

		if data.PipelineNameContains != "" {
			pipelineSearch = pipelines.ListPipelines{Filter: fmt.Sprintf("name LIKE '%s'", data.PipelineNameContains), MaxResults: 100}
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
