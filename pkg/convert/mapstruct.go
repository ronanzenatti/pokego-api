package convert

import "github.com/mitchellh/mapstructure"

// MapStruct - Convert maps to structs and structs to maps.
func MapStruct(input, output interface{}) {
	mapstructure.Decode(input, &output)
}
