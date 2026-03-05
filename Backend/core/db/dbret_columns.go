package db

import (
	"fmt"
	"reflect"
)

func DBRetColumnsFromType[T any]() ([]string, error) {
	var zero T
	return DBRetColumns(zero)
}

func DBRetColumns(entity any) ([]string, error) {
	v := reflect.ValueOf(entity)

	if v.Kind() == reflect.Pointer {
		if v.IsNil() {
			t := reflect.TypeOf(entity)
			if t == nil || t.Kind() != reflect.Pointer || t.Elem().Kind() != reflect.Struct {
				return nil, fmt.Errorf("entity must be struct or pointer to struct")
			}
			v = reflect.New(t.Elem()).Elem()
		} else {
			v = v.Elem()
		}
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("entity must be struct or pointer to struct, got %s", v.Kind())
	}

	t := v.Type()
	cols := make([]string, 0, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		if sf.PkgPath != "" {
			continue
		}

		rc := sf.Tag.Get("dbret")
		if rc == "" || rc == "-" {
			continue
		}

		cols = append(cols, rc)
	}

	if len(cols) == 0 {
		return nil, fmt.Errorf("no dbret columns for %T", entity)
	}

	return cols, nil
}
