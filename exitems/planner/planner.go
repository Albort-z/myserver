// Package planner 此包为直播计划服务。
package planner

import (
	"fmt"
	"github.com/Albort-z/myserver/core/exitem"
	"github.com/gin-gonic/gin"
)

var actions = map[string]*exitem.Action{
	"": {listPlan, ListPlanInput{}, new(string)},
}

func init() {
	var exItem exitem.ExItem
	exItem.Name = "Planner"
	exItem.Path = "Planner"
	exItem.Methods = "GET"
	exItem.Handlers = append([]gin.HandlerFunc{}, exitem.ActionsHandler(actions))
	exitem.RegisterExItem(&exItem)
	//log.Infof("Init Hello.")
}

type ListPlanInput struct {
	From string
}

func listPlan(ctx *gin.Context, input1 interface{}) interface{} {
	input := input1.(*ListPlanInput)
	fmt.Printf("####fuck! %+v\n", input.From)
	return "Fuck"
}
