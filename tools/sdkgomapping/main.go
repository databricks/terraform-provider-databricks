// Command sdkgomapping generates typed code that maps between the Terraform state of an
// SDKv2 resource and the structs of a modular sdk-go package, which common.StructToData
// and common.DataToStructPointer can't reflect over: sdk-go structs have no json tags
// and use pointers for scalars.
//
// It reads a <resource>_sdkgo_mapping.yaml spec and writes the file the spec names.
// Run it from a //go:generate line next to the spec:
//
//	//go:generate go run ../tools/sdkgomapping <resource>_sdkgo_mapping.yaml
//
// It reads the Go types from source and never imports the resource's package, so a
// broken generated file doesn't stop it from running.
package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// mappingSpec is a <resource>_sdkgo_mapping.yaml file. Its paths are relative to its
// directory.
type mappingSpec struct {
	// Output is the generated file, in the schema struct's package.
	Output string    `yaml:"output"`
	Schema schemaRef `yaml:"schema"`
	SdkGo  sdkGoRef  `yaml:"sdk_go"`
	// FromData are the functions that build an sdk-go request from the state.
	FromData []mappingFunc `yaml:"from_data"`
	// ToData are the functions that set the state from an sdk-go response.
	ToData []mappingFunc `yaml:"to_data"`
}

// schemaRef names the struct that the resource passes to common.StructToSchema. Its
// tags give the attribute names and whether each one is optional or computed.
type schemaRef struct {
	File   string `yaml:"file"`
	Struct string `yaml:"struct"`
}

// sdkGoRef names the sdk-go package and the name the generated file imports it as.
type sdkGoRef struct {
	Package string `yaml:"package"`
	Name    string `yaml:"name"`
}

// mappingFunc is one generated function, which maps attributes to the fields of the
// sdk-go struct Type with the same Go names as in the schema struct.
type mappingFunc struct {
	Function   string   `yaml:"function"`
	Type       string   `yaml:"type"`
	Attributes []string `yaml:"attributes"`
}

func main() {
	outDir := flag.String("out-dir", "", "directory to write to instead of the spec's")
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(),
			"usage: sdkgomapping [-out-dir dir] <resource>_sdkgo_mapping.yaml")
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(2)
	}
	if err := run(flag.Arg(0), *outDir); err != nil {
		fmt.Fprintf(os.Stderr, "sdkgomapping: %v\n", err)
		os.Exit(1)
	}
}

// run generates the file of the spec at specPath into outDir, or next to the spec when
// outDir is empty.
func run(specPath, outDir string) error {
	spec, err := readSpec(specPath)
	if err != nil {
		return err
	}
	specDir := filepath.Dir(specPath)
	src, err := generate(spec, specDir)
	if err != nil {
		return fmt.Errorf("%s: %w", specPath, err)
	}
	if outDir == "" {
		outDir = specDir
	}
	return os.WriteFile(filepath.Join(outDir, spec.Output), src, 0o644)
}

func readSpec(path string) (mappingSpec, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return mappingSpec{}, err
	}
	spec, err := decodeSpec(data)
	if err != nil {
		return mappingSpec{}, fmt.Errorf("%s: %w", path, err)
	}
	return spec, nil
}

// decodeSpec rejects unknown fields, because a misspelled one would silently drop a
// mapping.
func decodeSpec(data []byte) (mappingSpec, error) {
	decoder := yaml.NewDecoder(bytes.NewReader(data))
	decoder.KnownFields(true)
	var spec mappingSpec
	if err := decoder.Decode(&spec); err != nil {
		return mappingSpec{}, err
	}
	var extra any
	if err := decoder.Decode(&extra); !errors.Is(err, io.EOF) {
		return mappingSpec{}, errors.New("expected a single YAML document")
	}
	return spec, spec.validate()
}

func (s mappingSpec) validate() error {
	required := map[string]string{
		"output":         s.Output,
		"schema.file":    s.Schema.File,
		"schema.struct":  s.Schema.Struct,
		"sdk_go.package": s.SdkGo.Package,
		"sdk_go.name":    s.SdkGo.Name,
	}
	for field, value := range required {
		if value == "" {
			return fmt.Errorf("%s is required", field)
		}
	}
	if len(s.FromData)+len(s.ToData) == 0 {
		return errors.New("from_data or to_data must name a function")
	}
	functions := map[string]bool{}
	for _, f := range append(append([]mappingFunc{}, s.FromData...), s.ToData...) {
		if f.Function == "" || f.Type == "" || len(f.Attributes) == 0 {
			return fmt.Errorf("function %q needs function, type and attributes", f.Function)
		}
		if functions[f.Function] {
			return fmt.Errorf("function %s is named twice", f.Function)
		}
		functions[f.Function] = true
		attributes := map[string]bool{}
		for _, a := range f.Attributes {
			if attributes[a] {
				return fmt.Errorf("%s: attribute %s is named twice", f.Function, a)
			}
			attributes[a] = true
		}
	}
	return nil
}
