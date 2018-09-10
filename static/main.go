package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type A struct {
	B int
	C string
}

func main() {
	var s = new(A)

	bean := reflect.New(reflect.TypeOf(s)).Interface()

	//a := reflect.TypeOf(S)
	//b := reflect.New(a)
	//c := reflect.New(reflect.TypeOf(S)).Elem().Interface()
	//d := reflect.New(reflect.TypeOf(S)).Elem()

	fmt.Printf("a: %s\n", reflect.TypeOf(bean))
	bs := "{\"B\":2,\"C\":\"Fuck\"}"
	err := json.Unmarshal([]byte(bs), bean)
	if err != nil {
		fmt.Errorf("err: %s", err.Error())
	}
	fmt.Printf("%+v", *bean.(**A))
}

func convert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = convert(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convert(v)
		}
	}
	return i
}
