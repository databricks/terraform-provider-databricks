package clusters

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/mod/semver"
)

// SparkVersion - contains information about specific version
type SparkVersion struct {
	Version     string `json:"key"`
	Description string `json:"name"`
}

// SparkVersionsList - returns a list of all currently supported Spark Versions
// https://docs.databricks.com/dev-tools/api/latest/clusters.html#runtime-versions
type SparkVersionsList struct {
	SparkVersions []SparkVersion `json:"versions"`
}

// SparkVersionRequest - filtering request
type SparkVersionRequest struct {
	LongTermSupport bool   `json:"long_term_support,omitempty"`
	Beta            bool   `json:"beta,omitempty" tf:"conflicts:long_term_support"`
	Latest          bool   `json:"latest,omitempty" tf:"default:true"`
	ML              bool   `json:"ml,omitempty"`
	Genomics        bool   `json:"genomics,omitempty"`
	GPU             bool   `json:"gpu,omitempty"`
	Scala           string `json:"scala,omitempty" tf:"default:2.12"`
	SparkVersion    string `json:"spark_version,omitempty"`
	Photon          bool   `json:"photon,omitempty"`
	Graviton        bool   `json:"graviton,omitempty"`
}

// ListSparkVersions returns smallest (or default) node type id given the criteria
func (a ClustersAPI) ListSparkVersions() (SparkVersionsList, error) {
	var sparkVersions SparkVersionsList
	err := a.client.Get(a.context, "/clusters/spark-versions", nil, &sparkVersions)
	return sparkVersions, err
}

type sparkVersionsType []string

func (s sparkVersionsType) Len() int {
	return len(s)
}
func (s sparkVersionsType) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

var dbrVersionRegex = regexp.MustCompile(`^(\d+\.\d+)\.x-.*`)

func extractDbrVersions(s string) string {
	m := dbrVersionRegex.FindStringSubmatch(s)
	if len(m) > 1 {
		return m[1]
	}
	return s
}

func (s sparkVersionsType) Less(i, j int) bool {
	return semver.Compare("v"+extractDbrVersions(s[i]), "v"+extractDbrVersions(s[j])) > 0
}

// LatestSparkVersion returns latest version matching the request parameters
func (sparkVersions SparkVersionsList) LatestSparkVersion(req SparkVersionRequest) (string, error) {
	var versions []string

	for _, version := range sparkVersions.SparkVersions {
		if strings.Contains(version.Version, "-scala"+req.Scala) {
			matches := ((!strings.Contains(version.Version, "apache-spark-")) &&
				(strings.Contains(version.Version, "-ml-") == req.ML) &&
				(strings.Contains(version.Version, "-hls-") == req.Genomics) &&
				(strings.Contains(version.Version, "-gpu-") == req.GPU) &&
				(strings.Contains(version.Version, "-photon-") == req.Photon) &&
				(strings.Contains(version.Version, "-aarch64-") == req.Graviton) &&
				(strings.Contains(version.Description, "Beta") == req.Beta))
			if matches && req.LongTermSupport {
				matches = (matches && (strings.Contains(version.Description, "LTS") || strings.Contains(version.Version, "-esr-")))
			}
			if matches && len(req.SparkVersion) > 0 {
				matches = (matches && strings.Contains(version.Description, "Apache Spark "+req.SparkVersion))
			}
			if matches {
				versions = append(versions, version.Version)
			}
		}
	}
	if len(versions) < 1 {
		return "", fmt.Errorf("spark versions query returned no results. Please change your search criteria and try again")
	} else if len(versions) > 1 {
		if req.Latest {
			sort.Sort(sparkVersionsType(versions))
		} else {
			return "", fmt.Errorf("spark versions query returned multiple results. Please change your search criteria and try again")
		}
	}

	return versions[0], nil
}

// LatestSparkVersion returns latest version matching the request parameters
func (a ClustersAPI) LatestSparkVersion(svr SparkVersionRequest) (string, error) {
	sparkVersions, err := a.ListSparkVersions()
	if err != nil {
		return "", err
	}
	return sparkVersions.LatestSparkVersion(svr)
}

// LatestSparkVersionOrDefault returns Spark version matching the definition, or default in case of error
func (a ClustersAPI) LatestSparkVersionOrDefault(svr SparkVersionRequest) string {
	version, err := a.LatestSparkVersion(svr)
	if err != nil {
		return "7.3.x-scala2.12"
	}
	return version
}

// DataSourceSparkVersion returns DBR version matching to the specification
func DataSourceSparkVersion() common.Resource {
	s := common.StructToSchema(SparkVersionRequest{}, func(
		s map[string]*schema.Schema) map[string]*schema.Schema {

		s["photon"].Deprecated = "Specify runtime_engine=\"PHOTON\" in the cluster configuration"
		s["graviton"].Deprecated = "Not required anymore - it's automatically enabled on the Graviton-based node types"
		return s
	})

	return common.Resource{
		Schema: s,
		Read: func(ctx context.Context, d *schema.ResourceData, m *common.DatabricksClient) error {
			var this SparkVersionRequest
			common.DataToStructPointer(d, s, &this)
			version, err := NewClustersAPI(ctx, m).LatestSparkVersion(this)
			if err != nil {
				return err
			}
			d.SetId(version)
			return nil
		},
	}
}
