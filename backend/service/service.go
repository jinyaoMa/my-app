package service

import "context"

var instance *service

type service struct {
	ctx context.Context
}

func init() {
	instance = &service{}
}

func Service(ctxs ...context.Context) *service {
	if len(ctxs) > 0 {
		instance.ctx = ctxs[0]
	}
	return instance
}
