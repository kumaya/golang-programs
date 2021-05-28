package main

import (
	"fmt"
	"reflect"
)

type KeyValue struct {
	Key   interface{}
	Value interface{}
}

func stringer(key interface{}) string {
	kind := reflect.TypeOf(key).Kind()
	switch kind {
	case reflect.String:
		return key.(string)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", key)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", key)
	case reflect.Struct:
		return fmt.Sprintf("%v", key)
	default:
		panic("Unsupported key type")
	}
}
