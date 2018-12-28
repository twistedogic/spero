package tag

import (
	"reflect"
	"strings"
)

type Field struct {
	Tag   string
	Name  string
	Value reflect.Value
}

func GetTaggedFields(s interface{}, tagName string) []Field {
	fields := make([]Field, 0)
	ifValue := reflect.ValueOf(s)
	ifType := reflect.TypeOf(s)
	for i := 0; i < ifType.NumField(); i++ {
		t := ifType.Field(i)
		v := ifValue.Field(i)
		if t.Type.Kind() == reflect.Struct {
			fields = append(fields, GetTaggedFields(v.Interface(), tagName)...)
		}
		if tag, ok := t.Tag.Lookup(tagName); ok && v.IsValid() && v.CanInterface() && reflect.Zero(v.Type()).Interface() != v.Interface() {
			fields = append(fields, Field{tag, t.Name, v})
		}
	}
	return fields
}

func ParseTag(tags string) map[string]string {
	parsedTag := make(map[string]string)
	for _, tag := range strings.Split(tags, ",") {
		if strings.Contains(tag, "=") {
			tagValue := strings.TrimSpace(tag)
			kv := strings.Split(tagValue, "=")
			parsedTag[kv[0]] = kv[1]
		}
	}
	return parsedTag
}
