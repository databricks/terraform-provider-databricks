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
		errorMessage := fmt.Errorf("no pipelines found")
		pipelineSearch := pipelines.ListPipelines{MaxResults: 100}

		if data.Name != "" {
			searchPattern := fmt.Sprintf("name LIKE '%s'", data.Name)
			errorMessage = fmt.Errorf("there is no pipeline with name LIKE '%s'; you need to specify `pipeline_name` as the full pipeline name or with percent wildcards", data.Name)
			pipelineSearch = pipelines.ListPipelines{Filter: searchPattern, MaxResults: 100}
		}

		pipelines, err := w.Pipelines.ListPipelinesAll(ctx, pipelineSearch)
		if err != nil {
			return err
		} else if len(pipelines) == 0 {
			return errorMessage
		}

		data.Ids = make(map[string]string)
		for _, p := range pipelines {
			data.Ids[p.Name] = p.PipelineId
		}

		return nil

	})
}
