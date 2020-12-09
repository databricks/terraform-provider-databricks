package compute

import (
	"strings"
	"testing"

	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func commonFixtures() []qa.HTTPFixture {
	return []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/spark-versions",
			Response: SparkVersionsList{
				SparkVersions: []SparkVersion{
					{
						Version:     "7.1.x-scala2.12",
						Description: "7.1 (includes Apache Spark 3.0.0, Scala 2.12)",
					},
					{
						Version:     "7.1.x-gpu-ml-scala2.12",
						Description: "7.1 ML (includes Apache Spark 3.0.0, Scala 2.12)",
					},
					{
						Version:     "apache-spark-2.4.x-scala2.11",
						Description: "Light 2.4 (includes Apache Spark 2.4, Scala 2.11)",
					},
					{
						Version:     "7.3.x-hls-scala2.12",
						Description: "7.3 LTS Genomics (includes Apache Spark 3.0.1, Scala 2.12)",
					},
					{
						Version:     "6.4.x-scala2.11",
						Description: "6.4 (includes Apache Spark 2.4.5, Scala 2.11)",
					},
					{
						Version:     "7.3.x-scala2.12",
						Description: "7.3 LTS (includes Apache Spark 3.0.1, Scala 2.12)",
					},
					{
						Version:     "7.4.x-scala2.12",
						Description: "7.4 (includes Apache Spark 3.0.1, Scala 2.12)",
					},
					{
						Version:     "7.5.x-scala2.12",
						Description: "7.5 Beta (includes Apache Spark 3.0.1, Scala 2.12)",
					},
				},
			},
		},
	}
}

func TestSparkVersionLatest(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures:    commonFixtures(),
		Read:        true,
		Resource:    DataSourceSparkVersion(),
		NonWritable: true,
		State:       map[string]interface{}{},
		ID:          ".",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "7.4.x-scala2.12", d.Id())
}

func TestSparkVersionLTS(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures:    commonFixtures(),
		Read:        true,
		Resource:    DataSourceSparkVersion(),
		NonWritable: true,
		State: map[string]interface{}{
			"long_term_support": true,
		},
		ID: ".",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "7.3.x-scala2.12", d.Id())
}

func TestSparkVersionGpuMl(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures:    commonFixtures(),
		Read:        true,
		Resource:    DataSourceSparkVersion(),
		NonWritable: true,
		State: map[string]interface{}{
			"gpu": true,
			"ml":  true,
		},
		ID: ".",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "7.1.x-gpu-ml-scala2.12", d.Id())
}

func TestSparkVersionGenomics(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures:    commonFixtures(),
		Read:        true,
		Resource:    DataSourceSparkVersion(),
		NonWritable: true,
		State: map[string]interface{}{
			"genomics": true,
		},
		ID: ".",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "7.3.x-hls-scala2.12", d.Id())
}

func TestSparkVersion300(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures:    commonFixtures(),
		Read:        true,
		Resource:    DataSourceSparkVersion(),
		NonWritable: true,
		State: map[string]interface{}{
			"spark_version": "3.0.0",
		},
		ID: ".",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "7.1.x-scala2.12", d.Id())
}

func TestSparkVersionBeta(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures:    commonFixtures(),
		Read:        true,
		Resource:    DataSourceSparkVersion(),
		NonWritable: true,
		State: map[string]interface{}{
			"beta": true,
		},
		ID: ".",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "7.5.x-scala2.12", d.Id())
}

func TestSparkVersionErrorNoResults(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures:    commonFixtures(),
		Read:        true,
		Resource:    DataSourceSparkVersion(),
		NonWritable: true,
		State: map[string]interface{}{
			"beta":              true,
			"long_term_support": true,
		},
		ID: ".",
	}.Apply(t)
	assert.Error(t, err)
	assert.Equal(t, true, strings.Contains(err.Error(), "query returned no results"))
}

func TestSparkVersionErrorMultipleResults(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures:    commonFixtures(),
		Read:        true,
		Resource:    DataSourceSparkVersion(),
		NonWritable: true,
		State: map[string]interface{}{
			"latest": false,
		},
		ID: ".",
	}.Apply(t)
	assert.Error(t, err)
	assert.Equal(t, true, strings.Contains(err.Error(), "query returned multiple results"))
}

func TestSparkVersionErrorBadAnswer(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/spark-versions",
				Response: "{garbage....",
			},
		},
		Read:        true,
		Resource:    DataSourceSparkVersion(),
		NonWritable: true,
		State: map[string]interface{}{
			"latest": false,
		},
		ID: ".",
	}.Apply(t)
	assert.Error(t, err)
	require.Equal(t, true, strings.Contains(err.Error(), "Invalid JSON received"))
}
