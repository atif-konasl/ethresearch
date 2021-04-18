package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func analyzeStructValue() {
	type T struct {
		A int
		B string
	}
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	fmt.Printf("%d: %s = %v\n", s, s.Type(), s.Interface())
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		s.Field(0).SetInt(77)
		s.Field(1).SetString("Sunset Strip")

		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
		fmt.Println("settability of f: ", f.CanSet())
	}
}

// Value contains concrete value as well type information and Type contains only type information.
func analyzeTypeAndValue() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value: ", reflect.ValueOf(x).String())

	var z float64 = 3.4
	v := reflect.ValueOf(z)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	alt := reflect.ValueOf(&z)
	alt.Elem().SetFloat(7.0)
	fmt.Printf("z = %f alt elem value = %f " +
		"alt type = %v and alt elem type = %v \n",
		z, alt.Elem().Interface(), alt.Type(), alt.Elem().Type())

	type MyInt int
	var y MyInt = 7
	v = reflect.ValueOf(y)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is int: ", v.Kind() == reflect.Int)
	fmt.Println("interface: ", v.Interface())
}

func analyzeElemAndInterface() {
	var i int = 3
	var p *int = &i
	fmt.Println(p, i)

	v := reflect.ValueOf(p)
	fmt.Println(v.Interface()) // This is the p pointer

	v2 := v.Elem()
	fmt.Println(v2.Interface()) // This is i's value: 3
}

func analyzePointerInInterface() {
	var r io.Reader = os.Stdin // os.Stdin is of type *os.File which implements io.Reader

	v := reflect.ValueOf(r) // r is interface wrapping *os.File value
	fmt.Println(v.Type())   // *os.File

	v2 := reflect.ValueOf(&r)            // pointer passed, will be wrapped in interface{}
	fmt.Println(v2.Type())               // *io.Reader
	fmt.Println(v2.Elem().Type())        // navigate to pointed: io.Reader (interface type)
	fmt.Println(v2.Elem().Elem().Type()) // 2nd Elem(): get concrete value in interface: *os.File
}

func main() {
	analyzeTypeAndValue()
	println("=================")
	analyzeStructValue()
	println("=================")
	analyzeElemAndInterface()
	println("=================")
	analyzePointerInInterface()
}