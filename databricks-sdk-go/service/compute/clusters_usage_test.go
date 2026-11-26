// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package compute_test

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/iam"
)

func ExampleClustersAPI_ChangeOwner_clustersApiIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	latest, err := w.Clusters.SelectSparkVersion(ctx, compute.SparkVersionRequest{
		Latest:          true,
		LongTermSupport: true,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", latest)

	clusterName := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	otherOwner, err := w.Users.Create(ctx, iam.User{
		UserName: fmt.Sprintf("sdk-%x@example.com", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", otherOwner)

	clstr, err := w.Clusters.CreateAndWait(ctx, compute.CreateCluster{
		ClusterName:            clusterName,
		SparkVersion:           latest,
		InstancePoolId:         os.Getenv("TEST_INSTANCE_POOL_ID"),
		AutoterminationMinutes: 15,
		NumWorkers:             1,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", clstr)

	err = w.Clusters.ChangeOwner(ctx, compute.ChangeClusterOwner{
		ClusterId:     clstr.ClusterId,
		OwnerUsername: otherOwner.UserName,
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Users.DeleteById(ctx, otherOwner.Id)
	if err != nil {
		panic(err)
	}
	err = w.Clusters.PermanentDeleteByClusterId(ctx, clstr.ClusterId)
	if err != nil {
		panic(err)
	}

}

func ExampleClustersAPI_Create_clustersApiIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	latest, err := w.Clusters.SelectSparkVersion(ctx, compute.SparkVersionRequest{
		Latest:          true,
		LongTermSupport: true,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", latest)

	clusterName := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	clstr, err := w.Clusters.CreateAndWait(ctx, compute.CreateCluster{
		ClusterName:            clusterName,
		SparkVersion:           latest,
		InstancePoolId:         os.Getenv("TEST_INSTANCE_POOL_ID"),
		AutoterminationMinutes: 15,
		NumWorkers:             1,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", clstr)

	// cleanup

	err = w.Clusters.PermanentDeleteByClusterId(ctx, clstr.ClusterId)
	if err != nil {
		panic(err)
	}

}

func ExampleClustersAPI_Delete_clustersApiIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	latest, err := w.Clusters.SelectSparkVersion(ctx, compute.SparkVersionRequest{
		Latest:          true,
		LongTermSupport: true,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", latest)

	clusterName := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	clstr, err := w.Clusters.CreateAndWait(ctx, compute.CreateCluster{
		ClusterName:            clusterName,
		SparkVersion:           latest,
		InstancePoolId:         os.Getenv("TEST_INSTANCE_POOL_ID"),
		AutoterminationMinutes: 15,
		NumWorkers:             1,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", clstr)

	_, err = w.Clusters.DeleteByClusterIdAndWait(ctx, clstr.ClusterId)
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Clusters.PermanentDeleteByClusterId(ctx, clstr.ClusterId)
	if err != nil {
		panic(err)
	}

}

func ExampleClustersAPI_Edit_clustersApiIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	clusterName := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	latest, err := w.Clusters.SelectSparkVersion(ctx, compute.SparkVersionRequest{
		Latest:          true,
		LongTermSupport: true,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", latest)

	clstr, err := w.Clusters.CreateAndWait(ctx, compute.CreateCluster{
		ClusterName:            clusterName,
		SparkVersion:           latest,
		InstancePoolId:         os.Getenv("TEST_INSTANCE_POOL_ID"),
		AutoterminationMinutes: 15,
		NumWorkers:             1,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", clstr)

	_, err = w.Clusters.EditAndWait(ctx, compute.EditCluster{
		ClusterId:              clstr.ClusterId,
		SparkVersion:           latest,
		ClusterName:            clusterName,
		InstancePoolId:         os.Getenv("TEST_INSTANCE_POOL_ID"),
		AutoterminationMinutes: 10,
		NumWorkers:             2,
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Clusters.PermanentDeleteByClusterId(ctx, clstr.ClusterId)
	if err != nil {
		panic(err)
	}

}

func ExampleClustersAPI_EnsureClusterIsRunning_commandsDirectUsage() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	clusterId := os.Getenv("TEST_DEFAULT_CLUSTER_ID")

	context, err := w.CommandExecution.CreateAndWait(ctx, compute.CreateContext{
		ClusterId: clusterId,
		Language:  compute.LanguagePython,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", context)

	err = w.Clusters.EnsureClusterIsRunning(ctx, clusterId)
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.CommandExecution.Destroy(ctx, compute.DestroyContext{
		ClusterId: clusterId,
		ContextId: context.Id,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleClustersAPI_Events_clustersApiIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	latest, err := w.Clusters.SelectSparkVersion(ctx, compute.SparkVersionRequest{
		Latest:          true,
		LongTermSupport: true,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", latest)

	clusterName := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	clstr, err := w.Clusters.CreateAndWait(ctx, compute.CreateCluster{
		ClusterName:            clusterName,
		SparkVersion:           latest,
		InstancePoolId:         os.Getenv("TEST_INSTANCE_POOL_ID"),
		AutoterminationMinutes: 15,
		NumWorkers:             1,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", clstr)

	events, err := w.Clusters.EventsAll(ctx, compute.GetEvents{
		ClusterId: clstr.ClusterId,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", events)

	// cleanup

	err = w.Clusters.PermanentDeleteByClusterId(ctx, clstr.ClusterId)
	if err != nil {
		panic(err)
	}

}

func ExampleClustersAPI_Get_clustersApiIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	latest, err := w.Clusters.SelectSparkVersion(ctx, compute.SparkVersionRequest{
		Latest:          true,
		LongTermSupport: true,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", latest)

	clusterName := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	clstr, err := w.Clusters.CreateAndWait(ctx, compute.CreateCluster{
		ClusterName:            clusterName,
		SparkVersion:           latest,
		InstancePoolId:         os.Getenv("TEST_INSTANCE_POOL_ID"),
		AutoterminationMinutes: 15,
		NumWorkers:             1,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", clstr)

	byId, err := w.Clusters.GetByClusterId(ctx, clstr.ClusterId)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

	// cleanup

	err = w.Clusters.PermanentDeleteByClusterId(ctx, clstr.ClusterId)
	if err != nil {
		panic(err)
	}

}

func ExampleClustersAPI_List_clustersApiIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.Clusters.ListAll(ctx, compute.ListClustersRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

func ExampleClustersAPI_ListNodeTypes_clustersApiIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	nodes, err := w.Clusters.ListNodeTypes(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", nodes)

}

func ExampleClustersAPI_Pin_clustersApiIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	latest, err := w.Clusters.SelectSparkVersion(ctx, compute.SparkVersionRequest{
		Latest:          true,
		LongTermSupport: true,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", latest)

	clusterName := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	clstr, err := w.Clusters.CreateAndWait(ctx, compute.CreateCluster{
		ClusterName:            clusterName,
		SparkVersion:           latest,
		InstancePoolId:         os.Getenv("TEST_INSTANCE_POOL_ID"),
		AutoterminationMinutes: 15,
		NumWorkers:             1,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", clstr)

	err = w.Clusters.PinByClusterId(ctx, clstr.ClusterId)
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Clusters.PermanentDeleteByClusterId(ctx, clstr.ClusterId)
	if err != nil {
		panic(err)
	}

}

func ExampleClustersAPI_Resize_clustersApiIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	latest, err := w.Clusters.SelectSparkVersion(ctx, compute.SparkVersionRequest{
		Latest:          true,
		LongTermSupport: true,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", latest)

	clusterName := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	clstr, err := w.Clusters.CreateAndWait(ctx, compute.CreateCluster{
		ClusterName:            clusterName,
		SparkVersion:           latest,
		InstancePoolId:         os.Getenv("TEST_INSTANCE_POOL_ID"),
		AutoterminationMinutes: 15,
		NumWorkers:             1,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", clstr)

	byId, err := w.Clusters.ResizeAndWait(ctx, compute.ResizeCluster{
		ClusterId:  clstr.ClusterId,
		NumWorkers: 1,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

	// cleanup

	err = w.Clusters.PermanentDeleteByClusterId(ctx, clstr.ClusterId)
	if err != nil {
		panic(err)
	}

}

func ExampleClustersAPI_Restart_clustersApiIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	latest, err := w.Clusters.SelectSparkVersion(ctx, compute.SparkVersionRequest{
		Latest:          true,
		LongTermSupport: true,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", latest)

	clusterName := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	clstr, err := w.Clusters.CreateAndWait(ctx, compute.CreateCluster{
		ClusterName:            clusterName,
		SparkVersion:           latest,
		InstancePoolId:         os.Getenv("TEST_INSTANCE_POOL_ID"),
		AutoterminationMinutes: 15,
		NumWorkers:             1,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", clstr)

	_, err = w.Clusters.RestartAndWait(ctx, compute.RestartCluster{
		ClusterId: clstr.ClusterId,
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Clusters.PermanentDeleteByClusterId(ctx, clstr.ClusterId)
	if err != nil {
		panic(err)
	}

}

func ExampleClustersAPI_SelectNodeType_instancePools() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	smallest, err := w.Clusters.SelectNodeType(ctx, compute.NodeTypeRequest{
		LocalDisk: true,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", smallest)

}

func ExampleClustersAPI_SelectSparkVersion_clustersApiIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	latest, err := w.Clusters.SelectSparkVersion(ctx, compute.SparkVersionRequest{
		Latest:          true,
		LongTermSupport: true,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", latest)

}

func ExampleClustersAPI_Start_clustersApiIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	latest, err := w.Clusters.SelectSparkVersion(ctx, compute.SparkVersionRequest{
		Latest:          true,
		LongTermSupport: true,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", latest)

	clusterName := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	clstr, err := w.Clusters.CreateAndWait(ctx, compute.CreateCluster{
		ClusterName:            clusterName,
		SparkVersion:           latest,
		InstancePoolId:         os.Getenv("TEST_INSTANCE_POOL_ID"),
		AutoterminationMinutes: 15,
		NumWorkers:             1,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", clstr)

	_, err = w.Clusters.StartByClusterIdAndWait(ctx, clstr.ClusterId)
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Clusters.PermanentDeleteByClusterId(ctx, clstr.ClusterId)
	if err != nil {
		panic(err)
	}

}

func ExampleClustersAPI_Unpin_clustersApiIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	latest, err := w.Clusters.SelectSparkVersion(ctx, compute.SparkVersionRequest{
		Latest:          true,
		LongTermSupport: true,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", latest)

	clusterName := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	clstr, err := w.Clusters.CreateAndWait(ctx, compute.CreateCluster{
		ClusterName:            clusterName,
		SparkVersion:           latest,
		InstancePoolId:         os.Getenv("TEST_INSTANCE_POOL_ID"),
		AutoterminationMinutes: 15,
		NumWorkers:             1,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", clstr)

	err = w.Clusters.UnpinByClusterId(ctx, clstr.ClusterId)
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Clusters.PermanentDeleteByClusterId(ctx, clstr.ClusterId)
	if err != nil {
		panic(err)
	}

}
