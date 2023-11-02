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
	res, ok := mapCache[structType]
	mutex.Unlock()
	if ok {
		return res
	}
	res = map[string]string{}
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		fieldName := chooseFieldName(field)
		if fieldName != "-" {
			res[fieldName] = field.Name
		}
	}
	mutex.Lock()
	mapCache[structType] = res
	mutex.Unlock()
	return res
}
