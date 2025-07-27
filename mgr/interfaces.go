package mgr

import "context"

type Validator interface {
	//Validate checks if an object has mandatory config fields
	Validate(ctx context.Context, data interface{}) error
}

type Transformer interface {
	PreProcess(ctx context.Context, data interface{}) (interface{}, error)
	Process(ctx context.Context, data interface{}) (interface{}, error)
	PostProcess(ctx context.Context, data interface{}) (interface{}, error)
}

type Notifier interface {
	AddChangeNotifiy(ctx context.Context, data interface{}) interface{}
	RemoveNotify(ctx context.Context, id interface{}) interface{}
}

type Operator interface {
	Validator
	Transformer
	Notifier
}

