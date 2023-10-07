package plugin

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func DecodeMapToStruct(m map[string]any, s any) error {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("a non-nil pointer is required")
	}

	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("a struct pointer is required")
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag.Get("field")
		if tag == "" {
			tag = field.Name
		}

		var value interface{}
		var ok bool
		for k, v := range m {
			if strings.EqualFold(k, tag) {
				value = v
				ok = true
				break
			}
		}

		if !ok {
			continue
		}

		fieldValue := v.Field(i)
		if fieldValue.Kind() == reflect.Struct {
			if subMap, ok := value.(map[string]interface{}); ok {
				err := DecodeMapToStruct(subMap, fieldValue.Addr().Interface())
				if err != nil {
					return err
				}
			}
		} else {
			jsonData, err := json.Marshal(value)
			if err != nil {
				return err
			}

			err = json.Unmarshal(jsonData, fieldValue.Addr().Interface())
			if err != nil {
				return err
			}
		}
	}

	return nil
}
