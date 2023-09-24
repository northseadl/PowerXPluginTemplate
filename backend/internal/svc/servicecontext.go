package svc

import (
	"PluginTemplate/internal/config"
	"PluginTemplate/internal/middleware"
	"PluginTemplate/pkg/powerx/client"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config           config.Config
	PluginMiddleware rest.Middleware
	*client.PClient
}

func NewServiceContext(c config.Config, client *client.PClient) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		PluginMiddleware: middleware.NewPluginMiddleware().Handle,
		PClient:          client,
	}
}
