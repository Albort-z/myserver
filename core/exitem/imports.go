package exitem

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"runtime/debug"
)

// 注册条目暂存
var RegExItem *ExItem

// 所有注册的条目
var ExItems []*ExItem

//////////////////////// global type ///////////////////////////////////

type ExItem struct {
	Name     string
	Path     string
	Methods  string
	Handlers []gin.HandlerFunc
}

//////////////////////////// global function ///////////////////////////////

func RegisterExItem(ei *ExItem) error {
	if ei != nil {
		RegExItem = ei
		ExItems = append(ExItems, ei)
		return nil
	} else {
		return errors.New("注册失败！")
	}
}

type Action struct {
	Function func(*gin.Context, interface{}) interface{}
	Input    interface{}
	Output   interface{}
}

type ApiError struct {
	Code    string `json:",omitempty"`
	Message string `json:",omitempty"`
}

func ActionsHandler(actions map[string]*Action) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				debug.PrintStack()
				//ctx.AbortWithStatusJSON(500, ApiError{"InternalError", "Server encountered an internal error: " + err.(error).Error()})
			}
		}()
		ctx.Header("Cache-Control", "no-cache, must-revalidate")
		ctx.Header("Pragma", "no-cache")
		ctx.Header("Expires", "Sat, 26 Jul 1997 05:00:00 GMT")
		//handler.SetNoCache(ctx)
		//formValues := ctx.FormValues
		//headers := ctx.Request().Header

		//log.Debugf("parameters=%v, formValues=%v, headers=%v, json=%v", params, formValues, headers, json)

		var action string
		if vs := ctx.Query("Action"); vs != "" {
			action = vs
		}

		if h, ok := actions[action]; ok {
			input := reflect.New(reflect.TypeOf(h.Input)).Interface()
			if err := ctx.BindJSON(input); err != nil {
				panic(err)
			}
			output := h.Function(ctx, input)
			ctx.JSON(200, output)
		} else {
			ctx.AbortWithStatusJSON(400, ApiError{"InvalidAction", fmt.Sprintf("Invalid action '%s'", action)})
		}
	}
}
