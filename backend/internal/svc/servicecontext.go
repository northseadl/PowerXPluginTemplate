package svc

import (
	"PluginTemplate/internal/config"
	"PluginTemplate/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config           config.Config
	PluginMiddleware rest.Middleware
	*server.PClient
}

func NewServiceContext(c config.Config, client *server.PClient) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		PluginMiddleware: middleware.NewPluginMiddleware().Handle,
		PClient:          client,
	}
}
