package powerx

import "PluginTemplate/pkg/powerx/client"

type PowerX struct {
	*client.PClient
}

func NewPowerX(endpoint string) *PowerX {
	return &PowerX{
		PClient: client.NewPClient(endpoint),
	}
}
