package main

import (
	"fmt"
	"reflect"
	"unicode"
)

type Point struct {
	X       int
	Y       int
	private string
}

func main() {
	fmt.Println(fields(Point{})) // [X Y]
}

func fields(t interface{}) []string {
	var exportedFields []string

	structFields := reflect.ValueOf(t)

	for i := 0; i < structFields.Type().NumField(); i++ {
        structField := structFields.Type().Field(i)
		structFieldName := structField.Name

		if len(structFieldName) > 0 && unicode.IsUpper([]rune(structFieldName)[0]) {
			exportedFields = append(exportedFields, structField.Name)
		}

    }

	return exportedFields
}