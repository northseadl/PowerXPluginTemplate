package powerx

import (
	"PluginTemplate/pkg/powerx/powerxtypes"
	"context"

	"net/http"
)

type WebCustomerAuthOa struct {
	*PowerX
}

func (c *WebCustomerAuthOa) OALogin(ctx context.Context, req *powerxtypes.OACustomerLoginRequest) (*powerxtypes.OACustomerLoginAuthReply, error) {
	res := &powerxtypes.OACustomerLoginAuthReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/web/customer/oa/login").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *WebCustomerAuthOa) AuthByPhone(ctx context.Context, req *powerxtypes.OACustomerAuthRequest) (*powerxtypes.OACustomerLoginAuthReply, error) {
	res := &powerxtypes.OACustomerLoginAuthReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/web/customer/oa/authByPhone").
		Json(req).
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *WebCustomerAuthOa) AuthByProfile(ctx context.Context) (*powerxtypes.OACustomerLoginAuthReply, error) {
	res := &powerxtypes.OACustomerLoginAuthReply{}
	err := c.H.Df().Method(http.MethodPost).
		WithContext(ctx).
		Uri("/api/v1/web/customer/oa/authByProfile").
		Result(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
