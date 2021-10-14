package svc

import "github.com/pinkhello/hathor/core/conf"

type ServiceContext struct {
	Config conf.Config
}

func NewServiceContext(c conf.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
