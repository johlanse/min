package main

import (
	"flag"
	"fmt"

	"min/mins"
)

var port int

func init() {
	flag.IntVar(&port, "p", 8088, "listen a port")
}

func main() {
	startActivity()
	engine := mins.Rout()
	err := engine.Run(fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		return
	}
}
