package util

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func reflectKind(k reflect.Kind) string {
	switch k {
	case reflect.Bool:
		return "Bool"
	case reflect.Int:
		return "Int"
	case reflect.Int8:
		return "Int8"
	case reflect.Int16:
		return "Int16"
	case reflect.Int32:
		return "Int32"
	case reflect.Int64:
		return "Int64"
	case reflect.Uint:
		return "Uint"
	case reflect.Uint8:
		return "Uint8"
	case reflect.Uint16:
		return "Uint16"
	case reflect.Uint32:
		return "Uint32"
	case reflect.Uint64:
		return "Uint64"
	case reflect.Uintptr:
		return "Uintptr"
	case reflect.Float32:
		return "Float32"
	case reflect.Float64:
		return "Float64"
	case reflect.Complex64:
		return "Complex64"
	case reflect.Complex128:
		return "Complex128"
	case reflect.Array:
		return "Array"
	case reflect.Chan:
		return "Chan"
	case reflect.Func:
		return "Func"
	case reflect.Interface:
		return "Interface"
	case reflect.Ptr:
		return "Ptr"
	case reflect.Slice:
		return "Slice"
	case reflect.String:
		return "String"
	case reflect.Struct:
		return "Struct"
	case reflect.UnsafePointer:
		return "UnsafePointer"
	default:
		return "other"
	}
}

// SchemaPath helps to navigate
func SchemaPath(s map[string]*schema.Schema, path ...string) (*schema.Schema, error) {
	cs := s
	for _, p := range path {
		v, ok := cs[p]
		if !ok {
			return nil, fmt.Errorf("Missing key %s", p)
		}
		if p == path[len(path)-1] {
			return v, nil
		}
		cv, ok := v.Elem.(*schema.Resource)
		if !ok {
			return nil, fmt.Errorf("%s is not nested resource", p)
		}
		cs = cv.Schema
	}
	return nil, fmt.Errorf("%v does not compute", path)
}

// StructToSchema makes schema from a struct type & applies customizations from callback given
func StructToSchema(v interface{}, customize func(map[string]*schema.Schema) map[string]*schema.Schema) map[string]*schema.Schema {
	rv := reflect.ValueOf(v)
	scm := typeToSchema(rv, rv.Type())
	scm = customize(scm)
	return scm
}

func typeToSchema(v reflect.Value, t reflect.Type) map[string]*schema.Schema {
	scm := map[string]*schema.Schema{}
	rk := v.Kind()
	if rk != reflect.Struct {
		panic(fmt.Errorf("Schema value of Struct is expected, but got %s: %#v", reflectKind(rk), v))
	}
	for i := 0; i < v.NumField(); i++ {
		typeField := t.Field(i)

		jsonTag := typeField.Tag.Get("json")
		tfTag := typeField.Tag.Get("tf")
		jsonFieldName := strings.Split(jsonTag, ",")[0]
		if jsonFieldName == "-" {
			continue
		}
		// TODO: conflicts with - tf:"group:a"
		// TODO: tf:"default:name"
		// TODO: tf:"conflicts:instance_pool_id"
		// TODO:  tf:"max_items:10"
		// TODO: tf:"alias:library"
		scm[jsonFieldName] = &schema.Schema{}
		if strings.Contains(jsonTag, "omitempty") {
			scm[jsonFieldName].Optional = true
		} else {
			scm[jsonFieldName].Required = true
		}
		switch typeField.Type.Kind() {
		case reflect.Int, reflect.Int32, reflect.Int64:
			scm[jsonFieldName].Type = schema.TypeInt
		case reflect.Float64:
			scm[jsonFieldName].Type = schema.TypeFloat
		case reflect.Bool:
			scm[jsonFieldName].Type = schema.TypeBool
		case reflect.String:
			scm[jsonFieldName].Type = schema.TypeString
		case reflect.Map:
			scm[jsonFieldName].Type = schema.TypeMap
		case reflect.Ptr:
			scm[jsonFieldName].MaxItems = 1
			scm[jsonFieldName].Type = schema.TypeList
			elem := typeField.Type.Elem()
			sv := reflect.New(elem).Elem()
			scm[jsonFieldName].Elem = &schema.Resource{
				Schema: typeToSchema(sv, elem),
			}
		// case reflect.Struct:
		// 	// this is required and pointer is no required :P
		// 	scm[jsonFieldName].MaxItems = 1
		// 	scm[jsonFieldName].Type = schema.TypeList
		// 	elem := typeField.Type.Elem()
		// 	sv := reflect.New(elem).Elem()
		// 	ns := typeToSchema(sv, elem)
		// 	scm[jsonFieldName].Elem = schema.Resource{
		// 		Schema: ns,
		// 	}
		case reflect.Slice:
			ft := schema.TypeList
			if strings.Contains(tfTag, "slice_set") {
				ft = schema.TypeSet
			}
			scm[jsonFieldName].Type = ft
			elem := typeField.Type.Elem()
			switch elem.Kind() {
			case reflect.Int:
				scm[jsonFieldName].Elem = &schema.Schema{Type: schema.TypeInt}
			case reflect.Float64:
				scm[jsonFieldName].Elem = &schema.Schema{Type: schema.TypeFloat}
			case reflect.Bool:
				scm[jsonFieldName].Elem = &schema.Schema{Type: schema.TypeBool}
			case reflect.String:
				scm[jsonFieldName].Elem = &schema.Schema{Type: schema.TypeString}
			case reflect.Struct:
				sv := reflect.New(elem).Elem()
				scm[jsonFieldName].Elem = &schema.Resource{
					Schema: typeToSchema(sv, elem),
				}
			}
		default:
			panic(fmt.Errorf("Unknown type for %s: %s", jsonFieldName, reflectKind(typeField.Type.Kind())))
		}
	}
	return scm
}

func iterFields(rv reflect.Value, path []string, s map[string]*schema.Schema,
	cb func(fieldSchema *schema.Schema, path []string, valueField *reflect.Value) error) error {
	rk := rv.Kind()
	if rk != reflect.Struct {
		return fmt.Errorf("Value of Struct is expected, but got %s: %#v", reflectKind(rk), rv)
	}
	if !rv.IsValid() {
		return fmt.Errorf("%s: got invalid reflect value %#v", path, rv)
	}
	for i := 0; i < rv.NumField(); i++ {
		typeField := rv.Type().Field(i)
		jsonTag := typeField.Tag.Get("json")
		jsonFieldName := strings.Split(jsonTag, ",")[0]
		if jsonFieldName == "-" {
			continue
		}
		fieldSchema, ok := s[jsonFieldName]
		if !ok {
			continue
		}
		omitEmpty := strings.Contains(jsonTag, "omitempty")
		if omitEmpty && !fieldSchema.Optional {
			return fmt.Errorf("Inconsistency: %s has omitempty, but is not optional", jsonFieldName)
		}
		defaultEmpty := reflect.ValueOf(fieldSchema.Default).Kind() == reflect.Invalid
		if fieldSchema.Optional && defaultEmpty && !omitEmpty {
			return fmt.Errorf("Inconsistency: %s is optional, default is empty, but has no omitempty", jsonFieldName)
		}
		valueField := rv.Field(i)
		err := cb(fieldSchema, append(path, jsonFieldName), &valueField)
		if err != nil {
			return fmt.Errorf("%s: %s", jsonFieldName, err)
		}
	}
	return nil
}

func collectionToMaps(v interface{}, s *schema.Schema) ([]interface{}, error) {
	resultList := []interface{}{}
	if sl, ok := v.([]string); ok {
		// most likely list of parameters to job task
		for _, str := range sl {
			resultList = append(resultList, str)
		}
		return resultList, nil
	}
	r, ok := s.Elem.(*schema.Resource)
	if !ok {
		return nil, fmt.Errorf("not resource")
	}
	var allItems []reflect.Value
	if s.MaxItems == 1 {
		allItems = append(allItems, reflect.ValueOf(v))
	} else {
		vs := reflect.ValueOf(v)
		for i := 0; i < vs.Len(); i++ {
			allItems = append(allItems, vs.Index(i))
		}
	}
	for _, v := range allItems {
		data := map[string]interface{}{}
		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				continue
			}
			v = v.Elem()
		}
		err := iterFields(v, []string{}, r.Schema, func(fieldSchema *schema.Schema,
			path []string, valueField *reflect.Value) error {
			fieldName := path[len(path)-1]
			fieldValue := valueField.Interface()
			switch fieldSchema.Type {
			case schema.TypeList, schema.TypeSet:
				nv, err := collectionToMaps(fieldValue, fieldSchema)
				if err != nil {
					return fmt.Errorf("%s: %v", path, err)
				}
				data[fieldName] = nv
			default:
				if s, ok := fieldValue.(string); ok && s == "" {
					return nil
				}
				data[fieldName] = fieldValue
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
		if len(data) == 0 {
			continue
		}
		resultList = append(resultList, data)
	}
	return resultList, nil
}

// StructToData reads result using schema onto resource data
func StructToData(result interface{}, s map[string]*schema.Schema, d *schema.ResourceData) error {
	return iterFields(reflect.ValueOf(result), []string{}, s, func(
		fieldSchema *schema.Schema, path []string, valueField *reflect.Value) error {
		fieldValue := valueField.Interface()
		if fieldValue == nil {
			return nil
		}
		fieldPath := strings.Join(path, ".")
		if fieldSchema.Optional {
			switch valueField.Kind() {
			case reflect.Ptr:
				if valueField.IsNil() {
					log.Printf("[DEBUG] skipping empty %s %#v", fieldPath, valueField)
					return nil
				}
			case reflect.Array, reflect.Map, reflect.String, reflect.Slice:
				if valueField.Len() == 0 {
					log.Printf("[DEBUG] skipping empty %s %#v", fieldPath, valueField)
					return nil
				}
			}
		}
		_, configured := d.GetOk(fieldPath)
		if !d.IsNewResource() && !fieldSchema.Computed && !configured {
			log.Printf("[DEBUG] Removing default fields sent back by server: %s - %#v",
				fieldPath, fieldValue)
			return nil
		}
		switch fieldSchema.Type {
		case schema.TypeList, schema.TypeSet:
			es, ok := fieldSchema.Elem.(*schema.Schema)
			if ok {
				switch es.Type {
				case schema.TypeString:
					v, ok := fieldValue.([]string)
					if !ok {
						return fmt.Errorf("%s[%v] is not a string",
							fieldPath, fieldValue)
					}
					return d.Set(fieldPath, v)
				case schema.TypeInt:
					v, ok := fieldValue.([]int)
					if !ok {
						return fmt.Errorf("%s[%v] is not a string",
							fieldPath, fieldValue)
					}
					return d.Set(fieldPath, v)
				}
				return fmt.Errorf("%s[%v] supported schema detected",
					fieldPath, fieldValue)
			}
			nv, err := collectionToMaps(fieldValue, fieldSchema)
			if err != nil {
				return fmt.Errorf("%s: %v", fieldPath, err)
			}
			if len(nv) == 0 {
				return nil
			}
			log.Printf("[DEBUG] set %s %#v", fieldPath, nv)
			return d.Set(fieldPath, nv)
		default:
			log.Printf("[DEBUG] set %s %#v", fieldPath, fieldValue)
			return d.Set(fieldPath, fieldValue)
		}
	})
}

// DataToStructPointer reads resource data with given schema onto result pointer
func DataToStructPointer(d *schema.ResourceData, scm map[string]*schema.Schema, result interface{}) error {
	rv := reflect.ValueOf(result)
	rk := rv.Kind()
	if rk != reflect.Ptr {
		return fmt.Errorf("Pointer is expected, but got %s: %#v", reflectKind(rk), result)
	}
	rv = rv.Elem()
	return readReflectValueFromData([]string{}, d, rv, scm)
}

// DataToReflectValue reads reflect value from data
func DataToReflectValue(d *schema.ResourceData, r *schema.Resource, rv reflect.Value) error {
	return readReflectValueFromData([]string{}, d, rv, r.Schema)
}

func readReflectValueFromData(path []string, d *schema.ResourceData,
	rv reflect.Value, s map[string]*schema.Schema) error {
	return iterFields(rv, path, s, func(fieldSchema *schema.Schema,
		path []string, valueField *reflect.Value) error {
		fieldPath := strings.Join(path, ".")
		raw, ok := d.GetOk(fieldPath)
		if !ok {
			return nil
		}
		switch fieldSchema.Type {
		case schema.TypeInt:
			v, ok := raw.(int)
			if !ok {
				return fmt.Errorf("%s is not int", fieldPath)
			}
			valueField.SetInt(int64(v))
		case schema.TypeString:
			v, ok := raw.(string)
			if !ok {
				return fmt.Errorf("%s is not string", fieldPath)
			}
			valueField.SetString(v)
		case schema.TypeBool:
			v, ok := raw.(bool)
			if !ok {
				return fmt.Errorf("%s is not bool", fieldPath)
			}
			valueField.SetBool(v)
		case schema.TypeFloat:
			v, ok := raw.(float64)
			if !ok {
				return fmt.Errorf("%s is not float", fieldPath)
			}
			valueField.SetFloat(v)
		case schema.TypeMap:
			mapValueKind := valueField.Type().Elem().Kind()
			valueField.Set(reflect.MakeMap(valueField.Type()))
			for key, ivalue := range raw.(map[string]interface{}) {
				kv := reflect.ValueOf(key)
				switch mapValueKind {
				case reflect.String:
					valueField.SetMapIndex(
						kv, reflect.ValueOf(fmt.Sprintf("%v", ivalue)))
				case reflect.Float32:
					v, ok := ivalue.(float32)
					if !ok {
						return fmt.Errorf("%s[%s] '%v' is not float32",
							fieldPath, key, ivalue)
					}
					valueField.SetMapIndex(kv, reflect.ValueOf(v))
				case reflect.Float64:
					v, ok := ivalue.(float64)
					if !ok {
						return fmt.Errorf("%s[%s] '%v' is not float64",
							fieldPath, key, ivalue)
					}
					valueField.SetMapIndex(kv, reflect.ValueOf(v))
				case reflect.Int:
					v, ok := ivalue.(int)
					if !ok {
						return fmt.Errorf("%s[%s] '%v' is not int",
							fieldPath, key, ivalue)
					}
					valueField.SetMapIndex(kv, reflect.ValueOf(v))
				case reflect.Bool:
					v, ok := ivalue.(bool)
					if !ok {
						return fmt.Errorf("%s[%s] '%v' is not bool",
							fieldPath, key, ivalue)
					}
					valueField.SetMapIndex(kv, reflect.ValueOf(v))
				default:
					return fmt.Errorf("%s[%s] '%v' is not valid primitive",
						fieldPath, key, ivalue)
				}
			}
		case schema.TypeSet:
			rawSet, ok := raw.(*schema.Set)
			if !ok {
				return fmt.Errorf("%s[%v] is not set", fieldPath, raw)
			}
			rawList := rawSet.List()
			return readListFromData(path, d, rawList, valueField,
				fieldSchema, func(i int) string {
					return strconv.Itoa(rawSet.F(rawList[i]))
				})
		case schema.TypeList:
			rawList, ok := raw.([]interface{})
			if !ok {
				return fmt.Errorf("%s[%v] is not list", fieldPath, raw)
			}
			return readListFromData(path, d, rawList, valueField, fieldSchema, strconv.Itoa)
		default:
			return fmt.Errorf("%s[%v] unsupported field type", fieldPath, raw)
		}
		return nil
	})
}

func readListFromData(path []string, d *schema.ResourceData,
	rawList []interface{}, valueField *reflect.Value, fieldSchema *schema.Schema,
	offsetConverter func(i int) string) error {
	if len(rawList) == 0 {
		return nil
	}
	fieldPath := strings.Join(path, ".")
	switch valueField.Type().Kind() {
	case reflect.Ptr:
		vpointer := reflect.New(valueField.Type().Elem())
		valueField.Set(vpointer)
		ve := vpointer.Elem()
		nestedResource, ok := fieldSchema.Elem.(*schema.Resource)
		if !ok {
			return fmt.Errorf("%s[%v] is not a resource", fieldPath, rawList[0])
		}
		nestedPath := append(path, offsetConverter(0))
		return readReflectValueFromData(nestedPath, d, ve, nestedResource.Schema)
	case reflect.Slice:
		k := valueField.Type().Elem().Kind()
		newSlice := reflect.MakeSlice(valueField.Type(), len(rawList), len(rawList))
		valueField.Set(newSlice)
		for i, elem := range rawList {
			switch k {
			case reflect.String:
				v, ok := elem.(string)
				if !ok {
					return fmt.Errorf("%s[%v] is not a string", fieldPath, elem)
				}
				newSlice.Index(i).SetString(v)
			case reflect.Struct:
				nestedResource, ok := fieldSchema.Elem.(*schema.Resource)
				if !ok {
					return fmt.Errorf("%s[%v] is not a resource", fieldPath, elem)
				}
				nestedPath := append(path, offsetConverter(i))
				vpointer := reflect.New(valueField.Type().Elem())
				ve := vpointer.Elem()
				err := readReflectValueFromData(nestedPath, d, ve, nestedResource.Schema)
				if err != nil {
					return err
				}
				newSlice.Index(i).Set(ve)
			default:
				return fmt.Errorf("%s[%v] is not valid slice elem", fieldPath, elem)
			}
		}
	default:
		return fmt.Errorf("%s[%v] unknown collection field", fieldPath, rawList)
	}
	return nil
}
