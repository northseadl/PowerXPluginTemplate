package plugin

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"log/slog"
	"path/filepath"
)

type Config interface{}

func CreateRestConf(mode, name string) rest.RestConf {
	if mode == "dev" {
		return DefaultRestConfForDev(name)
	}

	restConf, err := DefaultRestConfWithRandomPort(name)
	if err != nil {
		slog.Error(fmt.Sprintf("Get random port failed: %s", err.Error()))
	}
	return restConf
}

func LoadConfigFromEtcYaml(configPath string, c any) {
	absPath, err := filepath.Abs(configPath)
	if err != nil {
		slog.Error(fmt.Sprintf("Get abs path of %s failed: %s", configPath, err.Error()))
	}
	conf.MustLoad(absPath, c)
}

func RegisterPlugin(host string, c any, restPlugin *Plugin) {
	etcMap, err := RegisterPluginToHost(host, restPlugin)
	if err != nil {
		slog.Error(fmt.Sprintf("Register plugin to host %s failed: %s", host, err.Error()))
	}
	if err = DecodeMapToStruct(etcMap, c); err != nil {
		return
	}
}
