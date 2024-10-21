package main

import (
	"flag"
	"loon/core"
	"loon/pkg/log"
	"os"
	"os/signal"
	"syscall"
)

var confPath = flag.String("conf", "./config.json", "path/to/your/config.json")

func main() {
	flag.Parse()
	sigs := make(chan os.Signal, 1)

	// 捕获信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// 创建核心实例
	c, err := core.NewCore(*confPath)
	if err != nil {
		panic(err)
	}
	c.Run()
	<-sigs
	log.Info("server shutdowning ....")
	c.Stop()
}
