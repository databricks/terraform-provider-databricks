package exporter

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

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/compute"
	"github.com/databrickslabs/terraform-provider-databricks/identity"
	"github.com/databrickslabs/terraform-provider-databricks/provider"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
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
	Context     context.Context
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
	mountMap    map[string]mount
	variables   map[string]string

	debug               bool
	mounts              bool
	services            string
	listing             string
	match               string
	lastActiveDays      int64
	generateDeclaration bool
	meAdmin             bool
	prefix              string
}

type mount struct {
	URL             string
	InstanceProfile string
	ClusterID       string
}

func newImportContext(c *common.DatabricksClient) *importContext {
	p := provider.DatabricksProvider()
	p.TerraformVersion = "importer"
	p.SetMeta(c)
	ctx := context.WithValue(context.Background(), common.Provider, p)
	c.WithCommandExecutor(func(
		ctx context.Context,
		c *common.DatabricksClient) common.CommandExecutor {
		return compute.NewCommandsAPI(ctx, c)
	})
	return &importContext{
		Client:      c,
		Context:     ctx,
		State:       stateApproximation{},
		Importables: resourcesMap,
		Resources:   p.ResourcesMap,
		Files:       map[string]*hclwrite.File{},
		Scope:       []*resource{},
		importing:   map[string]bool{},
		nameFixes: []regexFix{
			{regexp.MustCompile(`[0-9a-f]{8}[_-][0-9a-f]{4}[_-][0-9a-f]{4}` +
				`[_-][0-9a-f]{4}[_-][0-9a-f]{12}[_-]`), ""},
			{regexp.MustCompile(`[_-][0-9]+[\._-][0-9]+[\._-].*\.([a-z0-9]{1,4})`), "_$1"},
			{regexp.MustCompile(`@.*$`), ""},
			{regexp.MustCompile(`[-\s\.\|]`), "_"},
			{regexp.MustCompile(`\W+`), ""},
			{regexp.MustCompile(`[_]{2,}`), "_"},
		},
		hclFixes: []regexFix{
			{regexp.MustCompile(`\{ `), "{\n\t\t"},
			{regexp.MustCompile(`, `), ",\n\t\t"},
		},
		allUsers:  []identity.ScimUser{},
		variables: map[string]string{},
	}
}

func (ic *importContext) Run() error {
	if len(ic.services) == 0 {
		return fmt.Errorf("no services to import")
	}
	log.Printf("[INFO] Importing %s module into %s directory Databricks resources of %s services",
		ic.Module, ic.Directory, ic.services)

	info, err := os.Stat(ic.Directory)
	if os.IsNotExist(err) {
		err = os.MkdirAll(ic.Directory, 0755)
		if err != nil {
			return fmt.Errorf("can't create directory %s", ic.Directory)
		}
	} else if !info.IsDir() {
		return fmt.Errorf("the path %s is not a directory", ic.Directory)
	}
	usersAPI := identity.NewUsersAPI(ic.Context, ic.Client)
	me, err := usersAPI.Me()
	if err != nil {
		return err
	}
	for _, g := range me.Groups {
		if g.Display == "admins" {
			ic.meAdmin = true
			break
		}
	}
	for resourceName, ir := range ic.Importables {
		if ir.List == nil {
			continue
		}
		if !strings.Contains(ic.listing, ir.Service) {
			log.Printf("[DEBUG] %s (%s service) is not part of listing",
				resourceName, ir.Service)
			continue
		}
		if err := ir.List(ic); err != nil {
			return err
		}
	}
	if len(ic.Scope) == 0 {
		return fmt.Errorf("no resources to import")
	}
	sh, err := os.OpenFile(fmt.Sprintf("%s/import.sh", ic.Directory), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer sh.Close()
	// nolint
	sh.WriteString("#!/bin/sh\n\n")

	if ic.generateDeclaration {
		dcfile, err := os.Create(fmt.Sprintf("%s/databricks.tf", ic.Directory))
		if err != nil {
			return err
		}
		// nolint
		dcfile.WriteString(
			`terraform {
				required_providers {
			  		databricks = {
						source = "databrickslabs/databricks"
						version = "` + common.Version() + `"
				  	}
				}
		  	}

		  	provider "databricks" {
		  	}
		  	`)
		dcfile.Close()
	}

	sort.Sort(ic.Scope)
	scopeSize := len(ic.Scope)
	log.Printf("[INFO] Generating configuration for %d resources", scopeSize)
	for i, r := range ic.Scope {
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
				log.Printf("[ERROR] %s", err.Error())
			}
		} else {
			resourceBlock := body.AppendNewBlock("resource", []string{r.Resource, r.Name})
			err := ic.dataToHcl(ir, []string{}, ic.Resources[r.Resource],
				r.Data, resourceBlock.Body())
			if err != nil {
				log.Printf("[ERROR] %s", err.Error())
			}
		}
		if i%50 == 0 {
			log.Printf("[INFO] Generated %d of %d resources", i, scopeSize)
		}
		if r.Mode != "data" {
			// nolint
			sh.WriteString(r.ImportCommand(ic) + "\n")
		}
	}
	for service, f := range ic.Files {
		formatted := hclwrite.Format(f.Bytes())
		// fix some formatting in a hacky way instead of writing 100 lines
		// of HCL AST writer code
		formatted = []byte(ic.regexFix(string(formatted), ic.hclFixes))
		log.Printf("[DEBUG] %s", formatted)
		generatedFile := fmt.Sprintf("%s/%s.tf", ic.Directory, service)
		if tf, err := os.Create(generatedFile); err == nil {
			defer tf.Close()
			if _, err = tf.Write(formatted); err != nil {
				return err
			}
		}
		log.Printf("[INFO] Created %s", generatedFile)
	}
	if len(ic.variables) > 0 {
		vf, err := os.Create(fmt.Sprintf("%s/vars.tf", ic.Directory))
		if err != nil {
			return err
		}
		defer vf.Close()
		f := hclwrite.NewEmptyFile()
		body := f.Body()
		for k, v := range ic.variables {
			b := body.AppendNewBlock("variable", []string{k}).Body()
			b.SetAttributeValue("description", cty.StringVal(v))
		}
		// nolint
		vf.Write(f.Bytes())
		log.Printf("[INFO] Written %d variables", len(ic.variables))
	}
	cmd := exec.CommandContext(context.Background(), "terraform", "fmt")
	cmd.Dir = ic.Directory
	err = cmd.Run()
	if err != nil {
		return err
	}
	log.Printf("[INFO] Done. Please edit the files and roll out new environment.")
	return nil
}

func (ic *importContext) MatchesName(n string) bool {
	if ic.match == "" {
		return true
	}
	return strings.Contains(strings.ToLower(n), strings.ToLower(ic.match))
}

func (ic *importContext) Find(r *resource, pick string) hcl.Traversal {
	for _, sr := range ic.State.Resources {
		if sr.Type != r.Resource {
			continue
		}
		for _, i := range sr.Instances {
			if i.Attributes[r.Attribute].(string) == r.Value {
				if sr.Mode == "data" {
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
				return true
			}
		}
	}
	return false
}

func (ic *importContext) Add(r *resource) {
	if ic.Has(r) {
		return
	}
	state := r.Data.State()
	if state == nil {
		log.Printf("[ERROR] state is nil for %s", r)
		return
	}
	inst := instanceApproximation{
		Attributes: map[string]interface{}{},
	}
	for k, v := range state.Attributes {
		inst.Attributes[k] = v
	}
	if r.Mode == "" {
		r.Mode = "managed"
	}
	inst.Attributes["id"] = r.ID
	ic.State.Resources = append(ic.State.Resources, resourceApproximation{
		Mode:      r.Mode,
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

func (ic *importContext) ResourceName(r *resource) string {
	name := r.Name
	if name == "" && ic.Importables[r.Resource].Name != nil {
		name = ic.Importables[r.Resource].Name(r.Data)
	}
	if name == "" {
		name = r.ID
	}
	name = ic.prefix + name
	name = strings.ToLower(name)
	name = ic.regexFix(name, ic.nameFixes)
	// this is either numeric id or all-non-ascii
	if regexp.MustCompile(`^\d`).MatchString(name) || name == "" {
		if name == "" {
			name = r.ID
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
		if err := ir.Search(ic, r); err != nil {
			log.Printf("[ERROR] Cannot search for a resource %s: %v", err, r)
			return
		}
		if r.ID == "" {
			log.Printf("[INFO] Cannot find %s", r)
			return
		}
	}
	if r.Data == nil {
		// empty data with resource schema
		r.Data = pr.Data(&terraform.InstanceState{
			Attributes: map[string]string{},
			ID:         r.ID,
		})
		r.Data.MarkNewResource()

		if pr.Read != nil {
			if err := pr.Read(r.Data, ic.Client); err != nil {
				log.Printf("[ERROR] Error reading %s#%s: %v", r.Resource, r.ID, err)
				return
			}
		} else {
			if dia := pr.ReadContext(context.Background(), r.Data, ic.Client); dia != nil {
				log.Printf("[ERROR] Error reading %s#%s: %v", r.Resource, r.ID, dia)
				return
			}
		}
		if r.Data.Id() == "" {
			r.Data.SetId(r.ID)
		}
	}
	r.Name = ic.ResourceName(r)
	if ir.Import != nil {
		if err := ir.Import(ic, r); err != nil {
			log.Printf("[ERROR] Failed custom import of %s: %s", r, err)
			return
		}
	}
	ic.Add(r)
}

// TODO: move to IC
var dependsRe = regexp.MustCompile(`(\.[\d]+)`)

func (ic *importContext) reference(i importable, path []string, value string) hclwrite.Tokens {
	match := dependsRe.ReplaceAllString(strings.Join(path, "."), "")
	for _, d := range i.Depends {
		if d.Path != match {
			continue
		}
		attr := "id"
		if d.Match != "" {
			attr = d.Match
		}
		traversal := ic.Find(&resource{
			Resource:  d.Resource,
			Attribute: attr,
			Value:     value,
		}, attr)

		if traversal == nil {
			break
		}
		return hclwrite.TokensForTraversal(traversal)
	}
	return hclwrite.TokensForValue(cty.StringVal(value))
}

func (ic *importContext) variable(name, desc string) hclwrite.Tokens {
	ic.variables[name] = desc
	return hclwrite.TokensForTraversal(hcl.Traversal{
		hcl.TraverseRoot{Name: "var"},
		hcl.TraverseAttr{Name: name},
	})
}

type fieldTuple struct {
	Field  string
	Schema *schema.Schema
}

func (ic *importContext) dataToHcl(i importable, path []string,
	pr *schema.Resource, d *schema.ResourceData, body *hclwrite.Body) error {
	ss := []fieldTuple{}
	for a, as := range pr.Schema {
		ss = append(ss, fieldTuple{a, as})
	}
	sort.Slice(ss, func(i, j int) bool {
		// it just happens that reverse field order
		// makes the most beautiful configs
		return ss[i].Field > ss[j].Field
	})
	for _, tuple := range ss {
		a, as := tuple.Field, tuple.Schema
		if as.Computed {
			continue
		}
		raw, ok := d.GetOk(strings.Join(append(path, a), "."))
		if !ok {
			continue
		}
		switch as.Type {
		case schema.TypeString:
			body.SetAttributeRaw(a, ic.reference(i, append(path, a), raw.(string)))
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
			ov := map[string]cty.Value{}
			for key, iv := range raw.(map[string]interface{}) {
				v := cty.StringVal(fmt.Sprintf("%v", iv))
				ov[key] = v
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
			return fmt.Errorf("unsupported schema type: %v", path)
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
		toks := hclwrite.Tokens{}
		toks = append(toks, &hclwrite.Token{
			Type:  hclsyntax.TokenOBrack,
			Bytes: []byte{'['},
		})
		for _, raw := range rawList {
			if len(toks) != 1 {
				toks = append(toks, &hclwrite.Token{
					Type:  hclsyntax.TokenComma,
					Bytes: []byte{','},
				})
			}
			switch x := raw.(type) {
			case string:
				toks = append(toks, ic.reference(i, path, x)...)
			case int:
				// probably we don't even use integer lists?...
				toks = append(toks, hclwrite.TokensForValue(
					cty.NumberIntVal(int64(x)))...)
			default:
				return fmt.Errorf("unsupported primitive list: %#v", path)
			}
		}
		toks = append(toks, &hclwrite.Token{
			Type:  hclsyntax.TokenCBrack,
			Bytes: []byte{']'},
		})
		body.SetAttributeRaw(name, toks)
	}
	return nil
}
