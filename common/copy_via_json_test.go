package common

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/stretchr/testify/assert"
)

func TestCopyViaJSON_Flat(t *testing.T) {
	src := compute.ClusterSpec{
		PolicyId:                 "policy-id",
		ApplyPolicyDefaultValues: true,
	}

	dst := CopyViaJSON(src, []string{"policy_id"})
	assert.Equal(t, "policy-id", dst.PolicyId)
	assert.Equal(t, false, dst.ApplyPolicyDefaultValues)
}

func TestCopyViaJSON_Nested_Nil(t *testing.T) {
	src := compute.ClusterSpec{
		PolicyId:  "policy-id",
		Autoscale: nil,
	}

	dst := CopyViaJSON(src, []string{"policy_id", "autoscale.min_workers"})
	assert.Equal(t, "policy-id", dst.PolicyId)
	assert.Nil(t, dst.Autoscale)
}

func TestCopyViaJSON_Nested_Empty(t *testing.T) {
	src := compute.ClusterSpec{
		PolicyId:  "policy-id",
		Autoscale: &compute.AutoScale{},
	}

	dst := CopyViaJSON(src, []string{"policy_id", "autoscale.min_workers"})
	assert.Equal(t, "policy-id", dst.PolicyId)
	assert.NotNil(t, dst.Autoscale)
}

func TestCopyViaJSON_Nested_IncludeField(t *testing.T) {
	src := compute.ClusterSpec{
		PolicyId: "policy-id",
		Autoscale: &compute.AutoScale{
			MinWorkers: 1,
			MaxWorkers: 10,
		},
	}

	dst := CopyViaJSON(src, []string{"policy_id", "autoscale.min_workers"})
	assert.Equal(t, "policy-id", dst.PolicyId)
	assert.Equal(t, 1, dst.Autoscale.MinWorkers)
	assert.Equal(t, 0, dst.Autoscale.MaxWorkers)

	// Only the included fields will be sent to the API.
	assert.Equal(t, []string{"MinWorkers"}, dst.Autoscale.ForceSendFields)
}

func TestCopyViaJSON_Nested_IncludeContainer(t *testing.T) {
	src := compute.ClusterSpec{
		PolicyId: "policy-id",
		Autoscale: &compute.AutoScale{
			MinWorkers: 1,
			MaxWorkers: 10,
		},
	}

	dst := CopyViaJSON(src, []string{"policy_id", "autoscale"})
	assert.Equal(t, "policy-id", dst.PolicyId)
	assert.NotNil(t, dst.Autoscale)
	assert.Equal(t, 1, dst.Autoscale.MinWorkers)
	assert.Equal(t, 10, dst.Autoscale.MaxWorkers)

	// Only the included fields will be sent to the API.
	assert.ElementsMatch(t, []string{"MinWorkers", "MaxWorkers"}, dst.Autoscale.ForceSendFields)
}

func TestCopyViaJSON_Nested_IncludeMultipleFields(t *testing.T) {
	src := compute.ClusterSpec{
		PolicyId: "policy-id",
		Autoscale: &compute.AutoScale{
			MinWorkers: 1,
			MaxWorkers: 10,
		},
	}

	dst := CopyViaJSON(src, []string{"policy_id", "autoscale.min_workers", "autoscale.max_workers"})
	assert.Equal(t, "policy-id", dst.PolicyId)
	assert.NotNil(t, dst.Autoscale)
	assert.Equal(t, 1, dst.Autoscale.MinWorkers)
	assert.Equal(t, 10, dst.Autoscale.MaxWorkers)

	// Only the included fields will be sent to the API.
	assert.ElementsMatch(t, []string{"MinWorkers", "MaxWorkers"}, dst.Autoscale.ForceSendFields)
}

func TestCopyViaJSON_Map_Nil(t *testing.T) {
	src := compute.ClusterSpec{
		PolicyId:  "policy-id",
		SparkConf: nil,
	}

	dst := CopyViaJSON(src, []string{"policy_id", "spark_conf"})
	assert.Equal(t, "policy-id", dst.PolicyId)
	assert.Nil(t, dst.SparkConf)
}

func TestCopyViaJSON_Map_NonNil(t *testing.T) {
	src := compute.ClusterSpec{
		PolicyId:  "policy-id",
		SparkConf: map[string]string{"key": "value"},
	}

	dst := CopyViaJSON(src, []string{"policy_id", "spark_conf"})
	assert.Equal(t, "policy-id", dst.PolicyId)
	assert.NotNil(t, dst.SparkConf)
	assert.Equal(t, map[string]string{"key": "value"}, dst.SparkConf)
}

func TestCopyViaJSON_Slice_Nil(t *testing.T) {
	src := compute.ClusterSpec{
		PolicyId:    "policy-id",
		InitScripts: nil,
	}

	dst := CopyViaJSON(src, []string{"policy_id", "init_scripts"})
	assert.Equal(t, "policy-id", dst.PolicyId)
	assert.Nil(t, dst.InitScripts)
}

func TestCopyViaJSON_Slice_NonNil(t *testing.T) {
	src := compute.ClusterSpec{
		PolicyId: "policy-id",
		InitScripts: []compute.InitScriptInfo{
			{
				Dbfs: &compute.DbfsStorageInfo{
					Destination: "dbfs:/scripts/init.sh",
				},
			},
		},
	}

	dst := CopyViaJSON(src, []string{"policy_id", "init_scripts"})
	assert.Equal(t, "policy-id", dst.PolicyId)
	assert.NotNil(t, dst.InitScripts)
	assert.Len(t, dst.InitScripts, 1)
	assert.Equal(t, "dbfs:/scripts/init.sh", dst.InitScripts[0].Dbfs.Destination)
}
