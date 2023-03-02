package helper

import (
	"reflect"
	"strings"
)

func GetJSONTagName(field reflect.StructField) string {
	jsonTag := field.Tag.Get("json")
	if jsonTag == "" {
		return strings.ToLower(field.Name)
	}
	jsonTagParts := strings.Split(jsonTag, ",")
	return jsonTagParts[0]
}