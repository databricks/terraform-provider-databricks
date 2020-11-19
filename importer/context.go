package importer

import (
	"context"
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/identity"
	"github.com/databrickslabs/databricks-terraform/provider"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/zclconf/go-cty/cty"
)

/** High level overview of importer design:

                                    +----------+  Add  +--------------------+
                                    | resource +-------> stateApproximation |
                                    +--^-------+       +----|-------------|-+
                                       |                    |             |
  +------------------------+           | Emit               |             |
  | "normal provider flow" |    +------^-----+        +-----V-----+   +---V---+
  +------------^-----------+    | importable +--------> reference |   | scope |
               |                +------^-----+        +--------|--+   +---V---+
+--------------+--------------+        |                       |          |
|terraform-provider-databricks|        |                       |          |
+--------------+--------------+        |                       |          |
               |                       | List               +--V----------V---+
    +----------v---------+        +----^------------+       |                 |
    |                    |        |                 |       |    Generated    |
    |  importer command  +-------->  importContext  |       |    HCL Files    |
    |                    |        |                 |       |                 |
    +--------------------+        +-----------------+       +-----------------+
*/

type importContext struct {
	Module      string
	Client      *common.DatabricksClient
	State       stateApproximation
	Importables map[string]importable
	Resources   map[string]*schema.Resource
	Scope       importedResources
	Files       map[string]*hclwrite.File
	Directory   string
	importing   map[string]bool
	nameFixes   []regexFix
	hclFixes    []regexFix
	allUsers    []identity.ScimUser
	allGroups   []identity.ScimGroup

	debug          bool
	services       string
	lastActiveDays int64
}

func newImportContext(c *common.DatabricksClient) *importContext {
	p := provider.DatabricksProvider()
	return &importContext{
		Module:      "",
		Client:      c,
		State:       stateApproximation{},
		Importables: resourcesMap,
		Resources:   p.ResourcesMap,
		Files:       map[string]*hclwrite.File{},
		Scope:       []*resource{},
		Directory:   "/tmp/importer",
		importing:   map[string]bool{},
		nameFixes: []regexFix{
			{regexp.MustCompile(`[0-9a-f]{8}[_-][0-9a-f]{4}[_-][0-9a-f]{4}` +
				`[_-][0-9a-f]{4}[_-][0-9a-f]{12}[_-]`), ""},
			{regexp.MustCompile(`[_-][0-9]+[\._-][0-9]+[\._-].*\.(whl|jar|egg)`), "_$1"},
			{regexp.MustCompile(`@.*$`), ""},
			{regexp.MustCompile(`[-\s\.\|]`), "_"},
			{regexp.MustCompile(`\W+`), ""},
			{regexp.MustCompile(`[_]{2,}`), "_"},
		},
		hclFixes: []regexFix{
			{regexp.MustCompile(`\{ "`), "{\n\t\t\""},
			{regexp.MustCompile(`", "`), "\",\n\t\t\""},
			{regexp.MustCompile(`" \}`), "\"\t\n}"},
		},
		allUsers: []identity.ScimUser{},
	}
}

func (ic *importContext) Run() error {
	if len(ic.services) == 0 {
		return fmt.Errorf("No services to import")
	}
	log.Printf("[INFO] Importing %s module into %s directory Databricks resources of %s services",
		ic.Module, ic.Directory, ic.services)
	for resourceName, ir := range ic.Importables {
		if ir.List == nil {
			continue
		}
		if !strings.Contains(ic.services, ir.Service) {
			log.Printf("[DEBUG] %s (%s service) is not part of the import",
				resourceName, ir.Service)
			continue
		}
		if err := ir.List(ic); err != nil {
			return err
		}
	}
	if len(ic.Scope) == 0 {
		return fmt.Errorf("No resources to import")
	}
	sh, err := os.Create(fmt.Sprintf("%s/import.sh", ic.Directory))
	if err != nil {
		return err
	}
	defer sh.Close()

	sort.Sort(ic.Scope)
	for _, r := range ic.Scope {
		ir := ic.Importables[r.Resource]
		f, ok := ic.Files[ir.Service]
		if !ok {
			f = hclwrite.NewEmptyFile()
			ic.Files[ir.Service] = f
		}
		body := f.Body()
		if ir.Body != nil {
			err := ir.Body(ic, body, r)
			if err != nil {
				return err
			}
		} else {
			pr := ic.Resources[r.Resource]
			resourceBlock := body.AppendNewBlock("resource", []string{r.Resource, r.Name})
			err := ic.dataToHcl(ir, []string{}, pr, pr.Data(ic.InstanceState(r)), resourceBlock.Body())
			if err != nil {
				return err
			}
		}

		// nolint
		sh.WriteString(r.ImportCommand(ic) + "\n")
	}
	for service, f := range ic.Files {
		formatted := hclwrite.Format(f.Bytes())
		// fix some formatting in a hacky way instead of writing 100 lines
		// of HCL AST writer code
		formatted = []byte(ic.regexFix(string(formatted), ic.hclFixes))
		log.Printf("[INFO] %s", formatted)
		tf, err := os.Create(fmt.Sprintf("%s/%s.tf", ic.Directory, service))
		if err != nil {
			return err
		}
		defer tf.Close()
		_, err = tf.Write(formatted)
		if err != nil {
			return err
		}
	}
	cmd := exec.CommandContext(context.Background(), "terraform", "fmt")
	cmd.Dir = ic.Directory
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (ic *importContext) Find(r *resource, pick string) hcl.Traversal {
	for _, sr := range ic.State.Resources {
		if sr.Type != r.Resource {
			continue
		}
		for _, i := range sr.Instances {
			if i.Attributes[r.Attribute].(string) == r.Value {
				if "data" == sr.Mode {
					return hcl.Traversal{
						hcl.TraverseRoot{Name: "data"},
						hcl.TraverseAttr{Name: sr.Type},
						hcl.TraverseAttr{Name: sr.Name},
						hcl.TraverseAttr{Name: pick},
					}
				}
				return hcl.Traversal{
					hcl.TraverseRoot{Name: sr.Type},
					hcl.TraverseAttr{Name: sr.Name},
					hcl.TraverseAttr{Name: pick},
				}
			}
		}
	}
	return nil
}

func (ic *importContext) Has(r *resource) bool {
	if _, visiting := ic.importing[r.String()]; visiting {
		return true
	}
	k, v := r.MatchPair()
	for _, sr := range ic.State.Resources {
		if sr.Type != r.Resource {
			continue
		}
		for _, i := range sr.Instances {
			if i.Attributes[k].(string) == v {
				if "data" == sr.Mode && ic.Module != sr.Module {
					return false
				}
				return true
			}
		}
	}
	return false
}

func (ic *importContext) InstanceState(r *resource) *terraform.InstanceState {
	k, v := r.MatchPair()
	for _, sr := range ic.State.Resources {
		if sr.Type != r.Resource {
			continue
		}
		for _, i := range sr.Instances {
			if i.Attributes[k].(string) == v {
				attrs := map[string]string{}
				for k, v := range i.Attributes {
					if k == "id" {
						continue
					}
					attrs[k] = v.(string)
				}
				return &terraform.InstanceState{
					ID:         r.ID,
					Attributes: attrs,
				}
			}
		}
	}
	return nil
}

func (ic *importContext) Add(r *resource, attrs map[string]string) {
	if ic.Has(r) {
		return
	}
	inst := instanceApproximation{
		Attributes: map[string]interface{}{},
	}
	for k, v := range attrs {
		inst.Attributes[k] = v
	}
	inst.Attributes["id"] = r.ID
	ic.State.Resources = append(ic.State.Resources, resourceApproximation{
		Mode:      "managed",
		Module:    ic.Module,
		Type:      r.Resource,
		Name:      r.Name,
		Instances: []instanceApproximation{inst},
	})
	// in single-threaded scenario scope is toposorted
	ic.Scope = append(ic.Scope, r)
}

func (ic *importContext) regexFix(s string, fixes []regexFix) string {
	for _, x := range fixes {
		s = x.Regex.ReplaceAllString(s, x.Replacement)
	}
	return s
}

func (ic *importContext) ResourceName(r *resource, d *schema.ResourceData) string {
	name := r.Name
	if name == "" && ic.Importables[r.Resource].Name != nil {
		name = ic.Importables[r.Resource].Name(d)
	}
	if name == "" {
		name = d.Id()
	}
	name = strings.ToLower(name)
	name = ic.regexFix(name, ic.nameFixes)
	// this is either numeric id or all-non-ascii
	if regexp.MustCompile(`^\d`).MatchString(name) || "" == name {
		if "" == name {
			name = d.Id()
		}
		name = fmt.Sprintf("r%x", md5.Sum([]byte(name)))[0:12]
	}
	return name
}

func (ic *importContext) Emit(r *resource) {
	// TODO: change into channels, if stack trace depth issues would surface
	_, v := r.MatchPair()
	if v == "" {
		log.Printf("[DEBUG] %s has got empty identifier", r)
		return
	}
	if ic.Has(r) {
		log.Printf("[DEBUG] %s already imported", r)
		return
	}
	ic.importing[r.String()] = true
	pr, ok := ic.Resources[r.Resource]
	if !ok {
		log.Printf("[ERROR] %s is not available in provider", r)
		return
	}
	ir, ok := ic.Importables[r.Resource]
	if !ok {
		log.Printf("[ERROR] %s is not available for import", r)
		return
	}
	if !strings.Contains(ic.services, ir.Service) {
		log.Printf("[DEBUG] %s (%s service) is not part of the import",
			r.Resource, ir.Service)
		return
	}
	if r.ID == "" {
		if ir.Search == nil {
			log.Printf("[ERROR] Searching %s is not available", r)
			return
		}
		err := ir.Search(ic, r)
		if err != nil {
			log.Printf("[ERROR] Cannot search for a resource %s: %v", err, r)
			return
		}
		if r.ID == "" {
			log.Printf("[INFO] Cannot find %s", r)
			return
		}
	}
	// empty data with resource schema
	d := pr.Data(&terraform.InstanceState{
		Attributes: map[string]string{},
		ID:         r.ID,
	})
	d.MarkNewResource()

	if pr.Read != nil {
		err := pr.Read(d, ic.Client)
		if err != nil {
			return
		}
	} else {
		dia := pr.ReadContext(context.Background(), d, ic.Client)
		if dia != nil {
			log.Printf("[ERROR] Error reading %s#%s: %v", r.Resource, r.ID, dia)
			return
		}
	}
	r.Name = ic.ResourceName(r, d)
	if ir.Import != nil {
		err := ir.Import(ic, d)
		if err != nil {
			log.Printf("[ERROR] Failed custom import of %s: %s", r, err)
			return
		}
	}
	state := d.State()
	if state == nil {
		log.Printf("[ERROR] state is nil for %s", r)
		return
	}
	ic.Add(r, state.Attributes)
}

// TODO: move to IC
var dependsRe = regexp.MustCompile(`(\.[\d]+)`)

func (ic *importContext) reference(i importable, path []string, key, value string, body *hclwrite.Body) {
	match := dependsRe.ReplaceAllString(strings.Join(append(path, key), "."), "")
	for _, d := range i.Depends {
		if d.Path != match {
			continue
		}
		attr := "id"
		if d.Match != "" {
			attr = d.Match
		}
		if d.Pick == "" {
			d.Pick = "id"
		}
		traversal := ic.Find(&resource{
			Resource:  d.Resource,
			Attribute: attr,
			Value:     value,
		}, d.Pick)
		if traversal != nil {
			body.SetAttributeTraversal(key, traversal)
			return
		}
	}
	body.SetAttributeValue(key, cty.StringVal(value))
}

func (ic *importContext) dataToHcl(i importable, path []string,
	pr *schema.Resource, d *schema.ResourceData, body *hclwrite.Body) error {
	for a, as := range pr.Schema {
		if as.Computed {
			continue
		}
		raw, ok := d.GetOk(strings.Join(append(path, a), "."))
		if !ok {
			continue
		}
		switch as.Type {
		case schema.TypeString:
			ic.reference(i, path, a, raw.(string), body)
		case schema.TypeBool:
			body.SetAttributeValue(a, cty.BoolVal(raw.(bool)))
		case schema.TypeInt:
			switch iv := raw.(type) {
			case int:
				body.SetAttributeValue(a, cty.NumberIntVal(int64(iv)))
			case int32:
				body.SetAttributeValue(a, cty.NumberIntVal(int64(iv)))
			case int64:
				body.SetAttributeValue(a, cty.NumberIntVal(iv))
			}
		case schema.TypeFloat:
			body.SetAttributeValue(a, cty.NumberFloatVal(raw.(float64)))
		case schema.TypeMap:
			// mapBlock := body.AppendNewBlock(a, []string{}).Body()
			ov := map[string]cty.Value{}
			for key, iv := range raw.(map[string]interface{}) {
				v := cty.StringVal(fmt.Sprintf("%v", iv))
				ov[key] = v
				// mapBlock.SetAttributeValue(key, v)
			}
			body.SetAttributeValue(a, cty.ObjectVal(ov))
		case schema.TypeSet:
			if rawSet, ok := raw.(*schema.Set); ok {
				rawList := rawSet.List()
				err := ic.readListFromData(i, append(path, a), d, rawList, body, as, func(i int) string {
					return strconv.Itoa(rawSet.F(rawList[i]))
				})
				if err != nil {
					return err
				}
			}
		case schema.TypeList:
			if rawList, ok := raw.([]interface{}); ok {
				err := ic.readListFromData(i, append(path, a), d, rawList, body, as, strconv.Itoa)
				if err != nil {
					return err
				}
			}
		default:
			return fmt.Errorf("Unsupported schema type: %v", path)
		}
	}
	return nil
}

func (ic *importContext) readListFromData(i importable, path []string, d *schema.ResourceData,
	rawList []interface{}, body *hclwrite.Body, as *schema.Schema,
	offsetConverter func(i int) string) error {
	if len(rawList) == 0 {
		return nil
	}
	name := path[len(path)-1]
	switch elem := as.Elem.(type) {
	case *schema.Resource:
		if as.MaxItems == 1 {
			nestedPath := append(path, offsetConverter(0))
			confBlock := body.AppendNewBlock(name, []string{})
			return ic.dataToHcl(i, nestedPath, elem, d, confBlock.Body())
		}
		for offset := range rawList {
			confBlock := body.AppendNewBlock(name, []string{})
			nestedPath := append(path, offsetConverter(offset))
			err := ic.dataToHcl(i, nestedPath, elem, d, confBlock.Body())
			if err != nil {
				return err
			}
		}
	case *schema.Schema:
		primitiveValues := []cty.Value{}
		for _, raw := range rawList {
			switch x := raw.(type) {
			case string:
				primitiveValues = append(primitiveValues, cty.StringVal(x))
			case int:
				primitiveValues = append(primitiveValues, cty.NumberIntVal(int64(x)))
			default:
				return fmt.Errorf("Unsupported primitive list: %#v", path)
			}
		}
		body.SetAttributeValue(name, cty.ListVal(primitiveValues))
	}
	return nil
}
