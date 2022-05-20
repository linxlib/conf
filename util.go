package conf

import (
	"os"
	"reflect"
	"strings"
	"time"
)

func stringSlice(s string) []string {
	s = strings.TrimSuffix(strings.TrimPrefix(s, "["), "]")
	return strings.Split(s, ",")
}
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func isStructPtr(i interface{}) bool {
	v := reflect.ValueOf(i)
	return v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.Slice, reflect.Array:
		return v.Len() == 0
	case reflect.Struct:
		if t, ok := v.Interface().(time.Time); ok {
			return t.IsZero()
		}
		return false
	case reflect.Invalid:
		return true
	default:
		return v.IsZero()
	}
}
