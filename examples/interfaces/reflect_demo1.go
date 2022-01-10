package interfaces

import (
	"fmt"
	"reflect"
)

type S struct {
	age int8
	mark bool
}

func (s S) GetAge() int8 {
	return s.age + 1
}

func (s *S) IsMark() bool {
	return s.mark
}

func reflectDemo() {
	myMap := make(map[int32]string, 10)
	myMap[1] = "abc"
	myMap[2] = "def"
	myMap[3] = "ghi"

	t := reflect.TypeOf(myMap)
	fmt.Println("type: ", t)

	v := reflect.ValueOf(myMap)
	fmt.Println("value: ", v)

	fmt.Println("==============")
	s := S{age: 20}
	sv := reflect.ValueOf(s)
	for i := 0; i < sv.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, sv.Field(i))
	}

	for i := 0; i < sv.NumMethod(); i++ {
		fmt.Printf("Method %d: %v\n", i, sv.Method(i))
	}

	fmt.Println("S valueOf:", sv, sv.NumMethod())

	st := reflect.TypeOf(s)
	fmt.Println("S typeof:", st, st.NumMethod())

	res := sv.Method(0).Call(nil)
	fmt.Println("result: ", res, "typeOf:", reflect.TypeOf(res), res[0].Int())

}
