package service

import (
	"reflect"
	"strconv"
	"strings"
)

func Mask(obj interface{}) interface{} {
	// Wrap the original in a reflect.Value
	original := reflect.ValueOf(obj)

	copy := reflect.New(original.Type()).Elem()
	maskRecursive(copy, original, false)

	// Remove the reflection wrapper
	return copy.Interface()
}

func maskRecursive(copy, original reflect.Value, mask bool) {
	switch original.Kind() {
	// The first cases handle nested structures and Mask them recursively

	// If it is a pointer we need to unwrap and call once again
	case reflect.Ptr:
		// To get the actual value of the original we have to call Elem()
		// At the same time this unwraps the pointer so we don't end up in
		// an infinite recursion
		if !original.IsZero() {
			originalValue := original.Elem()
			// Check if the pointer is nil
			if !originalValue.IsValid() {
				return
			}
			// Allocate a new object and set the pointer to it
			copy.Set(reflect.New(originalValue.Type()))
			// Unwrap the newly created pointer
			maskRecursive(copy.Elem(), originalValue, false)
		}

	// If it is an interface (which is very similar to a pointer), do basically the
	// same as for the pointer. Though a pointer is not the same as an interface so
	// note that we have to call Elem() after creating a new object because otherwise
	// we would end up with an actual pointer
	case reflect.Interface:
		// Get rid of the wrapping interface
		if !original.IsZero(){
			originalValue := original.Elem()
			// Create a new object. Now new gives us a pointer, but we want the value it
			// points to, so we have to call Elem() to unwrap it
			copyValue := reflect.New(originalValue.Type()).Elem()
			maskRecursive(copyValue, originalValue, false)
			copy.Set(copyValue)
		}

	// If it is a struct we Mask each field
	case reflect.Struct:
		for i := 0; i < original.NumField(); i += 1 {
			//log.Println()
			maskValue, maskInStruct := original.Type().Field(i).Tag.Lookup("mask")
			maskIsTrue, _  := strconv.ParseBool(maskValue)
			maskRecursive(copy.Field(i), original.Field(i), maskInStruct && maskIsTrue)
		}

	// If it is a slice we create a new slice and Mask each element
	case reflect.Slice:
		copy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
		for i := 0; i < original.Len(); i += 1 {
			maskRecursive(copy.Index(i), original.Index(i), false)
		}

	// If it is a map we create a new map and Mask each value
	case reflect.Map:
		copy.Set(reflect.MakeMap(original.Type()))
		for _, key := range original.MapKeys() {
			originalValue := original.MapIndex(key)
			// New gives us a pointer, but again we want the value
			copyValue := reflect.New(originalValue.Type()).Elem()
			maskRecursive(copyValue, originalValue, false)
			copy.SetMapIndex(key, copyValue)
		}

	// Otherwise we cannot traverse anywhere so this finishes the the recursion

	// If it is a string Mask it (yay finally we're doing what we came for)
	case reflect.String:
		switch mask {
		case true:
			copy.SetString("[REDACTED]")
		default:
			copy.Set(original)
		}
	// And everything else will simply be taken from the original
	default:
		copy.Set(original)
	}

}

type SecretsMask struct {
	Secrets []string
}

func (a SecretsMask) MaskString(str string) string {
	placeHolder := str
	for _, secret := range a.Secrets {
		placeHolder = strings.ReplaceAll(placeHolder, secret, "[REDACTED]")
	}
	return placeHolder
}