package helper

import (
	"fmt"
	"reflect"

	"gopkg.in/mgo.v2/bson"
)

func ToMap(in interface{}, tag string) (bson.M, error) {
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

func ToMap2(in interface{}, tag string) bson.M {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return nil
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
	return out
}

type Foo struct {
	ID     int    `m:"id"`
	MyName string `m:"my_name"`
	MyAge  string `m:"my_age"`
}

func main2() {
	f := Foo{ID: 24, MyName: "hello", MyAge: "world"}
	fmt.Printf("%[1]T, %+[1]v\n", f)

	g := ToMap2(f, "m")
	fmt.Printf("%v", g)
}
