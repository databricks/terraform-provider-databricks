package compute

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
)

func TestListSparkVersions(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/spark-versions",
				Response: SparkVersionsInfo{
					SparkVerions: []SparkVersion{
						{
							Key:  "7.3.x-cpu-ml-scala2.12",
							Name: "7.3 ML (includes Apache Spark 3.0.1, Scala 2.12)",
						},
						{
							Key:  "7.3.x-hls-scala2.12",
							Name: "7.3 Genomics (includes Apache Spark 3.0.1, Scala 2.12)",
						},
						{
							Key:  "7.3.x-gpu-ml-scala2.12",
							Name: "7.3 ML (includes Apache Spark 3.0.1, GPU, Scala 2.12)",
						},
						{
							Key:  "7.0.x-cpu-ml-scala2.12",
							Name: "7.0 ML (includes Apache Spark 3.0.0, Scala 2.12)",
						},
						{
							Key:  "5.5.x-gpu-ml-scala2.12",
							Name: "5.5 ML (includes Apache Spark 2.4.3, GPU, Scala 2.11)",
						},
					},
				},
			},
		},
		Read:        true,
		Resource:    DataSourceClusterSparkVersions(),
		NonWritable: true,
		ID:          ".",
		HCL: `
		latest = true # default
    	long_term_support = false # default
    	ml = true # default - false
    	gpu = true # default - false
		scala = "2.12" # default
		genomics = false
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "7.3.x-gpu-ml-scala2.12", d.Get("version"))
}

func TestAccListSparkVersions(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Tests skipped unless env 'CLOUD_ENV' is set")
	}

	client := common.CommonEnvironmentClient()
	sparkVersionsInfo, err := NewClustersAPI(client).ListSparkVersions()
	assert.NoError(t, err, err)

	for _, version := range sparkVersionsInfo.SparkVerions {
		if strings.Contains(version.Key, "scala2.12") {
			fmt.Printf(version.Key)
		}
	}
}
