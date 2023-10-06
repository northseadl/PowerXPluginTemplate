package svc

import (
	"PluginTemplate/internal/config"
	"PluginTemplate/internal/middleware"
	"PluginTemplate/pkg/powerx"
	"PluginTemplate/pkg/powerx/client"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config           config.Config
	PluginMiddleware rest.Middleware
	*client.PClient
	PowerX *powerx.PowerX
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		PluginMiddleware: middleware.NewPluginMiddleware().Handle,
	}
}

func (c *ServiceContext) Setup() {
	c.PowerX = powerx.NewPowerX(c.Config.Endpoint, c.Config.Mode == "dev")
}
