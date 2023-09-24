package main

import (
	"PluginTemplate/internal/config"
	"PluginTemplate/internal/handler"
	"PluginTemplate/internal/svc"
	"PluginTemplate/pkg/plugin"
	"PluginTemplate/pkg/powerx/client"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest"
	"log/slog"
)

// 当前插件的插件名
var name = flag.String("n", "plugin", "the name")

// 主程序的地址
var host = flag.String("h", "localhost:8888", "the host")

func main() {
	flag.Parse()
	// 从主程序加载配置，如果时开发环境可以从etc加载

	// 对于插件，这些选项由程序定义
	restConf, err := plugin.DefaultRestConfWithRandomPort(*name)
	server := rest.MustNewServer(restConf)
	defer server.Stop()

	var c config.Config
	c.RestConf = restConf

	// 创建主程序请求客户端
	client := client.NewPClient(fmt.Sprintf("http://%s", *host))

	ctx := svc.NewServiceContext(c, client)
	handler.RegisterHandlers(server, ctx)

	// 解析插件信息并注册到主程序
	restPlugin := plugin.ParsePlugin(server, c.RestConf)
	err = plugin.RegisterPluginToHost(*host, restPlugin)
	if err != nil {
		slog.Error(fmt.Sprintf("Register plugin %s to host %s failed: %s", *name, *host, err.Error()))
	}

	slog.Info(fmt.Sprintf("Plugin %s start at %s", *name, restPlugin.Data()["addr"]))
	server.Start()
}
