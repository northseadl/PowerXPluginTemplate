package powerx

import "PluginTemplate/pkg/powerx/client"

type PowerX struct {
	*client.PClient
	AdminEmployee *AdminEmployee
}

func NewPowerX(endpoint string, debug bool) *PowerX {
	power := &PowerX{
		PClient: client.NewPClient(endpoint, debug),
	}
	power.AdminEmployee = &AdminEmployee{power}
	return power
}
