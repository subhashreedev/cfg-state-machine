package main

import (
	"fmt"

	"github.com/shubhashreeb/cfg-state-machine/mgr"
	"github.com/shubhashreeb/cfg-state-machine/pkg/routes"
)

func main() {
	fmt.Println("Starting the application...")
	engine := mgr.NewEngine(routes.NewRoute())
	engine.Start()

}
