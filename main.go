package main

import (
	"min/mins"
)

func main() {
	startActivity()
	engine := mins.Rout()
	err := engine.Run(":13450")
	if err != nil {
		return
	}
}
