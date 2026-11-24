// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package compute_test

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/compute"
)

func ExampleClusterPoliciesAPI_Create_clusterPolicies() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.ClusterPolicies.Create(ctx, compute.CreatePolicy{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Definition: `{
            "spark_conf.spark.databricks.delta.preview.enabled": {
                "type": "fixed",
                "value": true
            }
        }
`,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	err = w.ClusterPolicies.DeleteByPolicyId(ctx, created.PolicyId)
	if err != nil {
		panic(err)
	}

}

func ExampleClusterPoliciesAPI_Edit_clusterPolicies() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.ClusterPolicies.Create(ctx, compute.CreatePolicy{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Definition: `{
            "spark_conf.spark.databricks.delta.preview.enabled": {
                "type": "fixed",
                "value": true
            }
        }
`,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	policy, err := w.ClusterPolicies.GetByPolicyId(ctx, created.PolicyId)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", policy)

	err = w.ClusterPolicies.Edit(ctx, compute.EditPolicy{
		PolicyId: policy.PolicyId,
		Name:     policy.Name,
		Definition: `{
            "spark_conf.spark.databricks.delta.preview.enabled": {
                "type": "fixed",
                "value": false
            }
        }
`,
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.ClusterPolicies.DeleteByPolicyId(ctx, created.PolicyId)
	if err != nil {
		panic(err)
	}

}

func ExampleClusterPoliciesAPI_Get_clusterPolicies() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.ClusterPolicies.Create(ctx, compute.CreatePolicy{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Definition: `{
            "spark_conf.spark.databricks.delta.preview.enabled": {
                "type": "fixed",
                "value": true
            }
        }
`,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	policy, err := w.ClusterPolicies.GetByPolicyId(ctx, created.PolicyId)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", policy)

	// cleanup

	err = w.ClusterPolicies.DeleteByPolicyId(ctx, created.PolicyId)
	if err != nil {
		panic(err)
	}

}

func ExampleClusterPoliciesAPI_List_clusterPolicies() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.ClusterPolicies.ListAll(ctx, compute.ListClusterPoliciesRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}
