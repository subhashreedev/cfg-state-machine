package mgr

import (
	"context"
	"fmt"
)

type Engine struct {
	Route Operator
}

func NewEngine(o Operator) *Engine {
	return &Engine{
		Route: o,
	}
}

func (e *Engine) Start() {
	fmt.Println("Starting the engine")
	e.Route.Validate(context.Background(), nil)
	e.Route.PreProcess(context.Background(), nil)
	e.Route.Process(context.Background(), nil)
	e.Route.PostProcess(context.Background(), nil)
	e.Route.AddChangeNotifiy(context.Background(), nil)
	e.Route.RemoveNotify(context.Background(), nil)
	fmt.Println("Engine started successfully")

}
