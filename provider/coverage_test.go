package provider

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

type CoverageReport struct {
	Resources []ResourceCoverage
}

type ResourceCoverage struct {
	Name    string
	Data    bool
	Docs    bool
	Readme  bool
	AccTest bool
	AccFile bool
	ResFile bool
	ResTest bool
	Fields  []FieldCoverage
}

func (rc ResourceCoverage) Prefixless() string {
	return strings.TrimPrefix(rc.Name, "databricks_")
}

func (rc ResourceCoverage) DocLocation() string {
	if rc.Data {
		return path.Join("../docs/data-sources", rc.Prefixless()+".md")
	}
	return path.Join("../docs/resources", rc.Prefixless()+".md")
}

func (rc ResourceCoverage) ResourceFilename() string {
	if rc.Data {
		return fmt.Sprintf("data_%s.go", rc.Prefixless())
	}
	return fmt.Sprintf("resource_%s.go", rc.Prefixless())
}

func (rc ResourceCoverage) ResourceFn() string {
	return strings.ReplaceAll(
		strings.Title(
			strings.ReplaceAll(
				rc.Prefixless(), "_", " ")), " ", "")
}

func (rc ResourceCoverage) TestFilename() string {
	if rc.Data {
		return fmt.Sprintf("data_%s_test.go", rc.Prefixless())
	}
	return fmt.Sprintf("resource_%s_test.go", rc.Prefixless())
}

func (rc ResourceCoverage) AccFilename() string {
	return fmt.Sprintf("acceptance/%s_test.go", rc.Prefixless())
}

func (rc ResourceCoverage) coverage(cb func(FieldCoverage) bool, green, yellow int) string {
	var x float32
	for _, v := range rc.Fields {
		if cb(v) {
			x++
		}
	}
	coverage := int(100 * x / float32(len(rc.Fields)))
	coverageStr := fmt.Sprintf("(%d%%)", coverage)
	if coverage > green {
		return fmt.Sprintf("✅ %6s", coverageStr)
	}
	if coverage > yellow {
		return fmt.Sprintf("👎 %6s", coverageStr)
	}
	return fmt.Sprintf("❌ %6s", coverageStr)
}

func (rc ResourceCoverage) DocCoverage() string {
	return rc.coverage(func(fc FieldCoverage) bool {
		return fc.Docs
	}, 80, 50)
}

func (rc ResourceCoverage) AccCoverage() string {
	return rc.coverage(func(fc FieldCoverage) bool {
		return fc.AccTest
	}, 40, 20)
}

func (rc ResourceCoverage) UnitCoverage() string {
	return rc.coverage(func(fc FieldCoverage) bool {
		return fc.UnitTest
	}, 40, 20)
}

type FieldCoverage struct {
	Name     string
	Docs     bool
	AccTest  bool
	UnitTest bool
}

func (fc FieldCoverage) EverythingCovered() bool {
	return fc.Docs && fc.AccTest && fc.UnitTest
}

func newResourceCoverage(files FileSet, name string, s map[string]*schema.Schema, data bool) ResourceCoverage {
	r := ResourceCoverage{
		Name:    name,
		Data:    data,
		Readme:  files.Exists("../README.md", name),
		AccTest: files.Exists(`acceptance/.*_test.go`, fmt.Sprintf(`"%s"`, name)),
	}
	r.Docs = fileExists(r.DocLocation())
	// acceptance test file with a correct name
	r.AccFile = files.Exists(r.AccFilename(), r.Name)
	// resource file with a correct name
	r.ResFile = files.Exists(r.ResourceFilename(), r.ResourceFn())
	// resource unit test file with a correct name
	r.ResTest = files.Exists(r.TestFilename(), r.ResourceFn())
	r.Fields = fields(r, s, files)
	sort.Slice(r.Fields, func(i, j int) bool {
		return r.Fields[i].Name < r.Fields[j].Name
	})
	return r
}

func TestCoverageReport(t *testing.T) {
	if _, ok := os.LookupEnv("VSCODE_PID"); !ok {
		t.Skip("Cleaning up tests only from IDE")
	}
	files, err := recursiveChildren("..")
	assert.NoError(t, err)

	p := DatabricksProvider()
	var cr CoverageReport
	var longestResourceName, longestFieldName int

	for k, v := range p.ResourcesMap {
		if len(k) > longestResourceName {
			longestResourceName = len(k)
		}
		r := newResourceCoverage(files, k, v.Schema, false)
		cr.Resources = append(cr.Resources, r)
	}
	for k, v := range p.DataSourcesMap {
		if len(k) > longestResourceName {
			longestResourceName = len(k)
		}
		r := newResourceCoverage(files, k, v.Schema, true)
		cr.Resources = append(cr.Resources, r)
	}
	sort.Slice(cr.Resources, func(i, j int) bool {
		return cr.Resources[i].Name < cr.Resources[j].Name
	})
	report, err := os.OpenFile("completeness.md", os.O_CREATE|os.O_WRONLY, 0755)
	assert.NoError(t, err)
	defer report.Close()

	report.WriteString("| Resource | Readme | Docs | Acceptance Test | Acceptance File | Resource File | Unit test |\n")
	report.WriteString("| --- | --- | --- | --- | --- | --- | --- |\n")
	resSummaryFormat := "| %" + fmt.Sprint(longestResourceName) + "s | %s | %s | %s | %s | %s | %s |\n"
	for _, r := range cr.Resources {
		for _, field := range r.Fields {
			if len(field.Name) > longestFieldName {
				longestFieldName = len(field.Name)
			}
		}
		name := r.Name
		if r.Data {
			name = "* " + name
		}
		report.WriteString(fmt.Sprintf(resSummaryFormat, name,
			checkbox(r.Readme),
			r.DocCoverage(),
			r.AccCoverage(),
			checkbox(r.AccFile),
			checkbox(r.ResFile),
			r.UnitCoverage(),
		))
	}
	report.WriteString("\n\n| Resource | Field | Docs | Acceptance Test | Unit Test |\n")
	report.WriteString("| --- | --- | --- | --- | --- |\n")
	fieldSummaryFormat := "| %" + fmt.Sprint(longestResourceName) + "s | %" +
		fmt.Sprint(longestFieldName) + "s | %s | %s | %s |\n"
	for _, r := range cr.Resources {
		for _, field := range r.Fields {
			if field.EverythingCovered() {
				continue
			}
			report.WriteString(fmt.Sprintf(fieldSummaryFormat,
				r.Name,
				field.Name,
				checkbox(field.Docs),
				checkbox(field.AccTest),
				checkbox(field.UnitTest),
			))
		}
	}
}

func fields(r ResourceCoverage, s map[string]*schema.Schema, files FileSet) (fields []FieldCoverage) {
	type pathWrapper struct {
		r    *schema.Resource
		path []string
	}
	queue := []pathWrapper{
		{
			r: &schema.Resource{Schema: s},
		},
	}
	doc := File{Absolute: r.DocLocation()}

	noisyDuplicates := map[string]bool{
		"new_cluster": true,
		"task":        true,
	}
	for {
		head := queue[0]
		queue = queue[1:]
		for field, v := range head.r.Schema {
			if noisyDuplicates[field] {
				continue
			}
			path := append(head.path, field)
			if nested, ok := v.Elem.(*schema.Resource); ok {
				queue = append(queue, pathWrapper{
					r:    nested,
					path: path,
				})
			}
			fc := FieldCoverage{
				Name: strings.Join(path, "."),
			}
			if v.Computed {
				fc.Name += " (computed)"
			}
			if r.Docs {
				fc.Docs = doc.MustMatch(field)
			}
			if r.AccTest {
				fc.AccTest = files.Exists(`acceptance/.*_test.go`, field)
			}
			if r.ResTest {
				fc.UnitTest = files.Exists(r.TestFilename(), field)
			}
			fields = append(fields, fc)
		}
		if len(queue) == 0 {
			break
		}
	}
	return fields
}

func checkbox(b bool) string {
	if b {
		return "✅"
	}
	return "❌"
}

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

type FileSet []File

func (fi FileSet) Exists(pathRegex, needleRegex string) bool {
	path := regexp.MustCompile(pathRegex)
	needle := regexp.MustCompile(needleRegex)
	for _, v := range fi {
		if !path.MatchString(v.Absolute) {
			continue
		}
		if v.Match(needle) {
			return true
		}
	}
	return false
}

type File struct {
	fs.DirEntry
	Absolute string
}

func (fi File) MustMatch(needle string) bool {
	return fi.Match(regexp.MustCompile(needle))
}

func (fi File) Match(needle *regexp.Regexp) bool {
	raw, err := fi.Raw()
	if err != nil {
		log.Printf("[ERROR] read %s: %s", fi.Absolute, err)
		return false
	}
	return needle.Match(raw)
}

func (fi File) Raw() ([]byte, error) {
	f, err := os.Open(fi.Absolute)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(f)
}

func recursiveChildren(dir string) (found FileSet, err error) {
	queue, err := readDir(dir)
	if err != nil {
		return nil, err
	}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if !current.IsDir() {
			found = append(found, current)
			continue
		}
		if current.Name() == "vendor" {
			continue
		}
		if current.Name() == "scripts" {
			continue
		}
		children, err := readDir(current.Absolute)
		if err != nil {
			return nil, err
		}
		queue = append(queue, children...)
	}
	return found, nil
}

func readDir(dir string) (queue []File, err error) {
	f, err := os.Open(dir)
	if err != nil {
		return
	}
	defer f.Close()
	dirs, err := f.ReadDir(-1)
	if err != nil {
		return
	}
	for _, v := range dirs {
		absolute := path.Join(dir, v.Name())
		queue = append(queue, File{v, absolute})
	}
	return
}
