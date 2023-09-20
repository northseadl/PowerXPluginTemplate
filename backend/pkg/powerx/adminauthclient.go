package powerx

import (
	"PluginTemplate/pkg/powerx/types"
	"fmt"
	"net/http"
)

type AdminAuth struct {
	*PowerX
}

func (c *AdminAuth) Login(req *types.LoginRequest) (*types.LoginReply, error) {
	res := &types.LoginReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/auth/access/actions/basic-login").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminAuth) Exchange(req *types.ExchangeRequest) (*types.ExchangeReply, error) {
	res := &types.ExchangeReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri(fmt.Sprintf("/api/v1/admin/auth/access/actions/exchange-token", req.Type)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
