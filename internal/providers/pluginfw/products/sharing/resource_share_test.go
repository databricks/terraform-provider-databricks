package sharing

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/stretchr/testify/assert"
)

func TestShareSyncEffectiveFields(t *testing.T) {
	shareName := "test-share-name"
	ctx := context.Background()
	shares := ShareResource{}

	tests := []struct {
		name       string
		planGoSDK  sharing.ShareInfo
		stateGoSDK sharing.ShareInfo
	}{
		{
			name: "plan with less objects",
			planGoSDK: sharing.ShareInfo{
				Name: shareName,
				Objects: []sharing.SharedDataObject{
					{
						Name: "obj-1",
						Partitions: []sharing.Partition{
							{Values: []sharing.PartitionValue{{Value: "part-1"}}},
						},
					},
					{
						Name: "obj-3",
						Partitions: []sharing.Partition{
							{Values: []sharing.PartitionValue{{Value: "part-3"}}},
						},
					},
				},
			},
			stateGoSDK: sharing.ShareInfo{
				Name: shareName,
				Objects: []sharing.SharedDataObject{
					{
						Name: "obj-1",
						Partitions: []sharing.Partition{
							{Values: []sharing.PartitionValue{{Value: "part-1"}}},
						},
					},
					{
						Name: "obj-2",
						Partitions: []sharing.Partition{
							{Values: []sharing.PartitionValue{{Value: "part-2"}}},
						},
					},
					{
						Name: "obj-3",
						Partitions: []sharing.Partition{
							{Values: []sharing.PartitionValue{{Value: "part-3"}}},
						},
					},
				},
			},
		},
		{
			name: "plan with more objects",
			planGoSDK: sharing.ShareInfo{
				Name: shareName,
				Objects: []sharing.SharedDataObject{
					{
						Name: "obj-1",
						Partitions: []sharing.Partition{
							{Values: []sharing.PartitionValue{{Value: "part-1"}}},
						},
					},
					{
						Name: "obj-2",
						Partitions: []sharing.Partition{
							{Values: []sharing.PartitionValue{{Value: "part-2"}}},
						},
					},
					{
						Name: "obj-3",
						Partitions: []sharing.Partition{
							{Values: []sharing.PartitionValue{{Value: "part-3"}}},
						},
					},
				},
			},
			stateGoSDK: sharing.ShareInfo{
				Name: shareName,
				Objects: []sharing.SharedDataObject{
					{
						Name: "obj-1",
						Partitions: []sharing.Partition{
							{Values: []sharing.PartitionValue{{Value: "part-1"}}},
						},
					},
					{
						Name: "obj-3",
						Partitions: []sharing.Partition{
							{Values: []sharing.PartitionValue{{Value: "part-3"}}},
						},
					},
				},
			},
		},
		{
			name: "empty plan",
			planGoSDK: sharing.ShareInfo{
				Name:    shareName,
				Objects: []sharing.SharedDataObject{},
			},
			stateGoSDK: sharing.ShareInfo{
				Name: shareName,
				Objects: []sharing.SharedDataObject{
					{
						Name: "obj-1",
						Partitions: []sharing.Partition{
							{Values: []sharing.PartitionValue{{Value: "part-1"}}},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var planTFSDK ShareInfoExtended
			diagnostics := converters.GoSdkToTfSdkStruct(ctx, tt.planGoSDK, &planTFSDK)
			assert.False(t, diagnostics.HasError())

			var stateTFSDK ShareInfoExtended
			diagnostics = converters.GoSdkToTfSdkStruct(ctx, tt.stateGoSDK, &stateTFSDK)
			assert.False(t, diagnostics.HasError())

			_, diagnostics = shares.syncEffectiveFields(ctx, planTFSDK, stateTFSDK, effectiveFieldsActionRead{})
			assert.False(t, diagnostics.HasError())
		})
	}
}
