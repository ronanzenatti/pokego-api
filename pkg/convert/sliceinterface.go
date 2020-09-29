package convert

import (
	"reflect"
)

// SliceInterface - Convert slice of any type to slice of interfaces.
func SliceInterface(input interface{}) []interface{} {
	if reflect.TypeOf(input).Kind() != reflect.Slice {
		panic("Pkg: convert.SliceInterface - Expected value type: slice | Received value type: " + reflect.TypeOf(input).String())
	}

	slice := reflect.Indirect(reflect.ValueOf(input))

	var output []interface{} = make([]interface{}, slice.Len())
	for i := 0; i < slice.Len(); i++ {
		output[i] = slice.Index(i).Interface()
	}
	return output
}
