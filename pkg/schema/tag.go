package schema

import (
	"reflect"
)

const (
	TAG_NAME = "store"
	INDEX    = "index"
	SEARCH   = "search"
	RANGE    = "range"
)

//get inverse index
func GetTaggedFieldNames(s interface{}) []string {
	out := []string{}
	t := reflect.TypeOf(s)
	quene := []reflect.StructField{}
	for i := 0; i < t.NumField(); i++ {
		quene = append(quene, t.Field(i))
	}
	for len(quene) != 0 {
		var c reflect.StructField
		c, quene = quene[0], quene[1:]
		switch {
		case c.Type.Kind() == reflect.Struct:
			nested := reflect.TypeOf(c)
			for i := 0; i < nested.NumField(); i++ {
				quene = append(quene, nested.Field(i))
			}
		case c.Tag.Get(TAG_NAME) != "":
			out = append(out, c.Name)
		}
	}
	return out
}
