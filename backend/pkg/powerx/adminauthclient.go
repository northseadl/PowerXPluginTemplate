package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"
	"fmt"
	"net/http"
)

type AdminAuth struct {
	*PowerX
}

func (c *AdminAuth) Login(ctx context.Context, req *powerxtypes.LoginRequest) (*powerxtypes.LoginReply, error) {
	res := &powerxtypes.LoginReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/admin/auth/access/actions/basic-login").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AdminAuth) Exchange(ctx context.Context, req *powerxtypes.ExchangeRequest) (*powerxtypes.ExchangeReply, error) {
	res := &powerxtypes.ExchangeReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri(fmt.Sprintf("/api/v1/admin/auth/access/actions/exchange-token", req.Type)).
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
