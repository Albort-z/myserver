package provider

import (
	"context"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/Albort-z/myserver/core/exitem"
	"github.com/gin-gonic/gin"
)

type ServerConfig struct {
	ListenHost string `yaml:"server"`
	ListenPort int    `yaml:"port"`
}

type Server interface {
	Load(configFile string)
	Start()
	Stop()
	Status()
	Reload(configFile string)
}

func NewServer(t string) Server {
	var s Server
	if t == "gin" {
		s = &GinServer{}
		s.Load("conf/web.yaml")
		return s
	}
	return nil
}

type GoServer struct {
	app    *http.Server
	config string
	status string
}

func (w *GoServer) Load(configFile string) {

}

func (w *GoServer) Start() {
	w.status = "starting"
	w.app = &http.Server{
		Addr:    ":8000",
		Handler: http.DefaultServeMux,
	}
	//exitem.HandleTree(&w.app.RouterGroup)
	w.app.ListenAndServe()
	//w.app.Run("0.0.0.0:" + strconv.Itoa(8080))
	w.status = "started"
}

func (w *GoServer) Stop() {
	w.status = "stopping"
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	w.app.Shutdown(ctx)
	w.status = "stopped"
}

func (w *GoServer) Status() {

}

func (w *GoServer) Reload(configFile string) {

}

// GinServer Gin框架作服务器
type GinServer struct {
	app    *gin.Engine
	config ServerConfig
	status string
}

func (w *GinServer) Load(configFile string) {
	conf, err := ioutil.ReadFile(configFile)
	if err != nil {
		//log.Errorf("读取配置文件出错。")
		panic(err)
	}
	err = yaml.Unmarshal(conf, &w.config)
	if err != nil {
		//log.Errorf("解析配置文件出错。")
		panic(err)
	}
}

func (w *GinServer) Start() {
	w.status = "starting"
	w.app = gin.Default()
	w.app.RedirectFixedPath = false
	gin.SetMode(gin.DebugMode)
	exitem.HandleTree(&w.app.RouterGroup)
	w.app.Run(w.config.ListenHost + ":" + strconv.Itoa(w.config.ListenPort))
	w.status = "started"
}

func (w *GinServer) Stop() {
	w.status = "stopping"
	w.status = "stopped"
}

func (w *GinServer) Status() {

}

func (w *GinServer) Reload(configFile string) {

}
