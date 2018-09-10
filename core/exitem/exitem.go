package exitem

import (
	"errors"
	"github.com/gin-gonic/gin"
	"plugin"
	"strings"
)

const configPath = "conf/config.yaml"

func loadPlugin(eiPath string) (*ExItem, error) {
	// 保证不会同路经多次导入(避免模块init方法不执行导致的注册失败)
	for _, ei := range ExItems {
		if ei.Path == eiPath {
			return ei, nil
		}
	}
	if _, err := plugin.Open(eiPath); err != nil {
		return nil, err
	}
	if RegExItem != nil {
		defer func() { RegExItem = nil }()
		RegExItem.Path = eiPath
		return RegExItem, nil
	}
	return nil, errors.New("WebPlugin导入失败！")
}

func loadPackage(eiPath string) (*ExItem, error) {
	for _, ei := range ExItems {
		if ei.Path == eiPath {
			return ei, nil
		}
	}
	return nil, errors.New("WebPackage未导入！请先import要导入的包(" + eiPath + ")。")
}

func HandleTree(api *gin.RouterGroup) {
	root := ReadConfig(configPath)
	var err error
	if root.Url != "" {
		err = subTree(api.Group(root.Url), root)
	} else {
		err = subTree(api, root)
	}
	if err != nil {
		panic(err)
	}
}

// 整个分支是插件
var isPlugins bool

func subTree(api *gin.RouterGroup, node *Branch) error {
	isPlugins = node.Types != "Package"
	// 单个拓展插件
	var isPlugin bool
	for _, worm := range node.Wormes {
		//log.Infof("树枝(%s)发现一只虫子(%s)", node.Name, worm.Name)
		if worm.Type == "" {
			isPlugin = isPlugins
		} else {
			isPlugin = worm.Type != "Package"
		}
		var ei *ExItem
		var err error
		if isPlugin {
			ei, err = loadPlugin(worm.Path)
		} else {
			ei, err = loadPackage(worm.Path)
		}
		if err != nil {
			return err
		} else {
			api.Use(ei.Handlers...)
		}
	}
	for _, leaf := range node.Leaves {
		//log.Infof("树枝(%s)长了一片叶子(%s)", node.Name, leaf.Name)
		if leaf.Type == "" {
			isPlugin = isPlugins
		} else {
			isPlugin = leaf.Type != "Package"
		}

		if leaf.Handle == "" || leaf.Handle == Handle {
			var ei *ExItem
			var err error
			if isPlugin {
				ei, err = loadPlugin(leaf.Path)
			} else {
				ei, err = loadPackage(leaf.Path)
			}
			if err != nil {
				return err
			} else {
				api.Handle(ei.Methods, leaf.Url, ei.Handlers...)
			}
		} else if leaf.Handle == StaticWeb {
			if strings.HasSuffix(leaf.Url, "/") {
				api.Static(leaf.Url, leaf.Path)
			} else {
				api.StaticFile(leaf.Url, leaf.Path)
			}
		}
	}
	for _, branch := range node.Branches {
		//log.Infof("树枝(%s)分出了一支树枝(%s)", node.Name, branch.Name)
		_api := api.Group(branch.Url)
		branch.Types = node.Types
		if err := subTree(_api, &branch); err != nil {
			return err
		}
	}
	return nil
}
