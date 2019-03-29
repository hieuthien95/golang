package helper

import (
	"fmt"
	"reflect"
)

// ToMap converts a struct to a map using the struct's tags.
//
// ToMap uses tags on struct fields to decide which fields to add to the
// returned map.
func ToMap(in interface{}, tag string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts structs; got %T", v)
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		if tagv := fi.Tag.Get(tag); tagv != "" {
			// set key of map to value in struct field
			out[tagv] = v.Field(i).Interface()
		}
	}
	return out, nil
}

type Foo struct {
	ID     int    `m:"id"`
	MyName string `m:"my_name"`
	MyAge  string `m:"my_age"`
}

func main() {
	f := Foo{1, "hello", "world"}
	fmt.Printf("%[1]T, %+[1]v\n", f)

	g, err := ToMap(f, "m")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%[1]T, %+[1]v\n", g)
}
