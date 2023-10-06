package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	PowerXPlugin
	Endpoint string
	Mode     string
}

type PowerXPlugin struct {
	AppKey string
}
