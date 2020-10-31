package handler

import (
	"pokego-api/pkg/convert"
)

type genContext interface {
	JSON(code int, i interface{}) error
	Param(name string) string
}

func mapStruct(input, output interface{}) {
	convert.MapStruct(input, &output)
}

func sliceInterface(input interface{}) []interface{} {
	return convert.SliceInterface(input)
}
