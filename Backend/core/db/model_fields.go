package db

import (
	"fmt"
	"reflect"
)

func ExtractDBColumnsAndArgs(entity any) ([]string, []any, error) {
	v := reflect.ValueOf(entity)
	if v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return nil, nil, fmt.Errorf("nil pointer entity")
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, nil, fmt.Errorf("entity must be struct, got %s", v.Kind())
	}

	t := v.Type()
	cols := make([]string, 0, t.NumField())
	args := make([]any, 0, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)

		if sf.PkgPath != "" {
			continue
		}

		tag := sf.Tag.Get("db")
		if tag == "" || tag == "-" {
			continue
		}

		fv := v.Field(i)

		if fv.Kind() == reflect.Pointer {
			cols = append(cols, tag)
			if fv.IsNil() {
				args = append(args, nil)
			} else {
				args = append(args, fv.Elem().Interface())
			}
			continue
		}

		cols = append(cols, tag)
		args = append(args, fv.Interface())
	}

	if len(cols) == 0 {
		return nil, nil, fmt.Errorf("no db columns found for %T", entity)
	}

	return cols, args, nil
}
