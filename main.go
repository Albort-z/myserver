package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "./exitems"
	"github.com/Albort-z/myserver/core/provider"
)

var (
	VERSION   = "2000.01.01.release"
	BUILDTIME = "2000-01-01T00:00:00+0800"
)

var (
	help    bool
	version bool
	cmd     string
	c       string
)

func usage() {
	fmt.Fprintf(os.Stderr, `%s version: %s
Usage: %s [version] [-c ./config.yaml|http://localhost:8080/config.yaml] [-s start|stop|status]

Options:
`, VERSION)
	flag.PrintDefaults()

}

func init() {
	flag.BoolVar(&help, "h", false, "show this help")
	flag.BoolVar(&help, "help", false, "show this help")
	flag.BoolVar(&version, "v", false, "show version and exit")
	flag.BoolVar(&version, "version", false, "show version and exit")
	flag.StringVar(&cmd, "s", "", "send a cmd to a master process: stop, quit, reopen, reload")
	flag.StringVar(&c, "c", "conf/nginx.conf", "set configuration `file`")

	flag.Usage = usage
}

func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	flag.Parse()

	se := provider.NewServer("gin")
	go se.Start()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
	fmt.Println("The main thread has started.")

	for s := range c {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			fmt.Println("退出", s)
			se.Stop()
			exit()
		case syscall.SIGUSR1:
			fmt.Println("usr1", s)
		case syscall.SIGUSR2:
			fmt.Println("热重启", s)
		default:
			fmt.Println("other", s)
		}
	}
}

func exit() {
	fmt.Println("开始退出...")
	fmt.Println("执行清理...")
	fmt.Println("结束退出...")
	os.Exit(0)
}
