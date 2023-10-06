package plugin

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
)

// DefaultRestConf 返回一个默认的 gozero rest 配置
func DefaultRestConf(name string, port int) rest.RestConf {
	return rest.RestConf{
		ServiceConf: service.ServiceConf{
			Name: name,
		},
		Host: "0.0.0.0",
		Port: port,
	}
}

// DefaultRestConfWithRandomPort 返回一个默认的 gozero rest 配置，端口随机
func DefaultRestConfWithRandomPort(name string) (rest.RestConf, error) {
	port, err := GetRandomPort()
	if err != nil {
		return rest.RestConf{}, err
	}
	return DefaultRestConf(name, port), nil
}

// DefaultRestConfForDev 返回一个默认的 gozero rest 配置，端口 8999
func DefaultRestConfForDev(name string) rest.RestConf {
	return DefaultRestConf(name, 8999)
}
