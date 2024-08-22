//go:build ignore
// +build ignore

// Usage: go run provider/gen/main.go -name mws_workspaces -package mws -is-data -dry-run=false
package main

import (
	"errors"
	"flag"
	"fmt"
	"html/template"
	"os"
	"strings"
)

var ctx Context

func main() {
	flag.StringVar(&ctx.Name, "name", "", "name of the resource, like `mws_workspaces`")
	flag.StringVar(&ctx.Package, "package", "", "package of the resource, like `clusters`")
	flag.StringVar(&ctx.DocCategory, "category", "Other", "documentation category")
	flag.BoolVar(&ctx.DataSource, "is-data", false, "is this a data resource")
	flag.BoolVar(&ctx.DryRun, "dry-run", true, "print to stdout instead of real files")

	flag.Parse()

	if ctx.Name == "" || ctx.Package == "" {
		println("USAGE: go run provider/gen/main.go -name res -package pkg")
		flag.PrintDefaults()
		os.Exit(1)
	}

	err := ctx.Generate()
	if err != nil {
		fmt.Printf("ERROR: %s\n\n", err)
		os.Exit(1)
	}
}

type Context struct {
	Package     string
	Name        string
	DocCategory string
	DataSource  bool
	DryRun      bool
	BT          string

	tmpl *template.Template
}

func (c *Context) Generate() error {
	if !c.DataSource {
		return errors.New("only data sources are supported now")
	}
	c.BT = "`"
	c.tmpl = template.Must(template.ParseGlob("provider/gen/*.go.tmpl"))
	return c.FileSet(map[string]string{
		"data_x.go.tmpl":         "{{.Package}}/data_{{.Name}}.go",
		"data_x_test.go.tmpl":    "{{.Package}}/data_{{.Name}}_test.go",
		"data_x_acctest.go.tmpl": "{{.Package}}/acceptance/data_{{.Name}}_test.go",
		"data_x.md.go.tmpl":      "docs/data-sources/{{.Name}}.md",
		"data_x_manual.go.tmpl":  "stdout",
	})
}

func (c *Context) FileSet(m map[string]string) (err error) {
	for k, v := range m {
		err = c.File(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Context) File(contentTRef, nameT string) error {
	nt, err := template.New("filename").Parse(nameT)
	if err != nil {
		return fmt.Errorf("parse %s: %w", nameT, err)
	}
	var filename, contents strings.Builder
	err = nt.Execute(&filename, c)
	if err != nil {
		return fmt.Errorf("exec %s: %w", nameT, err)
	}
	err = c.tmpl.ExecuteTemplate(&contents, contentTRef, &ctx)
	if err != nil {
		return fmt.Errorf("exec %s: %w", contentTRef, err)
	}
	if c.DryRun {
		fmt.Printf("\n---\nDRY RUN: %s\n---\n%s", &filename, &contents)
		return nil
	}
	if nameT == "stdout" {
		println(contents.String())
		return nil
	}
	file, err := os.OpenFile(filename.String(), os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return fmt.Errorf("open %s: %w", &filename, err)
	}
	_, err = file.WriteString(contents.String())
	if err != nil {
		return fmt.Errorf("write %s: %w", &filename, err)
	}
	return file.Close()
}

func (c *Context) CamelName() string {
	return strings.ReplaceAll(
		strings.Title(
			strings.ReplaceAll(
				c.Name, "_", " ")), " ", "")
}

func (c *Context) LowerCamel() string {
	cc := c.CamelName()
	return strings.ToLower(cc[0:1]) + cc[1:]
}
