package pipelines

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourcePipelines() *schema.Resource {
	type pipelineData struct {
		Name string            `json:"pipeline_name,omitempty"`
		Ids  map[string]string `json:"ids,omitempty" tf:"computed"`
	}
	return common.WorkspaceData(func(ctx context.Context, data *pipelineData, w *databricks.WorkspaceClient) error {
		pipelineSearch := pipelines.ListPipelines{MaxResults: 100}

		if data.Name != "" {
			pipelineSearch = pipelines.ListPipelines{Filter: fmt.Sprintf("name LIKE '%s'", data.Name), MaxResults: 100}
		}

		pipelines, err := w.Pipelines.ListPipelinesAll(ctx, pipelineSearch)
		if err != nil {
			return err
		}

		data.Ids = make(map[string]string)
		for _, p := range pipelines {
			data.Ids[p.PipelineId] = p.Name
		}

		return nil

	})
}
