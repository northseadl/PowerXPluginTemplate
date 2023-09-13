package plugin

import (
	"fmt"
	"net"
	"strings"

	"github.com/zeromicro/go-zero/rest"
)

type Plugin struct {
	Routes []Route `json:"routes,omitempty"`
	Port   int     `json:"port,omitempty"`
	Name   string  `json:"name,omitempty"`
}

type Route struct {
	Method string `json:"method,omitempty"`
	Path   string `json:"path,omitempty"`
}

func (p *Plugin) Data() map[string]interface{} {
	return map[string]interface{}{
		"name":   strings.ToLower(p.Name),
		"addr":   fmt.Sprintf("localhost:%d", p.Port),
		"routes": p.Routes,
	}
}

func (p *Plugin) SetName(name string) {
	p.Name = name
}

func (p *Plugin) SetPort(port int) {
	p.Port = port
}

func ParsePlugin(server *rest.Server, c rest.RestConf) *Plugin {
	var routes []Route
	for _, route := range server.Routes() {
		routes = append(routes, Route{
			Method: route.Method,
			Path:   route.Path,
		})
	}

	return &Plugin{
		Routes: routes,
		Port:   c.Port,
		Name:   c.Name,
	}
}

func GetRandomPort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}

	port := listener.Addr().(*net.TCPAddr).Port
	_ = listener.Close()

	return port, nil
}
