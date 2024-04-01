package main

import (
	"fmt"
	"reflect"
	"strings"
)

func tryCast(input interface{}, expected reflect.Type) interface{} {
	t := reflect.ValueOf(input)
	if !t.CanConvert(expected) {
		return nil
	}
	baseConvert := t.Convert(expected)
	switch expected.Kind() {
	case reflect.Float32:
		return float32(baseConvert.Float())
	case reflect.Float64:
		return float64(baseConvert.Float())
	case reflect.Int8:
		return int8(baseConvert.Int())
	case reflect.Int32:
		return int32(baseConvert.Int())
	case reflect.Int64:
		return int64(baseConvert.Int())
	}
	return nil
}

func parseField(keyName string, value interface{}, out interface{}) error {
	structValue := reflect.ValueOf(out).Elem()
	structFieldValue := structValue.FieldByName(keyName)

	if !structFieldValue.IsValid() {
		structFieldValue = structValue.FieldByName(strings.Title(keyName))
		if !structFieldValue.IsValid() {
			return fmt.Errorf("No such field %s in out obj.", keyName)
		}
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set value %s, maybe use a pointer &out", keyName)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		val = reflect.ValueOf(tryCast(value, structFieldType))
		if structFieldType != val.Type() {
			return fmt.Errorf("Types don't match. Kind: %s, Type: %s, structField: %s", val.Kind(), val.Type(), structFieldType)
		}
	}

	structFieldValue.Set(val)

	return nil
}

func FromMap(input map[string]interface{}, output interface{}) error {
	for k, v := range input {
		err := parseField(k, v, output)
		if err != nil {
			return err
		}
	}

	return nil
}
