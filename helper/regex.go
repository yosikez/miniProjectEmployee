package helper

import (
	"miniProject/input"
	"reflect"
	"regexp"

	"github.com/go-playground/validator/v10"
)


func CheckDiveField(err validator.FieldError) (map[string]string, bool) {
	pattern := `^(\w+)\.(\w+)\[(\d+)\]\.(\w+)$`
	re := regexp.MustCompile(pattern)
	str := err.Namespace()

	isTrue := re.MatchString(str)

	if !isTrue {
		return nil, false
	}

	result := re.FindStringSubmatch(str)
	mapStr := make(map[string]string)

	fieldJson, _ :=  reflect.TypeOf(input.JsonDataOpportunity{}).FieldByName(result[2])
	field := GetJSONTagName(fieldJson)

	attributJson, _ := reflect.TypeOf(input.Resource{}).FieldByName(result[4])
	attrField := GetJSONTagName(attributJson)
	
	mapStr["field"] = field
	mapStr["index"] = result[3]
	mapStr["attribute"] = attrField

	return mapStr, true
}