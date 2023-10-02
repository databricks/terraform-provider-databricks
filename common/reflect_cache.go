package common

import (
	"reflect"
	"sync"
)

var mutex sync.Mutex

var mapCache = map[reflect.Type]map[string]string{}

func getJsonToFieldNameMap(structType reflect.Type) map[string]string {
	if structType.Kind() != reflect.Ptr && structType.Kind() != reflect.Struct {
		return nil
	}
	mutex.Lock()
	defer mutex.Unlock()
	if res, ok := mapCache[structType]; ok {
		return res
	}
	res := map[string]string{}
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		fieldName := chooseFieldName(field)
		if fieldName != "-" {
			res[fieldName] = field.Name
		}
	}
	mapCache[structType] = res
	return res
}
