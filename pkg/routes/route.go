package routes

import (
	"context"
	"fmt"

	"github.com/shubhashreeb/cfg-state-machine/mgr"
)

type Route struct {
}

func NewRoute() mgr.Operator {
	return &Route{}
}

// Validate checks if an object has mandatory config fields
func (r *Route) Validate(ctx context.Context, data interface{}) error {
	fmt.Println("Validating data:")
	return nil
}

func (r *Route) PreProcess(ctx context.Context, data interface{}) (interface{}, error) {
	fmt.Println("Pre-processing data:")
	return data, nil
}
func (r *Route) Process(ctx context.Context, data interface{}) (interface{}, error) {
	fmt.Println("Processing data:")
	return data, nil
}
func (r *Route) PostProcess(ctx context.Context, data interface{}) (interface{}, error) {
	fmt.Println("Post-processing data:")
	return data, nil
}

func (r *Route) AddChangeNotifiy(ctx context.Context, data interface{}) interface{} {
	fmt.Println("Adding change notification for data:")
	return data
}
func (r *Route) RemoveNotify(ctx context.Context, id interface{}) interface{} {
	fmt.Println("Removing notification for id:")
	return id

}
