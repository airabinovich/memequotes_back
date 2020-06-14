package main

import (
	"github.com/airabinovich/memequotes_back/router"
)

func main() {
	engine := router.Route()
	if err := engine.Run(":9000"); err != nil {
		println("Backend service could not be started")
		panic(err)
	}
}
