package filc

import "reflect"

// Look - look for struct fields that are nil
func Look(value interface{}) []string {
	return look(value, "")
}

func look(value interface{}, parent string) []string {
	nilFields := make([]string, 0)

	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Struct:
		// ok
	case reflect.Ptr:
		if rv.IsNil() {
			return append(nilFields, parent)
		}
	default:
		return nilFields
	}

	rv = reflect.Indirect(rv)
	rt := rv.Type()
	structName := rt.Name()
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		fieldName := f.Name

		field := rv.FieldByName(fieldName)
		if !field.IsValid() {
			continue
		}

		key := structName + "." + fieldName
		if parent != "" {
			key = parent + "." + key
		}

		switch field.Kind() {
		case reflect.Ptr, reflect.Interface:
			if field.IsNil() {
				nilFields = append(nilFields, key)
				continue
			}
		}

		if f.Anonymous {
			nilFields = append(nilFields, look(field.Interface(), structName)...)
		}

		continue
	}

	return nilFields
}
