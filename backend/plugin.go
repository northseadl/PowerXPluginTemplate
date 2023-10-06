package main

import (
	"PluginTemplate/internal/config"
	"PluginTemplate/internal/handler"
	"PluginTemplate/internal/svc"
	"PluginTemplate/pkg/plugin"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest"
	"log/slog"
)

var (
	name       = flag.String("n", "plugin", "the name")
	host       = flag.String("h", "localhost:8888", "the host")
	mode       = flag.String("m", "dev", "the mode")
	configPath = flag.String("c", "etc/etc.yaml", "the config file")
)

func main() {
	flag.Parse()

	restConf := plugin.CreateRestConf(*mode, *name)
	server := rest.MustNewServer(restConf)
	defer server.Stop()

	c := config.Config{RestConf: restConf}
	c.Endpoint = fmt.Sprintf("http://%s", *host)
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	restPlugin := plugin.ParsePlugin(server, c.RestConf)

	if *mode == "dev" {
		plugin.LoadConfigFromEtcYaml(*configPath, &c.PowerXPlugin)
	} else {
		plugin.RegisterPlugin(*host, &c.PowerXPlugin, restPlugin)
	}

	ctx.Setup()

	slog.Info(fmt.Sprintf("Plugin %s start at %s", *name, restPlugin.Data()["addr"]))
	server.Start()
}
