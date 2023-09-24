// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"PluginTemplate/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.PluginMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/example",
					Handler: ExampleHandler(serverCtx),
				},
			}...,
		),
	)
}
