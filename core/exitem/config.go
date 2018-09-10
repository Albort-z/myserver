package exitem

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const (
	Handle    = "Handle"
	StaticWeb = "StaticWeb" // 静态资源
)

type Leaf struct {
	Name   string
	Url    string
	Path   string
	Handle string
	Type   string
}

type Worm struct {
	Name string
	Path string
	Type string
}

type Branch struct {
	Name     string
	Url      string
	Wormes   []Worm
	Branches []Branch
	Leaves   []Leaf
	Types    string
}

func ReadConfig(path string) *Branch {
	conf, err := ioutil.ReadFile(path)
	if err != nil {
		//log.Errorf("读取配置文件出错。")
		panic(err)
	}
	var c Branch
	err = yaml.Unmarshal(conf, &c)
	if err != nil {
		//log.Errorf("解析配置文件出错。")
		panic(err)
	}
	return &c
}
