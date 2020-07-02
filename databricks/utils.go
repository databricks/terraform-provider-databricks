package databricks

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
)

func convertListInterfaceToString(m []interface{}) []string {
	response := []string{}
	for _, v := range m {
		if v != nil {
			response = append(response, v.(string))
		}
	}
	return response
}

func getMapFromOneItemList(input interface{}) map[string]interface{} {
	inputList := input.([]interface{})
	if len(inputList) >= 1 {
		return inputList[0].(map[string]interface{})
	}
	return nil
}

func changeClusterIntoRunningState(clusterID string, client *service.DBApiClient) error {
	clusterInfo, err := client.Clusters().Get(clusterID)
	if err != nil {
		return err
	}
	//.. simplify it a bit
	switch clusterInfo.State {
	case model.ClusterStateRunning:
		time.Sleep(5 * time.Second)
		return nil
	case model.ClusterStatePending, model.ClusterStateResizing, model.ClusterStateRestarting:
		err := client.Clusters().WaitForClusterRunning(clusterID)
		if err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
		return nil
	case model.ClusterStateTerminating, model.ClusterStateTerminated:
		err := client.Clusters().WaitForClusterTerminated(clusterID)
		if err != nil {
			return err
		}
		err = client.Clusters().Start(clusterID)
		if err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
		return nil
	}
	return fmt.Errorf("cluster is in a non recoverable state: %s", clusterInfo.State)
}

// PackagedMWSIds is a struct that contains both the MWS acct id and the ResourceId (resources are networks, creds, etc.)
type PackagedMWSIds struct {
	MwsAcctID  string
	ResourceID string
}

// Helps package up MWSAccountId with another id such as credentials id or network id
// uses format mwsAcctId/otherId
func packMWSAccountID(idsToPackage PackagedMWSIds) string {
	return fmt.Sprintf("%s/%s", idsToPackage.MwsAcctID, idsToPackage.ResourceID)
}

// Helps unpackage MWSAccountId from another id such as credentials id or network id
func unpackMWSAccountID(combined string) (PackagedMWSIds, error) {
	var packagedMWSIds PackagedMWSIds
	parts := strings.Split(combined, "/")
	if len(parts) != 2 {
		return packagedMWSIds, fmt.Errorf("unpacked account has more than or less than two parts, combined id: %s", combined)
	}
	packagedMWSIds.MwsAcctID = parts[0]
	packagedMWSIds.ResourceID = parts[1]
	return packagedMWSIds, nil
}

// ValidateInstanceProfileARN is a ValidateFunc that ensures the role id is a valid aws iam instance profile arn
func ValidateInstanceProfileARN(val interface{}, key string) (warns []string, errs []error) {
	v := val.(string)

	if v == "" {
		return nil, []error{fmt.Errorf("%s is empty got: %s, must be an aws instance profile arn", key, v)}
	}

	// Parse and verify instance profiles
	instanceProfileArn, err := arn.Parse(v)
	if err != nil {
		return nil, []error{fmt.Errorf("%s is invalid got: %s received error: %w", key, v, err)}
	}
	// Verify instance profile resource type, Resource gets parsed as instance-profile/<profile-name>
	if !strings.HasPrefix(instanceProfileArn.Resource, "instance-profile") {
		return nil, []error{fmt.Errorf("%s must be an instance profile resource, got: %s in %s",
			key, instanceProfileArn.Resource, v)}
	}
	// TODO: later check if this ARN actually can launch stuff...
	return nil, nil
}

// func typeToSchema(t reflect.Type) map[string]*schema.Schema {
// 	scm := map[string]*schema.Schema{}
// 	for i := 0; i < t.NumField(); i++ {
// 		typeField := t.Field(i)
// 		jsonTag := typeField.Tag.Get("json")
// 		tfTag := typeField.Tag.Get("tf")
// 		jsonFieldName := strings.Split(jsonTag, ",")[0]
// 		if jsonFieldName == "-" {
// 			continue
// 		}
// 		scm[jsonFieldName] = &schema.Schema{}
// 		if strings.Contains(jsonTag, "omitempty") {
// 			scm[jsonFieldName].Optional = true
// 		} else {
// 			scm[jsonFieldName].Required = true
// 		}
// 		switch typeField.Type.Kind() {
// 		case reflect.Int:
// 			scm[jsonFieldName].Type = schema.TypeInt
// 		case reflect.Float64:
// 			scm[jsonFieldName].Type = schema.TypeFloat
// 		case reflect.Bool:
// 			scm[jsonFieldName].Type = schema.TypeBool
// 		case reflect.String:
// 			scm[jsonFieldName].Type = schema.TypeString
// 		case reflect.Map:
// 			scm[jsonFieldName].Type = schema.TypeMap
// 		case reflect.Struct:
// 			scm[jsonFieldName].MaxItems = 1
// 			scm[jsonFieldName].Type = schema.TypeList
// 			scm[jsonFieldName].Elem = schema.Resource{
// 				Schema: typeToSchema(typeField.Type),
// 			}
// 		case reflect.Slice:
// 			ft := schema.TypeList
// 			if strings.Contains(tfTag, "slice_set") {
// 				ft = schema.TypeSet
// 			}
// 			scm[jsonFieldName].Type = ft
// 			elem := typeField.Type.Elem()
// 			switch elem.Kind() {
// 			case reflect.Int:
// 				scm[jsonFieldName].Elem = &schema.Schema{Type: schema.TypeInt}
// 			case reflect.Float64:
// 				scm[jsonFieldName].Elem = &schema.Schema{Type: schema.TypeFloat}
// 			case reflect.Bool:
// 				scm[jsonFieldName].Elem = &schema.Schema{Type: schema.TypeBool}
// 			case reflect.String:
// 				scm[jsonFieldName].Elem = &schema.Schema{Type: schema.TypeString}
// 			case reflect.Struct:
// 				scm[jsonFieldName].Elem = schema.Resource{
// 					Schema: typeToSchema(typeField.Type),
// 				}
// 			}
// 		}
// 	}
// 	return scm
// }

func iterFields(v interface{}, path []string, r *schema.Resource,
	cb func(fieldSchema *schema.Schema, path []string, valueField *reflect.Value) error) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if !rv.IsValid() {
		return nil
	}
	for i := 0; i < rv.NumField(); i++ {
		typeField := rv.Type().Field(i)
		jsonTag := typeField.Tag.Get("json")
		jsonFieldName := strings.Split(jsonTag, ",")[0]
		if jsonFieldName == "-" {
			continue
		}
		fieldSchema, ok := r.Schema[jsonFieldName]
		if !ok {
			continue
		}
		valueField := rv.Field(i)
		err := cb(fieldSchema, append(path, jsonFieldName), &valueField)
		if err != nil {
			return err
		}
	}
	return nil
}

func collectionToMaps(v interface{}, s *schema.Schema) ([]interface{}, error) {
	r, ok := s.Elem.(*schema.Resource)
	if !ok {
		return nil, fmt.Errorf("not resource")
	}
	var allItems []interface{}
	if s.MaxItems == 1 {
		allItems = append(allItems, v)
	} else {
		vs := reflect.ValueOf(v)
		for i := 0; i < vs.Len(); i++ {
			allItems = append(allItems, vs.Index(i).Interface())
		}
	}
	resultList := []interface{}{}
	for _, v := range allItems {
		data := map[string]interface{}{}
		err := iterFields(v, []string{}, r, func(fieldSchema *schema.Schema,
			path []string, valueField *reflect.Value) error {
			fieldName := path[len(path)-1]
			fieldValue := valueField.Interface()
			switch fieldSchema.Type {
			case schema.TypeList, schema.TypeSet:
				nv, err := collectionToMaps(fieldValue, fieldSchema)
				if err != nil {
					return err
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

func readResourceFromStruct(v interface{}, d *schema.ResourceData,
	path []string, r *schema.Resource) error {
	return iterFields(v, path, r, func(fieldSchema *schema.Schema,
		path []string, valueField *reflect.Value) error {
		fieldValue := valueField.Interface()
		if fieldValue == nil {
			return nil
		}
		fieldPath := strings.Join(path, ".")
		_, configured := d.GetOk(fieldPath)
		if !fieldSchema.Computed && !configured {
			log.Printf("[DEBUG] removing default fields sent back by server")
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
				return err
			}
			if len(nv) == 0 {
				return nil
			}
			return d.Set(fieldPath, nv)
		default:
			return d.Set(fieldPath, fieldValue)
		}
	})
}

func readStructFromData(path []string, d *schema.ResourceData,
	result interface{}, r *schema.Resource) error {
	return iterFields(result, path, r, func(fieldSchema *schema.Schema,
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
		izv := vpointer.Interface()
		nestedResource, ok := fieldSchema.Elem.(*schema.Resource)
		if !ok {
			return fmt.Errorf("%s[%v] is not a resource", fieldPath, rawList[0])
		}
		nestedPath := append(path, offsetConverter(0))
		return readStructFromData(nestedPath, d, izv, nestedResource)
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
				izv := vpointer.Interface()
				err := readStructFromData(nestedPath, d, izv, nestedResource)
				if err != nil {
					return err
				}
				newSlice.Index(i).Set(reflect.ValueOf(izv).Elem())
			default:
				return fmt.Errorf("%s[%v] is not valid slice elem", fieldPath, elem)
			}
		}
	default:
		return fmt.Errorf("%s[%v] unknown collection field", fieldPath, rawList)
	}
	return nil
}
