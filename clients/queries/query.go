package queries

import "reflect"

type FieldParser func(value reflect.Value) string

func getReflectedKeyValues(q any, fieldParser FieldParser) map[string]string {
	val := reflect.ValueOf(q)
	keyValues := make(map[string]string)
	numFields := val.NumField()
	for i := 0; i < numFields; i++ {
		typeField := val.Type().Field(i)
		key, exists := typeField.Tag.Lookup("key")
		if !exists {
			continue
		}

		stringValue := fieldParser(val.Field(i))
		if stringValue == "" {
			continue
		}
		keyValues[key] = stringValue
	}
	return keyValues
}
