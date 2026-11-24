package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"reflect"
	"strconv"
	"strings"

	"golang.org/x/tools/go/packages"
)

// schemaField is a field of the schema struct, with the rules that
// common.StructToSchema derives from its tags.
type schemaField struct {
	goName string
	kind   schemaKind
	// omitEmpty is true when the json tag has omitempty.
	omitEmpty bool
	optional  bool
	computed  bool
}

type schemaKind int

const (
	schemaUnsupported schemaKind = iota
	schemaString
	// schemaNamed is a named type, such as catalog.VolumeType, which can only map to an
	// sdk-go string enum.
	schemaNamed
)

// readSchemaStruct returns the package name of the file at path and the fields of its
// struct, keyed by attribute name.
func readSchemaStruct(path, name string) (string, map[string]schemaField, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.SkipObjectResolution)
	if err != nil {
		return "", nil, err
	}
	st := findStruct(file, name)
	if st == nil {
		return "", nil, fmt.Errorf("%s has no struct %s", path, name)
	}
	fields := map[string]schemaField{}
	for _, f := range st.Fields.List {
		// Embedded fields, such as common.Namespace, are declared in other files, and
		// fields without tags are not in the schema.
		if len(f.Names) == 0 || f.Tag == nil {
			continue
		}
		tag, err := strconv.Unquote(f.Tag.Value)
		if err != nil {
			return "", nil, fmt.Errorf("%s: tag %s: %w", name, f.Tag.Value, err)
		}
		attribute, field := schemaFieldOf(reflect.StructTag(tag), f.Type)
		if attribute == "-" || attribute == "" {
			continue
		}
		for _, n := range f.Names {
			field.goName = n.Name
			fields[attribute] = field
		}
	}
	return file.Name.Name, fields, nil
}

func findStruct(file *ast.File, name string) *ast.StructType {
	for _, decl := range file.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.TYPE {
			continue
		}
		for _, s := range gen.Specs {
			ts := s.(*ast.TypeSpec)
			if st, ok := ts.Type.(*ast.StructType); ok && ts.Name.Name == name {
				return st
			}
		}
	}
	return nil
}

// schemaFieldOf reads a field's attribute name and rules from its tags as
// common.StructToSchema does.
func schemaFieldOf(tag reflect.StructTag, typ ast.Expr) (string, schemaField) {
	jsonTag := tag.Get("json")
	attribute, _, _ := strings.Cut(jsonTag, ",")
	field := schemaField{omitEmpty: strings.Contains(jsonTag, "omitempty")}
	field.optional = field.omitEmpty
	for t := range strings.SplitSeq(tag.Get("tf"), ",") {
		switch {
		case t == "optional":
			field.optional = true
		case t == "computed":
			field.computed = true
		case strings.HasPrefix(t, "alias:"):
			attribute = strings.TrimPrefix(t, "alias:")
		}
	}
	switch t := typ.(type) {
	case *ast.Ident:
		field.kind = schemaNamed
		if isStringIdent(t) {
			field.kind = schemaString
		} else if types.Universe.Lookup(t.Name) != nil {
			field.kind = schemaUnsupported
		}
	case *ast.SelectorExpr:
		field.kind = schemaNamed
	}
	return attribute, field
}

// sdkGoPackage is the parsed source of an sdk-go package.
type sdkGoPackage struct {
	path string
	// types are the types of the package's type declarations, by name.
	types map[string]ast.Expr
}

// loadSdkGoPackage parses the source of the sdk-go package, which it finds through the
// module of dir. It doesn't type-check the package, so it doesn't depend on the export
// data format of the Go toolchain.
func loadSdkGoPackage(dir, path string) (sdkGoPackage, error) {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedSyntax,
		Dir:  dir,
	}
	pkgs, err := packages.Load(cfg, path)
	if err != nil {
		return sdkGoPackage{}, err
	}
	if len(pkgs) != 1 {
		return sdkGoPackage{}, fmt.Errorf("%s matches %d packages", path, len(pkgs))
	}
	if len(pkgs[0].Errors) > 0 {
		return sdkGoPackage{}, fmt.Errorf("load %s: %v", path, pkgs[0].Errors[0])
	}
	pkg := sdkGoPackage{path: path, types: map[string]ast.Expr{}}
	for _, file := range pkgs[0].Syntax {
		for _, decl := range file.Decls {
			gen, ok := decl.(*ast.GenDecl)
			if !ok || gen.Tok != token.TYPE {
				continue
			}
			for _, s := range gen.Specs {
				ts := s.(*ast.TypeSpec)
				pkg.types[ts.Name.Name] = ts.Type
			}
		}
	}
	return pkg, nil
}

// sdkGoField is the type of an sdk-go struct field.
type sdkGoField struct {
	kind sdkGoKind
	// enum is the name of the string enum type of an sdkGoEnum field.
	enum string
}

type sdkGoKind int

const (
	sdkGoUnsupported sdkGoKind = iota
	sdkGoStringPointer
	sdkGoEnum
)

// fields returns the fields of the struct name, keyed by Go name.
func (p sdkGoPackage) fields(name string) (map[string]sdkGoField, error) {
	t, ok := p.types[name]
	if !ok {
		return nil, fmt.Errorf("%s has no type %s", p.path, name)
	}
	st, ok := t.(*ast.StructType)
	if !ok {
		return nil, fmt.Errorf("%s.%s is not a struct", p.path, name)
	}
	fields := map[string]sdkGoField{}
	for _, f := range st.Fields.List {
		for _, n := range f.Names {
			fields[n.Name] = p.fieldOf(f.Type)
		}
	}
	return fields, nil
}

func (p sdkGoPackage) fieldOf(t ast.Expr) sdkGoField {
	switch t := t.(type) {
	case *ast.StarExpr:
		if isStringIdent(t.X) {
			return sdkGoField{kind: sdkGoStringPointer}
		}
	case *ast.Ident:
		if underlying, ok := p.types[t.Name]; ok && isStringIdent(underlying) {
			return sdkGoField{kind: sdkGoEnum, enum: t.Name}
		}
	}
	return sdkGoField{}
}

func isStringIdent(e ast.Expr) bool {
	id, ok := e.(*ast.Ident)
	return ok && id.Name == "string"
}
