package compute

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"golang.org/x/mod/semver"
)

// SparkVersionRequest - filtering request
type SparkVersionRequest struct {
	Id              string `json:"id,omitempty"`
	LongTermSupport bool   `json:"long_term_support,omitempty"`
	Beta            bool   `json:"beta,omitempty"`
	Latest          bool   `json:"latest,omitempty"`
	ML              bool   `json:"ml,omitempty"`
	Genomics        bool   `json:"genomics,omitempty"`
	GPU             bool   `json:"gpu,omitempty"`
	Scala           string `json:"scala,omitempty"`
	SparkVersion    string `json:"spark_version,omitempty"`
	Photon          bool   `json:"photon,omitempty"`
	Graviton        bool   `json:"graviton,omitempty"`
}

type sparkVersionsType []string

func (s sparkVersionsType) Len() int {
	return len(s)
}
func (s sparkVersionsType) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

var dbrVersionRegex = regexp.MustCompile(`^(\d+(?:\.\d+)?)\.x-.*`)

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

// Strip -scala2.12 and -scala2.13 from the version string for uniqueness check
func stripScalaVersion(version string) string {
	v := strings.ReplaceAll(version, "-scala2.12", "")
	v = strings.ReplaceAll(v, "-scala2.13", "")
	return v
}

func (sv GetSparkVersionsResponse) Select(req SparkVersionRequest) (string, error) {
	var versions []string
	for _, version := range sv.Versions {
		if strings.Contains(version.Key, "-scala"+req.Scala) {
			matches := ((!strings.Contains(version.Key, "apache-spark-")) &&
				(strings.Contains(version.Key, "-ml-") == req.ML) &&
				(strings.Contains(version.Key, "-hls-") == req.Genomics) &&
				(strings.Contains(version.Key, "-gpu-") == req.GPU) &&
				(strings.Contains(version.Key, "-photon-") == req.Photon) &&
				(strings.Contains(version.Key, "-aarch64-") == req.Graviton) &&
				(strings.Contains(version.Name, "Beta") == req.Beta))
			if matches && req.LongTermSupport {
				matches = (matches && (strings.Contains(version.Name, "LTS") || strings.Contains(version.Key, "-esr-")))
			}
			if matches && len(req.SparkVersion) > 0 {
				matches = (matches && strings.Contains(version.Name, "Apache Spark "+req.SparkVersion))
			}
			if matches {
				versions = append(versions, version.Key)
			}
		}
	}
	if len(versions) < 1 {
		return "", fmt.Errorf("spark versions query returned no results. Please change your search criteria and try again")
	} else if len(versions) > 1 {
		if len(versions) == 2 {
			if stripScalaVersion(versions[0]) == stripScalaVersion(versions[1]) {
				sort.Strings(versions) // Sort versions to ensure we return 2.12 first
				return versions[0], nil
			}
			// If the versions are different, we fallthrough to the default behavior below
		}
		if req.Latest {
			sort.Sort(sparkVersionsType(versions))
		} else {
			return "", fmt.Errorf("spark versions query returned multiple results %#v. Please change your search criteria and try again", versions)
		}
	}
	return versions[0], nil
}

// SelectSparkVersion returns latest DBR version matching the request parameters.
// If there are multiple versions matching the request, it will error (if latest = false)
// or return the latest version.
// Possible parameters are:
// - LongTermSupport: LTS versions only
// - Beta: Beta versions only
// - ML: ML versions only
// - Genomics: Genomics versions only
// - GPU: GPU versions only
// - Scala: Scala version
// - SparkVersion: Apache Spark version
// - Photon: Photon versions only (deprecated)
// - Graviton: Graviton versions only (deprecated)
func (a *ClustersAPI) SelectSparkVersion(ctx context.Context, r SparkVersionRequest) (string, error) {
	sv, err := a.SparkVersions(ctx)
	if err != nil {
		return "", err
	}
	return sv.Select(r)
}
