package pipelines

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourcePipeline() common.Resource {
	type pipelineData struct {
		common.Namespace
		PipelineName string `json:"pipeline_name,omitempty"`
		Id           string `json:"id,omitempty" tf:"computed"`

		State         string `json:"state,omitempty" tf:"computed"`
		RunAsUserName string `json:"run_as_user_name,omitempty" tf:"computed"`
		ClusterId     string `json:"cluster_id,omitempty" tf:"computed"`
	}
	return common.WorkspaceData(func(ctx context.Context, data *pipelineData, w *databricks.WorkspaceClient) error {

		if data.Id == "" && data.PipelineName == "" {
			return fmt.Errorf("either 'Id' or 'PipelineName' is required")
		}

		if data.Id == "" {
			pipelineSearch := pipelines.ListPipelinesRequest{Filter: fmt.Sprintf("name LIKE '%s'", data.PipelineName), MaxResults: 100}
			pipelines, err := w.Pipelines.ListPipelinesAll(ctx, pipelineSearch)
			if err != nil {
				return err
			}

			for _, p := range pipelines {
				if p.Name == data.PipelineName {
					data.Id = p.PipelineId
					break
				}
			}

		}

		pipelineInfo, err := w.Pipelines.Get(ctx, pipelines.GetPipelineRequest{PipelineId: data.Id})
		if err != nil {
			return err
		}

		data.ClusterId = pipelineInfo.ClusterId
		data.RunAsUserName = pipelineInfo.RunAsUserName
		data.State = string(pipelineInfo.State)

		return nil

	})
}
