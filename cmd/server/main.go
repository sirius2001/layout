package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"loon/core"
)

var confPath = flag.String("conf", "./config.json", "path/to/your/config.json")

func main() {
	flag.Parse()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	// 创建一个通道用于阻塞主协程
	done := make(chan bool, 1)
	c, err := core.NewCore(*confPath)
	if err != nil {
		panic(err)
	}
	c.Run()
	<-done
}
