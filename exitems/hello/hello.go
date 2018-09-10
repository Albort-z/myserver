package hello

import (
	"encoding/json"
	"github.com/Albort-z/myserver/core/exitem"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

var actions = map[string]*exitem.Action{
	"hello":     {sayHello, sayHelloInput{}, new(string)},
	"Json2Yaml": {json2Yaml, json2YamlInput{}, new(string)},
	"Yaml2Json": {yaml2Json, yaml2JsonInput{}, new(string)},
}

func init() {
	var exItem exitem.ExItem
	exItem.Name = "HelloWorld"
	exItem.Path = "Hello"
	exItem.Methods = "POST"
	exItem.Handlers = append([]gin.HandlerFunc{}, exitem.ActionsHandler(actions))
	exitem.RegisterExItem(&exItem)
	//log.Infof("Init Hello.")
}

type sayHelloInput struct {
	Msg string `json:"msg"`
}

type json2YamlInput struct {
	Json string
}
type yaml2JsonInput struct {
	HHH  string
	Yaml string
}

func sayHello(ctx *gin.Context, input1 interface{}) interface{} {
	input := input1.(*sayHelloInput)
	return input.Msg
}

func yaml2Json(ctx *gin.Context, input1 interface{}) interface{} {
	input := input1.(*yaml2JsonInput)

	var body interface{}
	if err := yaml.Unmarshal([]byte(input.Yaml), &body); err != nil {
		panic(err)
	}
	body = convert(body)
	b, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	return string(b)
}

//json2Yaml json转yaml
func json2Yaml(ctx *gin.Context, input1 interface{}) interface{} {
	input := input1.(*json2YamlInput)

	var body interface{}
	if err := json.Unmarshal([]byte(input.Json), &body); err != nil {
		panic(err)
	}

	//body = convert(body)
	b, err := yaml.Marshal(body)
	if err != nil {
		panic(err)
	}
	return string(b)
}

// convert yaml2Json need convert.
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
