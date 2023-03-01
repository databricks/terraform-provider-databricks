package pipelines

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourcePipeline() *schema.Resource {
	type pipelineData struct {
		Name string            `json:"pipeline_name,omitempty"`
		Ids  map[string]string `json:"ids,omitempty" tf:"computed"`
	}
	return common.DataResource(pipelineData{}, func(ctx context.Context, e interface{}, c *common.DatabricksClient) error {
		data := e.(*pipelineData)
		pipelinesAPI := NewPipelinesAPI(ctx, c)
		searchPattern := ""
		errorMessage := fmt.Errorf("no pipelines found")

		if data.Name != "" {
			searchPattern = fmt.Sprintf("name LIKE '%s'", data.Name)
			errorMessage = fmt.Errorf("there is no pipeline with name LIKE '%s'; you need to specify `pipeline_name` as an `exact_name` or with wildcards `%s`", data.Name, "%name%")
		}

		pipelines, err := pipelinesAPI.List(100, searchPattern)
		if err != nil {
			return err
		}

		if len(pipelines) == 0 {
			return errorMessage
		}

		data.Ids = make(map[string]string)
		for _, p := range pipelines {
			data.Ids[p.Name] = p.PipelineID
		}

		return nil

	})
}
