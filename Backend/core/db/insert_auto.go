package db

import (
	"fmt"
	"reflect"
	"strings"
)

func BuildInsertAutoReturning(table string, model any) (query string, args []any, err error) {
	if table == "" {
		return "", nil, fmt.Errorf("table is empty")
	}

	v := reflect.ValueOf(model)
	if v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return "", nil, fmt.Errorf("nil model pointer")
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return "", nil, fmt.Errorf("model must be struct, got %s", v.Kind())
	}

	t := v.Type()

	cols := make([]string, 0, t.NumField())
	args = make([]any, 0, t.NumField())
	ret := make([]string, 0, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		if sf.PkgPath != "" {
			continue
		}

		if rc := sf.Tag.Get("dbret"); rc != "" && rc != "-" {
			ret = append(ret, rc)
		}

		col := sf.Tag.Get("db")
		if col == "" || col == "-" {
			continue
		}

		cols = append(cols, col)

		fv := v.Field(i)
		if fv.Kind() == reflect.Pointer {
			if fv.IsNil() {
				args = append(args, nil)
			} else {
				args = append(args, fv.Elem().Interface())
			}
		} else {
			args = append(args, fv.Interface())
		}
	}

	if len(cols) == 0 {
		return "", nil, fmt.Errorf("no insertable columns for %T", model)
	}

	ph := make([]string, len(cols))
	for i := range cols {
		ph[i] = fmt.Sprintf("$%d", i+1)
	}

	q := fmt.Sprintf(
		`INSERT INTO %s (%s) VALUES (%s)`,
		table,
		strings.Join(cols, ", "),
		strings.Join(ph, ", "),
	)

	if len(ret) > 0 {
		q += "\nRETURNING " + strings.Join(ret, ", ")
	}

	return q, args, nil
}

func ScanTargetsByDBRet(destPtr any) ([]any, error) {
	v := reflect.ValueOf(destPtr)
	if v.Kind() != reflect.Pointer || v.IsNil() {
		return nil, fmt.Errorf("destPtr must be non-nil pointer to struct")
	}
	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("destPtr must point to struct, got %s", v.Kind())
	}

	t := v.Type()
	targets := make([]any, 0, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		if sf.PkgPath != "" {
			continue
		}
		rc := sf.Tag.Get("dbret")
		if rc == "" || rc == "-" {
			continue
		}

		fv := v.Field(i)
		if !fv.CanAddr() {
			return nil, fmt.Errorf("field %s not addressable", sf.Name)
		}
		targets = append(targets, fv.Addr().Interface())
	}

	if len(targets) == 0 {
		return nil, fmt.Errorf("no dbret targets for %T", destPtr)
	}

	return targets, nil
}
