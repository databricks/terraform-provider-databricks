package databricks

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
)

func changeClusterIntoRunningState(clusterID string, client *service.DatabricksClient) error {
	//return nil
	clusterInfo, err := client.Clusters().Get(clusterID)
	if err != nil {
		return err
	}
	currentState := clusterInfo.State

	if model.ContainsClusterState([]model.ClusterState{model.ClusterStateRunning}, currentState) {
		return nil
	}

	if model.ContainsClusterState([]model.ClusterState{model.ClusterStatePending, model.ClusterStateResizing, model.ClusterStateRestarting}, currentState) {
		err := client.Clusters().WaitForClusterRunning(clusterID, 5, 180)
		if err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
		return nil
	}

	if model.ContainsClusterState([]model.ClusterState{model.ClusterStateTerminating}, currentState) {
		err := client.Clusters().WaitForClusterTerminated(clusterID, 5, 180)
		if err != nil {
			return err
		}
		err = client.Clusters().Start(clusterID)
		if err != nil {
			if !strings.Contains(err.Error(), fmt.Sprintf("Cluster %s is in unexpected state Pending.", clusterID)) {
				return err
			}
		}
		err = client.Clusters().WaitForClusterRunning(clusterID, 5, 180)
		if err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
		return nil
	}

	if model.ContainsClusterState([]model.ClusterState{model.ClusterStateTerminated}, currentState) {
		err = client.Clusters().Start(clusterID)
		if err != nil {
			if !strings.Contains(err.Error(), fmt.Sprintf("Cluster %s is in unexpected state Pending.", clusterID)) {
				return err
			}
		}

		err = client.Clusters().WaitForClusterRunning(clusterID, 5, 180)
		if err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
		return nil
	}

	return fmt.Errorf("cluster is in a non recoverable state: %s", currentState)
}

// PackagedMWSIds is a struct that contains both the MWS acct id and the ResourceId (resources are networks, creds, etc.)
type PackagedMWSIds struct {
	MwsAcctID  string
	ResourceID string
}

// Helps package up MWSAccountId with another id such as credentials id or network id
// uses format mwsAcctID/otherId
func packMWSAccountID(idsToPackage PackagedMWSIds) string {
	return fmt.Sprintf("%s/%s", idsToPackage.MwsAcctID, idsToPackage.ResourceID)
}

// Helps unpackage MWSAccountId from another id such as credentials id or network id
func unpackMWSAccountID(combined string) (PackagedMWSIds, error) {
	var packagedMWSIds PackagedMWSIds
	parts := strings.Split(combined, "/")
	if len(parts) != 2 {
		// TODO: set id to "" if invalid format
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
	return nil, nil
}

var PathEmptyError error = errors.New("provided path is empty")

// we would never want to handle root directories in regards to creating them
var DirPathRootDirError error = errors.New("dir path is root directory")

// Os libraries behave bizarely on windows as they will replace slashes with other values.
// This causes issues & errors when submitting the request
func GetParentDirPath(filePath string) (string, error) {
	if filePath == "" {
		return "", PathEmptyError
	}

	pathParts := strings.Split(filePath, "/")

	// if length of pathParts is just two items then the parent should be the root directory
	if len(pathParts) == 2 {
		return "", DirPathRootDirError
	}

	dirPath := strings.Join(pathParts[0:len(pathParts)-1], "/")

	return dirPath, nil
}

func reflectKind(k reflect.Kind) string {
	switch k {
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

func iterFields(rv reflect.Value, path []string, r *schema.Resource,
	cb func(fieldSchema *schema.Schema, path []string, valueField *reflect.Value) error) error {
	rk := rv.Kind()
	if rk != reflect.Struct {
		return fmt.Errorf("Value of Struct is expected, but got %s: %#v", reflectKind(rk), rv)
	}
	if !rv.IsValid() {
		return fmt.Errorf("Got invalid reflect value %#v", rv)
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

// func collectionToMaps(v interface{}, s *schema.Schema) ([]interface{}, error) {
// 	r, ok := s.Elem.(*schema.Resource)
// 	if !ok {
// 		return nil, fmt.Errorf("not resource")
// 	}
// 	var allItems []reflect.Value
// 	if s.MaxItems == 1 {
// 		allItems = append(allItems, reflect.ValueOf(v))
// 	} else {
// 		vs := reflect.ValueOf(v)
// 		for i := 0; i < vs.Len(); i++ {
// 			allItems = append(allItems, vs.Index(i))
// 		}
// 	}
// 	resultList := []interface{}{}
// 	for _, v := range allItems {
// 		data := map[string]interface{}{}
// 		err := iterFields(v, []string{}, r, func(fieldSchema *schema.Schema,
// 			path []string, valueField *reflect.Value) error {
// 			fieldName := path[len(path)-1]
// 			fieldValue := valueField.Interface()
// 			switch fieldSchema.Type {
// 			case schema.TypeList, schema.TypeSet:
// 				nv, err := collectionToMaps(fieldValue, fieldSchema)
// 				if err != nil {
// 					return err
// 				}
// 				data[fieldName] = nv
// 			default:
// 				if s, ok := fieldValue.(string); ok && s == "" {
// 					return nil
// 				}
// 				data[fieldName] = fieldValue
// 			}
// 			return nil
// 		})
// 		if err != nil {
// 			return nil, err
// 		}
// 		if len(data) == 0 {
// 			continue
// 		}
// 		resultList = append(resultList, data)
// 	}
// 	return resultList, nil
// }

// func readResourceFromStruct(v interface{}, d *schema.ResourceData,
// 	path []string, r *schema.Resource) error {
// 	return iterFields(reflect.Value(v), path, r, func(fieldSchema *schema.Schema,
// 		path []string, valueField *reflect.Value) error {
// 		fieldValue := valueField.Interface()
// 		if fieldValue == nil {
// 			return nil
// 		}
// 		fieldPath := strings.Join(path, ".")
// 		_, configured := d.GetOk(fieldPath)
// 		if !fieldSchema.Computed && !configured {
// 			log.Printf("[DEBUG] removing default fields sent back by server: %s", fieldPath)
// 			return nil
// 		}
// 		switch fieldSchema.Type {
// 		case schema.TypeList, schema.TypeSet:
// 			es, ok := fieldSchema.Elem.(*schema.Schema)
// 			if ok {
// 				switch es.Type {
// 				case schema.TypeString:
// 					v, ok := fieldValue.([]string)
// 					if !ok {
// 						return fmt.Errorf("%s[%v] is not a string",
// 							fieldPath, fieldValue)
// 					}
// 					return d.Set(fieldPath, v)
// 				case schema.TypeInt:
// 					v, ok := fieldValue.([]int)
// 					if !ok {
// 						return fmt.Errorf("%s[%v] is not a string",
// 							fieldPath, fieldValue)
// 					}
// 					return d.Set(fieldPath, v)
// 				}
// 				return fmt.Errorf("%s[%v] supported schema detected",
// 					fieldPath, fieldValue)
// 			}
// 			nv, err := collectionToMaps(fieldValue, fieldSchema)
// 			if err != nil {
// 				return err
// 			}
// 			if len(nv) == 0 {
// 				return nil
// 			}
// 			return d.Set(fieldPath, nv)
// 		default:
// 			return d.Set(fieldPath, fieldValue)
// 		}
// 	})
// }

func readStructFromData(path []string, d *schema.ResourceData,
	result interface{}, r *schema.Resource) error {
	return readReflectValueFromData(path, d, reflect.ValueOf(result), r)
}

func readReflectValueFromData(path []string, d *schema.ResourceData,
	rv reflect.Value, r *schema.Resource) error {
	return iterFields(rv, path, r, func(fieldSchema *schema.Schema,
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
