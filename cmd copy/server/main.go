package main

import (
	"flag"

	"cloudVideo/core"
)

var confPath = flag.String("conf", "./config.json", "path/to/your/config.json")

func main() {
	flag.Parse()
	c, err := core.NewCore(*confPath)
	if err != nil {
		panic(err)
	}
	c.Run()
}
