package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"fmt"
	"net/http"
)

type AdminAuth struct {
	*PowerX
}

func (c *AdminAuth) Login(req *powerxtypes.LoginRequest) (*powerxtypes.LoginReply, error) {
	res := &powerxtypes.LoginReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri("/api/v1/admin/auth/access/actions/basic-login").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminAuth) Exchange(req *powerxtypes.ExchangeRequest) (*powerxtypes.ExchangeReply, error) {
	res := &powerxtypes.ExchangeReply{}
	err := c.H.Df().Method(http.MethodPost).
		Uri(fmt.Sprintf("/api/v1/admin/auth/access/actions/exchange-token", req.Type)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
