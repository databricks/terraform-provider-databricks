package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const repoRoot = "../.."

// TestGeneratedFilesAreUpToDate fails when a spec in the repository generates a file that
// differs from the checked-in one.
func TestGeneratedFilesAreUpToDate(t *testing.T) {
	specs := findSpecs(t)
	require.NotEmpty(t, specs)
	for _, specPath := range specs {
		name, err := filepath.Rel(repoRoot, specPath)
		require.NoError(t, err)
		t.Run(name, func(t *testing.T) {
			spec, err := readSpec(specPath)
			require.NoError(t, err)
			outDir := t.TempDir()
			require.NoError(t, run(specPath, outDir))
			want, err := os.ReadFile(filepath.Join(filepath.Dir(specPath), spec.Output))
			require.NoError(t, err)
			got, err := os.ReadFile(filepath.Join(outDir, spec.Output))
			require.NoError(t, err)
			assert.Equal(t, string(got), string(want),
				"%s is out of date: run go generate in its directory", spec.Output)
		})
	}
}

func findSpecs(t *testing.T) []string {
	var specs []string
	err := filepath.WalkDir(repoRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && path != repoRoot &&
			(d.Name() == "vendor" || strings.HasPrefix(d.Name(), ".")) {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, "_sdkgo_mapping.yaml") {
			specs = append(specs, path)
		}
		return nil
	})
	require.NoError(t, err)
	return specs
}

func TestGenerateRejectsAttributesItCannotMap(t *testing.T) {
	cases := map[string]struct {
		schemaStruct string
		sdkGoType    string
		attribute    string
		err          string
	}{
		"attribute missing from the schema struct": {
			attribute: "missing",
			err:       "f: attribute missing is not in testVolume",
		},
		"field missing from the sdk-go type": {
			attribute: "not_in_sdk_go",
			err:       "f: attribute not_in_sdk_go: VolumeInfo has no field NotInSdkGo",
		},
		"ignored field": {
			attribute: "Skipped",
			err:       "f: attribute Skipped is not in testVolume",
		},
		"unsupported schema struct type": {
			attribute: "browse_only",
			err:       "f: attribute browse_only: can't map testVolume.BrowseOnly",
		},
		"unsupported sdk-go type": {
			attribute: "created_at",
			err:       "f: attribute created_at: can't map testVolume.CreatedAt",
		},
		"missing sdk-go type": {
			sdkGoType: "NoSuchType",
			attribute: "name",
			err:       "uc/volumes/v1 has no type NoSuchType",
		},
		"missing schema struct": {
			schemaStruct: "noSuchStruct",
			attribute:    "name",
			err:          "testdata/schema.go has no struct noSuchStruct",
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			spec := mappingSpec{
				Output: "unused.go",
				Schema: schemaRef{File: "schema.go", Struct: "testVolume"},
				SdkGo: sdkGoRef{
					Package: "github.com/databricks/sdk-go/uc/volumes/v1",
					Name:    "sdkvolumes",
				},
				ToData: []mappingFunc{{
					Function:   "f",
					Type:       "VolumeInfo",
					Attributes: []string{c.attribute},
				}},
			}
			if c.schemaStruct != "" {
				spec.Schema.Struct = c.schemaStruct
			}
			if c.sdkGoType != "" {
				spec.ToData[0].Type = c.sdkGoType
			}
			_, err := generate(spec, "testdata")
			require.Error(t, err)
			assert.Contains(t, err.Error(), c.err)
		})
	}
}

func TestDecodeSpecRejectsInvalidSpecs(t *testing.T) {
	const valid = `
output: out.go
schema: {file: schema.go, struct: testVolume}
sdk_go: {package: github.com/databricks/sdk-go/uc/volumes/v1, name: sdkvolumes}
`
	cases := map[string]struct {
		spec string
		err  string
	}{
		"unknown field": {
			spec: valid + "to_data: [{function: f, type: VolumeInfo, attribute: [name]}]",
			err:  "field attribute not found",
		},
		"missing field": {
			spec: "schema: {file: schema.go, struct: testVolume}\n" +
				"sdk_go: {package: p, name: n}\n" +
				"to_data: [{function: f, type: T, attributes: [a]}]",
			err: "output is required",
		},
		"no functions": {
			spec: valid,
			err:  "from_data or to_data must name a function",
		},
		"function without attributes": {
			spec: valid + "to_data: [{function: f, type: VolumeInfo}]",
			err:  `function "f" needs function, type and attributes`,
		},
		"function named twice": {
			spec: valid + "from_data: [{function: f, type: T, attributes: [a]}]\n" +
				"to_data: [{function: f, type: T, attributes: [a]}]",
			err: "function f is named twice",
		},
		"attribute named twice": {
			spec: valid + "to_data: [{function: f, type: T, attributes: [a, a]}]",
			err:  "f: attribute a is named twice",
		},
		"two documents": {
			spec: valid + "to_data: [{function: f, type: T, attributes: [a]}]\n---\n{}",
			err:  "expected a single YAML document",
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := decodeSpec([]byte(c.spec))
			require.Error(t, err)
			assert.Contains(t, err.Error(), c.err)
		})
	}
}
